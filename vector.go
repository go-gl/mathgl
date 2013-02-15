package mathgl

import (
	"errors"
)

type VecType int8

const (
	INT32 = iota
	UINT32
	FLOAT32
	FLOAT64
)

type Vector struct {
	typ VecType
	dat []Scalar
}

func NewVector(t VecType) *Vector {
	return &Vector{typ: t, dat: make([]Scalar, 0, 2)}
}

func VectorOf(t VecType, el []Scalar) (v *Vector, err error) {
	for _, e := range el {
		if !checkType(t, e) {
			return nil, errors.New("Type of at least one element does not match declared type")
		}
	}

	return &Vector{t, el}, nil
}

func checkType(typ VecType, i interface{}) bool {
	switch typ {
	case INT32:
		_, ok := i.(ScalarInt32)
		return ok
	case UINT32:
		_, ok := i.(ScalarUint32)
		return ok
	case FLOAT32:
		_, ok := i.(ScalarFloat32)
		return ok
	case FLOAT64:
		_, ok := i.(ScalarFloat64)
		return ok
	}

	return false
}

func (v *Vector) AddElements(el []Scalar) error {
	for _, e := range el {
		if !checkType(v.typ, e) {
			return errors.New("Type of at least one element does not match vector's type")
		}
	}

	v.dat = append(v.dat, el...)
	return nil
}

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

func (v Vector) GetElement(loc int) Scalar {
	if loc < 0 || loc > len(v.dat)-1 {
		return nil
	}

	return v.dat[loc]
}

// Converts a 1-d vector to a scalar
func (v Vector) ToScalar() Scalar {
	if len(v.dat) != 1 {
		return nil
	}

	return v.dat[0]
}

func (v Vector) AsSlice() []Scalar {
	return v.dat
}

// Converts a vector of up to size 4 into the appropriately typed array
// Still must return an interface{} because of array size weirdness
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

// If row is true, it's a row vector (1xn) else a column vector (nx1)
func (v Vector) AsMatrix(row bool) (m Matrix, err error) {
	if row {
		mat, err := MatrixFromSlice(v.typ, v.dat, 1, len(v.dat))
		return *mat, err
	}

	mat, err := MatrixFromSlice(v.typ, v.dat, len(v.dat), 1)
	return *mat, err
}

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

// Should we allow 7-dimensional?
func (v1 Vector) Cross(v2 Vector) (v3 Vector) {
	if v1.typ != v2.typ || len(v1.dat) != len(v2.dat) || len(v1.dat) != 3 {
		return
	}

	v3.typ = v1.typ
	v3.dat = make([]Scalar, len(v3.dat))

	v3.dat[0] = v1.dat[1].Mul(v2.dat[2]).Sub(v1.dat[2].Mul(v2.dat[1]))
	v3.dat[1] = v1.dat[2].Mul(v2.dat[0]).Sub(v1.dat[0].Mul(v2.dat[2]))
	v3.dat[2] = v1.dat[0].Mul(v2.dat[1]).Sub(v1.dat[1].Mul(v2.dat[0]))

	return v3
}

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

// This is VECTOR LENGTH, a.k.a magnitude. For the number of elements, us Size()
func (v Vector) Len() float64 {

	dot := v.Dot(v)

	return dot.sqrt()
}

// This is the number of elements. For vector length or magnitude use Len()
func (v Vector) Size() int {
	return len(v.dat)
}

func (v Vector) Normalize() (v2 Vector) {
	return v.floatScale(float64(1.0) / v.Len())
}

func (v Vector) floatScale(c float64) (v2 Vector) {
	v2.typ = v.typ
	v2.dat = make([]Scalar, len(v.dat))

	for i := range v.dat {
		v2.dat[i] = v.dat[i].mulFl64(c)
	}

	return v2
}

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

// Assumes inner product, not out product.
func (v Vector) Mul(m MatrixMultiplyable) (out Matrix) {
	if v2, ok := m.(Vector); ok {
		if v.typ != v2.typ {
			return // We type check in Dot as well, but that will return a nil, I want to ensure we return a zero-val matrix
		}
		return *unsafeMatrixFromSlice(v.typ, []Scalar{v.Dot(v2)}, 1, 1)
	}
	mat := m.(Matrix)
	if v.typ != mat.typ {
		return
	}

	dat := make([]Scalar, 1*mat.n) // If v is a matrix then 1 is its "m"
	for j := 0; j < mat.n; j++ {   // Columns of m2 and m3
		//for i := 0; i < m1.m; i++ { // Rows of m1 and m3
		for k := 0; k < len(v.dat); k++ { // Columns of m1, rows of m2
			dat[j*mat.n] = dat[j*mat.n].Add(v.dat[k*mat.n].Mul(mat.dat[j*mat.n+k])) // I think, needs testing
		}
		//}
	}

	return *unsafeMatrixFromSlice(v.typ, dat, 1, mat.n)
}

func (v1 Vector) OuterProduct(v2 Vector) (m Matrix) {
	if v1.typ != v2.typ {
		return
	}
	
	// Should probably just spell it out
	m1,_ := v1.AsMatrix(false)
	m2,_ := v2.AsMatrix(true)
	
	return m1.Mul(m2)
}