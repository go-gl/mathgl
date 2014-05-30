package mathgl

import (
	"math"
)

type Vec2f [2]float32
type Vec3f [3]float32
type Vec4f [4]float32

func (v1 Vec2f) Add(v2 Vec2f) Vec2f {
	return Vec2f{v1[0] + v2[0], v1[1] + v2[1]}
}

func (v1 Vec3f) Add(v2 Vec3f) Vec3f {
	return Vec3f{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2]}
}

func (v1 Vec4f) Add(v2 Vec4f) Vec4f {
	return Vec4f{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2], v1[3] + v2[3]}
}

func (v1 Vec2f) Sub(v2 Vec2f) Vec2f {
	return Vec2f{v1[0] - v2[0], v1[1] - v2[1]}
}

func (v1 Vec3f) Sub(v2 Vec3f) Vec3f {
	return Vec3f{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2]}
}

func (v1 Vec4f) Sub(v2 Vec4f) Vec4f {
	return Vec4f{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2], v1[3] - v2[3]}
}

func (v1 Vec2f) Mul(c float32) Vec2f {
	return Vec2f{v1[0] * c, v1[1] * c}
}

func (v1 Vec3f) Mul(c float32) Vec3f {
	return Vec3f{v1[0] * c, v1[1] * c, v1[2] * c}
}

func (v1 Vec4f) Mul(c float32) Vec4f {
	return Vec4f{v1[0] * c, v1[1] * c, v1[2] * c, v1[3] * c}
}

func (v1 Vec2f) Dot(v2 Vec2f) float32 {
	return v1[0]*v2[0] + v1[1]*v2[1]
}

func (v1 Vec3f) Dot(v2 Vec3f) float32 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}

func (v1 Vec4f) Dot(v2 Vec4f) float32 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2] + v1[3]*v2[3]
}

func (v1 Vec2f) Len() float32 {
	return float32(math.Sqrt(float64(v1[0]*v1[0] + v1[1]*v1[1])))
}

func (v1 Vec3f) Len() float32 {
	return float32(math.Sqrt(float64(v1[0]*v1[0] + v1[1]*v1[1] + v1[2]*v1[2])))
}

func (v1 Vec4f) Len() float32 {
	return float32(math.Sqrt(float64(v1[0]*v1[0] + v1[1]*v1[1] + v1[2]*v1[2] + v1[3]*v1[3])))
}

func (v1 Vec2f) Normalize() Vec2f {
	l := 1.0 / math.Sqrt(float64(v1[0]*v1[0]+v1[1]*v1[1]))
	return Vec2f{float32(float64(v1[0]) * l), float32(float64(v1[1]) * l)}
}

func (v1 Vec3f) Normalize() Vec3f {
	l := 1.0 / math.Sqrt(float64(v1[0]*v1[0]+v1[1]*v1[1]+v1[2]*v1[2]))
	return Vec3f{float32(float64(v1[0]) * l), float32(float64(v1[1]) * l), float32(float64(v1[2]) * l)}
}

func (v1 Vec4f) Normalize() Vec4f {
	l := 1.0 / math.Sqrt(float64(v1[0]*v1[0]+v1[1]*v1[1]+v1[2]*v1[2]+v1[3]*v1[3]))
	return Vec4f{float32(float64(v1[0]) * l), float32(float64(v1[1]) * l), float32(float64(v1[2]) * l), float32(float64(v1[3]) * l)}
}

func (v1 Vec3f) Cross(v2 Vec3f) Vec3f {
	return Vec3f{v1[1]*v2[2] - v1[2]*v2[1], v1[2]*v2[0] - v1[0]*v2[2], v1[0]*v2[1] - v1[1]*v2[0]}
}

func (v1 Vec2f) ApproxEqual(v2 Vec2f) bool {
	for i := range v1 {
		if !FloatEqual32(v1[i], v2[i]) {
			return false
		}
	}
	return true
}

func (v1 Vec3f) ApproxEqual(v2 Vec3f) bool {
	for i := range v1 {
		if !FloatEqual32(v1[i], v2[i]) {
			return false
		}
	}
	return true
}

func (v1 Vec4f) ApproxEqual(v2 Vec4f) bool {
	for i := range v1 {
		if !FloatEqual32(v1[i], v2[i]) {
			return false
		}
	}
	return true
}

func (v1 Vec2f) ApproxEqualTheshold(v2 Vec2f, threshold float32) bool {
	for i := range v1 {
		if !FloatEqualThreshold32(v1[i], v2[i], threshold) {
			return false
		}
	}
	return true
}

func (v1 Vec3f) ApproxEqualTheshold(v2 Vec3f, threshold float32) bool {
	for i := range v1 {
		if !FloatEqualThreshold32(v1[i], v2[i], threshold) {
			return false
		}
	}
	return true
}

func (v1 Vec4f) ApproxEqualTheshold(v2 Vec4f, threshold float32) bool {
	for i := range v1 {
		if !FloatEqualThreshold32(v1[i], v2[i], threshold) {
			return false
		}
	}
	return true
}

func (v1 Vec2f) ApproxFuncEqual(v2 Vec2f, eq func(float32, float32) bool) bool {
	for i := range v1 {
		if !eq(v1[i], v2[i]) {
			return false
		}
	}
	return true
}

func (v1 Vec3f) ApproxFuncEqual(v2 Vec3f, eq func(float32, float32) bool) bool {
	for i := range v1 {
		if !eq(v1[i], v2[i]) {
			return false
		}
	}
	return true
}

func (v1 Vec4f) ApproxFuncEqual(v2 Vec4f, eq func(float32, float32) bool) bool {
	for i := range v1 {
		if !eq(v1[i], v2[i]) {
			return false
		}
	}
	return true
}
