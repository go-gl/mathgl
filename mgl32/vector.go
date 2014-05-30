package mgl32

import (
	"math"
)

type Vec2 [2]float32
type Vec3 [3]float32
type Vec4 [4]float32

func (v1 Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{v1[0] + v2[0], v1[1] + v2[1]}
}

func (v1 Vec3) Add(v2 Vec3) Vec3 {
	return Vec3{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2]}
}

func (v1 Vec4) Add(v2 Vec4) Vec4 {
	return Vec4{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2], v1[3] + v2[3]}
}

func (v1 Vec2) Sub(v2 Vec2) Vec2 {
	return Vec2{v1[0] - v2[0], v1[1] - v2[1]}
}

func (v1 Vec3) Sub(v2 Vec3) Vec3 {
	return Vec3{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2]}
}

func (v1 Vec4) Sub(v2 Vec4) Vec4 {
	return Vec4{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2], v1[3] - v2[3]}
}

func (v1 Vec2) Mul(c float32) Vec2 {
	return Vec2{v1[0] * c, v1[1] * c}
}

func (v1 Vec3) Mul(c float32) Vec3 {
	return Vec3{v1[0] * c, v1[1] * c, v1[2] * c}
}

func (v1 Vec4) Mul(c float32) Vec4 {
	return Vec4{v1[0] * c, v1[1] * c, v1[2] * c, v1[3] * c}
}

func (v1 Vec2) Dot(v2 Vec2) float32 {
	return v1[0]*v2[0] + v1[1]*v2[1]
}

func (v1 Vec3) Dot(v2 Vec3) float32 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}

func (v1 Vec4) Dot(v2 Vec4) float32 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2] + v1[3]*v2[3]
}

func (v1 Vec2) Len() float32 {
	return float32(math.Sqrt(float64(v1[0]*v1[0] + v1[1]*v1[1])))
}

func (v1 Vec3) Len() float32 {
	return float32(math.Sqrt(float64(v1[0]*v1[0] + v1[1]*v1[1] + v1[2]*v1[2])))
}

func (v1 Vec4) Len() float32 {
	return float32(math.Sqrt(float64(v1[0]*v1[0] + v1[1]*v1[1] + v1[2]*v1[2] + v1[3]*v1[3])))
}

func (v1 Vec2) Normalize() Vec2 {
	l := 1.0 / math.Sqrt(float64(v1[0]*v1[0]+v1[1]*v1[1]))
	return Vec2{float32(float64(v1[0]) * l), float32(float64(v1[1]) * l)}
}

func (v1 Vec3) Normalize() Vec3 {
	l := 1.0 / math.Sqrt(float64(v1[0]*v1[0]+v1[1]*v1[1]+v1[2]*v1[2]))
	return Vec3{float32(float64(v1[0]) * l), float32(float64(v1[1]) * l), float32(float64(v1[2]) * l)}
}

func (v1 Vec4) Normalize() Vec4 {
	l := 1.0 / math.Sqrt(float64(v1[0]*v1[0]+v1[1]*v1[1]+v1[2]*v1[2]+v1[3]*v1[3]))
	return Vec4{float32(float64(v1[0]) * l), float32(float64(v1[1]) * l), float32(float64(v1[2]) * l), float32(float64(v1[3]) * l)}
}

func (v1 Vec3) Cross(v2 Vec3) Vec3 {
	return Vec3{v1[1]*v2[2] - v1[2]*v2[1], v1[2]*v2[0] - v1[0]*v2[2], v1[0]*v2[1] - v1[1]*v2[0]}
}

func (v1 Vec2) ApproxEqual(v2 Vec2) bool {
	for i := range v1 {
		if !FloatEqual(v1[i], v2[i]) {
			return false
		}
	}
	return true
}

func (v1 Vec3) ApproxEqual(v2 Vec3) bool {
	for i := range v1 {
		if !FloatEqual(v1[i], v2[i]) {
			return false
		}
	}
	return true
}

func (v1 Vec4) ApproxEqual(v2 Vec4) bool {
	for i := range v1 {
		if !FloatEqual(v1[i], v2[i]) {
			return false
		}
	}
	return true
}

func (v1 Vec2) ApproxEqualThreshold(v2 Vec2, threshold float32) bool {
	for i := range v1 {
		if !FloatEqualThreshold(v1[i], v2[i], threshold) {
			return false
		}
	}
	return true
}

func (v1 Vec3) ApproxEqualThreshold(v2 Vec3, threshold float32) bool {
	for i := range v1 {
		if !FloatEqualThreshold(v1[i], v2[i], threshold) {
			return false
		}
	}
	return true
}

func (v1 Vec4) ApproxEqualThreshold(v2 Vec4, threshold float32) bool {
	for i := range v1 {
		if !FloatEqualThreshold(v1[i], v2[i], threshold) {
			return false
		}
	}
	return true
}

func (v1 Vec2) ApproxFuncEqual(v2 Vec2, eq func(float32, float32) bool) bool {
	for i := range v1 {
		if !eq(v1[i], v2[i]) {
			return false
		}
	}
	return true
}

func (v1 Vec3) ApproxFuncEqual(v2 Vec3, eq func(float32, float32) bool) bool {
	for i := range v1 {
		if !eq(v1[i], v2[i]) {
			return false
		}
	}
	return true
}

func (v1 Vec4) ApproxFuncEqual(v2 Vec4, eq func(float32, float32) bool) bool {
	for i := range v1 {
		if !eq(v1[i], v2[i]) {
			return false
		}
	}
	return true
}
