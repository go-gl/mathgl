package mathgl

import (
	"math"
)

type Vec2i [2]int32
type Vec3i [3]int32
type Vec4i [4]int32

func (v1 Vec2i) Add(v2 Vec2i) Vec2i {
	return Vec2i{v1[0] + v2[0], v1[1] + v2[1]}
}

func (v1 Vec3i) Add(v2 Vec3i) Vec3i {
	return Vec3i{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2]}
}

func (v1 Vec4i) Add(v2 Vec4i) Vec4i {
	return Vec4i{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2], v1[3] + v2[3]}
}

func (v1 Vec2i) Sub(v2 Vec2i) Vec2i {
	return Vec2i{v1[0] - v2[0], v1[1] - v2[1]}
}

func (v1 Vec3i) Sub(v2 Vec3i) Vec3i {
	return Vec3i{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2]}
}

func (v1 Vec4i) Sub(v2 Vec4i) Vec4i {
	return Vec4i{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2], v1[3] - v2[3]}
}

func (v1 Vec2i) Mul(c int32) Vec2i {
	return Vec2i{v1[0] * c, v1[1] * c}
}

func (v1 Vec3i) Mul(c int32) Vec3i {
	return Vec3i{v1[0] * c, v1[1] * c, v1[2] * c}
}

func (v1 Vec4i) Mul(c int32) Vec4i {
	return Vec4i{v1[0] * c, v1[1] * c, v1[2] * c, v1[3] * c}
}

func (v1 Vec2i) Dot(v2 Vec2i) int32 {
	return v1[0]*v2[0] + v1[1]*v2[1]
}

func (v1 Vec3i) Dot(v2 Vec3i) int32 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}

func (v1 Vec4i) Dot(v2 Vec4i) int32 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2] + v1[3]*v2[3]
}

func (v1 Vec2i) Len() int32 {
	return int32(math.Sqrt(float64(v1[0]*v1[0] + v1[1]*v1[1])))
}

func (v1 Vec3i) Len() int32 {
	return int32(math.Sqrt(float64(v1[0]*v1[0] + v1[1]*v1[1] + v1[2]*v1[2])))
}

func (v1 Vec4i) Len() int32 {
	return int32(math.Sqrt(float64(v1[0]*v1[0] + v1[1]*v1[1] + v1[2]*v1[2] + v1[3]*v1[3])))
}

func (v1 Vec2i) Normalize() Vec2i {
	l := 1.0 / math.Sqrt(float64(v1[0]*v1[0]+v1[1]*v1[1]))
	return Vec2i{int32(float64(v1[0]) * l), int32(float64(v1[1]) * l)}
}

func (v1 Vec3i) Normalize() Vec3i {
	l := 1.0 / math.Sqrt(float64(v1[0]*v1[0]+v1[1]*v1[1]+v1[2]*v1[2]))
	return Vec3i{int32(float64(v1[0]) * l), int32(float64(v1[1]) * l), int32(float64(v1[2]) * l)}
}

func (v1 Vec4i) Normalize() Vec4i {
	l := 1.0 / math.Sqrt(float64(v1[0]*v1[0]+v1[1]*v1[1]+v1[2]*v1[2]+v1[3]*v1[3]))
	return Vec4i{int32(float64(v1[0]) * l), int32(float64(v1[1]) * l), int32(float64(v1[2]) * l), int32(float64(v1[3]) * l)}
}
