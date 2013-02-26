package mathgl

import (
	"math"
)

type Vec2ul [2]uint64
type Vec3ul [3]uint64
type Vec4ul [4]uint64

func (v1 Vec2ul) Add(v2 Vec2ul) Vec2ul {
	return Vec2ul{v1[0] + v2[0], v1[1] + v2[1]}
}

func (v1 Vec3ul) Add(v2 Vec3ul) Vec3ul {
	return Vec3ul{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2]}
}

func (v1 Vec4ul) Add(v2 Vec4ul) Vec4ul {
	return Vec4ul{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2], v1[3] + v2[3]}
}

func (v1 Vec2ul) Sub(v2 Vec2ul) Vec2ul {
	return Vec2ul{v1[0] - v2[0], v1[1] - v2[1]}
}

func (v1 Vec3ul) Sub(v2 Vec3ul) Vec3ul {
	return Vec3ul{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2]}
}

func (v1 Vec4ul) Sub(v2 Vec4ul) Vec4ul {
	return Vec4ul{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2], v1[3] - v2[3]}
}

func (v1 Vec2ul) Mul(c uint64) Vec2ul {
	return Vec2ul{v1[0] * c, v1[1] * c}
}

func (v1 Vec3ul) Mul(c uint64) Vec3ul {
	return Vec3ul{v1[0] * c, v1[1] * c, v1[2] * c}
}

func (v1 Vec4ul) Mul(c uint64) Vec4ul {
	return Vec4ul{v1[0] * c, v1[1] * c, v1[2] * c, v1[3] * c}
}

func (v1 Vec2ul) Dot(v2 Vec2ul) uint64 {
	return v1[0]*v2[0] + v1[1]*v2[1]
}

func (v1 Vec3ul) Dot(v2 Vec3ul) uint64 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}

func (v1 Vec4ul) Dot(v2 Vec4ul) uint64 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2] + v1[3]*v2[3]
}

func (v1 Vec2ul) Len() uint64 {
	return uint64(math.Sqrt(float64(v1[0]*v1[0] + v1[1]*v1[1])))
}

func (v1 Vec3ul) Len() uint64 {
	return uint64(math.Sqrt(float64(v1[0]*v1[0] + v1[1]*v1[1] + v1[2]*v1[2])))
}

func (v1 Vec4ul) Len() uint64 {
	return uint64(math.Sqrt(float64(v1[0]*v1[0] + v1[1]*v1[1] + v1[2]*v1[2] + v1[3]*v1[3])))
}

func (v1 Vec2ul) Normalize() Vec2ul {
	l := 1.0 / math.Sqrt(float64(v1[0]*v1[0]+v1[1]*v1[1]))
	return Vec2ul{uint64(float64(v1[0]) * l), uint64(float64(v1[1]) * l)}
}

func (v1 Vec3ul) Normalize() Vec3ul {
	l := 1.0 / math.Sqrt(float64(v1[0]*v1[0]+v1[1]*v1[1]+v1[2]*v1[2]))
	return Vec3ul{uint64(float64(v1[0]) * l), uint64(float64(v1[1]) * l), uint64(float64(v1[2]) * l)}
}

func (v1 Vec4ul) Normalize() Vec4ul {
	l := 1.0 / math.Sqrt(float64(v1[0]*v1[0]+v1[1]*v1[1]+v1[2]*v1[2]+v1[3]*v1[3]))
	return Vec4ul{uint64(float64(v1[0]) * l), uint64(float64(v1[1]) * l), uint64(float64(v1[2]) * l), uint64(float64(v1[3]) * l)}
}
