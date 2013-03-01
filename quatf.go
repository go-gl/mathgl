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

	c, s := float32(math.Cos(float64(angle))), float32(math.Sin(float64(angle)))
	return Quatf{c / 2, axis.Mul(s / 2)}
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

func (q1 Quatf) Conjugate() Quatf {
	return Quatf{q1.w, q1.v.Mul(-1)}
}

func (q1 Quatf) Len() float32 {
	return q1.w*q1.w + q1.v[0]*q1.v[0] + q1.v[1]*q1.v[1] + q1.v[2]*q1.v[2]
}

func (q1 Quatf) Normalize() Quatf {
	length := float32(1.) / q1.Len()
	return Quatf{q1.w * length, q1.v.Mul(length)}
}

func (q1 Quatf) Mat4() Mat4f {
	w, x, y, z := q1.w, q1.v[0], q1.v[1], q1.v[2]
	return Mat4f{1 - 2*y*y - 2*z*z, 2*x*y + 2*w*z, 2*x*z - 2*w*y, 0, 2*x*y - 2*w*z, 1 - 2*x*x - 2*z*z, 2*y*z - 2*w*x, 0, 2*x*z + 2*w*y, 2*y*z + 2*w*z, 2*x*x - 2*y*y, 0, 0, 0, 0, 1}
}
