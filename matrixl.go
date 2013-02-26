package mathgl

import (
//"math"
)

type Matrix2l [4]int64
type Matrix2x3l [6]int64
type Matrix2x4l [8]int64
type Matrix3x2l [6]int64
type Matrix3l [9]int64
type Matrix3x4l [12]int64
type Matrix4x2l [8]int64
type Matrix4x3l [12]int64
type Matrix4l [16]int64

func (m1 Matrix2l) Add(m2 Matrix2l) Matrix2l {
	return Matrix2l{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3]}
}

func (m1 Matrix2x3l) Add(m2 Matrix2x3l) Matrix2x3l {
	return Matrix2x3l{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5]}
}

func (m1 Matrix2x4l) Add(m2 Matrix2x4l) Matrix2x4l {
	return Matrix2x4l{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7]}
}

func (m1 Matrix3x2l) Add(m2 Matrix3x2l) Matrix3x2l {
	return Matrix3x2l{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5]}
}

func (m1 Matrix3l) Add(m2 Matrix3l) Matrix3l {
	return Matrix3l{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7], m1[8] + m2[8]}
}

func (m1 Matrix3x4l) Add(m2 Matrix3x4l) Matrix3x4l {
	return Matrix3x4l{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7], m1[8] + m2[8], m1[9] + m2[9], m1[10] + m2[10], m1[11] + m2[11]}
}

func (m1 Matrix4x2l) Add(m2 Matrix4x2l) Matrix4x2l {
	return Matrix4x2l{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7]}
}

func (m1 Matrix4x3l) Add(m2 Matrix4x3l) Matrix4x3l {
	return Matrix4x3l{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7], m1[8] + m2[8], m1[9] + m2[9], m1[10] + m2[10], m1[11] + m2[11]}
}

func (m1 Matrix4l) Add(m2 Matrix4l) Matrix4l {
	return Matrix4l{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7], m1[8] + m2[8], m1[9] + m2[9], m1[10] + m2[10], m1[11] + m2[11], m1[12] + m2[12], m1[13] + m2[13], m1[14] + m2[14], m1[15] + m2[15]}
}

func (m1 Matrix2l) Sub(m2 Matrix2l) Matrix2l {
	return Matrix2l{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3]}
}

func (m1 Matrix2x3l) Sub(m2 Matrix2x3l) Matrix2x3l {
	return Matrix2x3l{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5]}
}

func (m1 Matrix2x4l) Sub(m2 Matrix2x4l) Matrix2x4l {
	return Matrix2x4l{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7]}
}

func (m1 Matrix3x2l) Sub(m2 Matrix3x2l) Matrix3x2l {
	return Matrix3x2l{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5]}
}

func (m1 Matrix3l) Sub(m2 Matrix3l) Matrix3l {
	return Matrix3l{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7], m1[8] - m2[8]}
}

func (m1 Matrix3x4l) Sub(m2 Matrix3x4l) Matrix3x4l {
	return Matrix3x4l{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7], m1[8] - m2[8], m1[9] - m2[9], m1[10] - m2[10], m1[11] - m2[11]}
}

func (m1 Matrix4x2l) Sub(m2 Matrix4x2l) Matrix4x2l {
	return Matrix4x2l{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7]}
}

func (m1 Matrix4x3l) Sub(m2 Matrix4x3l) Matrix4x3l {
	return Matrix4x3l{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7], m1[8] - m2[8], m1[9] - m2[9], m1[10] - m2[10], m1[11] - m2[11]}
}

func (m1 Matrix4l) Sub(m2 Matrix4l) Matrix4l {
	return Matrix4l{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7], m1[8] - m2[8], m1[9] - m2[9], m1[10] - m2[10], m1[11] - m2[11], m1[12] - m2[12], m1[13] - m2[13], m1[14] - m2[14], m1[15] - m2[15]}
}

func (m1 Matrix2l) Mul(c int64) Matrix2l {
	return Matrix2l{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c}
}

func (m1 Matrix2x3l) Mul(c int64) Matrix2x3l {
	return Matrix2x3l{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c}
}

func (m1 Matrix2x4l) Mul(c int64) Matrix2x4l {
	return Matrix2x4l{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c}
}

func (m1 Matrix3x2l) Mul(c int64) Matrix3x2l {
	return Matrix3x2l{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c}
}

func (m1 Matrix3l) Mul(c int64) Matrix3l {
	return Matrix3l{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c, m1[8] * c}
}

func (m1 Matrix3x4l) Mul(c int64) Matrix3x4l {
	return Matrix3x4l{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c, m1[8] * c, m1[9] * c, m1[10] * c, m1[11] * c}
}

func (m1 Matrix4x2l) Mul(c int64) Matrix4x2l {
	return Matrix4x2l{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c}
}

func (m1 Matrix4x3l) Mul(c int64) Matrix4x3l {
	return Matrix4x3l{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c, m1[8] * c, m1[9] * c, m1[10] * c, m1[11] * c}
}

func (m1 Matrix4l) Mul(c int64) Matrix4l {
	return Matrix4l{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c, m1[8] * c, m1[9] * c, m1[10] * c, m1[11] * c, m1[12] * c, m1[13] * c, m1[14] * c, m1[15] * c}
}

func (m1 Matrix2l) Mul2x1f(m2 Vec2l) Vec2l {
	return Vec2l{m1[0]*m2[0] + m1[2]*m2[1], m1[1]*m2[0] + m1[3]*m2[1]}
}

func (m1 Matrix2l) Mul2f(m2 Matrix2l) Matrix2l {
	return Matrix2l{m1[0]*m2[0] + m1[2]*m2[1], m1[1]*m2[0] + m1[3]*m2[1], m1[0]*m2[2] + m1[2]*m2[3], m1[1]*m2[2] + m1[3]*m2[3]}
}

func (m1 Matrix2l) Mul2x3f(m2 Matrix2x3l) Matrix2x3l {
	return Matrix2x3l{m1[0]*m2[0] + m1[2]*m2[1], m1[1]*m2[0] + m1[3]*m2[1], m1[0]*m2[2] + m1[2]*m2[3], m1[1]*m2[2] + m1[3]*m2[3], m1[0]*m2[4] + m1[2]*m2[5], m1[1]*m2[4] + m1[3]*m2[5]}
}

func (m1 Matrix2l) Mul2x4f(m2 Matrix2x4l) Matrix2x4l {
	return Matrix2x4l{m1[0]*m2[0] + m1[2]*m2[1], m1[1]*m2[0] + m1[3]*m2[1], m1[0]*m2[2] + m1[2]*m2[3], m1[1]*m2[2] + m1[3]*m2[3], m1[0]*m2[4] + m1[2]*m2[5], m1[1]*m2[4] + m1[3]*m2[5], m1[0]*m2[6] + m1[2]*m2[7], m1[1]*m2[6] + m1[3]*m2[7]}
}

func (m1 Matrix2x3l) Mul3x1f(m2 Vec3l) Vec2l {
	return Vec2l{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2]}
}

func (m1 Matrix2x3l) Mul3x2f(m2 Matrix3x2l) Matrix2l {
	return Matrix2l{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2], m1[0]*m2[3] + m1[2]*m2[4] + m1[4]*m2[5], m1[1]*m2[3] + m1[3]*m2[4] + m1[5]*m2[5]}
}

func (m1 Matrix2x3l) Mul3f(m2 Matrix3l) Matrix2x3l {
	return Matrix2x3l{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2], m1[0]*m2[3] + m1[2]*m2[4] + m1[4]*m2[5], m1[1]*m2[3] + m1[3]*m2[4] + m1[5]*m2[5], m1[0]*m2[6] + m1[2]*m2[7] + m1[4]*m2[8], m1[1]*m2[6] + m1[3]*m2[7] + m1[5]*m2[8]}
}

func (m1 Matrix2x3l) Mul3x4f(m2 Matrix3x4l) Matrix2x4l {
	return Matrix2x4l{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2], m1[0]*m2[3] + m1[2]*m2[4] + m1[4]*m2[5], m1[1]*m2[3] + m1[3]*m2[4] + m1[5]*m2[5], m1[0]*m2[6] + m1[2]*m2[7] + m1[4]*m2[8], m1[1]*m2[6] + m1[3]*m2[7] + m1[5]*m2[8], m1[0]*m2[9] + m1[2]*m2[10] + m1[4]*m2[11], m1[1]*m2[9] + m1[3]*m2[10] + m1[5]*m2[11]}
}

func (m1 Matrix2x4l) Mul4x1f(m2 Vec4l) Vec2l {
	return Vec2l{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2] + m1[6]*m2[3], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2] + m1[7]*m2[3]}
}

func (m1 Matrix2x4l) Mul4x2f(m2 Matrix4x2l) Matrix2l {
	return Matrix2l{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2] + m1[6]*m2[3], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[2]*m2[5] + m1[4]*m2[6] + m1[6]*m2[7], m1[1]*m2[4] + m1[3]*m2[5] + m1[5]*m2[6] + m1[7]*m2[7]}
}

func (m1 Matrix2x4l) Mul4x3f(m2 Matrix4x3l) Matrix2x3l {
	return Matrix2x3l{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2] + m1[6]*m2[3], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[2]*m2[5] + m1[4]*m2[6] + m1[6]*m2[7], m1[1]*m2[4] + m1[3]*m2[5] + m1[5]*m2[6] + m1[7]*m2[7], m1[0]*m2[8] + m1[2]*m2[9] + m1[4]*m2[10] + m1[6]*m2[11], m1[1]*m2[8] + m1[3]*m2[9] + m1[5]*m2[10] + m1[7]*m2[11]}
}

func (m1 Matrix2x4l) Mul4f(m2 Matrix4l) Matrix2x4l {
	return Matrix2x4l{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2] + m1[6]*m2[3], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[2]*m2[5] + m1[4]*m2[6] + m1[6]*m2[7], m1[1]*m2[4] + m1[3]*m2[5] + m1[5]*m2[6] + m1[7]*m2[7], m1[0]*m2[8] + m1[2]*m2[9] + m1[4]*m2[10] + m1[6]*m2[11], m1[1]*m2[8] + m1[3]*m2[9] + m1[5]*m2[10] + m1[7]*m2[11], m1[0]*m2[12] + m1[2]*m2[13] + m1[4]*m2[14] + m1[6]*m2[15], m1[1]*m2[12] + m1[3]*m2[13] + m1[5]*m2[14] + m1[7]*m2[15]}
}

func (m1 Matrix3x2l) Mul2x1f(m2 Vec2l) Vec3l {
	return Vec3l{m1[0]*m2[0] + m1[3]*m2[1], m1[1]*m2[0] + m1[4]*m2[1], m1[2]*m2[0] + m1[5]*m2[1]}
}

func (m1 Matrix3x2l) Mul2f(m2 Matrix2l) Matrix3x2l {
	return Matrix3x2l{m1[0]*m2[0] + m1[3]*m2[1], m1[1]*m2[0] + m1[4]*m2[1], m1[2]*m2[0] + m1[5]*m2[1], m1[0]*m2[2] + m1[3]*m2[3], m1[1]*m2[2] + m1[4]*m2[3], m1[2]*m2[2] + m1[5]*m2[3]}
}

func (m1 Matrix3x2l) Mul2x3f(m2 Matrix2x3l) Matrix3l {
	return Matrix3l{m1[0]*m2[0] + m1[3]*m2[1], m1[1]*m2[0] + m1[4]*m2[1], m1[2]*m2[0] + m1[5]*m2[1], m1[0]*m2[2] + m1[3]*m2[3], m1[1]*m2[2] + m1[4]*m2[3], m1[2]*m2[2] + m1[5]*m2[3], m1[0]*m2[4] + m1[3]*m2[5], m1[1]*m2[4] + m1[4]*m2[5], m1[2]*m2[4] + m1[5]*m2[5]}
}

func (m1 Matrix3x2l) Mul2x4f(m2 Matrix2x4l) Matrix3x4l {
	return Matrix3x4l{m1[0]*m2[0] + m1[3]*m2[1], m1[1]*m2[0] + m1[4]*m2[1], m1[2]*m2[0] + m1[5]*m2[1], m1[0]*m2[2] + m1[3]*m2[3], m1[1]*m2[2] + m1[4]*m2[3], m1[2]*m2[2] + m1[5]*m2[3], m1[0]*m2[4] + m1[3]*m2[5], m1[1]*m2[4] + m1[4]*m2[5], m1[2]*m2[4] + m1[5]*m2[5], m1[0]*m2[6] + m1[3]*m2[7], m1[1]*m2[6] + m1[4]*m2[7], m1[2]*m2[6] + m1[5]*m2[7]}
}

func (m1 Matrix3l) Mul3x1f(m2 Vec3l) Vec3l {
	return Vec3l{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2]}
}

func (m1 Matrix3l) Mul3x2f(m2 Matrix3x2l) Matrix3x2l {
	return Matrix3x2l{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2], m1[0]*m2[3] + m1[3]*m2[4] + m1[6]*m2[5], m1[1]*m2[3] + m1[4]*m2[4] + m1[7]*m2[5], m1[2]*m2[3] + m1[5]*m2[4] + m1[8]*m2[5]}
}

func (m1 Matrix3l) Mul3f(m2 Matrix3l) Matrix3l {
	return Matrix3l{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2], m1[0]*m2[3] + m1[3]*m2[4] + m1[6]*m2[5], m1[1]*m2[3] + m1[4]*m2[4] + m1[7]*m2[5], m1[2]*m2[3] + m1[5]*m2[4] + m1[8]*m2[5], m1[0]*m2[6] + m1[3]*m2[7] + m1[6]*m2[8], m1[1]*m2[6] + m1[4]*m2[7] + m1[7]*m2[8], m1[2]*m2[6] + m1[5]*m2[7] + m1[8]*m2[8]}
}

func (m1 Matrix3l) Mul3x4f(m2 Matrix3x4l) Matrix3x4l {
	return Matrix3x4l{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2], m1[0]*m2[3] + m1[3]*m2[4] + m1[6]*m2[5], m1[1]*m2[3] + m1[4]*m2[4] + m1[7]*m2[5], m1[2]*m2[3] + m1[5]*m2[4] + m1[8]*m2[5], m1[0]*m2[6] + m1[3]*m2[7] + m1[6]*m2[8], m1[1]*m2[6] + m1[4]*m2[7] + m1[7]*m2[8], m1[2]*m2[6] + m1[5]*m2[7] + m1[8]*m2[8], m1[0]*m2[9] + m1[3]*m2[10] + m1[6]*m2[11], m1[1]*m2[9] + m1[4]*m2[10] + m1[7]*m2[11], m1[2]*m2[9] + m1[5]*m2[10] + m1[8]*m2[11]}
}

func (m1 Matrix3x4l) Mul4x1f(m2 Vec4l) Vec3l {
	return Vec3l{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2] + m1[9]*m2[3], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2] + m1[10]*m2[3], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2] + m1[11]*m2[3]}
}

func (m1 Matrix3x4l) Mul4x2f(m2 Matrix4x2l) Matrix3x2l {
	return Matrix3x2l{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2] + m1[9]*m2[3], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2] + m1[10]*m2[3], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2] + m1[11]*m2[3], m1[0]*m2[4] + m1[3]*m2[5] + m1[6]*m2[6] + m1[9]*m2[7], m1[1]*m2[4] + m1[4]*m2[5] + m1[7]*m2[6] + m1[10]*m2[7], m1[2]*m2[4] + m1[5]*m2[5] + m1[8]*m2[6] + m1[11]*m2[7]}
}

func (m1 Matrix3x4l) Mul4x3f(m2 Matrix4x3l) Matrix3l {
	return Matrix3l{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2] + m1[9]*m2[3], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2] + m1[10]*m2[3], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2] + m1[11]*m2[3], m1[0]*m2[4] + m1[3]*m2[5] + m1[6]*m2[6] + m1[9]*m2[7], m1[1]*m2[4] + m1[4]*m2[5] + m1[7]*m2[6] + m1[10]*m2[7], m1[2]*m2[4] + m1[5]*m2[5] + m1[8]*m2[6] + m1[11]*m2[7], m1[0]*m2[8] + m1[3]*m2[9] + m1[6]*m2[10] + m1[9]*m2[11], m1[1]*m2[8] + m1[4]*m2[9] + m1[7]*m2[10] + m1[10]*m2[11], m1[2]*m2[8] + m1[5]*m2[9] + m1[8]*m2[10] + m1[11]*m2[11]}
}

func (m1 Matrix3x4l) Mul4f(m2 Matrix4l) Matrix3x4l {
	return Matrix3x4l{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2] + m1[9]*m2[3], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2] + m1[10]*m2[3], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2] + m1[11]*m2[3], m1[0]*m2[4] + m1[3]*m2[5] + m1[6]*m2[6] + m1[9]*m2[7], m1[1]*m2[4] + m1[4]*m2[5] + m1[7]*m2[6] + m1[10]*m2[7], m1[2]*m2[4] + m1[5]*m2[5] + m1[8]*m2[6] + m1[11]*m2[7], m1[0]*m2[8] + m1[3]*m2[9] + m1[6]*m2[10] + m1[9]*m2[11], m1[1]*m2[8] + m1[4]*m2[9] + m1[7]*m2[10] + m1[10]*m2[11], m1[2]*m2[8] + m1[5]*m2[9] + m1[8]*m2[10] + m1[11]*m2[11], m1[0]*m2[12] + m1[3]*m2[13] + m1[6]*m2[14] + m1[9]*m2[15], m1[1]*m2[12] + m1[4]*m2[13] + m1[7]*m2[14] + m1[10]*m2[15], m1[2]*m2[12] + m1[5]*m2[13] + m1[8]*m2[14] + m1[11]*m2[15]}
}

func (m1 Matrix4x2l) Mul2x1f(m2 Vec2l) Vec4l {
	return Vec4l{m1[0]*m2[0] + m1[4]*m2[1], m1[1]*m2[0] + m1[5]*m2[1], m1[2]*m2[0] + m1[6]*m2[1], m1[3]*m2[0] + m1[7]*m2[1]}
}

func (m1 Matrix4x2l) Mul2f(m2 Matrix2l) Matrix4x2l {
	return Matrix4x2l{m1[0]*m2[0] + m1[4]*m2[1], m1[1]*m2[0] + m1[5]*m2[1], m1[2]*m2[0] + m1[6]*m2[1], m1[3]*m2[0] + m1[7]*m2[1], m1[0]*m2[2] + m1[4]*m2[3], m1[1]*m2[2] + m1[5]*m2[3], m1[2]*m2[2] + m1[6]*m2[3], m1[3]*m2[2] + m1[7]*m2[3]}
}

func (m1 Matrix4x2l) Mul2x3f(m2 Matrix2x3l) Matrix4x3l {
	return Matrix4x3l{m1[0]*m2[0] + m1[4]*m2[1], m1[1]*m2[0] + m1[5]*m2[1], m1[2]*m2[0] + m1[6]*m2[1], m1[3]*m2[0] + m1[7]*m2[1], m1[0]*m2[2] + m1[4]*m2[3], m1[1]*m2[2] + m1[5]*m2[3], m1[2]*m2[2] + m1[6]*m2[3], m1[3]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[4]*m2[5], m1[1]*m2[4] + m1[5]*m2[5], m1[2]*m2[4] + m1[6]*m2[5], m1[3]*m2[4] + m1[7]*m2[5]}
}

func (m1 Matrix4x2l) Mul2x4f(m2 Matrix2x4l) Matrix4l {
	return Matrix4l{m1[0]*m2[0] + m1[4]*m2[1], m1[1]*m2[0] + m1[5]*m2[1], m1[2]*m2[0] + m1[6]*m2[1], m1[3]*m2[0] + m1[7]*m2[1], m1[0]*m2[2] + m1[4]*m2[3], m1[1]*m2[2] + m1[5]*m2[3], m1[2]*m2[2] + m1[6]*m2[3], m1[3]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[4]*m2[5], m1[1]*m2[4] + m1[5]*m2[5], m1[2]*m2[4] + m1[6]*m2[5], m1[3]*m2[4] + m1[7]*m2[5], m1[0]*m2[6] + m1[4]*m2[7], m1[1]*m2[6] + m1[5]*m2[7], m1[2]*m2[6] + m1[6]*m2[7], m1[3]*m2[6] + m1[7]*m2[7]}
}

func (m1 Matrix4x3l) Mul3x1f(m2 Vec3l) Vec4l {
	return Vec4l{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2]}
}

func (m1 Matrix4x3l) Mul3x2f(m2 Matrix3x2l) Matrix4x2l {
	return Matrix4x2l{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2], m1[0]*m2[3] + m1[4]*m2[4] + m1[8]*m2[5], m1[1]*m2[3] + m1[5]*m2[4] + m1[9]*m2[5], m1[2]*m2[3] + m1[6]*m2[4] + m1[10]*m2[5], m1[3]*m2[3] + m1[7]*m2[4] + m1[11]*m2[5]}
}

func (m1 Matrix4x3l) Mul3f(m2 Matrix3l) Matrix4x3l {
	return Matrix4x3l{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2], m1[0]*m2[3] + m1[4]*m2[4] + m1[8]*m2[5], m1[1]*m2[3] + m1[5]*m2[4] + m1[9]*m2[5], m1[2]*m2[3] + m1[6]*m2[4] + m1[10]*m2[5], m1[3]*m2[3] + m1[7]*m2[4] + m1[11]*m2[5], m1[0]*m2[6] + m1[4]*m2[7] + m1[8]*m2[8], m1[1]*m2[6] + m1[5]*m2[7] + m1[9]*m2[8], m1[2]*m2[6] + m1[6]*m2[7] + m1[10]*m2[8], m1[3]*m2[6] + m1[7]*m2[7] + m1[11]*m2[8]}
}

func (m1 Matrix4x3l) Mul3x4f(m2 Matrix3x4l) Matrix4l {
	return Matrix4l{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2], m1[0]*m2[3] + m1[4]*m2[4] + m1[8]*m2[5], m1[1]*m2[3] + m1[5]*m2[4] + m1[9]*m2[5], m1[2]*m2[3] + m1[6]*m2[4] + m1[10]*m2[5], m1[3]*m2[3] + m1[7]*m2[4] + m1[11]*m2[5], m1[0]*m2[6] + m1[4]*m2[7] + m1[8]*m2[8], m1[1]*m2[6] + m1[5]*m2[7] + m1[9]*m2[8], m1[2]*m2[6] + m1[6]*m2[7] + m1[10]*m2[8], m1[3]*m2[6] + m1[7]*m2[7] + m1[11]*m2[8], m1[0]*m2[9] + m1[4]*m2[10] + m1[8]*m2[11], m1[1]*m2[9] + m1[5]*m2[10] + m1[9]*m2[11], m1[2]*m2[9] + m1[6]*m2[10] + m1[10]*m2[11], m1[3]*m2[9] + m1[7]*m2[10] + m1[11]*m2[11]}
}

func (m1 Matrix4l) Mul4x1f(m2 Vec4l) Vec4l {
	return Vec4l{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2] + m1[12]*m2[3], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2] + m1[13]*m2[3], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2] + m1[14]*m2[3], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2] + m1[15]*m2[3]}
}

func (m1 Matrix4l) Mul4x2f(m2 Matrix4x2l) Matrix4x2l {
	return Matrix4x2l{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2] + m1[12]*m2[3], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2] + m1[13]*m2[3], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2] + m1[14]*m2[3], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2] + m1[15]*m2[3], m1[0]*m2[4] + m1[4]*m2[5] + m1[8]*m2[6] + m1[12]*m2[7], m1[1]*m2[4] + m1[5]*m2[5] + m1[9]*m2[6] + m1[13]*m2[7], m1[2]*m2[4] + m1[6]*m2[5] + m1[10]*m2[6] + m1[14]*m2[7], m1[3]*m2[4] + m1[7]*m2[5] + m1[11]*m2[6] + m1[15]*m2[7]}
}

func (m1 Matrix4l) Mul4x3f(m2 Matrix4x3l) Matrix4x3l {
	return Matrix4x3l{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2] + m1[12]*m2[3], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2] + m1[13]*m2[3], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2] + m1[14]*m2[3], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2] + m1[15]*m2[3], m1[0]*m2[4] + m1[4]*m2[5] + m1[8]*m2[6] + m1[12]*m2[7], m1[1]*m2[4] + m1[5]*m2[5] + m1[9]*m2[6] + m1[13]*m2[7], m1[2]*m2[4] + m1[6]*m2[5] + m1[10]*m2[6] + m1[14]*m2[7], m1[3]*m2[4] + m1[7]*m2[5] + m1[11]*m2[6] + m1[15]*m2[7], m1[0]*m2[8] + m1[4]*m2[9] + m1[8]*m2[10] + m1[12]*m2[11], m1[1]*m2[8] + m1[5]*m2[9] + m1[9]*m2[10] + m1[13]*m2[11], m1[2]*m2[8] + m1[6]*m2[9] + m1[10]*m2[10] + m1[14]*m2[11], m1[3]*m2[8] + m1[7]*m2[9] + m1[11]*m2[10] + m1[15]*m2[11]}
}

func (m1 Matrix4l) Mul4f(m2 Matrix4l) Matrix4l {
	return Matrix4l{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2] + m1[12]*m2[3], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2] + m1[13]*m2[3], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2] + m1[14]*m2[3], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2] + m1[15]*m2[3], m1[0]*m2[4] + m1[4]*m2[5] + m1[8]*m2[6] + m1[12]*m2[7], m1[1]*m2[4] + m1[5]*m2[5] + m1[9]*m2[6] + m1[13]*m2[7], m1[2]*m2[4] + m1[6]*m2[5] + m1[10]*m2[6] + m1[14]*m2[7], m1[3]*m2[4] + m1[7]*m2[5] + m1[11]*m2[6] + m1[15]*m2[7], m1[0]*m2[8] + m1[4]*m2[9] + m1[8]*m2[10] + m1[12]*m2[11], m1[1]*m2[8] + m1[5]*m2[9] + m1[9]*m2[10] + m1[13]*m2[11], m1[2]*m2[8] + m1[6]*m2[9] + m1[10]*m2[10] + m1[14]*m2[11], m1[3]*m2[8] + m1[7]*m2[9] + m1[11]*m2[10] + m1[15]*m2[11], m1[0]*m2[12] + m1[4]*m2[13] + m1[8]*m2[14] + m1[12]*m2[15], m1[1]*m2[12] + m1[5]*m2[13] + m1[9]*m2[14] + m1[13]*m2[15], m1[2]*m2[12] + m1[6]*m2[13] + m1[10]*m2[14] + m1[14]*m2[15], m1[3]*m2[12] + m1[7]*m2[13] + m1[11]*m2[14] + m1[15]*m2[15]}
}
