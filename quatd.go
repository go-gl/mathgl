package mathgl

import (
	"math"
)

type Quatd struct {
	w float64
	v Vec3d
}

func QuatIdentd() Quatd {
	return Quatd{1., Vec3d{0, 0, 0}}
}

func QuatRotated(angle float64, axis Vec3d) Quatd {
	angle = (float64(math.Pi) * angle) / 180.0

	c, s := float64(math.Cos(float64(angle))), float64(math.Sin(float64(angle)))
	return Quatd{c / 2, axis.Mul(s / 2)}
}

func (q Quatd) W() float64 {
	return q.w
}

func (q Quatd) V() Vec3d {
	return q.v
}

func (q Quatd) X() float64 {
	return q.v[0]
}

func (q Quatd) Y() float64 {
	return q.v[1]
}

func (q Quatd) Z() float64 {
	return q.v[2]
}

func (q1 Quatd) Add(q2 Quatd) Quatd {
	return Quatd{q1.w + q2.w, q1.v.Add(q2.v)}
}

func (q1 Quatd) Sub(q2 Quatd) Quatd {
	return Quatd{q1.w - q2.w, q1.v.Sub(q2.v)}
}

func (q1 Quatd) Mul(q2 Quatd) Quatd {
	return Quatd{q1.w * q1.v.Dot(q2.v), q1.v.Cross(q2.v).Add(q2.v.Mul(q1.w)).Add(q1.v.Mul(q2.w))}
}

func (q1 Quatd) Conjugate() Quatd {
	return Quatd{q1.w, q1.v.Mul(-1)}
}

func (q1 Quatd) Len() float64 {
	return q1.w*q1.w + q1.v[0]*q1.v[0] + q1.v[1]*q1.v[1] + q1.v[2]*q1.v[2]
}

func (q1 Quatd) Normalize() Quatd {
	length := float64(1.) / q1.Len()
	return Quatd{q1.w * length, q1.v.Mul(length)}
}

func (q1 Quatd) Mat4() Mat4d {
	w, x, y, z := q1.w, q1.v[0], q1.v[1], q1.v[2]
	return Mat4d{1 - 2*y*y - 2*z*z, 2*x*y + 2*w*z, 2*x*z - 2*w*y, 0, 2*x*y - 2*w*z, 1 - 2*x*x - 2*z*z, 2*y*z - 2*w*x, 0, 2*x*z + 2*w*y, 2*y*z + 2*w*z, 2*x*x - 2*y*y, 0, 0, 0, 0, 1}
}
