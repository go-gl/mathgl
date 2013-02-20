/*
Package mathgl is an open source 3D vector, matrix, and quaternion library. It is intended to serve as a useful helper to using OpenGL just as GLM works for C++, except being a tad bit more Go-ish.

That said, this package has to work around Go's lack of Generics, but generally the package works rather cleanly once you get your variables set up. It has helper functions that will convert
your data into an Array for use with functions size as glUniformMatrix[...] and so on.

The documentation was generally written at a level that assumes you understand basic math, and will not explain the USE of the functions (i.e. why you'd want a Determinant), but
it will explain exactly what the function does.

Generally in this package anything in the documentation enclosed by [Single Brackets] is a Vector

[a]
[b] = Column Vector

[a, b] = Row Vector

And anything enclosed by [[Double Brackets]] is a Matrix

[[ a, b ]]
[[ c, d ]]

(These notations don't work in the godoc, only in the actual source comments)
*/
package mathgl

import (
	"errors"
)

/*
A VecType is a simple ID for what type our vector(/matrix/quaternion) is
This is basically to sneak around lack of generics without using extreme code duplication tricks and
scripts

The associated consts should be self explanatory, as they correspond to the actual built-in go type, except in ALLCAPS
*/
type VecType int8

const (
	NOTYPE = iota
	INT32
	UINT32
	FLOAT32
	FLOAT64
)

/*
A Vector is essentially a wrapper for a slice of Scalars, with an extra VecType variable to signify the underlying type

A Vector is both a row and a column vector at once (whichever is convenient at the moment), the orientation only matters once converted to a Matrix
*/
type Vector struct {
	typ VecType
	dat []Scalar
}

// NewVector Returns an empty Vector of type t, ready to be added to
// It's recommended you use VectorOf() instead
func NewVector(t VecType) *Vector {
	return &Vector{typ: t, dat: make([]Scalar, 0, 2)}
}

// VectorOf takes in a list of Scalar objects and a type, and returns a pointer to a vector,
// if not all of the elements of el are of type t, an error is returned
func VectorOf(el []Scalar, t VecType) (v *Vector, err error) {
	for _, e := range el {
		if !checkType(t, e) {
			return nil, errors.New("Type of at least one element does not match declared type")
		}
	}

	return &Vector{t, el}, nil
}

// AddElements adds all the elements of the slice el to the vector in-place (as in, it modified the vector)
// It returns an error iff any element of el does not match v's type. In this scenario, the vector is not altered
func (v *Vector) AddElements(el []Scalar) error {
	for _, e := range el {
		if !checkType(v.typ, e) {
			return errors.New("Type of at least one element does not match vector's type")
		}
	}

	v.dat = append(v.dat, el...)
	return nil
}

// SetElement will change the element at zero-based index loc to Scalar el.
// It will return an error if loc is out of bounds or el is not of vec's type
func (v *Vector) SetElement(loc int, el Scalar) error {
	if !checkType(v.typ, el) {
		return errors.New("Element does not match vector's type")
	}

	if loc < 0 || loc > len(v.dat)-1 {
		return errors.New("Location out of bounds")
	}

	v.dat[loc] = el

	return nil
}

// GetElement returns the Scalar element at the zero-based index loc.
// nil is returned instead of loc is out of bounds
func (v Vector) GetElement(loc int) Scalar {
	if loc < 0 || loc > len(v.dat)-1 {
		return nil
	}

	return v.dat[loc]
}

// ToScalar is equivalent to GetElement(0), but works if and only if
// the vector is a Size 1 (or 1x1 or 1-d) Vector. It's effective an acknowledgment that
// a 1x1 vector or matrix is often treated as if it were a scalar. This function is implicitly called
// in most multiplication functions, so there is no need to convert to a scalar yourself if multiplying a 1x1 Vector and a Matrix
// or other Vector
func (v Vector) ToScalar() Scalar {
	if len(v.dat) != 1 {
		return nil
	}

	return v.dat[0]
}

// AsSlice simply returns the current underlying slice representation
func (v Vector) AsSlice() []Scalar {
	return v.dat
}

// AsArray converts a vector of up to size 4 into the appropriately typed array
// Because in Go Arrays have a static size that may never change, and this may return an array
// of any size(1-4) and any valid underlying vector type an interface{} in returned and the user is responsible
// for correctly casting it. Note that if the underlying type is, say, INT32, it will be an array of int32, not ScalarInt32.
// In other words, it will return the true underlying Go type, not its Scalar wrapper.
// Very useful for use with OpenGL packages that take in such arrays
func (v Vector) AsArray() interface{} {

	switch len(v.dat) {
	case 1:
		switch v.typ {
		case INT32:
			return [1]int32{int32(v.dat[0].(ScalarInt32))}
		case UINT32:
			return [1]uint32{uint32(v.dat[0].(ScalarUint32))}
		case FLOAT32:
			return [1]float32{float32(v.dat[0].(ScalarFloat32))}
		case FLOAT64:
			return [1]float64{float64(v.dat[0].(ScalarFloat64))}
		}
	case 2:
		switch v.typ {
		case INT32:
			return [2]int32{int32(v.dat[0].(ScalarInt32)), int32(v.dat[1].(ScalarInt32))}
		case UINT32:
			return [2]uint32{uint32(v.dat[0].(ScalarUint32)), uint32(v.dat[1].(ScalarUint32))}
		case FLOAT32:
			return [2]float32{float32(v.dat[0].(ScalarFloat32)), float32(v.dat[1].(ScalarFloat32))}
		case FLOAT64:
			return [2]float64{float64(v.dat[0].(ScalarFloat64)), float64(v.dat[1].(ScalarFloat64))}
		}
	case 3:
		switch v.typ {
		case INT32:
			return [3]int32{int32(v.dat[0].(ScalarInt32)), int32(v.dat[1].(ScalarInt32)), int32(v.dat[2].(ScalarInt32))}
		case UINT32:
			return [3]uint32{uint32(v.dat[0].(ScalarUint32)), uint32(v.dat[1].(ScalarUint32)), uint32(v.dat[2].(ScalarUint32))}
		case FLOAT32:
			return [3]float32{float32(v.dat[0].(ScalarFloat32)), float32(v.dat[1].(ScalarFloat32)), float32(v.dat[2].(ScalarFloat32))}
		case FLOAT64:
			return [3]float64{float64(v.dat[0].(ScalarFloat64)), float64(v.dat[1].(ScalarFloat64)), float64(v.dat[2].(ScalarFloat64))}
		}
	case 4:
		switch v.typ {
		case INT32:
			return [4]int32{int32(v.dat[0].(ScalarInt32)), int32(v.dat[1].(ScalarInt32)), int32(v.dat[2].(ScalarInt32)), int32(v.dat[3].(ScalarInt32))}
		case UINT32:
			return [4]uint32{uint32(v.dat[0].(ScalarUint32)), uint32(v.dat[1].(ScalarUint32)), uint32(v.dat[2].(ScalarUint32)), uint32(v.dat[3].(ScalarUint32))}
		case FLOAT32:
			return [4]float32{float32(v.dat[0].(ScalarFloat32)), float32(v.dat[1].(ScalarFloat32)), float32(v.dat[2].(ScalarFloat32)), float32(v.dat[3].(ScalarFloat32))}
		case FLOAT64:
			return [4]float64{float64(v.dat[0].(ScalarFloat64)), float64(v.dat[1].(ScalarFloat64)), float64(v.dat[2].(ScalarFloat64)), float64(v.dat[3].(ScalarFloat64))}
		}
	}

	return nil
}

// AsMatrix converts the vector to an nx1 or 1xn vector. Which one it converts to is
// determined by the argument "row". If it's true, it's treated as a 1xn ROW vector,
// else an nx1 COLUMN vector.
func (v Vector) AsMatrix(row bool) (m Matrix) {
	if row {
		return *unsafeMatrixFromSlice(v.dat, v.typ, 1, len(v.dat))
	}

	return *unsafeMatrixFromSlice(v.dat, v.typ, len(v.dat), 1)
}

// Add, obviously, adds two vectors in the normal fashion:
//
// [a]   [d]    [a+d]
// [b] + [e] =  [b+e]
// [c]   [f]    [c+f]
//
// It returns the zero-value for a vector (a nil slice and type of 0 (no type)
// if the two vectors either aren't the same Size or don't have the same VecType
func (v1 Vector) Add(v2 Vector) (v3 Vector) {
	if v1.typ != v2.typ || len(v1.dat) != len(v2.dat) {
		return
	}

	v3.typ = v1.typ
	v3.dat = make([]Scalar, len(v1.dat))

	for i := range v1.dat {
		v3.dat[i] = v1.dat[i].Add(v2.dat[i])
	}

	return v3
}

// Sub, like add, subtracts two vectors in the normal fashion:
//
// [a]   [d]    [a-d]
// [b] - [e] =  [b-e]
// [c]   [f]    [c-f]
//
// It returns the zero-value for a vector (a nil slice and type of 0 (no type)
// if the two vectors either aren't the same Size or don't have the same VecType
func (v1 Vector) Sub(v2 Vector) (v3 Vector) {
	if v1.typ != v2.typ || len(v1.dat) != len(v2.dat) {
		return
	}

	v3.typ = v1.typ
	v3.dat = make([]Scalar, len(v1.dat))

	for i := range v1.dat {
		v3.dat[i] = v1.dat[i].Sub(v2.dat[i])
	}

	return v3
}

// Dot returns the dot product of the two vectors
//
// [a] [d]
// [b] [e] = a*d+b*e+c*f
// [c].[f]
//
// If the two vectors Sizes don't match or their underlying VecTypes aren't the same
// it returns nil
func (v1 Vector) Dot(v2 Vector) (ret Scalar) {
	if v1.typ != v2.typ || len(v1.dat) != len(v2.dat) {
		return nil
	}

	ret = vecNumZero(v1.typ)

	for i := range v1.dat {
		ret = ret.Add(v1.dat[i].Mul(v2.dat[i]))
	}

	return ret
}

// The cross product is only defined in three dimensions (and does not currently support homogeneous vectors)
//
// [a] [d]   [b*f-c*e]
// [b]x[e] = [c*d-a*f]
// [c] [f]   [a*e-b*d]
//
// It returns the zero-type for a vector if any vector's Size is not 3, or the vector's underlying types
// don't match
//
// Note for the pedantic: the binary cross product is technically defined in 7D as well,
// but it's such a rare operation (especially for a library geared towards 3D math) that it's left unimplemented
func (v1 Vector) Cross(v2 Vector) (v3 Vector) {
	if v1.typ != v2.typ || len(v1.dat) != len(v2.dat) || len(v1.dat) != 3 {
		return
	}

	v3 = Vector{}
	v3.typ = v1.typ
	v3.dat = make([]Scalar, len(v1.dat))

	v3.dat[0] = v1.dat[1].Mul(v2.dat[2]).Sub(v1.dat[2].Mul(v2.dat[1]))
	v3.dat[1] = v1.dat[2].Mul(v2.dat[0]).Sub(v1.dat[0].Mul(v2.dat[2]))
	v3.dat[2] = v1.dat[0].Mul(v2.dat[1]).Sub(v1.dat[1].Mul(v2.dat[0]))

	return v3
}

// ScalarMul performs element-wise scalar multiplication on a vector
//
//   [x]   [c*x]
// c [y] = [c*y]
//   [z]   [c*z]
//
// If c's type doesn't match the vector's, it will return the Zero-type of a vector
func (v1 Vector) ScalarMul(c Scalar) (v2 Vector) {
	if !checkType(v1.typ, c) {
		return
	}

	v2.typ = v1.typ
	v2.dat = make([]Scalar, len(v1.dat))

	for i := range v1.dat {
		v2.dat[i] = v1.dat[i].Mul(c)
	}

	return v2
}

// Len returns the Vector Length. Also known as *magnitude*
// This is equivalent to sqrt(v.v) -- the square root of the dot product of v with itself
// If you want the dimension or number of elements of a vector, use Size()
//
// This returns a float64 because an integer length isn't very useful
func (v Vector) Len() float64 {

	dot := v.Dot(v)

	return dot.sqrt()
}

// Size simply returns the number of elements the vector has or it's "dimension"
func (v Vector) Size() int {
	return len(v.dat)
}

// If possible, Normalize will return a normalized version of the current vector -- aka a unit vector or a vector of Length 1
//
// ||v|| = 1/Len(v) * v = 1/sqrt(v.v) * v, or a vector multiplied with one divided by its magnitude
//
// If this is not possible (i.e. v is the zero vector), it simply returns v. It also returns v if normalization isn't necessary (Len(v) already is 1)
// This method works correctly on vectors of an integer type.
func (v Vector) Normalize() (v2 Vector) {
	length := v.Len()
	if FloatEqual(length, 0.) || FloatEqual(length, 1) { // compare to 0
		return v
	}
	return v.floatScale(float64(1.0) / length)
}

// INTERNAL: floatScale makes sure that every element is scaled correctly,
// converting the length to an int makes no sense, as by strict multiplication you cannot lower an int's value (beyond making it negative)
func (v Vector) floatScale(c float64) (v2 Vector) {
	v2.typ = v.typ
	v2.dat = make([]Scalar, len(v.dat))

	for i := range v.dat {
		v2.dat[i] = v.dat[i].mulFl64(c)
	}

	return v2
}

// Equal does an element-wise comparison of two vectors and returns true if they're equal.
//
// That is, if each element of the same ordinal in both vectors are equal, this function returns true.
//
// If the vectors are of different lengths, or different VecTypes, this returns false automaticaly
//
// Naturally, for float32/64 vectors, this is only approximately equal. See util.go:FloatEqual
// for more info on this.
func (v1 Vector) Equal(v2 Vector) (eq bool) {
	if v1.typ != v2.typ || len(v1.dat) != len(v2.dat) {
		return false
	}

	for i := 0; i < len(v1.dat); i++ {
		eq = v1.dat[i].Equal(v2.dat[i])
		if !eq {
			break
		}
	}

	return eq
}

// Mul multiplies a Vector and either another Vector, or a Matrix
// If m is another Vector, it performs the dot product and returns it as a 1x1 Matrix
//
// If m is a Matrix, it returns a 1xo Matrix (where o is the number of columns in m) as if we had
// multiplied a 1xn matrix with an nxo Matrix.
//
// In any case, if v is a size-1 vector, it will be treated as a Scalar, and the function will return a Matrix as if it had been multiplied by v.AsScalar()
// in the case m is a vector, it will returned converted to a row vector multiplied by AsScalar.
//
//No special privileges are given if m is 1x1
//
// The result will be the zero-type for a Matrix if any of the following conditions are met:
// v and m's underlying VecTypes don't match
// m is a vector and m and v's Sizes aren't the same
// m is a matrix and its number of rows is not equal to v's Size
func (v Vector) Mul(m MatrixMultiplyable) (out Matrix) {
	if v2, ok := m.(Vector); ok {
		if v.typ != v2.typ || len(v.dat) != len(v2.dat) {
			return // We type check in Dot as well, but that will return a nil, I want to ensure we return a zero-val matrix
		}
		if len(v.dat) == 1 {
			return v2.AsMatrix(true).ScalarMul(v.ToScalar())
		}
		return *unsafeMatrixFromSlice([]Scalar{v.Dot(v2)}, v.typ, 1, 1)
	}

	mat := m.(Matrix)
	if v.typ != mat.typ || len(v.dat) != mat.m {
		return
	}

	if len(v.dat) == 1 {
		return mat.ScalarMul(v.ToScalar())
	}

	dat := make([]Scalar, 1*mat.n) // If v is a matrix then 1 is its "m"
	for j := 0; j < mat.n; j++ {   // Columns of m2 and m3
		//for i := 0; i < m1.m; i++ { // Rows of m1 and m3
		for k := 0; k < len(v.dat); k++ { // Columns of m1, rows of m2
			if dat[j] == nil {
				dat[j] = MakeScalar(0, v.typ)
			}
			dat[j] = dat[j].Add(v.dat[k].Mul(mat.dat[k*mat.n+j])) // I think, needs testing
		}
		//}
	}

	return *unsafeMatrixFromSlice(dat, v.typ, 1, mat.n)
}

// The outer product of a vector is similar to the dot (or "inner") product, in that
// the dot product is equal to a row vector times a column vector, the outer product
// is a column vector times a row vector. This results in an nxo matrix, where
// n is the Size of the first vector and o is the Size of the second vector
//
// [a]            [[a*c, a*d]]
// [b] * [c, d] = [[b*c, b*d]]
//
// It returns the zero-type for a Matrix if v1 and v2's underlying VecTypes don't match
func (v1 Vector) OuterProduct(v2 Vector) (m Matrix) {
	if v1.typ != v2.typ {
		return
	}

	// Should probably just spell it out
	m1 := v1.AsMatrix(false)
	m2 := v2.AsMatrix(true)

	return m1.Mul(m2)
}
