package mathgl

import (
	"math"
)

type Quatd struct {
	W float64
	V Vec3d
}

func QuatIdentd() Quatd {
	return Quatd{1., Vec3d{0, 0, 0}}
}

func QuatRotated(angle float64, axis Vec3d) Quatd {
	angle = (float64(math.Pi) * angle) / 180.0

	c, s := float64(math.Cos(float64(angle/2))), float64(math.Sin(float64(angle/2)))

	return Quatd{c, axis.Mul(s)}
}

func (q Quatd) X() float64 {
	return q.V[0]
}

func (q Quatd) Y() float64 {
	return q.V[1]
}

func (q Quatd) Z() float64 {
	return q.V[2]
}

func (q1 Quatd) Add(q2 Quatd) Quatd {
	return Quatd{q1.W + q2.W, q1.V.Add(q2.V)}
}

func (q1 Quatd) Sub(q2 Quatd) Quatd {
	return Quatd{q1.W - q2.W, q1.V.Sub(q2.V)}
}

func (q1 Quatd) Mul(q2 Quatd) Quatd {
	return Quatd{q1.W*q2.W - q1.V.Dot(q2.V), q1.V.Cross(q2.V).Add(q2.V.Mul(q1.W)).Add(q1.V.Mul(q2.W))}
}

func (q1 Quatd) Scale(c float64) Quatd {
	return Quatd{q1.W * c, Vec3d{q1.V[0] * c, q1.V[1] * c, q1.V[2] * c}}
}

func (q1 Quatd) Conjugate() Quatd {
	return Quatd{q1.W, q1.V.Mul(-1)}
}

func (q1 Quatd) Len() float64 {
	return float64(math.Sqrt(float64(q1.W*q1.W + q1.V[0]*q1.V[0] + q1.V[1]*q1.V[1] + q1.V[2]*q1.V[2])))
}

func (q1 Quatd) Normalize() Quatd {
	length := q1.Len()

	if FloatEqual(1, float64(length)) {
		return q1
	}

	return Quatd{q1.W * 1 / length, q1.V.Mul(1 / length)}
}

func (q1 Quatd) Inverse() Quatd {
	leng := q1.Len()
	return Quatd{q1.W, q1.V.Mul(-1)}.Scale(1 / (leng * leng))
}

func (q1 Quatd) Rotate(v Vec3d) Vec3d {
	return q1.Mul(Quatd{0, v}).Mul(q1.Conjugate()).V
}

func (q1 Quatd) Mat4() Mat4d {
	w, x, y, z := q1.W, q1.V[0], q1.V[1], q1.V[2]
	return Mat4d{1 - 2*y*y - 2*z*z, 2*x*y + 2*w*z, 2*x*z - 2*w*y, 0, 2*x*y - 2*w*z, 1 - 2*x*x - 2*z*z, 2*y*z + 2*w*x, 0, 2*x*z + 2*w*y, 2*y*z - 2*w*x, 1 - 2*x*x - 2*y*y, 0, 0, 0, 0, 1}
}

func (q1 Quatd) Dot(q2 Quatd) float64 {
	return q1.W*q1.W + q1.V[0]*q1.V[0] + q1.V[1]*q1.V[1] + q1.V[2]*q1.V[2]
}

func QuatSlerpd(q1, q2 Quatd, amount float64) Quatd {
	q1, q2 = q1.Normalize(), q2.Normalize()
	dot := q1.Dot(q2)

	// This is here for precision errors, I'm perfectly aware the *technically* the dot is bound [-1,1], but since Acos will freak out if it's not (even if it's just a liiiiitle bit over due to normal error) we need to clamp it
	dot = Clampd(dot, -1, 1)

	theta := float64(math.Acos(float64(dot))) * amount
	c, s := float64(math.Cos(float64(theta))), float64(math.Sin(float64(theta)))
	rel := q2.Sub(q1.Scale(dot)).Normalize()

	return q2.Sub(q1.Scale(c)).Add(rel.Scale(s))
}

func QuatLerpd(q1, q2 Quatd, amount float64) Quatd {
	return q1.Add(q2.Sub(q1).Scale(amount))
}

func QuatNlerpd(q1, q2 Quatd, amount float64) Quatd {
	return QuatLerpd(q1, q2, amount).Normalize()
}
