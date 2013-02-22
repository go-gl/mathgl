package mathgl

import (
	"math"
)

type ScalarInt32 int32
type ScalarUint32 uint32
type ScalarFloat32 float32
type ScalarFloat64 float64

// Scalar is a wrapper that tiptoes around generics. By defining a few basic arithmetic operations
// on some aliases for things like int32, we can treat any type Vector (Matrix/Quaternion) interchangeably as long as the operations
// are performed between the same underlying type, there's no conflict and the code becomes much cleaner. (This prevents all the arguments and return vals from being
// interface{} everywhere in the code, which leads to a lot of really ugly casting and switch statements)
//
// Note that for purposs of helper functions (MakeScalar, ScalarSlice), an int is treated like an int32
type Scalar interface {
	Add(other Scalar) Scalar  // Adds two numbers, returns a Scalar of the same type
	Sub(other Scalar) Scalar  // Subtracts, returns a Scalar of the same type
	Mul(other Scalar) Scalar  // Multiplies, returns a Scalar of the same type
	Div(other Scalar) Scalar  // Divides, returns a Scalar of the same type
	Pow(toThe float64) Scalar // Does x^(toThe) of the Scalar, and returns a Scalar of the same type
	Equal(other Scalar) bool  // Equivalent to a==other for int types, util.go:FloatEquals(a, other) for float types
	Type() VecType            // Returns the VecType corresponding to the scalar (INT32 for ScalarInt32 etc)
	Fl64() float64            // Returns the underlying value as a float64, regardless of the underlying type
	Fl32() float32            // Returns the underlying value as a float32, regardless of the underlying type
	Int32() int32             // Returns the underlying value as an int32, regardless of the underlying type
	Uint32() uint32           // Returns the underlying value as a uint32, regardless of the underlying type

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
	return ScalarInt32(i.Fl64() * c)
}

func (i ScalarInt32) Type() VecType {
	return INT32
}

func (i ScalarInt32) Pow(toThe float64) Scalar {
	return ScalarInt32(int32(math.Pow(float64(i), toThe)))
}

func (i ScalarInt32) Fl64() float64 {
	return float64(i)
}

func (i ScalarInt32) Fl32() float32 {
	return float32(i)
}

func (i ScalarInt32) Int32() int32 {
	return int32(i)
}

func (i ScalarInt32) Uint32() uint32 {
	return uint32(i)
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
	return ScalarUint32(i.Fl64() * c)
}

func (i ScalarUint32) Type() VecType {
	return UINT32
}

func (i ScalarUint32) Pow(toThe float64) Scalar {
	return ScalarUint32(uint32(math.Pow(float64(i), toThe)))
}

func (i ScalarUint32) Fl64() float64 {
	return float64(i)
}

func (i ScalarUint32) Fl32() float32 {
	return float32(i)
}

func (i ScalarUint32) Int32() int32 {
	return int32(i)
}

func (i ScalarUint32) Uint32() uint32 {
	return uint32(i)
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
	f1 := float64(i)
	f2 := float64(other.(ScalarFloat32))

	return FloatEqual(f1, f2)
}

func (i ScalarFloat32) mulFl64(c float64) Scalar {
	return ScalarFloat32(i.Fl64() * c)
}

func (i ScalarFloat32) Type() VecType {
	return FLOAT32
}

func (i ScalarFloat32) Pow(toThe float64) Scalar {
	return ScalarFloat32(float32(math.Pow(float64(i), toThe)))
}

func (i ScalarFloat32) Fl64() float64 {
	return float64(i)
}

func (i ScalarFloat32) Fl32() float32 {
	return float32(i)
}

func (i ScalarFloat32) Int32() int32 {
	return int32(i)
}

func (i ScalarFloat32) Uint32() uint32 {
	return uint32(i)
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
	f1 := float64(i)
	f2 := float64(other.(ScalarFloat64))

	return FloatEqual(f1, f2)
}

func (i ScalarFloat64) mulFl64(c float64) Scalar {
	return ScalarFloat64(float64(i) * c)
}

func (i ScalarFloat64) Type() VecType {
	return FLOAT64
}

func (i ScalarFloat64) Pow(toThe float64) Scalar {
	return ScalarFloat64(math.Pow(float64(i), toThe))
}

func (i ScalarFloat64) Fl64() float64 {
	return float64(i)
}

func (i ScalarFloat64) Fl32() float32 {
	return float32(i)
}

func (i ScalarFloat64) Int32() int32 {
	return int32(i)
}

func (i ScalarFloat64) Uint32() uint32 {
	return uint32(i)
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

// MakeScalar converts an int/int32, uint32, float32, or float64 to a Scalar of type given by the second argument.
// If the number is not one of the allowed types (or typ is not a recognized value of VecType) it returns nil
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

// ScalarSlice Converts all elements of a slice to a Scalar of a given VecType
// All elements of the slice need not be of the same type, but MUST be of
// a Scalar-friendly type (int/int32, uint32, float32, or float64) or the function will return nil
//
//  All pieces of the scalar will be converted to the Scalar type specificed in the second argument, there is no need for them to be the "correct" type when passing in.
//
// Using the function can get ugly, my recommended syntax is:
//
//    ScalarSlice([]interface{}{
// 		 num, num, num...
// 		 num, num, num...
// 		 num, num, num... }, VECTYPE)
//
// Where the line of "nums" is in the shape of the matrix (vector etc) you're building
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

func InferScalar(num interface{}) (Scalar,VecType) {
	if n,ok := num.(float64); ok {
		return ScalarFloat64(n), FLOAT64
	} else if n,ok := num.(float32); ok {
		return ScalarFloat32(n), FLOAT32
	} else if n,ok := num.(int); ok {
		return ScalarInt32(n), INT32
	} else if n,ok := num.(int32); ok {
		return ScalarInt32(n), INT32
	} else if n,ok := num.(uint32); ok {
		return ScalarUint32(n), UINT32
	}
	
	return nil,NOTYPE
}

func InferScalarSlice(slice []interface{}) (out []Scalar,typ VecType) {
	out = make([]Scalar, len(slice))
	
	initial,typ := InferScalar(slice[0])
	out[0] = initial
	for i,el := range slice[1:] {
		out[i+1] = MakeScalar(el,typ)
		if out[i+1] == nil {
			return nil, NOTYPE
		}
	}

	return out, typ
}

func ScalarType(scalar Scalar) VecType {
	if _,ok := scalar.(ScalarFloat64); ok {
		return FLOAT64
	} else if _,ok := scalar.(ScalarFloat32); ok {
		return FLOAT32
	} else if _,ok := scalar.(ScalarInt32); ok {
		return INT32
	} else if _,ok := scalar.(ScalarUint32); ok {
		return UINT32
	}
	
	return NOTYPE
}
