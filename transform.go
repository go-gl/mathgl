package mathgl

import "math"

func Rotate2D(angle float64) Matrix {
	sin,cos := math.Sin(angle), math.Cos(angle)
	return *unsafeMatrixFromSlice(ScalarSlice([]interface{}{cos, -sin, sin, cos}, FLOAT64), FLOAT64, 2, 2)
}

// Rotates about X-axis
// [1 0 0]
// [0 c -s]
// [0 s c ]
func Rotate3DX(angle float64) Matrix {
	sin,cos := math.Sin(angle), math.Cos(angle)
	return *unsafeMatrixFromSlice(ScalarSlice([]interface{}{1,0,0,0,cos, -sin, 0, sin, cos}, FLOAT64), FLOAT64, 3, 3)
}

// Rotates around Y-axis
// [c 0 s]
// [0 1 0]
// [s 0 c ]
func Rotate3DY(angle float64) Matrix {
	sin,cos := math.Sin(angle), math.Cos(angle)
	return *unsafeMatrixFromSlice(ScalarSlice([]interface{}{cos,0,sin,0,1, 0, -sin, 0, cos}, FLOAT64), FLOAT64, 3, 3)
}

// Rotates about Z-axis
// [c -s 0]
// [s c 0]
// [0 0 1 ]
func Rotate3DZ(angle float64) Matrix {
	sin,cos := math.Sin(angle), math.Cos(angle)
	return *unsafeMatrixFromSlice(ScalarSlice([]interface{}{cos,-sin,0,-sin,cos,0, 0, 0, 1}, FLOAT64), FLOAT64, 3, 3)
}

// Returns a 2D homogeneous (3x3) transformation matrix
func Transform2D(Tx, Ty float64) Matrix {
	return *unsafeMatrixFromSlice(ScalarSlice([]interface{}{1,0,Tx,0,1,Ty,0,0, 1}, FLOAT64), FLOAT64, 3, 3)
}

// Returns a 3D homogeneous (4x4) transformation matrix
func Transform3D(Tx, Ty, Tz float64) Matrix {
	return *unsafeMatrixFromSlice(ScalarSlice([]interface{}{1,0,0,Tx,0,1,0,Ty,0,0, 1,Tz,0,0,0,1}, FLOAT64), FLOAT64, 3, 3)
}

// TODO: Quaternion rotation