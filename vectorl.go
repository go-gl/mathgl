package mathgl

import (
	"math"
)

type Vec2l [2]int64
type Vec3l [3]int64
type Vec4l [4]int64

func (v1 Vec2l) Add(v2 Vec2l) Vec2l {
	return Vec2l{v1[0] + v2[0], v1[1] + v2[1]}
}

func (v1 Vec3l) Add(v2 Vec3l) Vec3l {
	return Vec3l{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2]}
}

func (v1 Vec4l) Add(v2 Vec4l) Vec4l {
	return Vec4l{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2], v1[3] + v2[3]}
}

func (v1 Vec2l) Sub(v2 Vec2l) Vec2l {
	return Vec2l{v1[0] - v2[0], v1[1] - v2[1]}
}

func (v1 Vec3l) Sub(v2 Vec3l) Vec3l {
	return Vec3l{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2]}
}

func (v1 Vec4l) Sub(v2 Vec4l) Vec4l {
	return Vec4l{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2], v1[3] - v2[3]}
}

func (v1 Vec2l) Mul(c int64) Vec2l {
	return Vec2l{v1[0] * c, v1[1] * c}
}

func (v1 Vec3l) Mul(c int64) Vec3l {
	return Vec3l{v1[0] * c, v1[1] * c, v1[2] * c}
}

func (v1 Vec4l) Mul(c int64) Vec4l {
	return Vec4l{v1[0] * c, v1[1] * c, v1[2] * c, v1[3] * c}
}

func (v1 Vec2l) Dot(v2 Vec2l) int64 {
	return v1[0]*v2[0] + v1[1]*v2[1]
}

func (v1 Vec3l) Dot(v2 Vec3l) int64 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}

func (v1 Vec4l) Dot(v2 Vec4l) int64 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2] + v1[3]*v2[3]
}

func (v1 Vec2l) Len() int64 {
	return int64(math.Sqrt(float64(v1[0]*v1[0] + v1[1]*v1[1])))
}

func (v1 Vec3l) Len() int64 {
	return int64(math.Sqrt(float64(v1[0]*v1[0] + v1[1]*v1[1] + v1[2]*v1[2])))
}

func (v1 Vec4l) Len() int64 {
	return int64(math.Sqrt(float64(v1[0]*v1[0] + v1[1]*v1[1] + v1[2]*v1[2] + v1[3]*v1[3])))
}

func (v1 Vec2l) Normalize() Vec2l {
	l := 1.0 / math.Sqrt(float64(v1[0]*v1[0]+v1[1]*v1[1]))
	return Vec2l{int64(float64(v1[0]) * l), int64(float64(v1[1]) * l)}
}

func (v1 Vec3l) Normalize() Vec3l {
	l := 1.0 / math.Sqrt(float64(v1[0]*v1[0]+v1[1]*v1[1]+v1[2]*v1[2]))
	return Vec3l{int64(float64(v1[0]) * l), int64(float64(v1[1]) * l), int64(float64(v1[2]) * l)}
}

func (v1 Vec4l) Normalize() Vec4l {
	l := 1.0 / math.Sqrt(float64(v1[0]*v1[0]+v1[1]*v1[1]+v1[2]*v1[2]+v1[3]*v1[3]))
	return Vec4l{int64(float64(v1[0]) * l), int64(float64(v1[1]) * l), int64(float64(v1[2]) * l), int64(float64(v1[3]) * l)}
}
