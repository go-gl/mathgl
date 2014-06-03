// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mgl32

import (
	"math"
)

type Vec2 [2]float32
type Vec3 [3]float32
type Vec4 [4]float32

// Add performs element-wise addition between two vectors. It is equivalent to iterating
// over every element of v1 and adding the corresponding element of v2 to it.
func (v1 Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{v1[0] + v2[0], v1[1] + v2[1]}
}

// Add performs element-wise addition between two vectors. It is equivalent to iterating
// over every element of v1 and adding the corresponding element of v2 to it.
func (v1 Vec3) Add(v2 Vec3) Vec3 {
	return Vec3{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2]}
}

// Add performs element-wise addition between two vectors. It is equivalent to iterating
// over every element of v1 and adding the corresponding element of v2 to it.
func (v1 Vec4) Add(v2 Vec4) Vec4 {
	return Vec4{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2], v1[3] + v2[3]}
}

// Sub performs element-wise subtraction between two vectors. It is equivalent to iterating
// over every element of v1 and subtracting the corresponding element of v2 from it.
func (v1 Vec2) Sub(v2 Vec2) Vec2 {
	return Vec2{v1[0] - v2[0], v1[1] - v2[1]}
}

// Sub performs element-wise subtraction between two vectors. It is equivalent to iterating
// over every element of v1 and subtracting the corresponding element of v2 from it.
func (v1 Vec3) Sub(v2 Vec3) Vec3 {
	return Vec3{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2]}
}

// Sub performs element-wise subtraction between two vectors. It is equivalent to iterating
// over every element of v1 and subtracting the corresponding element of v2 from it.
func (v1 Vec4) Sub(v2 Vec4) Vec4 {
	return Vec4{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2], v1[3] - v2[3]}
}

// Mul performs a scalar multiplication between the vector and some constant value
// c. This is equivalent to iterating over every vector element and multiplying by c.
func (v1 Vec2) Mul(c float32) Vec2 {
	return Vec2{v1[0] * c, v1[1] * c}
}

// Mul performs a scalar multiplication between the vector and some constant value
// c. This is equivalent to iterating over every vector element and multiplying by c.
func (v1 Vec3) Mul(c float32) Vec3 {
	return Vec3{v1[0] * c, v1[1] * c, v1[2] * c}
}

// Mul performs a scalar multiplication between the vector and some constant value
// c. This is equivalent to iterating over every vector element and multiplying by c.
func (v1 Vec4) Mul(c float32) Vec4 {
	return Vec4{v1[0] * c, v1[1] * c, v1[2] * c, v1[3] * c}
}

// Dot returns the dot product of this vector with another. There are multiple ways
// to describe this value. One is the multiplication of their lengths and cos(theta) where
// theta is the angle between the vectors: v1.v2 = |v1||v2|cos(theta).
//
// The other (and what is actually done) is the sum of the element-wise multiplication of all
// elements. So for instance, two Vec3s would yield v1.x * v2.x + v1.y * v2.y + v1.z * v2.z.
//
// This means that the dot product of a vector and itself is the square of its Len (within
// the bounds of floating points error).
//
// The dot product is roughly a measure of how closely two vectors are to pointing in the same
// direction. If both vectors are normalized, the value will be -1 for opposite pointing,
// one for same pointing, and 0 for perpendicular vectors.
func (v1 Vec2) Dot(v2 Vec2) float32 {
	return v1[0]*v2[0] + v1[1]*v2[1]
}

// Dot returns the dot product of this vector with another. There are multiple ways
// to describe this value. One is the multiplication of their lengths and cos(theta) where
// theta is the angle between the vectors: v1.v2 = |v1||v2|cos(theta).
//
// The other (and what is actually done) is the sum of the element-wise multiplication of all
// elements. So for instance, two Vec3s would yield v1.x * v2.x + v1.y * v2.y + v1.z * v2.z.
//
// This means that the dot product of a vector and itself is the square of its Len (within
// the bounds of floating points error).
//
// The dot product is roughly a measure of how closely two vectors are to pointing in the same
// direction. If both vectors are normalized, the value will be -1 for opposite pointing,
// one for same pointing, and 0 for perpendicular vectors.
func (v1 Vec3) Dot(v2 Vec3) float32 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}

// Dot returns the dot product of this vector with another. There are multiple ways
// to describe this value. One is the multiplication of their lengths and cos(theta) where
// theta is the angle between the vectors: v1.v2 = |v1||v2|cos(theta).
//
// The other (and what is actually done) is the sum of the element-wise multiplication of all
// elements. So for instance, two Vec3s would yield v1.x * v2.x + v1.y * v2.y + v1.z * v2.z.
//
// This means that the dot product of a vector and itself is the square of its Len (within
// the bounds of floating points error).
//
// The dot product is roughly a measure of how closely two vectors are to pointing in the same
// direction. If both vectors are normalized, the value will be -1 for opposite pointing,
// one for same pointing, and 0 for perpendicular vectors.
func (v1 Vec4) Dot(v2 Vec4) float32 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2] + v1[3]*v2[3]
}

// Len returns the vector's length. Note that this is NOT the dimension of
// the vector (len(v)), but the mathematical length. This is equivalent to the square
// root of the sum of the squares of all elements. E.G. for a Vec2 it's
// math.Hypot(v[0], v[1]).
func (v1 Vec2) Len() float32 {
	return float32(math.Hypot(float64(v1[0]), float64(v1[1])))
}

// Len returns the vector's length. Note that this is NOT the dimension of
// the vector (len(v)), but the mathematical length. This is equivalent to the square
// root of the sum of the squares of all elements. E.G. for a Vec2 it's
// math.Hypot(v[0], v[1]).
func (v1 Vec3) Len() float32 {
	return float32(math.Sqrt(float64(v1[0]*v1[0] + v1[1]*v1[1] + v1[2]*v1[2])))
}

// Len returns the vector's length. Note that this is NOT the dimension of
// the vector (len(v)), but the mathematical length. This is equivalent to the square
// root of the sum of the squares of all elements. E.G. for a Vec2 it's
// math.Hypot(v[0], v[1]).
func (v1 Vec4) Len() float32 {
	return float32(math.Sqrt(float64(v1[0]*v1[0] + v1[1]*v1[1] + v1[2]*v1[2] + v1[3]*v1[3])))
}

// Normalize normalizes the vector. Normalization is (1/|v|)*v,
// making this equivalent to v.Scale(1/v.Len()). If the len is 0.0,
// this function will return an infinite value for all elements due
// to how floating point division works in Go (n/0.0 = math.Inf(Sign(n))).
//
// Normalization makes a vector's Len become 1.0 (within the margin of floating point error),
// while maintaining its directionality.
//
// (Can be seen here: http://play.golang.org/p/Aaj7SnbqIp )
func (v1 Vec2) Normalize() Vec2 {
	l := 1.0 / v1.Len()
	return Vec2{v1[0] * l, v1[1] * l}
}

// Normalize normalizes the vector. Normalization is (1/|v|)*v,
// making this equivalent to v.Scale(1/v.Len()). If the len is 0.0,
// this function will return an infinite value for all elements due
// to how floating point division works in Go (n/0.0 = math.Inf(Sign(n))).
//
// Normalization makes a vector's Len become 1.0 (within the margin of floating point error),
// while maintaining its directionality.
//
// (Can be seen here: http://play.golang.org/p/Aaj7SnbqIp )
func (v1 Vec3) Normalize() Vec3 {
	l := 1.0 / v1.Len()
	return Vec3{v1[0] * l, v1[1] * l, v1[2] * l}
}

// Normalize normalizes the vector. Normalization is (1/|v|)*v,
// making this equivalent to v.Scale(1/v.Len()). If the len is 0.0,
// this function will return an infinite value for all elements due
// to how floating point division works in Go (n/0.0 = math.Inf(Sign(n))).
//
// Normalization makes a vector's Len become 1.0 (within the margin of floating point error),
// while maintaining its directionality.
//
// (Can be seen here: http://play.golang.org/p/Aaj7SnbqIp )
func (v1 Vec4) Normalize() Vec4 {
	l := 1.0 / v1.Len()
	return Vec4{v1[0] * l, v1[1] * l, v1[2] * l, v1[3] * l}
}

// The vector cross product is an operation only defined on 3D vectors. It is equivalent to
// Vec3{v1[1]*v2[2]-v1[2]*v2[1], v1[2]*v2[0]-v1[0]*v2[2], v1[0]*v2[1] - v1[1]*v2[0]}.
// Another interpretation is |v1||v2|sin(theta) where there is the angle between v1 and v2.
//
// Technically, a generalized cross product exists as an "(N-1)ary" operation
// (that is, the 4D cross product requires 3 4D vectors). But the binary
// 3D (and 7D) cross product is the most important. It can be considered
// the area of a parallelograph with sides v1 and v2.
//
// Like the dot product, the cross product is roughly a measure of directionality.
// Two normalized perpendicular vectors will return a value of
// 1.0 or 0.0 and two parallel vectors will return a value of 0.
// The cross product is "anticommutative" meaning v1.Cross(v2) = -v2.Cross(v1),
// this property can be useful to know when finding normals,
// as taking the wrong cross product can lead to the opposite normal of the one you want.
func (v1 Vec3) Cross(v2 Vec3) Vec3 {
	return Vec3{v1[1]*v2[2] - v1[2]*v2[1], v1[2]*v2[0] - v1[0]*v2[2], v1[0]*v2[1] - v1[1]*v2[0]}
}

// ApproxEqual takes in a vector and does an element-wise
// approximate float comparison as if FloatEqual had been used
func (v1 Vec2) ApproxEqual(v2 Vec2) bool {
	for i := range v1 {
		if !FloatEqual(v1[i], v2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqual takes in a vector and does an element-wise
// approximate float comparison as if FloatEqual had been used
func (v1 Vec3) ApproxEqual(v2 Vec3) bool {
	for i := range v1 {
		if !FloatEqual(v1[i], v2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqual takes in a vector and does an element-wise
// approximate float comparison as if FloatEqual had been used
func (v1 Vec4) ApproxEqual(v2 Vec4) bool {
	for i := range v1 {
		if !FloatEqual(v1[i], v2[i]) {
			return false
		}
	}
	return true
}

// ApproxThresholdEq takes in a threshold for comparing two floats, and uses it to do an
// element-wise comparison of the vector to another.
func (v1 Vec2) ApproxEqualThreshold(v2 Vec2, threshold float32) bool {
	for i := range v1 {
		if !FloatEqualThreshold(v1[i], v2[i], threshold) {
			return false
		}
	}
	return true
}

// ApproxThresholdEq takes in a threshold for comparing two floats, and uses it to do an
// element-wise comparison of the vector to another.
func (v1 Vec3) ApproxEqualThreshold(v2 Vec3, threshold float32) bool {
	for i := range v1 {
		if !FloatEqualThreshold(v1[i], v2[i], threshold) {
			return false
		}
	}
	return true
}

// ApproxThresholdEq takes in a threshold for comparing two floats, and uses it to do an
// element-wise comparison of the vector to another.
func (v1 Vec4) ApproxEqualThreshold(v2 Vec4, threshold float32) bool {
	for i := range v1 {
		if !FloatEqualThreshold(v1[i], v2[i], threshold) {
			return false
		}
	}
	return true
}

// ApproxFuncEq takes in a func that compares two floats, and uses it to do an element-wise
// comparison of the vector to another. This is intended to be used with FloatEqualFunc
func (v1 Vec2) ApproxFuncEqual(v2 Vec2, eq func(float32, float32) bool) bool {
	for i := range v1 {
		if !eq(v1[i], v2[i]) {
			return false
		}
	}
	return true
}

// ApproxFuncEq takes in a func that compares two floats, and uses it to do an element-wise
// comparison of the vector to another. This is intended to be used with FloatEqualFunc
func (v1 Vec3) ApproxFuncEqual(v2 Vec3, eq func(float32, float32) bool) bool {
	for i := range v1 {
		if !eq(v1[i], v2[i]) {
			return false
		}
	}
	return true
}

// ApproxFuncEq takes in a func that compares two floats, and uses it to do an element-wise
// comparison of the vector to another. This is intended to be used with FloatEqualFunc
func (v1 Vec4) ApproxFuncEqual(v2 Vec4, eq func(float32, float32) bool) bool {
	for i := range v1 {
		if !eq(v1[i], v2[i]) {
			return false
		}
	}
	return true
}
