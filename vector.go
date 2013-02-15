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
	dat []VecNum
}

func NewVector(t VecType) *Vector {
	return &Vector{typ: t, dat: make([]VecNum, 0, 2)}
}

func VectorOf(t VecType, el []VecNum) (v *Vector, err error) {
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
		_,ok := i.(VecInt32)
		return ok
	case UINT32:
		_,ok := i.(VecUint32)
		return ok
	case FLOAT32:
		_,ok := i.(VecFloat32)
		return ok
	case FLOAT64:
		_,ok := i.(VecFloat64)
		return ok
	}

	return false
}

func (v *Vector) AddElements(el []VecNum) error {
	for _, e := range el {
		if !checkType(v.typ, e) {
			return errors.New("Type of at least one element does not match vector's type")
		}
	}

	v.dat = append(v.dat, el...)
	return nil
}

func (v *Vector) SetElement(loc int, el VecNum) error {
	if !checkType(v.typ, el) {
		return errors.New("Element does not match vector's type")
	}

	if loc < 0 || loc > len(v.dat)-1 {
		return errors.New("Location out of bounds")
	}

	v.dat[loc] = el

	return nil
}

func (v Vector) GetElement(loc int) VecNum {
	if loc < 0 || loc > len(v.dat)-1 {
		return nil
	}

	return v.dat[loc]
}

// Converts a 1-d vector to a scalar
func (v Vector) ToScalar() VecNum {
	if len(v.dat) != 1 {
		return nil
	}

	/*switch v.typ {
	case INT32:
		return v.dat[0].(int32)
	case UINT32:
		return v.dat[0].(uint32)
	case FLOAT32:
		return v.dat[0].(float32)
	case FLOAT64:
		return v.dat[0].(float64)
	}*/

	return v.dat[0]
}

func (v Vector) AsSlice() []VecNum {
	return v.dat
}

// Converts a vector of up to size 4 into the appropriately typed array
// Still must return an interface{} because of array size weirdness
func (v Vector) AsArray() interface{} {

	switch len(v.dat) {
	case 1:
		switch v.typ {
		case INT32:
			return [1]int32{int32(v.dat[0].(VecInt32))}
		case UINT32:
			return [1]uint32{uint32(v.dat[0].(VecUint32))}
		case FLOAT32:
			return [1]float32{float32(v.dat[0].(VecFloat32))}
		case FLOAT64:
			return [1]float64{float64(v.dat[0].(VecFloat64))}
		}
	case 2:
		switch v.typ {
		case INT32:
			return [2]int32{int32(v.dat[0].(VecInt32)), int32(v.dat[1].(VecInt32))}
		case UINT32:
			return [2]uint32{uint32(v.dat[0].(VecUint32)), uint32(v.dat[1].(VecUint32))}
		case FLOAT32:
			return [2]float32{float32(v.dat[0].(VecFloat32)), float32(v.dat[1].(VecFloat32))}
		case FLOAT64:
			return [2]float64{float64(v.dat[0].(VecFloat64)), float64(v.dat[1].(VecFloat64))}
		}
	case 3:
		switch v.typ {
		case INT32:
			return [3]int32{int32(v.dat[0].(VecInt32)), int32(v.dat[1].(VecInt32)), int32(v.dat[2].(VecInt32))}
		case UINT32:
			return [3]uint32{uint32(v.dat[0].(VecUint32)), uint32(v.dat[1].(VecUint32)), uint32(v.dat[2].(VecUint32))}
		case FLOAT32:
			return [3]float32{float32(v.dat[0].(VecFloat32)), float32(v.dat[1].(VecFloat32)), float32(v.dat[2].(VecFloat32))}
		case FLOAT64:
			return [3]float64{float64(v.dat[0].(VecFloat64)), float64(v.dat[1].(VecFloat64)), float64(v.dat[2].(VecFloat64))}
		}
	case 4:
		switch v.typ {
		case INT32:
			return [4]int32{int32(v.dat[0].(VecInt32)), int32(v.dat[1].(VecInt32)), int32(v.dat[2].(VecInt32)), int32(v.dat[3].(VecInt32))}
		case UINT32:
			return [4]uint32{uint32(v.dat[0].(VecUint32)), uint32(v.dat[1].(VecUint32)), uint32(v.dat[2].(VecUint32)), uint32(v.dat[3].(VecUint32))}
		case FLOAT32:
			return [4]float32{float32(v.dat[0].(VecFloat32)), float32(v.dat[1].(VecFloat32)), float32(v.dat[2].(VecFloat32)), float32(v.dat[3].(VecFloat32))}
		case FLOAT64:
			return [4]float64{float64(v.dat[0].(VecFloat64)), float64(v.dat[1].(VecFloat64)), float64(v.dat[2].(VecFloat64)), float64(v.dat[3].(VecFloat64))}
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
	v3.dat = make([]VecNum, len(v1.dat))

	for i := range v1.dat {
		v3.dat[i] = v1.dat[i].add(v2.dat[i])
		/*switch v1.typ {
		case INT32:
			v3.dat[i] = v1.dat[i].(int32) + v2.dat[i].(int32)
		case UINT32:
			v3.dat[i] = v1.dat[i].(uint32) + v2.dat[i].(uint32)
		case FLOAT32:
			v3.dat[i] = v1.dat[i].(float32) + v2.dat[i].(float32)
		case FLOAT64:
			v3.dat[i] = v1.dat[i].(float64) + v2.dat[i].(float64)
		}*/
	}

	return v3
}

func (v1 Vector) Sub(v2 Vector) (v3 Vector) {
	if v1.typ != v2.typ || len(v1.dat) != len(v2.dat) {
		return
	}

	v3.typ = v1.typ
	v3.dat = make([]VecNum, len(v1.dat))

	for i := range v1.dat {
		v3.dat[i] = v1.dat[i].sub(v2.dat[i])
		/*switch v1.typ {
		case INT32:
			v3.dat[i] = v1.dat[i].(int32) - v2.dat[i].(int32)
		case UINT32:
			v3.dat[i] = v1.dat[i].(uint32) - v2.dat[i].(uint32)
		case FLOAT32:
			v3.dat[i] = v1.dat[i].(float32) - v2.dat[i].(float32)
		case FLOAT64:
			v3.dat[i] = v1.dat[i].(float64) - v2.dat[i].(float64)
		}*/
	}

	return v3
}

func (v1 Vector) Dot(v2 Vector) VecNum {
	if v1.typ != v2.typ || len(v1.dat) != len(v2.dat) {
		return nil
	}

	ret := vecNumZero(v1.typ)

	for i := range v1.dat {
		ret = ret.add(v1.dat[i].mul(v2.dat[i]))
	}

	/*switch v1.typ {
	case INT32:
		ret := int32(0)
		for i := range v1.dat {
			ret = ret + v1.dat[i].(int32) * v2.dat[i].(int32)
		}
		return ret
	case UINT32:
		ret := uint32(0)
		for i := range v1.dat {
			ret = ret + v1.dat[i].(uint32) * v2.dat[i].(uint32)
		}
		return ret
	case FLOAT32:
		ret := float32(0)
		for i := range v1.dat {
			ret = ret + v1.dat[i].(float32) * v2.dat[i].(float32)
		}
		return ret
	case FLOAT64:
		ret := float64(0)
		for i := range v1.dat {
			ret = ret + v1.dat[i].(float64) * v2.dat[i].(float64)
		}
		return ret
	}*/

	return nil
}

// Should we allow 7-dimensional?
func (v1 Vector) Cross(v2 Vector) (v3 Vector) {
	if v1.typ != v2.typ || len(v1.dat) != len(v2.dat) || len(v1.dat) != 3 {
		return
	}

	v3.typ = v1.typ
	v3.dat = make([]VecNum, len(v3.dat))

	v3.dat[0] = v1.dat[1].mul(v2.dat[2]).sub(v1.dat[2].mul(v2.dat[1]))
	v3.dat[1] = v1.dat[2].mul(v2.dat[0]).sub(v1.dat[0].mul(v2.dat[2]))
	v3.dat[2] = v1.dat[0].mul(v2.dat[1]).sub(v1.dat[1].mul(v2.dat[0]))

	/*switch v1.typ {
	case INT32:
		v3.dat[0] = v1.dat[1].(int32) * v2.dat[2].(int32) - v1.dat[2].(int32) * v2.dat[1].(int32)
		v3.dat[1] = v1.dat[2].(int32) * v2.dat[0].(int32) - v1.dat[0].(int32) * v2.dat[2].(int32)
		v3.dat[2] = v1.dat[0].(int32) * v2.dat[1].(int32) - v1.dat[1].(int32) * v2.dat[0].(int32)
	case UINT32:
		v3.dat[0] = v1.dat[1].(uint32) * v2.dat[2].(uint32) - v1.dat[2].(uint32) * v2.dat[1].(uint32)
		v3.dat[1] = v1.dat[2].(uint32) * v2.dat[0].(uint32) - v1.dat[0].(uint32) * v2.dat[2].(uint32)
		v3.dat[2] = v1.dat[0].(uint32) * v2.dat[1].(uint32) - v1.dat[1].(uint32) * v2.dat[0].(uint32)
	case FLOAT32:
		v3.dat[0] = v1.dat[1].(float32) * v2.dat[2].(float32) - v1.dat[2].(float32) * v2.dat[1].(float32)
		v3.dat[1] = v1.dat[2].(float32) * v2.dat[0].(float32) - v1.dat[0].(float32) * v2.dat[2].(float32)
		v3.dat[2] = v1.dat[0].(float32) * v2.dat[1].(float32) - v1.dat[1].(float32) * v2.dat[0].(float32)
	case FLOAT64:
		v3.dat[0] = v1.dat[1].(float64) * v2.dat[2].(float64) - v1.dat[2].(float64) * v2.dat[1].(float64)
		v3.dat[1] = v1.dat[2].(float64) * v2.dat[0].(float64) - v1.dat[0].(float64) * v2.dat[2].(float64)
		v3.dat[2] = v1.dat[0].(float64) * v2.dat[1].(float64) - v1.dat[1].(float64) * v2.dat[0].(float64)
	}*/

	return v3
}

func (v1 Vector) ScalarMul(c VecNum) (v2 Vector) {
	if !checkType(v1.typ, c) {
		return
	}

	v2.typ = v1.typ
	v2.dat = make([]VecNum, len(v1.dat))

	for i := range v1.dat {
		v2.dat[i] = v1.dat[i].mul(c)
		/*switch v1.typ {
		case INT32:
			v2.dat[i] = v1.dat[i].(int32) * c.(int32)
		case UINT32:
			v2.dat[i] = v1.dat[i].(uint32) * c.(uint32)
		case FLOAT32:
			v2.dat[i] = v1.dat[i].(float32) * c.(float32)
		case FLOAT64:
			v2.dat[i] = v1.dat[i].(float64) * c.(float64)
		}*/
	}

	return v2
}

func (v Vector) Len() float64 {

	dot := v.Dot(v)
	/*switch v.typ {
	case INT32:
		return math.Sqrt( float64( dot.(int32)))
	case UINT32:
		return math.Sqrt( float64( dot.(uint32)))
	case FLOAT32:
		return math.Sqrt( float64( dot.(float32)))
	case FLOAT64:
		return math.Sqrt( float64( dot.(float64)))
	}*/

	return dot.sqrt()
}

func (v Vector) Normalize() (v2 Vector) {
	return v.ScalarMul(VecFloat64(float64(1.0) / v.Len()))
}
