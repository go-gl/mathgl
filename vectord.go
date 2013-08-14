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

func (v1 Vec3d) Cross(v2 Vec3d) Vec3d {
	return Vec3d{v1[1]*v2[2] - v1[2]*v2[1], v1[2]*v2[0] - v1[0]*v2[2], v1[0]*v2[1] - v1[1]*v2[0]}
}

func (v1 Vec2d) ApproxEqual(v2 Vec2d) bool {
	for i := range v1 {
		if !FloatEqual(v1[i], v2[i]) {
			return false
		}
	}
	return true
}

func (v1 Vec3d) ApproxEqual(v2 Vec3d) bool {
	for i := range v1 {
		if !FloatEqual(v1[i], v2[i]) {
			return false
		}
	}
	return true
}

func (v1 Vec4d) ApproxEqual(v2 Vec4d) bool {
	for i := range v1 {
		if !FloatEqual(v1[i], v2[i]) {
			return false
		}
	}
	return true
}

func (v1 Vec2d) ApproxEqualTheshold(v2 Vec2d, threshold float64) bool {
	for i := range v1 {
		if !FloatEqualThreshold(v1[i], v2[i], threshold) {
			return false
		}
	}
	return true
}

func (v1 Vec3d) ApproxEqualTheshold(v2 Vec3d, threshold float64) bool {
	for i := range v1 {
		if !FloatEqualThreshold(v1[i], v2[i], threshold) {
			return false
		}
	}
	return true
}

func (v1 Vec4d) ApproxEqualTheshold(v2 Vec4d, threshold float64) bool {
	for i := range v1 {
		if !FloatEqualThreshold(v1[i], v2[i], threshold) {
			return false
		}
	}
	return true
}

func (v1 Vec2d) ApproxFuncEqual(v2 Vec2d, eq func(float64, float64) bool) bool {
	for i := range v1 {
		if !eq(v1[i], v2[i]) {
			return false
		}
	}
	return true
}

func (v1 Vec3d) ApproxFuncEqual(v2 Vec3d, eq func(float64, float64) bool) bool {
	for i := range v1 {
		if !eq(v1[i], v2[i]) {
			return false
		}
	}
	return true
}

func (v1 Vec4d) ApproxFuncEqual(v2 Vec4d, eq func(float64, float64) bool) bool {
	for i := range v1 {
		if !eq(v1[i], v2[i]) {
			return false
		}
	}
	return true
}
