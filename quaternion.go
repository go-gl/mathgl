package mathgl

import (
	"math"
)

type Quaternion struct {
	w   Scalar
	v   Vector
	typ VecType
}

func NewQuaternion(w Scalar, v Vector, typ VecType) *Quaternion {
	if !checkType(typ, w) || v.typ != typ || v.Size() != 3 {
		return &Quaternion{}
	}
	return &Quaternion{w, v, typ}
}

func QuaternionIdentity(typ VecType) Quaternion {
	v, _ := VectorOf(ScalarSlice([]interface{}{0, 0, 0}, typ), typ)
	return Quaternion{MakeScalar(1, typ), *v, typ}
}

func (q Quaternion) ScalarComponent() Scalar {
	return q.w
}

func (q Quaternion) VectorComponent() Vector {
	return q.v
}

func (q1 Quaternion) Add(q2 Quaternion) (q3 Quaternion) {
	if q1.typ != q2.typ {
		return
	}

	return Quaternion{q1.w.Add(q2.w), q1.v.Add(q2.v), q1.typ}
}

func (q1 Quaternion) Mul(q2 Quaternion) (q3 Quaternion) {
	if q1.typ != q2.typ {
		return
	}

	return Quaternion{q1.w.Mul(q2.w).Sub(q1.v.Dot(q2.v)), q1.v.Cross(q2.v).Add(q2.v.ScalarMul(q1.w)).Add(q1.v.ScalarMul(q2.w)), q1.typ}
}

func (q Quaternion) Conjugate() Quaternion {
	return Quaternion{q.w, q.v.ScalarMul(MakeScalar(-1, q.typ)), q.typ}
}

func (q Quaternion) Len() float64 {
	return q.w.Fl64()*q.w.Fl64() + q.v.dat[0].Fl64()*q.v.dat[0].Fl64() + q.v.dat[1].Fl64()*q.v.dat[1].Fl64() + q.v.dat[2].Fl64()*q.v.dat[2].Fl64()
}

func (q Quaternion) Normalize() Quaternion {
	length := q.Len()
	if math.Abs(length) < 1e-7 {
		return q
	}

	return Quaternion{q.w.mulFl64(length), q.v.floatScale(length), q.typ}
}

func (q1 Quaternion) Equal(q2 Quaternion) bool {
	return q1.w.Equal(q2.w) && q1.v.Equal(q2.v)
}

// Always returns a matrix of type FLOAT64
func (q Quaternion) ToHomogRotationMatrix() Matrix {
	if math.Abs(q.Len()) > 1e-7 {
		return Matrix{}
	}

	w, x, y, z := q.w.Fl64(), q.v.dat[0].Fl64(), q.v.dat[1].Fl64(), q.v.dat[2].Fl64()

	dat := ScalarSlice([]interface{}{1. - 2.*y*y - 2.*z*z, 2.*x*y - 2*w*z, 2.*x*z + 2*w*y, 0.,
		2.*x*y + 2*w*z, 1. - 2.*x*x - 2*z*z, 2.*y*z + 2.*w*z, 0.,
		2.*x*z - 2*w*y, 2.*y*z - 2*w*x, 1. - 2.*x*x - 2.*y*y, 0.,
		0., 0., 0., 1.}, FLOAT64)
	return *unsafeMatrixFromSlice(dat, q.typ, 4, 4)
}

/*func (q Quaternion) ToHomogRotationMatrix() Matrix {
	if math.Abs(q.Len()) > 1e-7 {
		return Matrix{}
	}

	w, x, y, z := q.w, q.v.dat[0], q.v.dat[1], q.v.dat[2]
	zero, one, two := MakeScalar(0, q.typ), MakeScalar(1, q.typ), MakeScalar(2, q.typ)

	// Kinda ugly, hrm...
	dat := []Scalar{one.Sub(two.Mul(y.Mul(y))).Sub(two.Mul(z.Mul(z))), two.Mul(x.Mul(y)).Sub(two.Mul(w.Mul(z))), two.Mul(x.Mul(z)).Add(two.Mul(w.Mul(y))), zero,
		two.Mul(x.Mul(y)).Add(two.Mul(w.Mul(z))), one.Sub(two.Mul(x.Mul(x))).Sub(two.Mul(z.Mul(z))), two.Mul(y.Mul(z)).Add(two.Mul(w.Mul(x))), zero,
		two.Mul(x.Mul(z)).Sub(two.Mul(w.Mul(y))), two.Mul(y.Mul(z)).Sub(two.Mul(w.Mul(x))), one.Sub(two.Mul(x.Mul(x))).Sub(two.Mul(y.Mul(y))), zero,
		zero, zero, zero, one}
	return *unsafeMatrixFromSlice(dat, q.typ, 4, 4)
}*/
