package mathgl

import "math"

// Rotate2D returns a rotation Matrix of type FLOAT64 about a (radian) angle in 2-D space. Specifically about the origin.
// It is a 2x2 matrix, if you need a 3x3 for Homogeneous math (e.g. composition with a Translation matrix)
// see HomogRotate2D
func Rotate2D(angle float64) Matrix {
	sin, cos := math.Sin(angle), math.Cos(angle)
	return *unsafeMatrixFromSlice(ScalarSlice([]interface{}{cos, -sin,
		sin, cos}, FLOAT64), FLOAT64, 2, 2)
}

// Rotate3DX returns a 3x3 (non-homogeneous) Matrix of type FLOAT64 that rotates by (radian) angle about the X-axis
//
// Where c is cos(angle) and s is sin(angle)
//    [1 0 0]
//    [0 c -s]
//    [0 s c ]
func Rotate3DX(angle float64) Matrix {
	sin, cos := math.Sin(angle), math.Cos(angle)
	return *unsafeMatrixFromSlice(ScalarSlice([]interface{}{
		1, 0, 0,
		0, cos, -sin,
		0, sin, cos}, FLOAT64), FLOAT64, 3, 3)
}

// Rotate3DY returns a 3x3 (non-homogeneous) Matrix of type FLOAT64 that rotates by (radian) angle about the Y-axis
//
// Where c is cos(angle) and s is sin(angle)
//    [c 0 s]
//    [0 1 0]
//    [s 0 c ]
func Rotate3DY(angle float64) Matrix {
	sin, cos := math.Sin(angle), math.Cos(angle)
	return *unsafeMatrixFromSlice(ScalarSlice([]interface{}{
		cos, 0, sin,
		0, 1, 0,
		-sin, 0, cos}, FLOAT64), FLOAT64, 3, 3)
}

// Rotate3DZ returns a 3x3 (non-homogeneous) Matrix of type FLOAT64 that rotates by (radian) angle about the Z-axis
//
// Where c is cos(angle) and s is sin(angle)
//    [c -s 0]
//    [s c 0]
//    [0 0 1 ]
func Rotate3DZ(angle float64) Matrix {
	sin, cos := math.Sin(angle), math.Cos(angle)
	return *unsafeMatrixFromSlice(ScalarSlice([]interface{}{
		cos, -sin, 0,
		-sin, cos, 0,
		0, 0, 1}, FLOAT64), FLOAT64, 3, 3)
}

// Translate2D returns a homogeneous (3x3 for 2D-space) Translation matrix of type FLOAT64 that moves a point by Tx units in the x-direction and Ty units in the y-direction
//
//    [[1, 0, Tx]]
//    [[0, 1, Ty]]
//    [[0, 0, 1 ]]
func Translate2D(Tx, Ty float64) Matrix {
	return *unsafeMatrixFromSlice(ScalarSlice([]interface{}{
		1, 0, Tx,
		0, 1, Ty,
		0, 0, 1}, FLOAT64), FLOAT64, 3, 3)
}

// Translate3D returns a homogeneous (4x4 for 3D-space) Translation matrix of type FLOAT64 that moves a point by Tx units in the x-direction, Ty units in the y-direction,
// and Tz units in the z-direction
//
//    [[1, 0, 0, Tx]]
//    [[0, 1, 0, Ty]]
//    [[0, 0, 1, Tz]]
//    [[0, 0, 0, 1 ]]
func Translate3D(Tx, Ty, Tz float64) Matrix {
	return *unsafeMatrixFromSlice(ScalarSlice([]interface{}{
		1, 0, 0, Tx,
		0, 1, 0, Ty,
		0, 0, 1, Tz,
		0, 0, 0, 1}, FLOAT64), FLOAT64, 4, 4)
}

// Same as Rotate2D, except homogeneous (3x3 with the extra row/col being all zeroes with a one in the bottom right)
func HomogRotate2D(angle float64) Matrix {
	sin, cos := math.Sin(angle), math.Cos(angle)
	return *unsafeMatrixFromSlice(ScalarSlice([]interface{}{
		cos, -sin, 0,
		sin, cos, 0,
		0, 0, 1}, FLOAT64), FLOAT64, 3, 3)
}

// Same as Rotate3DX, except homogeneous (4x4 with the extra row/col being all zeroes with a one in the bottom right)
func HomogRotate3DX(angle float64) Matrix {
	sin, cos := math.Sin(angle), math.Cos(angle)
	return *unsafeMatrixFromSlice(ScalarSlice([]interface{}{
		1, 0, 0, 0,
		0, cos, -sin, 0,
		0, sin, cos, 0,
		0, 0, 0, 1}, FLOAT64), FLOAT64, 4, 4)
}

// Same as Rotate3DY, except homogeneous (4x4 with the extra row/col being all zeroes with a one in the bottom right)
func HomogRotate3DY(angle float64) Matrix {
	sin, cos := math.Sin(angle), math.Cos(angle)
	return *unsafeMatrixFromSlice(ScalarSlice([]interface{}{
		cos, 0, sin, 0,
		0, 1, 0, 0,
		-sin, 0, cos, 0,
		0, 0, 0, 1}, FLOAT64), FLOAT64, 4, 4)
}

// Same as Rotate3DZ, except homogeneous (4x4 with the extra row/col being all zeroes with a one in the bottom right)
func HomogRotate3DZ(angle float64) Matrix {
	sin, cos := math.Sin(angle), math.Cos(angle)
	return *unsafeMatrixFromSlice(ScalarSlice([]interface{}{
		cos, -sin, 0, 0,
		-sin, cos, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1}, FLOAT64), FLOAT64, 4, 4)
}

func Scale3D(scaleX, scaleY, scaleZ float64) Matrix {
	mp,_ := InferMatrixFromSlice([]interface{}{
		scaleX, 0., 0., 0.,
		0., scaleY, 0., 0.,
		0., 0., scaleZ, 0.,
		0., 0., 0., 1.}, 4, 4)
	return *mp
}

func Scale2D(scaleX, scaleY float64) Matrix {
	mp,_ := InferMatrixFromSlice([]interface{}{
		scaleX, 0., 0.,
		0., scaleY, 0.,
		0., 0., 1,}, 3, 3)
	return *mp
}

func ShearX2D(shear float64) Matrix {
	mp,_ := InferMatrixFromSlice([]interface{}{
		1., shear, 0.,
		0., 1., 0.,
		0., 0., 1,}, 3, 3)
	return *mp
}

func ShearY2D(shear float64) Matrix {
	mp,_ := InferMatrixFromSlice([]interface{}{
		1., 0., 0.,
		shear, 1., 0.,
		0., 0., 1,}, 3, 3)
	return *mp
}

func ScaleX3D(shear float64) Matrix {
	mp,_ := InferMatrixFromSlice([]interface{}{
		1., shear, shear, 0.,
		0., 1., 0., 0.,
		0., 0., 1., 0.,
		0., 0., 0., 1.}, 4, 4)
	return *mp
}

func ScaleY3D(shear float64) Matrix {
	mp,_ := InferMatrixFromSlice([]interface{}{
		1., 0., 0., 0.,
		shear, 1., shear, 0.,
		0., 0., 1., 0.,
		0., 0., 0., 1.}, 4, 4)
	return *mp
}

func ScaleZ3D(shear float64) Matrix {
	mp,_ := InferMatrixFromSlice([]interface{}{
		1., 0., 0., 0.,
		0., 1., 0., 0.,
		shear, shear, 1., 0.,
		0., 0., 0., 1.}, 4, 4)
	return *mp
}

// HomogRotate3D creates a 3D rotation Matrix of type FLOAT64 that rotates by (radian) angle about some arbitrary axis given by a Vector.
// It produces a homogeneous matrix (4x4)
//
// Where c is cos(angle) and s is sin(angle), and x, y, and z are the first, second, and third elements of the axis vector (respectively):
//
//    [[ x^2(c-1)+c, xy(c-1)-zs, xz(c-1)+ys, 0 ]]
//    [[ xy(c-1)+zs, y^2(c-1)+c, yz(c-1)-xs, 0 ]]
//    [[ xz(c-1)-ys, yz(c-1)+xs, z^2(c-1)+c, 0 ]]
//    [[ 0         , 0         , 0         , 1 ]]
//
// The axis vector's type must be FLOAT64 or you'll get the zero-type Matrix
func HomogRotate3D(angle float64, axis Vector) Matrix {
	if (axis.Size() != 4 && axis.Size() != 3) || axis.typ != FLOAT64 {
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

// QuaternionRotation creates a Quaternion of type FLOAT64 that represents a rotation about the axis given by a Vector by (radian) angle
//
// Where s is sin(angle) and c is cos(angle), and x,y,z and the first, second, and third elements of the axis vector (respectively), the Quaternion is represented by:
//
// (c/2 * 1) + (s/2 * x * i) + (s/2 * y * j) + (s/2 * z * k)
func QuaternionRotation(angle float64, axis Vector) Quaternion {
	if axis.typ != FLOAT64 {
		return Quaternion{}
	}

	sin, cos := math.Sin(angle), math.Cos(angle)
	return Quaternion{MakeScalar(cos/float64(2), FLOAT64), axis.ScalarMul(MakeScalar(sin/float64(2), FLOAT64)), FLOAT64}
}
