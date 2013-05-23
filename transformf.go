package mathgl

import "math"

// Rotate2D returns a rotation Matrix about a angle in 2-D space. Specifically about the origin.
// It is a 2x2 matrix, if you need a 3x3 for Homogeneous math (e.g. composition with a Translation matrix)
// see HomogRotate2D
func Rotate2D(angle float64) Mat2f {
	angle = (angle * math.Pi) / 180.0
	sin, cos := float32(math.Sin(angle)), float32(math.Cos(angle))
	return Mat2f{cos, sin, -sin, cos}
}

// Rotate3DX returns a 3x3 (non-homogeneous) Matrix that rotates by angle about the X-axis
//
// Where c is cos(angle) and s is sin(angle)
//    [1 0 0]
//    [0 c -s]
//    [0 s c ]
func Rotate3DX(angle float64) Mat3f {
	angle = (angle * math.Pi) / 180.0
	sin, cos := float32(math.Sin(angle)), float32(math.Cos(angle))

	return Mat3f{1, 0, 0, 0, cos, sin, 0, -sin, cos}
}

// Rotate3DY returns a 3x3 (non-homogeneous) Matrix that rotates by angle about the Y-axis
//
// Where c is cos(angle) and s is sin(angle)
//    [c 0 s]
//    [0 1 0]
//    [s 0 c ]
func Rotate3DY(angle float64) Mat3f {
	angle = (angle * math.Pi) / 180.0
	sin, cos := float32(math.Sin(angle)), float32(math.Cos(angle))

	return Mat3f{cos, 0, -sin, 0, 1, 0, sin, 0, cos}
}

// Rotate3DZ returns a 3x3 (non-homogeneous) Matrix that rotates by angle about the Z-axis
//
// Where c is cos(angle) and s is sin(angle)
//    [c -s 0]
//    [s c 0]
//    [0 0 1 ]
func Rotate3DZ(angle float64) Mat3f {
	angle = (angle * math.Pi) / 180.0
	sin, cos := float32(math.Sin(angle)), float32(math.Cos(angle))

	return Mat3f{cos, sin, 0, -sin, cos, 0, 0, 0, 1}
}

// Translate2D returns a homogeneous (3x3 for 2D-space) Translation matrix that moves a point by Tx units in the x-direction and Ty units in the y-direction
//
//    [[1, 0, Tx]]
//    [[0, 1, Ty]]
//    [[0, 0, 1 ]]
func Translate2D(Tx, Ty float64) Mat3f {
	return Mat3f{1, 0, 0, 0, 1, 0, float32(Tx), float32(Ty), 1}
}

// Translate3D returns a homogeneous (4x4 for 3D-space) Translation matrix that moves a point by Tx units in the x-direction, Ty units in the y-direction,
// and Tz units in the z-direction
//
//    [[1, 0, 0, Tx]]
//    [[0, 1, 0, Ty]]
//    [[0, 0, 1, Tz]]
//    [[0, 0, 0, 1 ]]
func Translate3D(Tx, Ty, Tz float64) Mat4f {
	return Mat4f{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, float32(Tx), float32(Ty), float32(Tz), 1}
}

// Same as Rotate2D, except homogeneous (3x3 with the extra row/col being all zeroes with a one in the bottom right)
func HomogRotate2D(angle float64) Mat3f {
	angle = (angle * math.Pi) / 180.0
	sin, cos := float32(math.Sin(angle)), float32(math.Cos(angle))
	return Mat3f{cos, sin, 0, -sin, cos, 0, 0, 0, 1}
}

// Same as Rotate3DX, except homogeneous (4x4 with the extra row/col being all zeroes with a one in the bottom right)
func HomogRotate3DX(angle float64) Mat4f {
	angle = (angle * math.Pi) / 180.0
	sin, cos := float32(math.Sin(angle)), float32(math.Cos(angle))

	return Mat4f{1, 0, 0, 0, 0, cos, sin, 0, 0, -sin, cos, 0, 0, 0, 0, 1}
}

// Same as Rotate3DY, except homogeneous (4x4 with the extra row/col being all zeroes with a one in the bottom right)
func HomogRotate3DY(angle float64) Mat4f {
	angle = (angle * math.Pi) / 180.0
	sin, cos := float32(math.Sin(angle)), float32(math.Cos(angle))
	return Mat4f{cos, 0, -sin, 0, 0, 1, 0, 0, sin, 0, cos, 0, 0, 0, 0, 1}
}

// Same as Rotate3DZ, except homogeneous (4x4 with the extra row/col being all zeroes with a one in the bottom right)
func HomogRotate3DZ(angle float64) Mat4f {
	angle = (angle * math.Pi) / 180.0
	sin, cos := float32(math.Sin(angle)), float32(math.Cos(angle))
	return Mat4f{cos, sin, 0, 0, -sin, cos, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
}

// Scale3D creates a homogeneous 3D scaling matrix
// [[ scaleX, 0     , 0     , 0 ]]
// [[ 0     , scaleY, 0     , 0 ]]
// [[ 0     , 0     , scaleZ, 0 ]]
// [[ 0     , 0     , 0     , 1 ]]
func Scale3D(scaleX, scaleY, scaleZ float64) Mat4f {

	return Mat4f{float32(scaleX), 0, 0, 0, 0, float32(scaleY), 0, 0, 0, 0, float32(scaleZ), 0, 0, 0, 0, 1}
}

// Scale2D creates a homogeneous 2D scaling matrix
// [[ scaleX, 0     , 0 ]]
// [[ 0     , scaleY, 0 ]]
// [[ 0     , 0     , 1 ]]
func Scale2D(scaleX, scaleY float64) Mat3f {
	return Mat3f{float32(scaleX), 0, 0, 0, float32(scaleY), 0, 0, 0, 1}
}

// ShearX2D creates a homogeneous 2D shear matrix along the X-axis
func ShearX2D(shear float64) Mat3f {
	return Mat3f{1, 0, 0, float32(shear), 1, 0, 0, 0, 1}
}

// ShearY2D creates a homogeneous 2D shear matrix along the Y-axis
func ShearY2D(shear float64) Mat3f {
	return Mat3f{1, float32(shear), 0, 0, 1, 0, 0, 0, 1}
}

// ShearX3D creates a homogeneous 3D shear matrix along the X-axis
func ShearX3D(shearY, shearZ float64) Mat4f {

	return Mat4f{1, float32(shearY), float32(shearZ), 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
}

// ShearY3D creates a homogeneous 3D shear matrix along the Y-axis
func ShearY3D(shearX, shearZ float64) Mat4f {
	return Mat4f{1, 0, 0, 0, float32(shearX), 1, float32(shearZ), 0, 0, 0, 1, 0, 0, 0, 0, 1}
}

// ShearZ3D creates a homogeneous 3D shear matrix along the Z-axis
func ShearZ3D(shearX, shearY float64) Mat4f {
	return Mat4f{1, 0, 0, 0, 0, 1, 0, 0, float32(shearX), float32(shearY), 1, 0, 0, 0, 0, 1}
}

// HomogRotate3D creates a 3D rotation Matrix that rotates by (radian) angle about some arbitrary axis given by a Vector.
// It produces a homogeneous matrix (4x4)
//
// Where c is cos(angle) and s is sin(angle), and x, y, and z are the first, second, and third elements of the axis vector (respectively):
//
//    [[ x^2(c-1)+c, xy(c-1)-zs, xz(c-1)+ys, 0 ]]
//    [[ xy(c-1)+zs, y^2(c-1)+c, yz(c-1)-xs, 0 ]]
//    [[ xz(c-1)-ys, yz(c-1)+xs, z^2(c-1)+c, 0 ]]
//    [[ 0         , 0         , 0         , 1 ]]
func HomogRotate3D(angle float64, axis Vec3f) Mat4f {
	x, y, z := axis[0], axis[1], axis[2]
	s, c := float32(math.Sin(angle)), float32(math.Cos(angle))
	k := 1 - c

	return Mat4f{x*x*k + c, x*y*k + z*s, x*z*k - y*s, 0, x*y*k - z*s, y*y*k + c, y*z*k + x*s, 0, x*z*k + y*s, y*z*k - x*s, z*z*k + c, 0, 0, 0, 0, 1}
}
