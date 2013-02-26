package mathgl

import (
	"math"
)

type Vec2ui [2]uint32
type Vec3ui [3]uint32
type Vec4ui [4]uint32

func (v1 Vec2ui) Add(v2 Vec2ui) Vec2ui {
	return Vec2ui{v1[0] + v2[0], v1[1] + v2[1]}
}

func (v1 Vec3ui) Add(v2 Vec3ui) Vec3ui {
	return Vec3ui{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2]}
}

func (v1 Vec4ui) Add(v2 Vec4ui) Vec4ui {
	return Vec4ui{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2], v1[3] + v2[3]}
}

func (v1 Vec2ui) Sub(v2 Vec2ui) Vec2ui {
	return Vec2ui{v1[0] - v2[0], v1[1] - v2[1]}
}

func (v1 Vec3ui) Sub(v2 Vec3ui) Vec3ui {
	return Vec3ui{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2]}
}

func (v1 Vec4ui) Sub(v2 Vec4ui) Vec4ui {
	return Vec4ui{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2], v1[3] - v2[3]}
}

func (v1 Vec2ui) Mul(c uint32) Vec2ui {
	return Vec2ui{v1[0] * c, v1[1] * c}
}

func (v1 Vec3ui) Mul(c uint32) Vec3ui {
	return Vec3ui{v1[0] * c, v1[1] * c, v1[2] * c}
}

func (v1 Vec4ui) Mul(c uint32) Vec4ui {
	return Vec4ui{v1[0] * c, v1[1] * c, v1[2] * c, v1[3] * c}
}

func (v1 Vec2ui) Dot(v2 Vec2ui) uint32 {
	return v1[0]*v2[0] + v1[1]*v2[1]
}

func (v1 Vec3ui) Dot(v2 Vec3ui) uint32 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}

func (v1 Vec4ui) Dot(v2 Vec4ui) uint32 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2] + v1[3]*v2[3]
}

func (v1 Vec2ui) Len() uint32 {
	return uint32(math.Sqrt(float64(v1[0]*v1[0] + v1[1]*v1[1])))
}

func (v1 Vec3ui) Len() uint32 {
	return uint32(math.Sqrt(float64(v1[0]*v1[0] + v1[1]*v1[1] + v1[2]*v1[2])))
}

func (v1 Vec4ui) Len() uint32 {
	return uint32(math.Sqrt(float64(v1[0]*v1[0] + v1[1]*v1[1] + v1[2]*v1[2] + v1[3]*v1[3])))
}

func (v1 Vec2ui) Normalize() Vec2ui {
	l := 1.0 / math.Sqrt(float64(v1[0]*v1[0]+v1[1]*v1[1]))
	return Vec2ui{uint32(float64(v1[0]) * l), uint32(float64(v1[1]) * l)}
}

func (v1 Vec3ui) Normalize() Vec3ui {
	l := 1.0 / math.Sqrt(float64(v1[0]*v1[0]+v1[1]*v1[1]+v1[2]*v1[2]))
	return Vec3ui{uint32(float64(v1[0]) * l), uint32(float64(v1[1]) * l), uint32(float64(v1[2]) * l)}
}

func (v1 Vec4ui) Normalize() Vec4ui {
	l := 1.0 / math.Sqrt(float64(v1[0]*v1[0]+v1[1]*v1[1]+v1[2]*v1[2]+v1[3]*v1[3]))
	return Vec4ui{uint32(float64(v1[0]) * l), uint32(float64(v1[1]) * l), uint32(float64(v1[2]) * l), uint32(float64(v1[3]) * l)}
}
