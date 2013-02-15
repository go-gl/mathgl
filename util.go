package mathgl

import (
	"math"
)

type ScalarInt32 int32
type ScalarUint32 uint32
type ScalarFloat32 float32
type ScalarFloat64 float64

// Before I did this, there were a lot of ugly switch statements and tons of code duplication.
// Go doesn't allow arbitrary math without casting, so some code duplication was inevitable, but with this wrapper
// I was able to localize all the duplication to one place and clean up the main package a great deal.
// This, unfortunately, makes it a pain to deal with anything that's not going to be used internally in the package
// i.e. doing math with the dot product of two vectors. Suggestions for how to improve this is welcome.
//
// For now, the basic mathematical operations are exported to make your lives a little bit easier
type Scalar interface {
	Add(other Scalar) Scalar
	Sub(other Scalar) Scalar
	Mul(other Scalar) Scalar
	Div(other Scalar) Scalar
	Equal(other Scalar) bool // Only "approximately equals" for float types, because of the minutae of floating point arithmetic
	Type() VecType

	// These remain unexported because they're basically shortcuts for internal benefit
	mulFl64(c float64) Scalar // This is for the rare case we need to Multiply by non-like types as for length
	sqrt() float64
}

// Begin Int
func (i ScalarInt32) Add(other Scalar) Scalar {
	return i + other.(ScalarInt32)
}

func (i ScalarInt32) Sub(other Scalar) Scalar {
	return i - other.(ScalarInt32)
}

func (i ScalarInt32) Mul(other Scalar) Scalar {
	return i * other.(ScalarInt32)
}

func (i ScalarInt32) Div(other Scalar) Scalar {
	return i / other.(ScalarInt32)
}

func (i ScalarInt32) sqrt() float64 {
	return math.Sqrt(float64(i))
}

func (i ScalarInt32) Equal(other Scalar) bool {
	return i == other.(ScalarInt32)
}

func (i ScalarInt32) mulFl64(c float64) Scalar {
	return ScalarInt32(int32(float64(i) * c))
}

func (i ScalarInt32) Type() VecType {
	return INT32
}

// Begin Uint
func (i ScalarUint32) Add(other Scalar) Scalar {
	return i + other.(ScalarUint32)
}

func (i ScalarUint32) Sub(other Scalar) Scalar {
	return i - other.(ScalarUint32)
}

func (i ScalarUint32) Mul(other Scalar) Scalar {
	return i * other.(ScalarUint32)
}

func (i ScalarUint32) Div(other Scalar) Scalar {
	return i / other.(ScalarUint32)
}

func (i ScalarUint32) sqrt() float64 {
	return math.Sqrt(float64(i))
}

func (i ScalarUint32) Equal(other Scalar) bool {
	return i == other.(ScalarUint32)
}

func (i ScalarUint32) mulFl64(c float64) Scalar {
	return ScalarUint32(uint32(float64(i) * c))
}

func (i ScalarUint32) Type() VecType {
	return UINT32
}

// Begin Float
func (i ScalarFloat32) Add(other Scalar) Scalar {
	return i + other.(ScalarFloat32)
}

func (i ScalarFloat32) Sub(other Scalar) Scalar {
	return i - other.(ScalarFloat32)
}

func (i ScalarFloat32) Mul(other Scalar) Scalar {
	return i * other.(ScalarFloat32)
}

func (i ScalarFloat32) Div(other Scalar) Scalar {
	return i / other.(ScalarFloat32)
}

func (i ScalarFloat32) sqrt() float64 {
	return math.Sqrt(float64(i))
}

func (i ScalarFloat32) Equal(other Scalar) bool {
	return math.Abs(float64(i-other.(ScalarFloat32))) < float64(.00000000001)
}

func (i ScalarFloat32) mulFl64(c float64) Scalar {
	return ScalarFloat32(float32(float64(i) * c))
}

func (i ScalarFloat32) Type() VecType {
	return FLOAT32
}

// Begin Float64
func (i ScalarFloat64) Add(other Scalar) Scalar {
	return i + other.(ScalarFloat64)
}

func (i ScalarFloat64) Sub(other Scalar) Scalar {
	return i - other.(ScalarFloat64)
}

func (i ScalarFloat64) Mul(other Scalar) Scalar {
	return i * other.(ScalarFloat64)
}

func (i ScalarFloat64) Div(other Scalar) Scalar {
	return i / other.(ScalarFloat64)
}

func (i ScalarFloat64) sqrt() float64 {
	return math.Sqrt(float64(i))
}

func (i ScalarFloat64) Equal(other Scalar) bool {
	return math.Abs(float64(i-other.(ScalarFloat64))) < float64(.00000000001)
}

func (i ScalarFloat64) mulFl64(c float64) Scalar {
	return ScalarFloat64(float64(i) * c)
}

func (i ScalarFloat64) Type() VecType {
	return FLOAT64
}

// Helper
func vecNumZero(typ VecType) Scalar {
	switch typ {
	case INT32:
		return ScalarInt32(0)
	case UINT32:
		return ScalarUint32(0)
	case FLOAT32:
		return ScalarFloat32(0)
	case FLOAT64:
		return ScalarFloat64(0)
	}

	return ScalarInt32(0)
}

// Converts an int/int32, uint32, float32, or float64 to the a Scalar of type given by the second argument.
// If the number is not one of these types, it returns nil
func MakeScalar(num interface{}, typ VecType) Scalar {

	if n, ok := num.(int); ok {
		switch typ {
		case INT32:
			return ScalarInt32(n)
		case UINT32:
			return ScalarUint32(n)
		case FLOAT32:
			return ScalarFloat32(n)
		case FLOAT64:
			return ScalarFloat64(n)
		}
	} else if n, ok := num.(int32); ok {
		switch typ {
		case INT32:
			return ScalarInt32(n)
		case UINT32:
			return ScalarUint32(n)
		case FLOAT32:
			return ScalarFloat32(n)
		case FLOAT64:
			return ScalarFloat64(n)
		}
	} else if n, ok := num.(uint32); ok {
		switch typ {
		case INT32:
			return ScalarInt32(n)
		case UINT32:
			return ScalarUint32(n)
		case FLOAT32:
			return ScalarFloat32(n)
		case FLOAT64:
			return ScalarFloat64(n)
		}
	} else if n, ok := num.(float32); ok {
		switch typ {
		case INT32:
			return ScalarInt32(n)
		case UINT32:
			return ScalarUint32(n)
		case FLOAT32:
			return ScalarFloat32(n)
		case FLOAT64:
			return ScalarFloat64(n)
		}
	} else if n, ok := num.(float64); ok {
		switch typ {
		case INT32:
			return ScalarInt32(n)
		case UINT32:
			return ScalarUint32(n)
		case FLOAT32:
			return ScalarFloat32(n)
		case FLOAT64:
			return ScalarFloat64(n)
		}
	}

	return nil
}

// Converts all elements of a slice to a Scalar of a given VecType
// All elements of the slice need not be of the same time, but MUST be of
// a Scalar-friendly type (int/int32, uint32, float32, float64) or the function will return nil
//
//  All pieces of the scalar will be converted to the Scalar type specificed in the second argument
func ScalarSlice(slice []interface{}, typ VecType) (out []Scalar) {
	out = make([]Scalar, len(slice))
	for i := range slice {
		out[i] = MakeScalar(slice[i], typ)
		if out[i] == nil {
			return nil
		}
	}

	return out
}
