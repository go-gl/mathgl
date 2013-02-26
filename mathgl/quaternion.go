package mathgl

import ()

// A Quaternion is a representation of a 3D rotation, it is an extension of the complex numbers
// The quaternion is represented by 4 characteristic "pieces": 1, i, j, k
//
// This can also be represented by {w,v} where w is some scalar value and v is a vector.
//
// Our implementation uses the scalar/vector form in, predictably, representing it with the types Scalar and Vector.
// In addition, it encodes the underlying VecType of the Quaternion.
type Quaternion struct {
	w   Scalar
	v   Vector
	typ VecType
}

// Returns a pointer to a Quaternion with Scalar part w, Vector part v, and type typ.
// If the scalar or vector is of the incorrect type, or the vector is not of length 3 you get nil
//
// Unlike Matrix and Vector, using this method for a new Quaternion is recommended.
func NewQuaternion(w Scalar, v Vector, typ VecType) *Quaternion {
	if !checkType(typ, w) || v.typ != typ || v.Size() != 3 {
		return nil
	}
	return &Quaternion{w, v, typ}
}

// QuaternionIdentity returns the Quaternion identity:
//
// 1 * 1 + 0 * i + 0 * j + 0 * k
func QuaternionIdentity(typ VecType) Quaternion {
	v, _ := VectorOf(ScalarSlice([]interface{}{0, 0, 0}, typ), typ)
	return Quaternion{MakeScalar(1, typ), *v, typ}
}

// ScalarComponent predictably returns w, the Scalar components of the Quaternion
func (q Quaternion) ScalarComponent() Scalar {
	return q.w
}

// VectorComponent predictably returns v, the Vector components of the Quaternion
func (q Quaternion) VectorComponent() Vector {
	return q.v
}

// Add does a traditional addition of two Quaternions
//
// Result = (w1+w2) 1 + (v1.x+v2.x) i + (v1.y+v2.y) j + (v1.z + v2.z) k = w1+w2,v1+v2
//
// It returns the zero-type for Quaternion if their types aren't the same
func (q1 Quaternion) Add(q2 Quaternion) (q3 Quaternion) {
	if q1.typ != q2.typ {
		return
	}

	return Quaternion{q1.w.Add(q2.w), q1.v.Add(q2.v), q1.typ}
}

// Add does a traditional addition of two Quaternions
//
// Result = (w1-w2) 1 + (v1.x-v2.x) i + (v1.y-v2.y) j + (v1.z-v2.z) k = {w1-w2,v1-v2}
//
// It returns the zero-type for Quaternion if their types aren't the same
func (q1 Quaternion) Sub(q2 Quaternion) (q3 Quaternion) {
	if q1.typ != q2.typ {
		return
	}

	return Quaternion{q1.w.Sub(q2.w), q1.v.Sub(q2.v), q1.typ}
}

// Mul Multiplies two Quaternions and returns the result, this works out to:
//
// {w1 * w2 - v1.v2, v1xv2 + w1v2 + w2v1}\
//
// It returns the zero-type for the Quaternion if the underlying types aren't equal
func (q1 Quaternion) Mul(q2 Quaternion) (q3 Quaternion) {
	if q1.typ != q2.typ {
		return
	}

	return Quaternion{q1.w.Mul(q2.w).Sub(q1.v.Dot(q2.v)), q1.v.Cross(q2.v).Add(q2.v.ScalarMul(q1.w)).Add(q1.v.ScalarMul(q2.w)), q1.typ}
}

// Conjugate returns the Quaternion's Conjugate,
//
// {w, -v}
func (q Quaternion) Conjugate() Quaternion {
	return Quaternion{q.w, q.v.ScalarMul(MakeScalar(-1, q.typ)), q.typ}
}

// This returns the Quaternion's Length as a float64 (for similar reasons to Vector)
//
// w^2 + x^2 + y^2 + z^2
func (q Quaternion) Len() float64 {
	return q.w.Fl64()*q.w.Fl64() + q.v.dat[0].Fl64()*q.v.dat[0].Fl64() + q.v.dat[1].Fl64()*q.v.dat[1].Fl64() + q.v.dat[2].Fl64()*q.v.dat[2].Fl64()
}

// Normalize returns a normalized Quaternion. Similarly to a vector, this is equivalent to dividing it through by its Length
// As is the case with Vector, if the Quaternion has a zero-length, or is of Length 1 (already Normalized), it shortcuts and returns itself
func (q Quaternion) Normalize() Quaternion {
	length := q.Len()
	if FloatEqual(length, 0.) || FloatEqual(length, 1.) {
		return q
	}

	return Quaternion{q.w.mulFl64(length), q.v.floatScale(length), q.typ}
}

// Equal does a piecewise comparison of two Vectors. This is equivalent to comparing its scalar part and vector parts directly and logically AND-ing the result
func (q1 Quaternion) Equal(q2 Quaternion) bool {
	return q1.w.Equal(q2.w) && q1.v.Equal(q2.v)
}

// ToHomogRotationMatrix will convert the Quaternion to the Homogeneous 3D (4x4) Rotation Matrix that this Quaternion is equivalent to.
// It is converted to a FLOAT64 matrix no matter what
//
//
// Where w = q.w, and x,y,z are the first, second, and third elements of the underlying vector (respectively), this gives us
//
//    [[ 1-2y^2-2z^2, 2xy-2wz    , 2xz+2wy    , 0 ]]
//    [[ 2xy+2wz    , 1-2x^2-2z^2, 2yz+2wz    , 0 ]]
//    [[ 2xz-2wy    , 2yz -2wx   , 1-2x^2-2y^2, 0 ]]
//    [[ 0          , 0          , 0          , 1 ]]
//
// The Quaternion must be Normalized before-hand. If len != 1, this will return the zero type for Matrix
func (q Quaternion) ToHomogRotationMatrix() Matrix {
	if FloatEqual(q.Len(), 0.) {
		return Matrix{}
	}

	w, x, y, z := q.w.Fl64(), q.v.dat[0].Fl64(), q.v.dat[1].Fl64(), q.v.dat[2].Fl64()

	dat := ScalarSlice([]interface{}{
		1. - 2.*y*y - 2.*z*z, 2.*x*y - 2.*w*z, 2.*x*z + 2.*w*y, 0.,
		2.*x*y + 2.*w*z, 1. - 2.*x*x - 2*z*z, 2.*y*z + 2.*w*z, 0.,
		2.*x*z - 2.*w*y, 2.*y*z - 2.*w*x, 1. - 2.*x*x - 2.*y*y, 0.,
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
