package mgl32

import (
//"math"
)

type Mat2 [4]float32
type Mat2x3 [6]float32
type Mat2x4 [8]float32
type Mat3x2 [6]float32
type Mat3 [9]float32
type Mat3x4 [12]float32
type Mat4x2 [8]float32
type Mat4x3 [12]float32
type Mat4 [16]float32

func Ident2() Mat2 {
	return Mat2{1, 0, 0, 1}
}

func Ident3() Mat3 {
	return Mat3{1, 0, 0, 0, 1, 0, 0, 0, 1}
}

func Ident4() Mat4 {
	return Mat4{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
}

func (m1 Mat2) Add(m2 Mat2) Mat2 {
	return Mat2{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3]}
}

func (m1 Mat2x3) Add(m2 Mat2x3) Mat2x3 {
	return Mat2x3{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5]}
}

func (m1 Mat2x4) Add(m2 Mat2x4) Mat2x4 {
	return Mat2x4{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7]}
}

func (m1 Mat3x2) Add(m2 Mat3x2) Mat3x2 {
	return Mat3x2{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5]}
}

func (m1 Mat3) Add(m2 Mat3) Mat3 {
	return Mat3{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7], m1[8] + m2[8]}
}

func (m1 Mat3x4) Add(m2 Mat3x4) Mat3x4 {
	return Mat3x4{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7], m1[8] + m2[8], m1[9] + m2[9], m1[10] + m2[10], m1[11] + m2[11]}
}

func (m1 Mat4x2) Add(m2 Mat4x2) Mat4x2 {
	return Mat4x2{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7]}
}

func (m1 Mat4x3) Add(m2 Mat4x3) Mat4x3 {
	return Mat4x3{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7], m1[8] + m2[8], m1[9] + m2[9], m1[10] + m2[10], m1[11] + m2[11]}
}

func (m1 Mat4) Add(m2 Mat4) Mat4 {
	return Mat4{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7], m1[8] + m2[8], m1[9] + m2[9], m1[10] + m2[10], m1[11] + m2[11], m1[12] + m2[12], m1[13] + m2[13], m1[14] + m2[14], m1[15] + m2[15]}
}

func (m1 Mat2) Sub(m2 Mat2) Mat2 {
	return Mat2{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3]}
}

func (m1 Mat2x3) Sub(m2 Mat2x3) Mat2x3 {
	return Mat2x3{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5]}
}

func (m1 Mat2x4) Sub(m2 Mat2x4) Mat2x4 {
	return Mat2x4{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7]}
}

func (m1 Mat3x2) Sub(m2 Mat3x2) Mat3x2 {
	return Mat3x2{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5]}
}

func (m1 Mat3) Sub(m2 Mat3) Mat3 {
	return Mat3{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7], m1[8] - m2[8]}
}

func (m1 Mat3x4) Sub(m2 Mat3x4) Mat3x4 {
	return Mat3x4{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7], m1[8] - m2[8], m1[9] - m2[9], m1[10] - m2[10], m1[11] - m2[11]}
}

func (m1 Mat4x2) Sub(m2 Mat4x2) Mat4x2 {
	return Mat4x2{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7]}
}

func (m1 Mat4x3) Sub(m2 Mat4x3) Mat4x3 {
	return Mat4x3{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7], m1[8] - m2[8], m1[9] - m2[9], m1[10] - m2[10], m1[11] - m2[11]}
}

func (m1 Mat4) Sub(m2 Mat4) Mat4 {
	return Mat4{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7], m1[8] - m2[8], m1[9] - m2[9], m1[10] - m2[10], m1[11] - m2[11], m1[12] - m2[12], m1[13] - m2[13], m1[14] - m2[14], m1[15] - m2[15]}
}

func (m1 Mat2) Mul(c float32) Mat2 {
	return Mat2{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c}
}

func (m1 Mat2x3) Mul(c float32) Mat2x3 {
	return Mat2x3{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c}
}

func (m1 Mat2x4) Mul(c float32) Mat2x4 {
	return Mat2x4{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c}
}

func (m1 Mat3x2) Mul(c float32) Mat3x2 {
	return Mat3x2{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c}
}

func (m1 Mat3) Mul(c float32) Mat3 {
	return Mat3{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c, m1[8] * c}
}

func (m1 Mat3x4) Mul(c float32) Mat3x4 {
	return Mat3x4{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c, m1[8] * c, m1[9] * c, m1[10] * c, m1[11] * c}
}

func (m1 Mat4x2) Mul(c float32) Mat4x2 {
	return Mat4x2{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c}
}

func (m1 Mat4x3) Mul(c float32) Mat4x3 {
	return Mat4x3{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c, m1[8] * c, m1[9] * c, m1[10] * c, m1[11] * c}
}

func (m1 Mat4) Mul(c float32) Mat4 {
	return Mat4{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c, m1[8] * c, m1[9] * c, m1[10] * c, m1[11] * c, m1[12] * c, m1[13] * c, m1[14] * c, m1[15] * c}
}

func (m1 Mat2) Mul2x1(m2 Vec2) Vec2 {
	return Vec2{m1[0]*m2[0] + m1[2]*m2[1], m1[1]*m2[0] + m1[3]*m2[1]}
}

func (m1 Mat2) Mul2(m2 Mat2) Mat2 {
	return Mat2{m1[0]*m2[0] + m1[2]*m2[1], m1[1]*m2[0] + m1[3]*m2[1], m1[0]*m2[2] + m1[2]*m2[3], m1[1]*m2[2] + m1[3]*m2[3]}
}

func (m1 Mat2) Mul2x3(m2 Mat2x3) Mat2x3 {
	return Mat2x3{m1[0]*m2[0] + m1[2]*m2[1], m1[1]*m2[0] + m1[3]*m2[1], m1[0]*m2[2] + m1[2]*m2[3], m1[1]*m2[2] + m1[3]*m2[3], m1[0]*m2[4] + m1[2]*m2[5], m1[1]*m2[4] + m1[3]*m2[5]}
}

func (m1 Mat2) Mul2x4(m2 Mat2x4) Mat2x4 {
	return Mat2x4{m1[0]*m2[0] + m1[2]*m2[1], m1[1]*m2[0] + m1[3]*m2[1], m1[0]*m2[2] + m1[2]*m2[3], m1[1]*m2[2] + m1[3]*m2[3], m1[0]*m2[4] + m1[2]*m2[5], m1[1]*m2[4] + m1[3]*m2[5], m1[0]*m2[6] + m1[2]*m2[7], m1[1]*m2[6] + m1[3]*m2[7]}
}

func (m1 Mat2x3) Mul3x1(m2 Vec3) Vec2 {
	return Vec2{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2]}
}

func (m1 Mat2x3) Mul3x2(m2 Mat3x2) Mat2 {
	return Mat2{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2], m1[0]*m2[3] + m1[2]*m2[4] + m1[4]*m2[5], m1[1]*m2[3] + m1[3]*m2[4] + m1[5]*m2[5]}
}

func (m1 Mat2x3) Mul3(m2 Mat3) Mat2x3 {
	return Mat2x3{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2], m1[0]*m2[3] + m1[2]*m2[4] + m1[4]*m2[5], m1[1]*m2[3] + m1[3]*m2[4] + m1[5]*m2[5], m1[0]*m2[6] + m1[2]*m2[7] + m1[4]*m2[8], m1[1]*m2[6] + m1[3]*m2[7] + m1[5]*m2[8]}
}

func (m1 Mat2x3) Mul3x4(m2 Mat3x4) Mat2x4 {
	return Mat2x4{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2], m1[0]*m2[3] + m1[2]*m2[4] + m1[4]*m2[5], m1[1]*m2[3] + m1[3]*m2[4] + m1[5]*m2[5], m1[0]*m2[6] + m1[2]*m2[7] + m1[4]*m2[8], m1[1]*m2[6] + m1[3]*m2[7] + m1[5]*m2[8], m1[0]*m2[9] + m1[2]*m2[10] + m1[4]*m2[11], m1[1]*m2[9] + m1[3]*m2[10] + m1[5]*m2[11]}
}

func (m1 Mat2x4) Mul4x1(m2 Vec4) Vec2 {
	return Vec2{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2] + m1[6]*m2[3], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2] + m1[7]*m2[3]}
}

func (m1 Mat2x4) Mul4x2(m2 Mat4x2) Mat2 {
	return Mat2{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2] + m1[6]*m2[3], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[2]*m2[5] + m1[4]*m2[6] + m1[6]*m2[7], m1[1]*m2[4] + m1[3]*m2[5] + m1[5]*m2[6] + m1[7]*m2[7]}
}

func (m1 Mat2x4) Mul4x3(m2 Mat4x3) Mat2x3 {
	return Mat2x3{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2] + m1[6]*m2[3], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[2]*m2[5] + m1[4]*m2[6] + m1[6]*m2[7], m1[1]*m2[4] + m1[3]*m2[5] + m1[5]*m2[6] + m1[7]*m2[7], m1[0]*m2[8] + m1[2]*m2[9] + m1[4]*m2[10] + m1[6]*m2[11], m1[1]*m2[8] + m1[3]*m2[9] + m1[5]*m2[10] + m1[7]*m2[11]}
}

func (m1 Mat2x4) Mul4(m2 Mat4) Mat2x4 {
	return Mat2x4{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2] + m1[6]*m2[3], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[2]*m2[5] + m1[4]*m2[6] + m1[6]*m2[7], m1[1]*m2[4] + m1[3]*m2[5] + m1[5]*m2[6] + m1[7]*m2[7], m1[0]*m2[8] + m1[2]*m2[9] + m1[4]*m2[10] + m1[6]*m2[11], m1[1]*m2[8] + m1[3]*m2[9] + m1[5]*m2[10] + m1[7]*m2[11], m1[0]*m2[12] + m1[2]*m2[13] + m1[4]*m2[14] + m1[6]*m2[15], m1[1]*m2[12] + m1[3]*m2[13] + m1[5]*m2[14] + m1[7]*m2[15]}
}

func (m1 Mat3x2) Mul2x1(m2 Vec2) Vec3 {
	return Vec3{m1[0]*m2[0] + m1[3]*m2[1], m1[1]*m2[0] + m1[4]*m2[1], m1[2]*m2[0] + m1[5]*m2[1]}
}

func (m1 Mat3x2) Mul2(m2 Mat2) Mat3x2 {
	return Mat3x2{m1[0]*m2[0] + m1[3]*m2[1], m1[1]*m2[0] + m1[4]*m2[1], m1[2]*m2[0] + m1[5]*m2[1], m1[0]*m2[2] + m1[3]*m2[3], m1[1]*m2[2] + m1[4]*m2[3], m1[2]*m2[2] + m1[5]*m2[3]}
}

func (m1 Mat3x2) Mul2x3(m2 Mat2x3) Mat3 {
	return Mat3{m1[0]*m2[0] + m1[3]*m2[1], m1[1]*m2[0] + m1[4]*m2[1], m1[2]*m2[0] + m1[5]*m2[1], m1[0]*m2[2] + m1[3]*m2[3], m1[1]*m2[2] + m1[4]*m2[3], m1[2]*m2[2] + m1[5]*m2[3], m1[0]*m2[4] + m1[3]*m2[5], m1[1]*m2[4] + m1[4]*m2[5], m1[2]*m2[4] + m1[5]*m2[5]}
}

func (m1 Mat3x2) Mul2x4(m2 Mat2x4) Mat3x4 {
	return Mat3x4{m1[0]*m2[0] + m1[3]*m2[1], m1[1]*m2[0] + m1[4]*m2[1], m1[2]*m2[0] + m1[5]*m2[1], m1[0]*m2[2] + m1[3]*m2[3], m1[1]*m2[2] + m1[4]*m2[3], m1[2]*m2[2] + m1[5]*m2[3], m1[0]*m2[4] + m1[3]*m2[5], m1[1]*m2[4] + m1[4]*m2[5], m1[2]*m2[4] + m1[5]*m2[5], m1[0]*m2[6] + m1[3]*m2[7], m1[1]*m2[6] + m1[4]*m2[7], m1[2]*m2[6] + m1[5]*m2[7]}
}

func (m1 Mat3) Mul3x1(m2 Vec3) Vec3 {
	return Vec3{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2]}
}

func (m1 Mat3) Mul3x2(m2 Mat3x2) Mat3x2 {
	return Mat3x2{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2], m1[0]*m2[3] + m1[3]*m2[4] + m1[6]*m2[5], m1[1]*m2[3] + m1[4]*m2[4] + m1[7]*m2[5], m1[2]*m2[3] + m1[5]*m2[4] + m1[8]*m2[5]}
}

func (m1 Mat3) Mul3(m2 Mat3) Mat3 {
	return Mat3{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2], m1[0]*m2[3] + m1[3]*m2[4] + m1[6]*m2[5], m1[1]*m2[3] + m1[4]*m2[4] + m1[7]*m2[5], m1[2]*m2[3] + m1[5]*m2[4] + m1[8]*m2[5], m1[0]*m2[6] + m1[3]*m2[7] + m1[6]*m2[8], m1[1]*m2[6] + m1[4]*m2[7] + m1[7]*m2[8], m1[2]*m2[6] + m1[5]*m2[7] + m1[8]*m2[8]}
}

func (m1 Mat3) Mul3x4(m2 Mat3x4) Mat3x4 {
	return Mat3x4{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2], m1[0]*m2[3] + m1[3]*m2[4] + m1[6]*m2[5], m1[1]*m2[3] + m1[4]*m2[4] + m1[7]*m2[5], m1[2]*m2[3] + m1[5]*m2[4] + m1[8]*m2[5], m1[0]*m2[6] + m1[3]*m2[7] + m1[6]*m2[8], m1[1]*m2[6] + m1[4]*m2[7] + m1[7]*m2[8], m1[2]*m2[6] + m1[5]*m2[7] + m1[8]*m2[8], m1[0]*m2[9] + m1[3]*m2[10] + m1[6]*m2[11], m1[1]*m2[9] + m1[4]*m2[10] + m1[7]*m2[11], m1[2]*m2[9] + m1[5]*m2[10] + m1[8]*m2[11]}
}

func (m1 Mat3x4) Mul4x1(m2 Vec4) Vec3 {
	return Vec3{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2] + m1[9]*m2[3], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2] + m1[10]*m2[3], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2] + m1[11]*m2[3]}
}

func (m1 Mat3x4) Mul4x2(m2 Mat4x2) Mat3x2 {
	return Mat3x2{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2] + m1[9]*m2[3], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2] + m1[10]*m2[3], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2] + m1[11]*m2[3], m1[0]*m2[4] + m1[3]*m2[5] + m1[6]*m2[6] + m1[9]*m2[7], m1[1]*m2[4] + m1[4]*m2[5] + m1[7]*m2[6] + m1[10]*m2[7], m1[2]*m2[4] + m1[5]*m2[5] + m1[8]*m2[6] + m1[11]*m2[7]}
}

func (m1 Mat3x4) Mul4x3(m2 Mat4x3) Mat3 {
	return Mat3{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2] + m1[9]*m2[3], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2] + m1[10]*m2[3], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2] + m1[11]*m2[3], m1[0]*m2[4] + m1[3]*m2[5] + m1[6]*m2[6] + m1[9]*m2[7], m1[1]*m2[4] + m1[4]*m2[5] + m1[7]*m2[6] + m1[10]*m2[7], m1[2]*m2[4] + m1[5]*m2[5] + m1[8]*m2[6] + m1[11]*m2[7], m1[0]*m2[8] + m1[3]*m2[9] + m1[6]*m2[10] + m1[9]*m2[11], m1[1]*m2[8] + m1[4]*m2[9] + m1[7]*m2[10] + m1[10]*m2[11], m1[2]*m2[8] + m1[5]*m2[9] + m1[8]*m2[10] + m1[11]*m2[11]}
}

func (m1 Mat3x4) Mul4(m2 Mat4) Mat3x4 {
	return Mat3x4{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2] + m1[9]*m2[3], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2] + m1[10]*m2[3], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2] + m1[11]*m2[3], m1[0]*m2[4] + m1[3]*m2[5] + m1[6]*m2[6] + m1[9]*m2[7], m1[1]*m2[4] + m1[4]*m2[5] + m1[7]*m2[6] + m1[10]*m2[7], m1[2]*m2[4] + m1[5]*m2[5] + m1[8]*m2[6] + m1[11]*m2[7], m1[0]*m2[8] + m1[3]*m2[9] + m1[6]*m2[10] + m1[9]*m2[11], m1[1]*m2[8] + m1[4]*m2[9] + m1[7]*m2[10] + m1[10]*m2[11], m1[2]*m2[8] + m1[5]*m2[9] + m1[8]*m2[10] + m1[11]*m2[11], m1[0]*m2[12] + m1[3]*m2[13] + m1[6]*m2[14] + m1[9]*m2[15], m1[1]*m2[12] + m1[4]*m2[13] + m1[7]*m2[14] + m1[10]*m2[15], m1[2]*m2[12] + m1[5]*m2[13] + m1[8]*m2[14] + m1[11]*m2[15]}
}

func (m1 Mat4x2) Mul2x1(m2 Vec2) Vec4 {
	return Vec4{m1[0]*m2[0] + m1[4]*m2[1], m1[1]*m2[0] + m1[5]*m2[1], m1[2]*m2[0] + m1[6]*m2[1], m1[3]*m2[0] + m1[7]*m2[1]}
}

func (m1 Mat4x2) Mul2(m2 Mat2) Mat4x2 {
	return Mat4x2{m1[0]*m2[0] + m1[4]*m2[1], m1[1]*m2[0] + m1[5]*m2[1], m1[2]*m2[0] + m1[6]*m2[1], m1[3]*m2[0] + m1[7]*m2[1], m1[0]*m2[2] + m1[4]*m2[3], m1[1]*m2[2] + m1[5]*m2[3], m1[2]*m2[2] + m1[6]*m2[3], m1[3]*m2[2] + m1[7]*m2[3]}
}

func (m1 Mat4x2) Mul2x3(m2 Mat2x3) Mat4x3 {
	return Mat4x3{m1[0]*m2[0] + m1[4]*m2[1], m1[1]*m2[0] + m1[5]*m2[1], m1[2]*m2[0] + m1[6]*m2[1], m1[3]*m2[0] + m1[7]*m2[1], m1[0]*m2[2] + m1[4]*m2[3], m1[1]*m2[2] + m1[5]*m2[3], m1[2]*m2[2] + m1[6]*m2[3], m1[3]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[4]*m2[5], m1[1]*m2[4] + m1[5]*m2[5], m1[2]*m2[4] + m1[6]*m2[5], m1[3]*m2[4] + m1[7]*m2[5]}
}

func (m1 Mat4x2) Mul2x4(m2 Mat2x4) Mat4 {
	return Mat4{m1[0]*m2[0] + m1[4]*m2[1], m1[1]*m2[0] + m1[5]*m2[1], m1[2]*m2[0] + m1[6]*m2[1], m1[3]*m2[0] + m1[7]*m2[1], m1[0]*m2[2] + m1[4]*m2[3], m1[1]*m2[2] + m1[5]*m2[3], m1[2]*m2[2] + m1[6]*m2[3], m1[3]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[4]*m2[5], m1[1]*m2[4] + m1[5]*m2[5], m1[2]*m2[4] + m1[6]*m2[5], m1[3]*m2[4] + m1[7]*m2[5], m1[0]*m2[6] + m1[4]*m2[7], m1[1]*m2[6] + m1[5]*m2[7], m1[2]*m2[6] + m1[6]*m2[7], m1[3]*m2[6] + m1[7]*m2[7]}
}

func (m1 Mat4x3) Mul3x1(m2 Vec3) Vec4 {
	return Vec4{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2]}
}

func (m1 Mat4x3) Mul3x2(m2 Mat3x2) Mat4x2 {
	return Mat4x2{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2], m1[0]*m2[3] + m1[4]*m2[4] + m1[8]*m2[5], m1[1]*m2[3] + m1[5]*m2[4] + m1[9]*m2[5], m1[2]*m2[3] + m1[6]*m2[4] + m1[10]*m2[5], m1[3]*m2[3] + m1[7]*m2[4] + m1[11]*m2[5]}
}

func (m1 Mat4x3) Mul3(m2 Mat3) Mat4x3 {
	return Mat4x3{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2], m1[0]*m2[3] + m1[4]*m2[4] + m1[8]*m2[5], m1[1]*m2[3] + m1[5]*m2[4] + m1[9]*m2[5], m1[2]*m2[3] + m1[6]*m2[4] + m1[10]*m2[5], m1[3]*m2[3] + m1[7]*m2[4] + m1[11]*m2[5], m1[0]*m2[6] + m1[4]*m2[7] + m1[8]*m2[8], m1[1]*m2[6] + m1[5]*m2[7] + m1[9]*m2[8], m1[2]*m2[6] + m1[6]*m2[7] + m1[10]*m2[8], m1[3]*m2[6] + m1[7]*m2[7] + m1[11]*m2[8]}
}

func (m1 Mat4x3) Mul3x4(m2 Mat3x4) Mat4 {
	return Mat4{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2], m1[0]*m2[3] + m1[4]*m2[4] + m1[8]*m2[5], m1[1]*m2[3] + m1[5]*m2[4] + m1[9]*m2[5], m1[2]*m2[3] + m1[6]*m2[4] + m1[10]*m2[5], m1[3]*m2[3] + m1[7]*m2[4] + m1[11]*m2[5], m1[0]*m2[6] + m1[4]*m2[7] + m1[8]*m2[8], m1[1]*m2[6] + m1[5]*m2[7] + m1[9]*m2[8], m1[2]*m2[6] + m1[6]*m2[7] + m1[10]*m2[8], m1[3]*m2[6] + m1[7]*m2[7] + m1[11]*m2[8], m1[0]*m2[9] + m1[4]*m2[10] + m1[8]*m2[11], m1[1]*m2[9] + m1[5]*m2[10] + m1[9]*m2[11], m1[2]*m2[9] + m1[6]*m2[10] + m1[10]*m2[11], m1[3]*m2[9] + m1[7]*m2[10] + m1[11]*m2[11]}
}

func (m1 Mat4) Mul4x1(m2 Vec4) Vec4 {
	return Vec4{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2] + m1[12]*m2[3], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2] + m1[13]*m2[3], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2] + m1[14]*m2[3], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2] + m1[15]*m2[3]}
}

func (m1 Mat4) Mul4x2(m2 Mat4x2) Mat4x2 {
	return Mat4x2{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2] + m1[12]*m2[3], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2] + m1[13]*m2[3], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2] + m1[14]*m2[3], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2] + m1[15]*m2[3], m1[0]*m2[4] + m1[4]*m2[5] + m1[8]*m2[6] + m1[12]*m2[7], m1[1]*m2[4] + m1[5]*m2[5] + m1[9]*m2[6] + m1[13]*m2[7], m1[2]*m2[4] + m1[6]*m2[5] + m1[10]*m2[6] + m1[14]*m2[7], m1[3]*m2[4] + m1[7]*m2[5] + m1[11]*m2[6] + m1[15]*m2[7]}
}

func (m1 Mat4) Mul4x3(m2 Mat4x3) Mat4x3 {
	return Mat4x3{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2] + m1[12]*m2[3], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2] + m1[13]*m2[3], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2] + m1[14]*m2[3], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2] + m1[15]*m2[3], m1[0]*m2[4] + m1[4]*m2[5] + m1[8]*m2[6] + m1[12]*m2[7], m1[1]*m2[4] + m1[5]*m2[5] + m1[9]*m2[6] + m1[13]*m2[7], m1[2]*m2[4] + m1[6]*m2[5] + m1[10]*m2[6] + m1[14]*m2[7], m1[3]*m2[4] + m1[7]*m2[5] + m1[11]*m2[6] + m1[15]*m2[7], m1[0]*m2[8] + m1[4]*m2[9] + m1[8]*m2[10] + m1[12]*m2[11], m1[1]*m2[8] + m1[5]*m2[9] + m1[9]*m2[10] + m1[13]*m2[11], m1[2]*m2[8] + m1[6]*m2[9] + m1[10]*m2[10] + m1[14]*m2[11], m1[3]*m2[8] + m1[7]*m2[9] + m1[11]*m2[10] + m1[15]*m2[11]}
}

func (m1 Mat4) Mul4(m2 Mat4) Mat4 {
	return Mat4{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2] + m1[12]*m2[3], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2] + m1[13]*m2[3], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2] + m1[14]*m2[3], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2] + m1[15]*m2[3], m1[0]*m2[4] + m1[4]*m2[5] + m1[8]*m2[6] + m1[12]*m2[7], m1[1]*m2[4] + m1[5]*m2[5] + m1[9]*m2[6] + m1[13]*m2[7], m1[2]*m2[4] + m1[6]*m2[5] + m1[10]*m2[6] + m1[14]*m2[7], m1[3]*m2[4] + m1[7]*m2[5] + m1[11]*m2[6] + m1[15]*m2[7], m1[0]*m2[8] + m1[4]*m2[9] + m1[8]*m2[10] + m1[12]*m2[11], m1[1]*m2[8] + m1[5]*m2[9] + m1[9]*m2[10] + m1[13]*m2[11], m1[2]*m2[8] + m1[6]*m2[9] + m1[10]*m2[10] + m1[14]*m2[11], m1[3]*m2[8] + m1[7]*m2[9] + m1[11]*m2[10] + m1[15]*m2[11], m1[0]*m2[12] + m1[4]*m2[13] + m1[8]*m2[14] + m1[12]*m2[15], m1[1]*m2[12] + m1[5]*m2[13] + m1[9]*m2[14] + m1[13]*m2[15], m1[2]*m2[12] + m1[6]*m2[13] + m1[10]*m2[14] + m1[14]*m2[15], m1[3]*m2[12] + m1[7]*m2[13] + m1[11]*m2[14] + m1[15]*m2[15]}
}

func (m1 Mat2) Transpose() Mat2 {
	return Mat2{m1[0], m1[2], m1[1], m1[3]}
}

func (m1 Mat2x3) Transpose() Mat3x2 {
	return Mat3x2{m1[0], m1[3], m1[1], m1[4], m1[2], m1[5]}
}

func (m1 Mat2x4) Transpose() Mat4x2 {
	return Mat4x2{m1[0], m1[4], m1[1], m1[5], m1[2], m1[6], m1[3], m1[7]}
}

func (m1 Mat3x2) Transpose() Mat2x3 {
	return Mat2x3{m1[0], m1[2], m1[4], m1[1], m1[3], m1[5]}
}

func (m1 Mat3) Transpose() Mat3 {
	return Mat3{m1[0], m1[3], m1[6], m1[1], m1[4], m1[7], m1[2], m1[5], m1[8]}
}

func (m1 Mat3x4) Transpose() Mat4x3 {
	return Mat4x3{m1[0], m1[4], m1[8], m1[1], m1[5], m1[9], m1[2], m1[6], m1[10], m1[3], m1[7], m1[11]}
}

func (m1 Mat4x2) Transpose() Mat2x4 {
	return Mat2x4{m1[0], m1[2], m1[4], m1[6], m1[1], m1[3], m1[5], m1[7]}
}

func (m1 Mat4x3) Transpose() Mat3x4 {
	return Mat3x4{m1[0], m1[3], m1[6], m1[9], m1[1], m1[4], m1[7], m1[10], m1[2], m1[5], m1[8], m1[11]}
}

func (m1 Mat4) Transpose() Mat4 {
	return Mat4{m1[0], m1[4], m1[8], m1[12], m1[1], m1[5], m1[9], m1[13], m1[2], m1[6], m1[10], m1[14], m1[3], m1[7], m1[11], m1[15]}
}

func (m Mat2) Det() float32 {
	return m[0]*m[2] - m[1]*m[3]
}

func (m Mat3) Det() float32 {
	return m[0]*m[4]*m[8] + m[3]*m[7]*m[2] + m[6]*m[1]*m[5] - m[6]*m[4]*m[2] - m[3]*m[1]*m[8] - m[0]*m[7]*m[5]
}

func (m Mat4) Det() float32 {
	return m[0]*m[5]*m[10]*m[15] - m[0]*m[5]*m[11]*m[14] - m[0]*m[6]*m[9]*m[15] + m[0]*m[6]*m[11]*m[13] + m[0]*m[7]*m[9]*m[14] - m[0]*m[7]*m[10]*m[13] - m[1]*m[4]*m[10]*m[15] + m[1]*m[4]*m[11]*m[14] + m[1]*m[6]*m[8]*m[15] - m[1]*m[6]*m[11]*m[12] - m[1]*m[7]*m[8]*m[14] + m[1]*m[7]*m[10]*m[12] + m[2]*m[4]*m[9]*m[15] - m[2]*m[4]*m[11]*m[13] - m[2]*m[5]*m[8]*m[15] + m[2]*m[5]*m[11]*m[12] + m[2]*m[7]*m[8]*m[13] - m[2]*m[7]*m[9]*m[12] - m[3]*m[4]*m[9]*m[14] + m[3]*m[4]*m[10]*m[13] + m[3]*m[5]*m[8]*m[14] - m[3]*m[5]*m[10]*m[12] - m[3]*m[6]*m[8]*m[13] + m[3]*m[6]*m[9]*m[12]
}

func (m Mat2) Inv() Mat2 {
	det := m.Det()
	if FloatEqual(det, float32(0.0)) {
		return Mat2{}
	}
	retMat := Mat2{m[3], -m[1], -m[2], m[0]}
	return retMat.Mul(1 / det)
}

func (m Mat3) Inv() Mat3 {
	det := m.Det()
	if FloatEqual(det, float32(0.0)) {
		return Mat3{}
	}
	retMat := Mat3{m[4]*m[8] - m[5]*m[7], m[2]*m[7] - m[1]*m[8], m[1]*m[5] - m[2]*m[4], m[5]*m[6] - m[3]*m[8], m[0]*m[8] - m[2]*m[6], m[2]*m[3] - m[0]*m[5], m[3]*m[7] - m[4]*m[6], m[1]*m[6] - m[0]*m[7], m[0]*m[4] - m[1]*m[3]}
	return retMat.Mul(1 / det)
}

func (m Mat4) Inv() Mat4 {
	det := m.Det()
	if FloatEqual(det, float32(0.0)) {
		return Mat4{}
	}
	retMat := Mat4{-m[7]*m[10]*m[13] + m[6]*m[11]*m[13] + m[7]*m[9]*m[14] - m[5]*m[11]*m[14] - m[6]*m[9]*m[15] + m[5]*m[10]*m[15], m[3]*m[10]*m[13] - m[2]*m[11]*m[13] - m[3]*m[9]*m[14] + m[1]*m[11]*m[14] + m[2]*m[9]*m[15] - m[1]*m[10]*m[15], -m[3]*m[6]*m[13] + m[2]*m[7]*m[13] + m[3]*m[5]*m[14] - m[1]*m[7]*m[14] - m[2]*m[5]*m[15] + m[1]*m[6]*m[15], m[3]*m[6]*m[9] - m[2]*m[7]*m[9] - m[3]*m[5]*m[10] + m[1]*m[7]*m[10] + m[2]*m[5]*m[11] - m[1]*m[6]*m[11], m[7]*m[10]*m[12] - m[6]*m[11]*m[12] - m[7]*m[8]*m[14] + m[4]*m[11]*m[14] + m[6]*m[8]*m[15] - m[4]*m[10]*m[15], -m[3]*m[10]*m[12] + m[2]*m[11]*m[12] + m[3]*m[8]*m[14] - m[0]*m[11]*m[14] - m[2]*m[8]*m[15] + m[0]*m[10]*m[15], m[3]*m[6]*m[12] - m[2]*m[7]*m[12] - m[3]*m[4]*m[14] + m[0]*m[7]*m[14] + m[2]*m[4]*m[15] - m[0]*m[6]*m[15], -m[3]*m[6]*m[8] + m[2]*m[7]*m[8] + m[3]*m[4]*m[10] - m[0]*m[7]*m[10] - m[2]*m[4]*m[11] + m[0]*m[6]*m[11], -m[7]*m[9]*m[12] + m[5]*m[11]*m[12] + m[7]*m[8]*m[13] - m[4]*m[11]*m[13] - m[5]*m[8]*m[15] + m[4]*m[9]*m[15], m[3]*m[9]*m[12] - m[1]*m[11]*m[12] - m[3]*m[8]*m[13] + m[0]*m[11]*m[13] + m[1]*m[8]*m[15] - m[0]*m[9]*m[15], -m[3]*m[5]*m[12] + m[1]*m[7]*m[12] + m[3]*m[4]*m[13] - m[0]*m[7]*m[13] - m[1]*m[4]*m[15] + m[0]*m[5]*m[15], m[3]*m[5]*m[8] - m[1]*m[7]*m[8] - m[3]*m[4]*m[9] + m[0]*m[7]*m[9] + m[1]*m[4]*m[11] - m[0]*m[5]*m[11], m[6]*m[9]*m[12] - m[5]*m[10]*m[12] - m[6]*m[8]*m[13] + m[4]*m[10]*m[13] + m[5]*m[8]*m[14] - m[4]*m[9]*m[14], -m[2]*m[9]*m[12] + m[1]*m[10]*m[12] + m[2]*m[8]*m[13] - m[0]*m[10]*m[13] - m[1]*m[8]*m[14] + m[0]*m[9]*m[14], m[2]*m[5]*m[12] - m[1]*m[6]*m[12] - m[2]*m[4]*m[13] + m[0]*m[6]*m[13] + m[1]*m[4]*m[14] - m[0]*m[5]*m[14], -m[2]*m[5]*m[8] + m[1]*m[6]*m[8] + m[2]*m[4]*m[9] - m[0]*m[6]*m[9] - m[1]*m[4]*m[10] + m[0]*m[5]*m[10]}
	return retMat.Mul(1 / det)
}

func (m1 Mat2) ApproxEqual(m2 Mat2) bool {
	for i := range m1 {
		if !FloatEqual(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat2x3) ApproxEqual(m2 Mat2x3) bool {
	for i := range m1 {
		if !FloatEqual(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat2x4) ApproxEqual(m2 Mat2x4) bool {
	for i := range m1 {
		if !FloatEqual(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat3x2) ApproxEqual(m2 Mat3x2) bool {
	for i := range m1 {
		if !FloatEqual(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat3) ApproxEqual(m2 Mat3) bool {
	for i := range m1 {
		if !FloatEqual(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat3x4) ApproxEqual(m2 Mat3x4) bool {
	for i := range m1 {
		if !FloatEqual(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat4x2) ApproxEqual(m2 Mat4x2) bool {
	for i := range m1 {
		if !FloatEqual(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat4x3) ApproxEqual(m2 Mat4x3) bool {
	for i := range m1 {
		if !FloatEqual(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat4) ApproxEqual(m2 Mat4) bool {
	for i := range m1 {
		if !FloatEqual(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat2) ApproxEqualThreshold(m2 Mat2, threshold float32) bool {
	for i := range m1 {
		if !FloatEqualThreshold(m1[i], m2[i], threshold) {
			return false
		}
	}
	return true
}

func (m1 Mat2x3) ApproxEqualThreshold(m2 Mat2x3, threshold float32) bool {
	for i := range m1 {
		if !FloatEqualThreshold(m1[i], m2[i], threshold) {
			return false
		}
	}
	return true
}

func (m1 Mat2x4) ApproxEqualThreshold(m2 Mat2x4, threshold float32) bool {
	for i := range m1 {
		if !FloatEqualThreshold(m1[i], m2[i], threshold) {
			return false
		}
	}
	return true
}

func (m1 Mat3x2) ApproxEqualThreshold(m2 Mat3x2, threshold float32) bool {
	for i := range m1 {
		if !FloatEqualThreshold(m1[i], m2[i], threshold) {
			return false
		}
	}
	return true
}

func (m1 Mat3) ApproxEqualThreshold(m2 Mat3, threshold float32) bool {
	for i := range m1 {
		if !FloatEqualThreshold(m1[i], m2[i], threshold) {
			return false
		}
	}
	return true
}

func (m1 Mat3x4) ApproxEqualThreshold(m2 Mat3x4, threshold float32) bool {
	for i := range m1 {
		if !FloatEqualThreshold(m1[i], m2[i], threshold) {
			return false
		}
	}
	return true
}

func (m1 Mat4x2) ApproxEqualThreshold(m2 Mat4x2, threshold float32) bool {
	for i := range m1 {
		if !FloatEqualThreshold(m1[i], m2[i], threshold) {
			return false
		}
	}
	return true
}

func (m1 Mat4x3) ApproxEqualThreshold(m2 Mat4x3, threshold float32) bool {
	for i := range m1 {
		if !FloatEqualThreshold(m1[i], m2[i], threshold) {
			return false
		}
	}
	return true
}

func (m1 Mat4) ApproxEqualThreshold(m2 Mat4, threshold float32) bool {
	for i := range m1 {
		if !FloatEqualThreshold(m1[i], m2[i], threshold) {
			return false
		}
	}
	return true
}

func (m1 Mat2) ApproxFuncEqual(m2 Mat2, eq func(float32, float32) bool) bool {
	for i := range m1 {
		if !eq(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat2x3) ApproxFuncEqual(m2 Mat2x3, eq func(float32, float32) bool) bool {
	for i := range m1 {
		if !eq(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat2x4) ApproxFuncEqual(m2 Mat2x4, eq func(float32, float32) bool) bool {
	for i := range m1 {
		if !eq(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat3x2) ApproxFuncEqual(m2 Mat3x2, eq func(float32, float32) bool) bool {
	for i := range m1 {
		if !eq(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat3) ApproxFuncEqual(m2 Mat3, eq func(float32, float32) bool) bool {
	for i := range m1 {
		if !eq(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat3x4) ApproxFuncEqual(m2 Mat3x4, eq func(float32, float32) bool) bool {
	for i := range m1 {
		if !eq(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat4x2) ApproxFuncEqual(m2 Mat4x2, eq func(float32, float32) bool) bool {
	for i := range m1 {
		if !eq(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat4x3) ApproxFuncEqual(m2 Mat4x3, eq func(float32, float32) bool) bool {
	for i := range m1 {
		if !eq(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

func (m1 Mat4) ApproxFuncEqual(m2 Mat4, eq func(float32, float32) bool) bool {
	for i := range m1 {
		if !eq(m1[i], m2[i]) {
			return false
		}
	}
	return true
}
