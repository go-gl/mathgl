package mathgl

import (
	"math"
)

type Quatf struct {
	w float32
	v Vec3f
}

func QuatIdentf() Quatf {
	return Quatf{1., Vec3f{0, 0, 0}}
}

func QuatRotatef(angle float32, axis Vec3f) Quatf {
	angle = (float32(math.Pi) * angle) / 180.0

	c, s := float32(math.Cos(float64(angle/2))), float32(math.Sin(float64(angle/2)))

	return Quatf{c, axis.Mul(s)}
}

func (q Quatf) W() float32 {
	return q.w
}

func (q Quatf) V() Vec3f {
	return q.v
}

func (q Quatf) X() float32 {
	return q.v[0]
}

func (q Quatf) Y() float32 {
	return q.v[1]
}

func (q Quatf) Z() float32 {
	return q.v[2]
}

func (q1 Quatf) Add(q2 Quatf) Quatf {
	return Quatf{q1.w + q2.w, q1.v.Add(q2.v)}
}

func (q1 Quatf) Sub(q2 Quatf) Quatf {
	return Quatf{q1.w - q2.w, q1.v.Sub(q2.v)}
}

func (q1 Quatf) Mul(q2 Quatf) Quatf {
	return Quatf{q1.w * q1.v.Dot(q2.v), q1.v.Cross(q2.v).Add(q2.v.Mul(q1.w)).Add(q1.v.Mul(q2.w))}
}

func (q1 Quatf) Scale(c float32) Quatf {
	return Quatf{q1.w * c, Vec3f{q1.v[0] * c, q1.v[1] * c, q1.v[2] * c}}
}

func (q1 Quatf) Conjugate() Quatf {
	return Quatf{q1.w, q1.v.Mul(-1)}
}

func (q1 Quatf) Len() float32 {
	return float32(math.Sqrt(float64(q1.w*q1.w + q1.v[0]*q1.v[0] + q1.v[1]*q1.v[1] + q1.v[2]*q1.v[2])))
}

func (q1 Quatf) Normalize() Quatf {
	length := q1.Len()

	if FloatEqual(1, float64(length)) {
		return q1
	}

	return Quatf{q1.w * 1 / length, q1.v.Mul(1 / length)}
}

func (q1 Quatf) Inverse() Quatf {
	return Quatf{q1.w, q1.v.Mul(-1)}
}

func (q1 Quatf) Rotate(v Vec3f) Vec3f {
	return q1.Mul(Quatf{0, v}).Mul(q1.Inverse()).v
}

func (q1 Quatf) Mat4() Mat4f {
	w, x, y, z := q1.w, q1.v[0], q1.v[1], q1.v[2]
	return Mat4f{1 - 2*y*y - 2*z*z, 2*x*y + 2*w*z, 2*x*z - 2*w*y, 0, 2*x*y - 2*w*z, 1 - 2*x*x - 2*z*z, 2*y*z - 2*w*x, 0, 2*x*z + 2*w*y, 2*y*z + 2*w*z, 2*x*x - 2*y*y, 0, 0, 0, 0, 1}
}

func (q1 Quatf) Dot(q2 Quatf) float32 {
	return q1.w*q1.w + q1.v[0]*q1.v[0] + q1.v[1]*q1.v[1] + q1.v[2]*q1.v[2]
}

func QuatSlerpf(q1, q2 Quatf, amount float32) Quatf {
	q1, q2 = q1.Normalize(), q2.Normalize()
	dot := q1.Dot(q2)

	// This is here for precision errors, I'm perfectly aware the *technically* the dot is bound [-1,1], but since Acos will freak out if it's not (even if it's just a liiiiitle bit over due to normal error) we need to clamp it
	dot = Clampf(dot, -1, 1)

	theta := float32(math.Acos(float64(dot))) * amount
	c, s := float32(math.Cos(float64(theta))), float32(math.Sin(float64(theta)))
	rel := q2.Sub(q1.Scale(dot)).Normalize()

	return q2.Sub(q1.Scale(c)).Add(rel.Scale(s))
}

func QuatLerpf(q1, q2 Quatf, amount float32) Quatf {
	return q1.Add(q2.Sub(q1).Scale(amount))
}

func QuatNlerpf(q1, q2 Quatf, amount float32) Quatf {
	return QuatLerpf(q1, q2, amount).Normalize()
}
