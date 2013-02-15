package mathgl

import (
	"math"
)

type VecInt32 int32
type VecUint32 uint32
type VecFloat32 float32
type VecFloat64 float64

// Internal Note: Types still need to be checked (using the VecType enum, not reflection) before calling any of the VecNum functions or Go will panic due to an incorrect type assertion.

// Before I did this, there were a lot of ugly switch statements and tons of code duplication.
// Go doesn't allow arbitrary math without casting, so some code duplication was inevitable, but with this wrapper
// I was able to localize all the duplication to one place and clean up the main package a great deal.
type VecNum interface {
	add(other VecNum) VecNum
	sub(other VecNum) VecNum
	mul(other VecNum) VecNum
	mulFl64(c float64) VecNum // This is for the rare case we need to multiply by non-like types as for length
	div(other VecNum) VecNum
	equal(other VecNum) bool
	sqrt() float64
}

// Begin Int
func (i VecInt32) add(other VecNum) VecNum {
	return i + other.(VecInt32)
}

func (i VecInt32) sub(other VecNum) VecNum {
	return i - other.(VecInt32)
}

func (i VecInt32) mul(other VecNum) VecNum {
	return i * other.(VecInt32)
}

func (i VecInt32) div(other VecNum) VecNum {
	return i / other.(VecInt32)
}

func (i VecInt32) sqrt() float64 {
	return math.Sqrt(float64(i))
}

func (i VecInt32) equal(other VecNum) bool {
	return i == other.(VecInt32)
}

func (i VecInt32) mulFl64(c float64) VecNum {
	return VecInt32(int32(float64(i) * c))
}

// Begin Uint
func (i VecUint32) add(other VecNum) VecNum {
	return i + other.(VecUint32)
}

func (i VecUint32) sub(other VecNum) VecNum {
	return i - other.(VecUint32)
}

func (i VecUint32) mul(other VecNum) VecNum {
	return i * other.(VecUint32)
}

func (i VecUint32) div(other VecNum) VecNum {
	return i / other.(VecUint32)
}

func (i VecUint32) sqrt() float64 {
	return math.Sqrt(float64(i))
}

func (i VecUint32) equal(other VecNum) bool {
	return i == other.(VecUint32)
}

func (i VecUint32) mulFl64(c float64) VecNum {
	return VecUint32(uint32(float64(i) * c))
}

// Begin Float
func (i VecFloat32) add(other VecNum) VecNum {
	return i + other.(VecFloat32)
}

func (i VecFloat32) sub(other VecNum) VecNum {
	return i - other.(VecFloat32)
}

func (i VecFloat32) mul(other VecNum) VecNum {
	return i * other.(VecFloat32)
}

func (i VecFloat32) div(other VecNum) VecNum {
	return i / other.(VecFloat32)
}

func (i VecFloat32) sqrt() float64 {
	return math.Sqrt(float64(i))
}

func (i VecFloat32) equal(other VecNum) bool {
	return math.Abs(float64(i-other.(VecFloat32))) < float64(.00000000001)
}

func (i VecFloat32) mulFl64(c float64) VecNum {
	return VecFloat32(float32(float64(i) * c))
}

// Begin Float64
func (i VecFloat64) add(other VecNum) VecNum {
	return i + other.(VecFloat64)
}

func (i VecFloat64) sub(other VecNum) VecNum {
	return i - other.(VecFloat64)
}

func (i VecFloat64) mul(other VecNum) VecNum {
	return i * other.(VecFloat64)
}

func (i VecFloat64) div(other VecNum) VecNum {
	return i / other.(VecFloat64)
}

func (i VecFloat64) sqrt() float64 {
	return math.Sqrt(float64(i))
}

func (i VecFloat64) equal(other VecNum) bool {
	return math.Abs(float64(i-other.(VecFloat64))) < float64(.00000000001)
}

func (i VecFloat64) mulFl64(c float64) VecNum {
	return VecFloat64(float64(i) * c)
}

// Helper
func vecNumZero(typ VecType) VecNum {
	switch typ {
	case INT32:
		return VecInt32(0)
	case UINT32:
		return VecUint32(0)
	case FLOAT32:
		return VecFloat32(0)
	case FLOAT64:
		return VecFloat64(0)
	}

	return VecInt32(0)
}
