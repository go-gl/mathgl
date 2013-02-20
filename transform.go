package mathgl

import "math"

func Rotate2D(angle float64) Matrix {
	sin, cos := math.Sin(angle), math.Cos(angle)
	return *unsafeMatrixFromSlice(ScalarSlice([]interface{}{cos, -sin,
		sin, cos}, FLOAT64), FLOAT64, 2, 2)
}

// Rotates about X-axis
// [1 0 0]
// [0 c -s]
// [0 s c ]
func Rotate3DX(angle float64) Matrix {
	sin, cos := math.Sin(angle), math.Cos(angle)
	return *unsafeMatrixFromSlice(ScalarSlice([]interface{}{
		1, 0, 0,
		0, cos, -sin,
		0, sin, cos}, FLOAT64), FLOAT64, 3, 3)
}

// Rotates around Y-axis
// [c 0 s]
// [0 1 0]
// [s 0 c ]
func Rotate3DY(angle float64) Matrix {
	sin, cos := math.Sin(angle), math.Cos(angle)
	return *unsafeMatrixFromSlice(ScalarSlice([]interface{}{
		cos, 0, sin,
		0, 1, 0,
		-sin, 0, cos}, FLOAT64), FLOAT64, 3, 3)
}

// Rotates about Z-axis
// [c -s 0]
// [s c 0]
// [0 0 1 ]
func Rotate3DZ(angle float64) Matrix {
	sin, cos := math.Sin(angle), math.Cos(angle)
	return *unsafeMatrixFromSlice(ScalarSlice([]interface{}{
		cos, -sin, 0,
		-sin, cos, 0,
		0, 0, 1}, FLOAT64), FLOAT64, 3, 3)
}

// Returns a 2D homogeneous (3x3) transformation matrix
func Transform2D(Tx, Ty float64) Matrix {
	return *unsafeMatrixFromSlice(ScalarSlice([]interface{}{
		1, 0, Tx,
		0, 1, Ty,
		0, 0, 1}, FLOAT64), FLOAT64, 3, 3)
}

// Returns a 3D homogeneous (4x4) transformation matrix
func Transform3D(Tx, Ty, Tz float64) Matrix {
	return *unsafeMatrixFromSlice(ScalarSlice([]interface{}{
		1, 0, 0, Tx,
		0, 1, 0, Ty,
		0, 0, 1, Tz,
		0, 0, 0, 1}, FLOAT64), FLOAT64, 4, 4)
}

func HomogRotate2D(angle float64) Matrix {
	sin, cos := math.Sin(angle), math.Cos(angle)
	return *unsafeMatrixFromSlice(ScalarSlice([]interface{}{
		cos, -sin, 0,
		sin, cos, 0,
		0, 0, 1}, FLOAT64), FLOAT64, 3, 3)
}

// Rotates about X-axis
// [1 0 0]
// [0 c -s]
// [0 s c ]
func HomogRotate3DX(angle float64) Matrix {
	sin, cos := math.Sin(angle), math.Cos(angle)
	return *unsafeMatrixFromSlice(ScalarSlice([]interface{}{
		1, 0, 0, 0,
		0, cos, -sin, 0,
		0, sin, cos, 0,
		0, 0, 0, 1}, FLOAT64), FLOAT64, 4, 4)
}

// Rotates around Y-axis
// [c 0 s]
// [0 1 0]
// [s 0 c ]
func HomogRotate3DY(angle float64) Matrix {
	sin, cos := math.Sin(angle), math.Cos(angle)
	return *unsafeMatrixFromSlice(ScalarSlice([]interface{}{
		cos, 0, sin, 0,
		0, 1, 0, 0,
		-sin, 0, cos, 0,
		0, 0, 0, 1}, FLOAT64), FLOAT64, 4, 4)
}

// Rotates about Z-axis
// [c -s 0]
// [s c 0]
// [0 0 1 ]
func HomogRotate3DZ(angle float64) Matrix {
	sin, cos := math.Sin(angle), math.Cos(angle)
	return *unsafeMatrixFromSlice(ScalarSlice([]interface{}{
		cos, -sin, 0, 0, 
		-sin, cos, 0, 0, 
		0, 0, 1, 0, 
		0, 0, 0, 1}, FLOAT64), FLOAT64, 4, 4)
}

// Only accepts a homogeneous (size 4) 3D vector, if it's not returns the zero-type for Matrix. The Vector must also be of type Float64
// It will always return a Matrix of type Float64
func HomogRotate3D(angle float64, axis Vector) Matrix {
	if axis.Size() != 4 || axis.typ != FLOAT64 {
		return Matrix{}
	}
	x, y, z := axis.dat[0].Fl64(), axis.dat[1].Fl64(), axis.dat[2].Fl64()
	s, c := math.Sin(angle), math.Cos(angle)
	k := 1 - c
	return *unsafeMatrixFromSlice(ScalarSlice([]interface{}{
		x*x*k + c, x*y*k - z*s, x*z*k + y*s, 0,
		x*y*k + z*s, y*y*k + c, y*z*k - x*s, 0,
		x*z*k - y*s, y*z*k + x*s, z*z*k + c, 0,
		0, 0, 0, 1}, FLOAT64), FLOAT64, 4, 4)
}

// Note: Vector type must be FLOAT64
func QuaternionRotation(angle float64, axis Vector) Quaternion {
	if axis.typ != FLOAT64 {
		return Quaternion{}
	}

	sin, cos := math.Sin(angle), math.Cos(angle)
	return Quaternion{MakeScalar(cos/float64(2), FLOAT64), axis.ScalarMul(MakeScalar(sin/float64(2), FLOAT64)), FLOAT64}
}
