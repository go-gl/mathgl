package mathgl

import (
//"math"
)

type Matrix2f [4]float32
type Matrix2x3f [6]float32
type Matrix2x4f [8]float32
type Matrix3x2f [6]float32
type Matrix3f [9]float32
type Matrix3x4f [12]float32
type Matrix4x2f [8]float32
type Matrix4x3f [12]float32
type Matrix4f [16]float32

func (m1 Matrix2f) Add(m2 Matrix2f) Matrix2f {
	return Matrix2f{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3]}
}

func (m1 Matrix2x3f) Add(m2 Matrix2x3f) Matrix2x3f {
	return Matrix2x3f{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5]}
}

func (m1 Matrix2x4f) Add(m2 Matrix2x4f) Matrix2x4f {
	return Matrix2x4f{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7]}
}

func (m1 Matrix3x2f) Add(m2 Matrix3x2f) Matrix3x2f {
	return Matrix3x2f{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5]}
}

func (m1 Matrix3f) Add(m2 Matrix3f) Matrix3f {
	return Matrix3f{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7], m1[8] + m2[8]}
}

func (m1 Matrix3x4f) Add(m2 Matrix3x4f) Matrix3x4f {
	return Matrix3x4f{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7], m1[8] + m2[8], m1[9] + m2[9], m1[10] + m2[10], m1[11] + m2[11]}
}

func (m1 Matrix4x2f) Add(m2 Matrix4x2f) Matrix4x2f {
	return Matrix4x2f{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7]}
}

func (m1 Matrix4x3f) Add(m2 Matrix4x3f) Matrix4x3f {
	return Matrix4x3f{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7], m1[8] + m2[8], m1[9] + m2[9], m1[10] + m2[10], m1[11] + m2[11]}
}

func (m1 Matrix4f) Add(m2 Matrix4f) Matrix4f {
	return Matrix4f{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7], m1[8] + m2[8], m1[9] + m2[9], m1[10] + m2[10], m1[11] + m2[11], m1[12] + m2[12], m1[13] + m2[13], m1[14] + m2[14], m1[15] + m2[15]}
}

func (m1 Matrix2f) Sub(m2 Matrix2f) Matrix2f {
	return Matrix2f{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3]}
}

func (m1 Matrix2x3f) Sub(m2 Matrix2x3f) Matrix2x3f {
	return Matrix2x3f{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5]}
}

func (m1 Matrix2x4f) Sub(m2 Matrix2x4f) Matrix2x4f {
	return Matrix2x4f{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7]}
}

func (m1 Matrix3x2f) Sub(m2 Matrix3x2f) Matrix3x2f {
	return Matrix3x2f{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5]}
}

func (m1 Matrix3f) Sub(m2 Matrix3f) Matrix3f {
	return Matrix3f{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7], m1[8] - m2[8]}
}

func (m1 Matrix3x4f) Sub(m2 Matrix3x4f) Matrix3x4f {
	return Matrix3x4f{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7], m1[8] - m2[8], m1[9] - m2[9], m1[10] - m2[10], m1[11] - m2[11]}
}

func (m1 Matrix4x2f) Sub(m2 Matrix4x2f) Matrix4x2f {
	return Matrix4x2f{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7]}
}

func (m1 Matrix4x3f) Sub(m2 Matrix4x3f) Matrix4x3f {
	return Matrix4x3f{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7], m1[8] - m2[8], m1[9] - m2[9], m1[10] - m2[10], m1[11] - m2[11]}
}

func (m1 Matrix4f) Sub(m2 Matrix4f) Matrix4f {
	return Matrix4f{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7], m1[8] - m2[8], m1[9] - m2[9], m1[10] - m2[10], m1[11] - m2[11], m1[12] - m2[12], m1[13] - m2[13], m1[14] - m2[14], m1[15] - m2[15]}
}

func (m1 Matrix2f) Mul(c float32) Matrix2f {
	return Matrix2f{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c}
}

func (m1 Matrix2x3f) Mul(c float32) Matrix2x3f {
	return Matrix2x3f{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c}
}

func (m1 Matrix2x4f) Mul(c float32) Matrix2x4f {
	return Matrix2x4f{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c}
}

func (m1 Matrix3x2f) Mul(c float32) Matrix3x2f {
	return Matrix3x2f{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c}
}

func (m1 Matrix3f) Mul(c float32) Matrix3f {
	return Matrix3f{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c, m1[8] * c}
}

func (m1 Matrix3x4f) Mul(c float32) Matrix3x4f {
	return Matrix3x4f{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c, m1[8] * c, m1[9] * c, m1[10] * c, m1[11] * c}
}

func (m1 Matrix4x2f) Mul(c float32) Matrix4x2f {
	return Matrix4x2f{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c}
}

func (m1 Matrix4x3f) Mul(c float32) Matrix4x3f {
	return Matrix4x3f{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c, m1[8] * c, m1[9] * c, m1[10] * c, m1[11] * c}
}

func (m1 Matrix4f) Mul(c float32) Matrix4f {
	return Matrix4f{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c, m1[8] * c, m1[9] * c, m1[10] * c, m1[11] * c, m1[12] * c, m1[13] * c, m1[14] * c, m1[15] * c}
}

func (m1 Matrix2f) Mul2x1f(m2 Vec2f) Vec2f {
	return Vec2f{m1[0]*m2[0] + m1[2]*m2[1], m1[1]*m2[0] + m1[3]*m2[1]}
}

func (m1 Matrix2f) Mul2f(m2 Matrix2f) Matrix2f {
	return Matrix2f{m1[0]*m2[0] + m1[2]*m2[1], m1[1]*m2[0] + m1[3]*m2[1], m1[0]*m2[2] + m1[2]*m2[3], m1[1]*m2[2] + m1[3]*m2[3]}
}

func (m1 Matrix2f) Mul2x3f(m2 Matrix2x3f) Matrix2x3f {
	return Matrix2x3f{m1[0]*m2[0] + m1[2]*m2[1], m1[1]*m2[0] + m1[3]*m2[1], m1[0]*m2[2] + m1[2]*m2[3], m1[1]*m2[2] + m1[3]*m2[3], m1[0]*m2[4] + m1[2]*m2[5], m1[1]*m2[4] + m1[3]*m2[5]}
}

func (m1 Matrix2f) Mul2x4f(m2 Matrix2x4f) Matrix2x4f {
	return Matrix2x4f{m1[0]*m2[0] + m1[2]*m2[1], m1[1]*m2[0] + m1[3]*m2[1], m1[0]*m2[2] + m1[2]*m2[3], m1[1]*m2[2] + m1[3]*m2[3], m1[0]*m2[4] + m1[2]*m2[5], m1[1]*m2[4] + m1[3]*m2[5], m1[0]*m2[6] + m1[2]*m2[7], m1[1]*m2[6] + m1[3]*m2[7]}
}

func (m1 Matrix2x3f) Mul3x1f(m2 Vec3f) Vec2f {
	return Vec2f{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2]}
}

func (m1 Matrix2x3f) Mul3x2f(m2 Matrix3x2f) Matrix2f {
	return Matrix2f{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2], m1[0]*m2[3] + m1[2]*m2[4] + m1[4]*m2[5], m1[1]*m2[3] + m1[3]*m2[4] + m1[5]*m2[5]}
}

func (m1 Matrix2x3f) Mul3f(m2 Matrix3f) Matrix2x3f {
	return Matrix2x3f{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2], m1[0]*m2[3] + m1[2]*m2[4] + m1[4]*m2[5], m1[1]*m2[3] + m1[3]*m2[4] + m1[5]*m2[5], m1[0]*m2[6] + m1[2]*m2[7] + m1[4]*m2[8], m1[1]*m2[6] + m1[3]*m2[7] + m1[5]*m2[8]}
}

func (m1 Matrix2x3f) Mul3x4f(m2 Matrix3x4f) Matrix2x4f {
	return Matrix2x4f{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2], m1[0]*m2[3] + m1[2]*m2[4] + m1[4]*m2[5], m1[1]*m2[3] + m1[3]*m2[4] + m1[5]*m2[5], m1[0]*m2[6] + m1[2]*m2[7] + m1[4]*m2[8], m1[1]*m2[6] + m1[3]*m2[7] + m1[5]*m2[8], m1[0]*m2[9] + m1[2]*m2[10] + m1[4]*m2[11], m1[1]*m2[9] + m1[3]*m2[10] + m1[5]*m2[11]}
}

func (m1 Matrix2x4f) Mul4x1f(m2 Vec4f) Vec2f {
	return Vec2f{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2] + m1[6]*m2[3], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2] + m1[7]*m2[3]}
}

func (m1 Matrix2x4f) Mul4x2f(m2 Matrix4x2f) Matrix2f {
	return Matrix2f{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2] + m1[6]*m2[3], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[2]*m2[5] + m1[4]*m2[6] + m1[6]*m2[7], m1[1]*m2[4] + m1[3]*m2[5] + m1[5]*m2[6] + m1[7]*m2[7]}
}

func (m1 Matrix2x4f) Mul4x3f(m2 Matrix4x3f) Matrix2x3f {
	return Matrix2x3f{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2] + m1[6]*m2[3], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[2]*m2[5] + m1[4]*m2[6] + m1[6]*m2[7], m1[1]*m2[4] + m1[3]*m2[5] + m1[5]*m2[6] + m1[7]*m2[7], m1[0]*m2[8] + m1[2]*m2[9] + m1[4]*m2[10] + m1[6]*m2[11], m1[1]*m2[8] + m1[3]*m2[9] + m1[5]*m2[10] + m1[7]*m2[11]}
}

func (m1 Matrix2x4f) Mul4f(m2 Matrix4f) Matrix2x4f {
	return Matrix2x4f{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2] + m1[6]*m2[3], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[2]*m2[5] + m1[4]*m2[6] + m1[6]*m2[7], m1[1]*m2[4] + m1[3]*m2[5] + m1[5]*m2[6] + m1[7]*m2[7], m1[0]*m2[8] + m1[2]*m2[9] + m1[4]*m2[10] + m1[6]*m2[11], m1[1]*m2[8] + m1[3]*m2[9] + m1[5]*m2[10] + m1[7]*m2[11], m1[0]*m2[12] + m1[2]*m2[13] + m1[4]*m2[14] + m1[6]*m2[15], m1[1]*m2[12] + m1[3]*m2[13] + m1[5]*m2[14] + m1[7]*m2[15]}
}

func (m1 Matrix3x2f) Mul2x1f(m2 Vec2f) Vec3f {
	return Vec3f{m1[0]*m2[0] + m1[3]*m2[1], m1[1]*m2[0] + m1[4]*m2[1], m1[2]*m2[0] + m1[5]*m2[1]}
}

func (m1 Matrix3x2f) Mul2f(m2 Matrix2f) Matrix3x2f {
	return Matrix3x2f{m1[0]*m2[0] + m1[3]*m2[1], m1[1]*m2[0] + m1[4]*m2[1], m1[2]*m2[0] + m1[5]*m2[1], m1[0]*m2[2] + m1[3]*m2[3], m1[1]*m2[2] + m1[4]*m2[3], m1[2]*m2[2] + m1[5]*m2[3]}
}

func (m1 Matrix3x2f) Mul2x3f(m2 Matrix2x3f) Matrix3f {
	return Matrix3f{m1[0]*m2[0] + m1[3]*m2[1], m1[1]*m2[0] + m1[4]*m2[1], m1[2]*m2[0] + m1[5]*m2[1], m1[0]*m2[2] + m1[3]*m2[3], m1[1]*m2[2] + m1[4]*m2[3], m1[2]*m2[2] + m1[5]*m2[3], m1[0]*m2[4] + m1[3]*m2[5], m1[1]*m2[4] + m1[4]*m2[5], m1[2]*m2[4] + m1[5]*m2[5]}
}

func (m1 Matrix3x2f) Mul2x4f(m2 Matrix2x4f) Matrix3x4f {
	return Matrix3x4f{m1[0]*m2[0] + m1[3]*m2[1], m1[1]*m2[0] + m1[4]*m2[1], m1[2]*m2[0] + m1[5]*m2[1], m1[0]*m2[2] + m1[3]*m2[3], m1[1]*m2[2] + m1[4]*m2[3], m1[2]*m2[2] + m1[5]*m2[3], m1[0]*m2[4] + m1[3]*m2[5], m1[1]*m2[4] + m1[4]*m2[5], m1[2]*m2[4] + m1[5]*m2[5], m1[0]*m2[6] + m1[3]*m2[7], m1[1]*m2[6] + m1[4]*m2[7], m1[2]*m2[6] + m1[5]*m2[7]}
}

func (m1 Matrix3f) Mul3x1f(m2 Vec3f) Vec3f {
	return Vec3f{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2]}
}

func (m1 Matrix3f) Mul3x2f(m2 Matrix3x2f) Matrix3x2f {
	return Matrix3x2f{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2], m1[0]*m2[3] + m1[3]*m2[4] + m1[6]*m2[5], m1[1]*m2[3] + m1[4]*m2[4] + m1[7]*m2[5], m1[2]*m2[3] + m1[5]*m2[4] + m1[8]*m2[5]}
}

func (m1 Matrix3f) Mul3f(m2 Matrix3f) Matrix3f {
	return Matrix3f{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2], m1[0]*m2[3] + m1[3]*m2[4] + m1[6]*m2[5], m1[1]*m2[3] + m1[4]*m2[4] + m1[7]*m2[5], m1[2]*m2[3] + m1[5]*m2[4] + m1[8]*m2[5], m1[0]*m2[6] + m1[3]*m2[7] + m1[6]*m2[8], m1[1]*m2[6] + m1[4]*m2[7] + m1[7]*m2[8], m1[2]*m2[6] + m1[5]*m2[7] + m1[8]*m2[8]}
}

func (m1 Matrix3f) Mul3x4f(m2 Matrix3x4f) Matrix3x4f {
	return Matrix3x4f{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2], m1[0]*m2[3] + m1[3]*m2[4] + m1[6]*m2[5], m1[1]*m2[3] + m1[4]*m2[4] + m1[7]*m2[5], m1[2]*m2[3] + m1[5]*m2[4] + m1[8]*m2[5], m1[0]*m2[6] + m1[3]*m2[7] + m1[6]*m2[8], m1[1]*m2[6] + m1[4]*m2[7] + m1[7]*m2[8], m1[2]*m2[6] + m1[5]*m2[7] + m1[8]*m2[8], m1[0]*m2[9] + m1[3]*m2[10] + m1[6]*m2[11], m1[1]*m2[9] + m1[4]*m2[10] + m1[7]*m2[11], m1[2]*m2[9] + m1[5]*m2[10] + m1[8]*m2[11]}
}

func (m1 Matrix3x4f) Mul4x1f(m2 Vec4f) Vec3f {
	return Vec3f{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2] + m1[9]*m2[3], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2] + m1[10]*m2[3], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2] + m1[11]*m2[3]}
}

func (m1 Matrix3x4f) Mul4x2f(m2 Matrix4x2f) Matrix3x2f {
	return Matrix3x2f{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2] + m1[9]*m2[3], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2] + m1[10]*m2[3], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2] + m1[11]*m2[3], m1[0]*m2[4] + m1[3]*m2[5] + m1[6]*m2[6] + m1[9]*m2[7], m1[1]*m2[4] + m1[4]*m2[5] + m1[7]*m2[6] + m1[10]*m2[7], m1[2]*m2[4] + m1[5]*m2[5] + m1[8]*m2[6] + m1[11]*m2[7]}
}

func (m1 Matrix3x4f) Mul4x3f(m2 Matrix4x3f) Matrix3f {
	return Matrix3f{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2] + m1[9]*m2[3], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2] + m1[10]*m2[3], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2] + m1[11]*m2[3], m1[0]*m2[4] + m1[3]*m2[5] + m1[6]*m2[6] + m1[9]*m2[7], m1[1]*m2[4] + m1[4]*m2[5] + m1[7]*m2[6] + m1[10]*m2[7], m1[2]*m2[4] + m1[5]*m2[5] + m1[8]*m2[6] + m1[11]*m2[7], m1[0]*m2[8] + m1[3]*m2[9] + m1[6]*m2[10] + m1[9]*m2[11], m1[1]*m2[8] + m1[4]*m2[9] + m1[7]*m2[10] + m1[10]*m2[11], m1[2]*m2[8] + m1[5]*m2[9] + m1[8]*m2[10] + m1[11]*m2[11]}
}

func (m1 Matrix3x4f) Mul4f(m2 Matrix4f) Matrix3x4f {
	return Matrix3x4f{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2] + m1[9]*m2[3], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2] + m1[10]*m2[3], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2] + m1[11]*m2[3], m1[0]*m2[4] + m1[3]*m2[5] + m1[6]*m2[6] + m1[9]*m2[7], m1[1]*m2[4] + m1[4]*m2[5] + m1[7]*m2[6] + m1[10]*m2[7], m1[2]*m2[4] + m1[5]*m2[5] + m1[8]*m2[6] + m1[11]*m2[7], m1[0]*m2[8] + m1[3]*m2[9] + m1[6]*m2[10] + m1[9]*m2[11], m1[1]*m2[8] + m1[4]*m2[9] + m1[7]*m2[10] + m1[10]*m2[11], m1[2]*m2[8] + m1[5]*m2[9] + m1[8]*m2[10] + m1[11]*m2[11], m1[0]*m2[12] + m1[3]*m2[13] + m1[6]*m2[14] + m1[9]*m2[15], m1[1]*m2[12] + m1[4]*m2[13] + m1[7]*m2[14] + m1[10]*m2[15], m1[2]*m2[12] + m1[5]*m2[13] + m1[8]*m2[14] + m1[11]*m2[15]}
}

func (m1 Matrix4x2f) Mul2x1f(m2 Vec2f) Vec4f {
	return Vec4f{m1[0]*m2[0] + m1[4]*m2[1], m1[1]*m2[0] + m1[5]*m2[1], m1[2]*m2[0] + m1[6]*m2[1], m1[3]*m2[0] + m1[7]*m2[1]}
}

func (m1 Matrix4x2f) Mul2f(m2 Matrix2f) Matrix4x2f {
	return Matrix4x2f{m1[0]*m2[0] + m1[4]*m2[1], m1[1]*m2[0] + m1[5]*m2[1], m1[2]*m2[0] + m1[6]*m2[1], m1[3]*m2[0] + m1[7]*m2[1], m1[0]*m2[2] + m1[4]*m2[3], m1[1]*m2[2] + m1[5]*m2[3], m1[2]*m2[2] + m1[6]*m2[3], m1[3]*m2[2] + m1[7]*m2[3]}
}

func (m1 Matrix4x2f) Mul2x3f(m2 Matrix2x3f) Matrix4x3f {
	return Matrix4x3f{m1[0]*m2[0] + m1[4]*m2[1], m1[1]*m2[0] + m1[5]*m2[1], m1[2]*m2[0] + m1[6]*m2[1], m1[3]*m2[0] + m1[7]*m2[1], m1[0]*m2[2] + m1[4]*m2[3], m1[1]*m2[2] + m1[5]*m2[3], m1[2]*m2[2] + m1[6]*m2[3], m1[3]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[4]*m2[5], m1[1]*m2[4] + m1[5]*m2[5], m1[2]*m2[4] + m1[6]*m2[5], m1[3]*m2[4] + m1[7]*m2[5]}
}

func (m1 Matrix4x2f) Mul2x4f(m2 Matrix2x4f) Matrix4f {
	return Matrix4f{m1[0]*m2[0] + m1[4]*m2[1], m1[1]*m2[0] + m1[5]*m2[1], m1[2]*m2[0] + m1[6]*m2[1], m1[3]*m2[0] + m1[7]*m2[1], m1[0]*m2[2] + m1[4]*m2[3], m1[1]*m2[2] + m1[5]*m2[3], m1[2]*m2[2] + m1[6]*m2[3], m1[3]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[4]*m2[5], m1[1]*m2[4] + m1[5]*m2[5], m1[2]*m2[4] + m1[6]*m2[5], m1[3]*m2[4] + m1[7]*m2[5], m1[0]*m2[6] + m1[4]*m2[7], m1[1]*m2[6] + m1[5]*m2[7], m1[2]*m2[6] + m1[6]*m2[7], m1[3]*m2[6] + m1[7]*m2[7]}
}

func (m1 Matrix4x3f) Mul3x1f(m2 Vec3f) Vec4f {
	return Vec4f{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2]}
}

func (m1 Matrix4x3f) Mul3x2f(m2 Matrix3x2f) Matrix4x2f {
	return Matrix4x2f{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2], m1[0]*m2[3] + m1[4]*m2[4] + m1[8]*m2[5], m1[1]*m2[3] + m1[5]*m2[4] + m1[9]*m2[5], m1[2]*m2[3] + m1[6]*m2[4] + m1[10]*m2[5], m1[3]*m2[3] + m1[7]*m2[4] + m1[11]*m2[5]}
}

func (m1 Matrix4x3f) Mul3f(m2 Matrix3f) Matrix4x3f {
	return Matrix4x3f{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2], m1[0]*m2[3] + m1[4]*m2[4] + m1[8]*m2[5], m1[1]*m2[3] + m1[5]*m2[4] + m1[9]*m2[5], m1[2]*m2[3] + m1[6]*m2[4] + m1[10]*m2[5], m1[3]*m2[3] + m1[7]*m2[4] + m1[11]*m2[5], m1[0]*m2[6] + m1[4]*m2[7] + m1[8]*m2[8], m1[1]*m2[6] + m1[5]*m2[7] + m1[9]*m2[8], m1[2]*m2[6] + m1[6]*m2[7] + m1[10]*m2[8], m1[3]*m2[6] + m1[7]*m2[7] + m1[11]*m2[8]}
}

func (m1 Matrix4x3f) Mul3x4f(m2 Matrix3x4f) Matrix4f {
	return Matrix4f{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2], m1[0]*m2[3] + m1[4]*m2[4] + m1[8]*m2[5], m1[1]*m2[3] + m1[5]*m2[4] + m1[9]*m2[5], m1[2]*m2[3] + m1[6]*m2[4] + m1[10]*m2[5], m1[3]*m2[3] + m1[7]*m2[4] + m1[11]*m2[5], m1[0]*m2[6] + m1[4]*m2[7] + m1[8]*m2[8], m1[1]*m2[6] + m1[5]*m2[7] + m1[9]*m2[8], m1[2]*m2[6] + m1[6]*m2[7] + m1[10]*m2[8], m1[3]*m2[6] + m1[7]*m2[7] + m1[11]*m2[8], m1[0]*m2[9] + m1[4]*m2[10] + m1[8]*m2[11], m1[1]*m2[9] + m1[5]*m2[10] + m1[9]*m2[11], m1[2]*m2[9] + m1[6]*m2[10] + m1[10]*m2[11], m1[3]*m2[9] + m1[7]*m2[10] + m1[11]*m2[11]}
}

func (m1 Matrix4f) Mul4x1f(m2 Vec4f) Vec4f {
	return Vec4f{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2] + m1[12]*m2[3], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2] + m1[13]*m2[3], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2] + m1[14]*m2[3], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2] + m1[15]*m2[3]}
}

func (m1 Matrix4f) Mul4x2f(m2 Matrix4x2f) Matrix4x2f {
	return Matrix4x2f{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2] + m1[12]*m2[3], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2] + m1[13]*m2[3], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2] + m1[14]*m2[3], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2] + m1[15]*m2[3], m1[0]*m2[4] + m1[4]*m2[5] + m1[8]*m2[6] + m1[12]*m2[7], m1[1]*m2[4] + m1[5]*m2[5] + m1[9]*m2[6] + m1[13]*m2[7], m1[2]*m2[4] + m1[6]*m2[5] + m1[10]*m2[6] + m1[14]*m2[7], m1[3]*m2[4] + m1[7]*m2[5] + m1[11]*m2[6] + m1[15]*m2[7]}
}

func (m1 Matrix4f) Mul4x3f(m2 Matrix4x3f) Matrix4x3f {
	return Matrix4x3f{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2] + m1[12]*m2[3], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2] + m1[13]*m2[3], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2] + m1[14]*m2[3], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2] + m1[15]*m2[3], m1[0]*m2[4] + m1[4]*m2[5] + m1[8]*m2[6] + m1[12]*m2[7], m1[1]*m2[4] + m1[5]*m2[5] + m1[9]*m2[6] + m1[13]*m2[7], m1[2]*m2[4] + m1[6]*m2[5] + m1[10]*m2[6] + m1[14]*m2[7], m1[3]*m2[4] + m1[7]*m2[5] + m1[11]*m2[6] + m1[15]*m2[7], m1[0]*m2[8] + m1[4]*m2[9] + m1[8]*m2[10] + m1[12]*m2[11], m1[1]*m2[8] + m1[5]*m2[9] + m1[9]*m2[10] + m1[13]*m2[11], m1[2]*m2[8] + m1[6]*m2[9] + m1[10]*m2[10] + m1[14]*m2[11], m1[3]*m2[8] + m1[7]*m2[9] + m1[11]*m2[10] + m1[15]*m2[11]}
}

func (m1 Matrix4f) Mul4f(m2 Matrix4f) Matrix4f {
	return Matrix4f{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2] + m1[12]*m2[3], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2] + m1[13]*m2[3], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2] + m1[14]*m2[3], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2] + m1[15]*m2[3], m1[0]*m2[4] + m1[4]*m2[5] + m1[8]*m2[6] + m1[12]*m2[7], m1[1]*m2[4] + m1[5]*m2[5] + m1[9]*m2[6] + m1[13]*m2[7], m1[2]*m2[4] + m1[6]*m2[5] + m1[10]*m2[6] + m1[14]*m2[7], m1[3]*m2[4] + m1[7]*m2[5] + m1[11]*m2[6] + m1[15]*m2[7], m1[0]*m2[8] + m1[4]*m2[9] + m1[8]*m2[10] + m1[12]*m2[11], m1[1]*m2[8] + m1[5]*m2[9] + m1[9]*m2[10] + m1[13]*m2[11], m1[2]*m2[8] + m1[6]*m2[9] + m1[10]*m2[10] + m1[14]*m2[11], m1[3]*m2[8] + m1[7]*m2[9] + m1[11]*m2[10] + m1[15]*m2[11], m1[0]*m2[12] + m1[4]*m2[13] + m1[8]*m2[14] + m1[12]*m2[15], m1[1]*m2[12] + m1[5]*m2[13] + m1[9]*m2[14] + m1[13]*m2[15], m1[2]*m2[12] + m1[6]*m2[13] + m1[10]*m2[14] + m1[14]*m2[15], m1[3]*m2[12] + m1[7]*m2[13] + m1[11]*m2[14] + m1[15]*m2[15]}
}
