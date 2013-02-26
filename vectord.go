package mathgl

import (
	"math"
)

type Vec2d [2]float64
type Vec3d [3]float64
type Vec4d [4]float64

func (v1 Vec2d) Add(v2 Vec2d) Vec2d {
	return Vec2d{v1[0] + v2[0], v1[1] + v2[1]}
}

func (v1 Vec3d) Add(v2 Vec3d) Vec3d {
	return Vec3d{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2]}
}

func (v1 Vec4d) Add(v2 Vec4d) Vec4d {
	return Vec4d{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2], v1[3] + v2[3]}
}

func (v1 Vec2d) Sub(v2 Vec2d) Vec2d {
	return Vec2d{v1[0] - v2[0], v1[1] - v2[1]}
}

func (v1 Vec3d) Sub(v2 Vec3d) Vec3d {
	return Vec3d{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2]}
}

func (v1 Vec4d) Sub(v2 Vec4d) Vec4d {
	return Vec4d{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2], v1[3] - v2[3]}
}

func (v1 Vec2d) Mul(c float64) Vec2d {
	return Vec2d{v1[0] * c, v1[1] * c}
}

func (v1 Vec3d) Mul(c float64) Vec3d {
	return Vec3d{v1[0] * c, v1[1] * c, v1[2] * c}
}

func (v1 Vec4d) Mul(c float64) Vec4d {
	return Vec4d{v1[0] * c, v1[1] * c, v1[2] * c, v1[3] * c}
}

func (v1 Vec2d) Dot(v2 Vec2d) float64 {
	return v1[0]*v2[0] + v1[1]*v2[1]
}

func (v1 Vec3d) Dot(v2 Vec3d) float64 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}

func (v1 Vec4d) Dot(v2 Vec4d) float64 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2] + v1[3]*v2[3]
}

func (v1 Vec2d) Len() float64 {
	return float64(math.Sqrt(float64(v1[0]*v1[0] + v1[1]*v1[1])))
}

func (v1 Vec3d) Len() float64 {
	return float64(math.Sqrt(float64(v1[0]*v1[0] + v1[1]*v1[1] + v1[2]*v1[2])))
}

func (v1 Vec4d) Len() float64 {
	return float64(math.Sqrt(float64(v1[0]*v1[0] + v1[1]*v1[1] + v1[2]*v1[2] + v1[3]*v1[3])))
}

func (v1 Vec2d) Normalize() Vec2d {
	l := 1.0 / math.Sqrt(float64(v1[0]*v1[0]+v1[1]*v1[1]))
	return Vec2d{float64(float64(v1[0]) * l), float64(float64(v1[1]) * l)}
}

func (v1 Vec3d) Normalize() Vec3d {
	l := 1.0 / math.Sqrt(float64(v1[0]*v1[0]+v1[1]*v1[1]+v1[2]*v1[2]))
	return Vec3d{float64(float64(v1[0]) * l), float64(float64(v1[1]) * l), float64(float64(v1[2]) * l)}
}

func (v1 Vec4d) Normalize() Vec4d {
	l := 1.0 / math.Sqrt(float64(v1[0]*v1[0]+v1[1]*v1[1]+v1[2]*v1[2]+v1[3]*v1[3]))
	return Vec4d{float64(float64(v1[0]) * l), float64(float64(v1[1]) * l), float64(float64(v1[2]) * l), float64(float64(v1[3]) * l)}
}
