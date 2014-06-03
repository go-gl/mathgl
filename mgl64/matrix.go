// Copyright 2012 The go-gl/mathgl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mgl32

import (
	"math"
)

type Mat2 [4]float64
type Mat2x3 [6]float64
type Mat2x4 [8]float64
type Mat3x2 [6]float64
type Mat3 [9]float64
type Mat3x4 [12]float64
type Mat4x2 [8]float64
type Mat4x3 [12]float64
type Mat4 [16]float64

// Ident<N> returns the NxN identity matrix.
// The identity matrix is a square matrix with the value 1 on its
// diagonals. The characteristic property of the identity matrix is that
// any matrix multiplied by it is itself. (MI = M; IN = N)
func Ident2() Mat2 {
	return Mat2{1, 0, 0, 1}
}

// Ident<N> returns the NxN identity matrix.
// The identity matrix is a square matrix with the value 1 on its
// diagonals. The characteristic property of the identity matrix is that
// any matrix multiplied by it is itself. (MI = M; IN = N)
func Ident3() Mat3 {
	return Mat3{1, 0, 0, 0, 1, 0, 0, 0, 1}
}

// Ident<N> returns the NxN identity matrix.
// The identity matrix is a square matrix with the value 1 on its
// diagonals. The characteristic property of the identity matrix is that
// any matrix multiplied by it is itself. (MI = M; IN = N)
func Ident4() Mat4 {
	return Mat4{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
}

// Add performs an element-wise addition of two matrices, this is
// equivalent to iterating over every element of m1 and adding the corresponding value of m2.
func (m1 Mat2) Add(m2 Mat2) Mat2 {
	return Mat2{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3]}
}

// Add performs an element-wise addition of two matrices, this is
// equivalent to iterating over every element of m1 and adding the corresponding value of m2.
func (m1 Mat2x3) Add(m2 Mat2x3) Mat2x3 {
	return Mat2x3{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5]}
}

// Add performs an element-wise addition of two matrices, this is
// equivalent to iterating over every element of m1 and adding the corresponding value of m2.
func (m1 Mat2x4) Add(m2 Mat2x4) Mat2x4 {
	return Mat2x4{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7]}
}

// Add performs an element-wise addition of two matrices, this is
// equivalent to iterating over every element of m1 and adding the corresponding value of m2.
func (m1 Mat3x2) Add(m2 Mat3x2) Mat3x2 {
	return Mat3x2{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5]}
}

// Add performs an element-wise addition of two matrices, this is
// equivalent to iterating over every element of m1 and adding the corresponding value of m2.
func (m1 Mat3) Add(m2 Mat3) Mat3 {
	return Mat3{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7], m1[8] + m2[8]}
}

// Add performs an element-wise addition of two matrices, this is
// equivalent to iterating over every element of m1 and adding the corresponding value of m2.
func (m1 Mat3x4) Add(m2 Mat3x4) Mat3x4 {
	return Mat3x4{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7], m1[8] + m2[8], m1[9] + m2[9], m1[10] + m2[10], m1[11] + m2[11]}
}

// Add performs an element-wise addition of two matrices, this is
// equivalent to iterating over every element of m1 and adding the corresponding value of m2.
func (m1 Mat4x2) Add(m2 Mat4x2) Mat4x2 {
	return Mat4x2{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7]}
}

// Add performs an element-wise addition of two matrices, this is
// equivalent to iterating over every element of m1 and adding the corresponding value of m2.
func (m1 Mat4x3) Add(m2 Mat4x3) Mat4x3 {
	return Mat4x3{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7], m1[8] + m2[8], m1[9] + m2[9], m1[10] + m2[10], m1[11] + m2[11]}
}

// Add performs an element-wise addition of two matrices, this is
// equivalent to iterating over every element of m1 and adding the corresponding value of m2.
func (m1 Mat4) Add(m2 Mat4) Mat4 {
	return Mat4{m1[0] + m2[0], m1[1] + m2[1], m1[2] + m2[2], m1[3] + m2[3], m1[4] + m2[4], m1[5] + m2[5], m1[6] + m2[6], m1[7] + m2[7], m1[8] + m2[8], m1[9] + m2[9], m1[10] + m2[10], m1[11] + m2[11], m1[12] + m2[12], m1[13] + m2[13], m1[14] + m2[14], m1[15] + m2[15]}
}

// Sub performs an element-wise subtraction of two matrices, this is
// equivalent to iterating over every element of m1 and subtracting the corresponding value of m2.
func (m1 Mat2) Sub(m2 Mat2) Mat2 {
	return Mat2{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3]}
}

// Sub performs an element-wise subtraction of two matrices, this is
// equivalent to iterating over every element of m1 and subtracting the corresponding value of m2.
func (m1 Mat2x3) Sub(m2 Mat2x3) Mat2x3 {
	return Mat2x3{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5]}
}

// Sub performs an element-wise subtraction of two matrices, this is
// equivalent to iterating over every element of m1 and subtracting the corresponding value of m2.
func (m1 Mat2x4) Sub(m2 Mat2x4) Mat2x4 {
	return Mat2x4{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7]}
}

// Sub performs an element-wise subtraction of two matrices, this is
// equivalent to iterating over every element of m1 and subtracting the corresponding value of m2.
func (m1 Mat3x2) Sub(m2 Mat3x2) Mat3x2 {
	return Mat3x2{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5]}
}

// Sub performs an element-wise subtraction of two matrices, this is
// equivalent to iterating over every element of m1 and subtracting the corresponding value of m2.
func (m1 Mat3) Sub(m2 Mat3) Mat3 {
	return Mat3{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7], m1[8] - m2[8]}
}

// Sub performs an element-wise subtraction of two matrices, this is
// equivalent to iterating over every element of m1 and subtracting the corresponding value of m2.
func (m1 Mat3x4) Sub(m2 Mat3x4) Mat3x4 {
	return Mat3x4{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7], m1[8] - m2[8], m1[9] - m2[9], m1[10] - m2[10], m1[11] - m2[11]}
}

// Sub performs an element-wise subtraction of two matrices, this is
// equivalent to iterating over every element of m1 and subtracting the corresponding value of m2.
func (m1 Mat4x2) Sub(m2 Mat4x2) Mat4x2 {
	return Mat4x2{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7]}
}

// Sub performs an element-wise subtraction of two matrices, this is
// equivalent to iterating over every element of m1 and subtracting the corresponding value of m2.
func (m1 Mat4x3) Sub(m2 Mat4x3) Mat4x3 {
	return Mat4x3{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7], m1[8] - m2[8], m1[9] - m2[9], m1[10] - m2[10], m1[11] - m2[11]}
}

// Sub performs an element-wise subtraction of two matrices, this is
// equivalent to iterating over every element of m1 and subtracting the corresponding value of m2.
func (m1 Mat4) Sub(m2 Mat4) Mat4 {
	return Mat4{m1[0] - m2[0], m1[1] - m2[1], m1[2] - m2[2], m1[3] - m2[3], m1[4] - m2[4], m1[5] - m2[5], m1[6] - m2[6], m1[7] - m2[7], m1[8] - m2[8], m1[9] - m2[9], m1[10] - m2[10], m1[11] - m2[11], m1[12] - m2[12], m1[13] - m2[13], m1[14] - m2[14], m1[15] - m2[15]}
}

// Mul performs a scalar multiplcation of the matrix. This is equivalent to iterating
// over every element of the matrix and multiply it by c.
func (m1 Mat2) Mul(c float64) Mat2 {
	return Mat2{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c}
}

// Mul performs a scalar multiplcation of the matrix. This is equivalent to iterating
// over every element of the matrix and multiply it by c.
func (m1 Mat2x3) Mul(c float64) Mat2x3 {
	return Mat2x3{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c}
}

// Mul performs a scalar multiplcation of the matrix. This is equivalent to iterating
// over every element of the matrix and multiply it by c.
func (m1 Mat2x4) Mul(c float64) Mat2x4 {
	return Mat2x4{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c}
}

// Mul performs a scalar multiplcation of the matrix. This is equivalent to iterating
// over every element of the matrix and multiply it by c.
func (m1 Mat3x2) Mul(c float64) Mat3x2 {
	return Mat3x2{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c}
}

// Mul performs a scalar multiplcation of the matrix. This is equivalent to iterating
// over every element of the matrix and multiply it by c.
func (m1 Mat3) Mul(c float64) Mat3 {
	return Mat3{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c, m1[8] * c}
}

// Mul performs a scalar multiplcation of the matrix. This is equivalent to iterating
// over every element of the matrix and multiply it by c.
func (m1 Mat3x4) Mul(c float64) Mat3x4 {
	return Mat3x4{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c, m1[8] * c, m1[9] * c, m1[10] * c, m1[11] * c}
}

// Mul performs a scalar multiplcation of the matrix. This is equivalent to iterating
// over every element of the matrix and multiply it by c.
func (m1 Mat4x2) Mul(c float64) Mat4x2 {
	return Mat4x2{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c}
}

// Mul performs a scalar multiplcation of the matrix. This is equivalent to iterating
// over every element of the matrix and multiply it by c.
func (m1 Mat4x3) Mul(c float64) Mat4x3 {
	return Mat4x3{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c, m1[8] * c, m1[9] * c, m1[10] * c, m1[11] * c}
}

// Mul performs a scalar multiplcation of the matrix. This is equivalent to iterating
// over every element of the matrix and multiply it by c.
func (m1 Mat4) Mul(c float64) Mat4 {
	return Mat4{m1[0] * c, m1[1] * c, m1[2] * c, m1[3] * c, m1[4] * c, m1[5] * c, m1[6] * c, m1[7] * c, m1[8] * c, m1[9] * c, m1[10] * c, m1[11] * c, m1[12] * c, m1[13] * c, m1[14] * c, m1[15] * c}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat2) Mul2x1(m2 Vec2) Vec2 {
	return Vec2{m1[0]*m2[0] + m1[2]*m2[1], m1[1]*m2[0] + m1[3]*m2[1]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat2) Mul2(m2 Mat2) Mat2 {
	return Mat2{m1[0]*m2[0] + m1[2]*m2[1], m1[1]*m2[0] + m1[3]*m2[1], m1[0]*m2[2] + m1[2]*m2[3], m1[1]*m2[2] + m1[3]*m2[3]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat2) Mul2x3(m2 Mat2x3) Mat2x3 {
	return Mat2x3{m1[0]*m2[0] + m1[2]*m2[1], m1[1]*m2[0] + m1[3]*m2[1], m1[0]*m2[2] + m1[2]*m2[3], m1[1]*m2[2] + m1[3]*m2[3], m1[0]*m2[4] + m1[2]*m2[5], m1[1]*m2[4] + m1[3]*m2[5]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat2) Mul2x4(m2 Mat2x4) Mat2x4 {
	return Mat2x4{m1[0]*m2[0] + m1[2]*m2[1], m1[1]*m2[0] + m1[3]*m2[1], m1[0]*m2[2] + m1[2]*m2[3], m1[1]*m2[2] + m1[3]*m2[3], m1[0]*m2[4] + m1[2]*m2[5], m1[1]*m2[4] + m1[3]*m2[5], m1[0]*m2[6] + m1[2]*m2[7], m1[1]*m2[6] + m1[3]*m2[7]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat2x3) Mul3x1(m2 Vec3) Vec2 {
	return Vec2{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat2x3) Mul3x2(m2 Mat3x2) Mat2 {
	return Mat2{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2], m1[0]*m2[3] + m1[2]*m2[4] + m1[4]*m2[5], m1[1]*m2[3] + m1[3]*m2[4] + m1[5]*m2[5]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat2x3) Mul3(m2 Mat3) Mat2x3 {
	return Mat2x3{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2], m1[0]*m2[3] + m1[2]*m2[4] + m1[4]*m2[5], m1[1]*m2[3] + m1[3]*m2[4] + m1[5]*m2[5], m1[0]*m2[6] + m1[2]*m2[7] + m1[4]*m2[8], m1[1]*m2[6] + m1[3]*m2[7] + m1[5]*m2[8]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat2x3) Mul3x4(m2 Mat3x4) Mat2x4 {
	return Mat2x4{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2], m1[0]*m2[3] + m1[2]*m2[4] + m1[4]*m2[5], m1[1]*m2[3] + m1[3]*m2[4] + m1[5]*m2[5], m1[0]*m2[6] + m1[2]*m2[7] + m1[4]*m2[8], m1[1]*m2[6] + m1[3]*m2[7] + m1[5]*m2[8], m1[0]*m2[9] + m1[2]*m2[10] + m1[4]*m2[11], m1[1]*m2[9] + m1[3]*m2[10] + m1[5]*m2[11]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat2x4) Mul4x1(m2 Vec4) Vec2 {
	return Vec2{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2] + m1[6]*m2[3], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2] + m1[7]*m2[3]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat2x4) Mul4x2(m2 Mat4x2) Mat2 {
	return Mat2{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2] + m1[6]*m2[3], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[2]*m2[5] + m1[4]*m2[6] + m1[6]*m2[7], m1[1]*m2[4] + m1[3]*m2[5] + m1[5]*m2[6] + m1[7]*m2[7]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat2x4) Mul4x3(m2 Mat4x3) Mat2x3 {
	return Mat2x3{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2] + m1[6]*m2[3], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[2]*m2[5] + m1[4]*m2[6] + m1[6]*m2[7], m1[1]*m2[4] + m1[3]*m2[5] + m1[5]*m2[6] + m1[7]*m2[7], m1[0]*m2[8] + m1[2]*m2[9] + m1[4]*m2[10] + m1[6]*m2[11], m1[1]*m2[8] + m1[3]*m2[9] + m1[5]*m2[10] + m1[7]*m2[11]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat2x4) Mul4(m2 Mat4) Mat2x4 {
	return Mat2x4{m1[0]*m2[0] + m1[2]*m2[1] + m1[4]*m2[2] + m1[6]*m2[3], m1[1]*m2[0] + m1[3]*m2[1] + m1[5]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[2]*m2[5] + m1[4]*m2[6] + m1[6]*m2[7], m1[1]*m2[4] + m1[3]*m2[5] + m1[5]*m2[6] + m1[7]*m2[7], m1[0]*m2[8] + m1[2]*m2[9] + m1[4]*m2[10] + m1[6]*m2[11], m1[1]*m2[8] + m1[3]*m2[9] + m1[5]*m2[10] + m1[7]*m2[11], m1[0]*m2[12] + m1[2]*m2[13] + m1[4]*m2[14] + m1[6]*m2[15], m1[1]*m2[12] + m1[3]*m2[13] + m1[5]*m2[14] + m1[7]*m2[15]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat3x2) Mul2x1(m2 Vec2) Vec3 {
	return Vec3{m1[0]*m2[0] + m1[3]*m2[1], m1[1]*m2[0] + m1[4]*m2[1], m1[2]*m2[0] + m1[5]*m2[1]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat3x2) Mul2(m2 Mat2) Mat3x2 {
	return Mat3x2{m1[0]*m2[0] + m1[3]*m2[1], m1[1]*m2[0] + m1[4]*m2[1], m1[2]*m2[0] + m1[5]*m2[1], m1[0]*m2[2] + m1[3]*m2[3], m1[1]*m2[2] + m1[4]*m2[3], m1[2]*m2[2] + m1[5]*m2[3]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat3x2) Mul2x3(m2 Mat2x3) Mat3 {
	return Mat3{m1[0]*m2[0] + m1[3]*m2[1], m1[1]*m2[0] + m1[4]*m2[1], m1[2]*m2[0] + m1[5]*m2[1], m1[0]*m2[2] + m1[3]*m2[3], m1[1]*m2[2] + m1[4]*m2[3], m1[2]*m2[2] + m1[5]*m2[3], m1[0]*m2[4] + m1[3]*m2[5], m1[1]*m2[4] + m1[4]*m2[5], m1[2]*m2[4] + m1[5]*m2[5]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat3x2) Mul2x4(m2 Mat2x4) Mat3x4 {
	return Mat3x4{m1[0]*m2[0] + m1[3]*m2[1], m1[1]*m2[0] + m1[4]*m2[1], m1[2]*m2[0] + m1[5]*m2[1], m1[0]*m2[2] + m1[3]*m2[3], m1[1]*m2[2] + m1[4]*m2[3], m1[2]*m2[2] + m1[5]*m2[3], m1[0]*m2[4] + m1[3]*m2[5], m1[1]*m2[4] + m1[4]*m2[5], m1[2]*m2[4] + m1[5]*m2[5], m1[0]*m2[6] + m1[3]*m2[7], m1[1]*m2[6] + m1[4]*m2[7], m1[2]*m2[6] + m1[5]*m2[7]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat3) Mul3x1(m2 Vec3) Vec3 {
	return Vec3{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat3) Mul3x2(m2 Mat3x2) Mat3x2 {
	return Mat3x2{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2], m1[0]*m2[3] + m1[3]*m2[4] + m1[6]*m2[5], m1[1]*m2[3] + m1[4]*m2[4] + m1[7]*m2[5], m1[2]*m2[3] + m1[5]*m2[4] + m1[8]*m2[5]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat3) Mul3(m2 Mat3) Mat3 {
	return Mat3{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2], m1[0]*m2[3] + m1[3]*m2[4] + m1[6]*m2[5], m1[1]*m2[3] + m1[4]*m2[4] + m1[7]*m2[5], m1[2]*m2[3] + m1[5]*m2[4] + m1[8]*m2[5], m1[0]*m2[6] + m1[3]*m2[7] + m1[6]*m2[8], m1[1]*m2[6] + m1[4]*m2[7] + m1[7]*m2[8], m1[2]*m2[6] + m1[5]*m2[7] + m1[8]*m2[8]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat3) Mul3x4(m2 Mat3x4) Mat3x4 {
	return Mat3x4{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2], m1[0]*m2[3] + m1[3]*m2[4] + m1[6]*m2[5], m1[1]*m2[3] + m1[4]*m2[4] + m1[7]*m2[5], m1[2]*m2[3] + m1[5]*m2[4] + m1[8]*m2[5], m1[0]*m2[6] + m1[3]*m2[7] + m1[6]*m2[8], m1[1]*m2[6] + m1[4]*m2[7] + m1[7]*m2[8], m1[2]*m2[6] + m1[5]*m2[7] + m1[8]*m2[8], m1[0]*m2[9] + m1[3]*m2[10] + m1[6]*m2[11], m1[1]*m2[9] + m1[4]*m2[10] + m1[7]*m2[11], m1[2]*m2[9] + m1[5]*m2[10] + m1[8]*m2[11]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat3x4) Mul4x1(m2 Vec4) Vec3 {
	return Vec3{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2] + m1[9]*m2[3], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2] + m1[10]*m2[3], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2] + m1[11]*m2[3]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat3x4) Mul4x2(m2 Mat4x2) Mat3x2 {
	return Mat3x2{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2] + m1[9]*m2[3], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2] + m1[10]*m2[3], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2] + m1[11]*m2[3], m1[0]*m2[4] + m1[3]*m2[5] + m1[6]*m2[6] + m1[9]*m2[7], m1[1]*m2[4] + m1[4]*m2[5] + m1[7]*m2[6] + m1[10]*m2[7], m1[2]*m2[4] + m1[5]*m2[5] + m1[8]*m2[6] + m1[11]*m2[7]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat3x4) Mul4x3(m2 Mat4x3) Mat3 {
	return Mat3{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2] + m1[9]*m2[3], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2] + m1[10]*m2[3], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2] + m1[11]*m2[3], m1[0]*m2[4] + m1[3]*m2[5] + m1[6]*m2[6] + m1[9]*m2[7], m1[1]*m2[4] + m1[4]*m2[5] + m1[7]*m2[6] + m1[10]*m2[7], m1[2]*m2[4] + m1[5]*m2[5] + m1[8]*m2[6] + m1[11]*m2[7], m1[0]*m2[8] + m1[3]*m2[9] + m1[6]*m2[10] + m1[9]*m2[11], m1[1]*m2[8] + m1[4]*m2[9] + m1[7]*m2[10] + m1[10]*m2[11], m1[2]*m2[8] + m1[5]*m2[9] + m1[8]*m2[10] + m1[11]*m2[11]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat3x4) Mul4(m2 Mat4) Mat3x4 {
	return Mat3x4{m1[0]*m2[0] + m1[3]*m2[1] + m1[6]*m2[2] + m1[9]*m2[3], m1[1]*m2[0] + m1[4]*m2[1] + m1[7]*m2[2] + m1[10]*m2[3], m1[2]*m2[0] + m1[5]*m2[1] + m1[8]*m2[2] + m1[11]*m2[3], m1[0]*m2[4] + m1[3]*m2[5] + m1[6]*m2[6] + m1[9]*m2[7], m1[1]*m2[4] + m1[4]*m2[5] + m1[7]*m2[6] + m1[10]*m2[7], m1[2]*m2[4] + m1[5]*m2[5] + m1[8]*m2[6] + m1[11]*m2[7], m1[0]*m2[8] + m1[3]*m2[9] + m1[6]*m2[10] + m1[9]*m2[11], m1[1]*m2[8] + m1[4]*m2[9] + m1[7]*m2[10] + m1[10]*m2[11], m1[2]*m2[8] + m1[5]*m2[9] + m1[8]*m2[10] + m1[11]*m2[11], m1[0]*m2[12] + m1[3]*m2[13] + m1[6]*m2[14] + m1[9]*m2[15], m1[1]*m2[12] + m1[4]*m2[13] + m1[7]*m2[14] + m1[10]*m2[15], m1[2]*m2[12] + m1[5]*m2[13] + m1[8]*m2[14] + m1[11]*m2[15]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat4x2) Mul2x1(m2 Vec2) Vec4 {
	return Vec4{m1[0]*m2[0] + m1[4]*m2[1], m1[1]*m2[0] + m1[5]*m2[1], m1[2]*m2[0] + m1[6]*m2[1], m1[3]*m2[0] + m1[7]*m2[1]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat4x2) Mul2(m2 Mat2) Mat4x2 {
	return Mat4x2{m1[0]*m2[0] + m1[4]*m2[1], m1[1]*m2[0] + m1[5]*m2[1], m1[2]*m2[0] + m1[6]*m2[1], m1[3]*m2[0] + m1[7]*m2[1], m1[0]*m2[2] + m1[4]*m2[3], m1[1]*m2[2] + m1[5]*m2[3], m1[2]*m2[2] + m1[6]*m2[3], m1[3]*m2[2] + m1[7]*m2[3]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat4x2) Mul2x3(m2 Mat2x3) Mat4x3 {
	return Mat4x3{m1[0]*m2[0] + m1[4]*m2[1], m1[1]*m2[0] + m1[5]*m2[1], m1[2]*m2[0] + m1[6]*m2[1], m1[3]*m2[0] + m1[7]*m2[1], m1[0]*m2[2] + m1[4]*m2[3], m1[1]*m2[2] + m1[5]*m2[3], m1[2]*m2[2] + m1[6]*m2[3], m1[3]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[4]*m2[5], m1[1]*m2[4] + m1[5]*m2[5], m1[2]*m2[4] + m1[6]*m2[5], m1[3]*m2[4] + m1[7]*m2[5]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat4x2) Mul2x4(m2 Mat2x4) Mat4 {
	return Mat4{m1[0]*m2[0] + m1[4]*m2[1], m1[1]*m2[0] + m1[5]*m2[1], m1[2]*m2[0] + m1[6]*m2[1], m1[3]*m2[0] + m1[7]*m2[1], m1[0]*m2[2] + m1[4]*m2[3], m1[1]*m2[2] + m1[5]*m2[3], m1[2]*m2[2] + m1[6]*m2[3], m1[3]*m2[2] + m1[7]*m2[3], m1[0]*m2[4] + m1[4]*m2[5], m1[1]*m2[4] + m1[5]*m2[5], m1[2]*m2[4] + m1[6]*m2[5], m1[3]*m2[4] + m1[7]*m2[5], m1[0]*m2[6] + m1[4]*m2[7], m1[1]*m2[6] + m1[5]*m2[7], m1[2]*m2[6] + m1[6]*m2[7], m1[3]*m2[6] + m1[7]*m2[7]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat4x3) Mul3x1(m2 Vec3) Vec4 {
	return Vec4{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat4x3) Mul3x2(m2 Mat3x2) Mat4x2 {
	return Mat4x2{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2], m1[0]*m2[3] + m1[4]*m2[4] + m1[8]*m2[5], m1[1]*m2[3] + m1[5]*m2[4] + m1[9]*m2[5], m1[2]*m2[3] + m1[6]*m2[4] + m1[10]*m2[5], m1[3]*m2[3] + m1[7]*m2[4] + m1[11]*m2[5]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat4x3) Mul3(m2 Mat3) Mat4x3 {
	return Mat4x3{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2], m1[0]*m2[3] + m1[4]*m2[4] + m1[8]*m2[5], m1[1]*m2[3] + m1[5]*m2[4] + m1[9]*m2[5], m1[2]*m2[3] + m1[6]*m2[4] + m1[10]*m2[5], m1[3]*m2[3] + m1[7]*m2[4] + m1[11]*m2[5], m1[0]*m2[6] + m1[4]*m2[7] + m1[8]*m2[8], m1[1]*m2[6] + m1[5]*m2[7] + m1[9]*m2[8], m1[2]*m2[6] + m1[6]*m2[7] + m1[10]*m2[8], m1[3]*m2[6] + m1[7]*m2[7] + m1[11]*m2[8]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat4x3) Mul3x4(m2 Mat3x4) Mat4 {
	return Mat4{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2], m1[0]*m2[3] + m1[4]*m2[4] + m1[8]*m2[5], m1[1]*m2[3] + m1[5]*m2[4] + m1[9]*m2[5], m1[2]*m2[3] + m1[6]*m2[4] + m1[10]*m2[5], m1[3]*m2[3] + m1[7]*m2[4] + m1[11]*m2[5], m1[0]*m2[6] + m1[4]*m2[7] + m1[8]*m2[8], m1[1]*m2[6] + m1[5]*m2[7] + m1[9]*m2[8], m1[2]*m2[6] + m1[6]*m2[7] + m1[10]*m2[8], m1[3]*m2[6] + m1[7]*m2[7] + m1[11]*m2[8], m1[0]*m2[9] + m1[4]*m2[10] + m1[8]*m2[11], m1[1]*m2[9] + m1[5]*m2[10] + m1[9]*m2[11], m1[2]*m2[9] + m1[6]*m2[10] + m1[10]*m2[11], m1[3]*m2[9] + m1[7]*m2[10] + m1[11]*m2[11]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat4) Mul4x1(m2 Vec4) Vec4 {
	return Vec4{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2] + m1[12]*m2[3], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2] + m1[13]*m2[3], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2] + m1[14]*m2[3], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2] + m1[15]*m2[3]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat4) Mul4x2(m2 Mat4x2) Mat4x2 {
	return Mat4x2{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2] + m1[12]*m2[3], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2] + m1[13]*m2[3], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2] + m1[14]*m2[3], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2] + m1[15]*m2[3], m1[0]*m2[4] + m1[4]*m2[5] + m1[8]*m2[6] + m1[12]*m2[7], m1[1]*m2[4] + m1[5]*m2[5] + m1[9]*m2[6] + m1[13]*m2[7], m1[2]*m2[4] + m1[6]*m2[5] + m1[10]*m2[6] + m1[14]*m2[7], m1[3]*m2[4] + m1[7]*m2[5] + m1[11]*m2[6] + m1[15]*m2[7]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat4) Mul4x3(m2 Mat4x3) Mat4x3 {
	return Mat4x3{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2] + m1[12]*m2[3], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2] + m1[13]*m2[3], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2] + m1[14]*m2[3], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2] + m1[15]*m2[3], m1[0]*m2[4] + m1[4]*m2[5] + m1[8]*m2[6] + m1[12]*m2[7], m1[1]*m2[4] + m1[5]*m2[5] + m1[9]*m2[6] + m1[13]*m2[7], m1[2]*m2[4] + m1[6]*m2[5] + m1[10]*m2[6] + m1[14]*m2[7], m1[3]*m2[4] + m1[7]*m2[5] + m1[11]*m2[6] + m1[15]*m2[7], m1[0]*m2[8] + m1[4]*m2[9] + m1[8]*m2[10] + m1[12]*m2[11], m1[1]*m2[8] + m1[5]*m2[9] + m1[9]*m2[10] + m1[13]*m2[11], m1[2]*m2[8] + m1[6]*m2[9] + m1[10]*m2[10] + m1[14]*m2[11], m1[3]*m2[8] + m1[7]*m2[9] + m1[11]*m2[10] + m1[15]*m2[11]}
}

// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
func (m1 Mat4) Mul4(m2 Mat4) Mat4 {
	return Mat4{m1[0]*m2[0] + m1[4]*m2[1] + m1[8]*m2[2] + m1[12]*m2[3], m1[1]*m2[0] + m1[5]*m2[1] + m1[9]*m2[2] + m1[13]*m2[3], m1[2]*m2[0] + m1[6]*m2[1] + m1[10]*m2[2] + m1[14]*m2[3], m1[3]*m2[0] + m1[7]*m2[1] + m1[11]*m2[2] + m1[15]*m2[3], m1[0]*m2[4] + m1[4]*m2[5] + m1[8]*m2[6] + m1[12]*m2[7], m1[1]*m2[4] + m1[5]*m2[5] + m1[9]*m2[6] + m1[13]*m2[7], m1[2]*m2[4] + m1[6]*m2[5] + m1[10]*m2[6] + m1[14]*m2[7], m1[3]*m2[4] + m1[7]*m2[5] + m1[11]*m2[6] + m1[15]*m2[7], m1[0]*m2[8] + m1[4]*m2[9] + m1[8]*m2[10] + m1[12]*m2[11], m1[1]*m2[8] + m1[5]*m2[9] + m1[9]*m2[10] + m1[13]*m2[11], m1[2]*m2[8] + m1[6]*m2[9] + m1[10]*m2[10] + m1[14]*m2[11], m1[3]*m2[8] + m1[7]*m2[9] + m1[11]*m2[10] + m1[15]*m2[11], m1[0]*m2[12] + m1[4]*m2[13] + m1[8]*m2[14] + m1[12]*m2[15], m1[1]*m2[12] + m1[5]*m2[13] + m1[9]*m2[14] + m1[13]*m2[15], m1[2]*m2[12] + m1[6]*m2[13] + m1[10]*m2[14] + m1[14]*m2[15], m1[3]*m2[12] + m1[7]*m2[13] + m1[11]*m2[14] + m1[15]*m2[15]}
}

// Transpose produces the transpose of this matrix. For any MxN matrix
// the transpose is an NxM matrix with the rows swapped with the columns. For instance
// the transpose of the Mat3x2 is a Mat2x3 like so:
//
//    [[a b]]    [[a c e]]
//    [[c d]] =  [[b d f]]
//    [[e f]]
func (m1 Mat2) Transpose() Mat2 {
	return Mat2{m1[0], m1[2], m1[1], m1[3]}
}

// Transpose produces the transpose of this matrix. For any MxN matrix
// the transpose is an NxM matrix with the rows swapped with the columns. For instance
// the transpose of the Mat3x2 is a Mat2x3 like so:
//
//    [[a b]]    [[a c e]]
//    [[c d]] =  [[b d f]]
//    [[e f]]
func (m1 Mat2x3) Transpose() Mat3x2 {
	return Mat3x2{m1[0], m1[3], m1[1], m1[4], m1[2], m1[5]}
}

// Transpose produces the transpose of this matrix. For any MxN matrix
// the transpose is an NxM matrix with the rows swapped with the columns. For instance
// the transpose of the Mat3x2 is a Mat2x3 like so:
//
//    [[a b]]    [[a c e]]
//    [[c d]] =  [[b d f]]
//    [[e f]]
func (m1 Mat2x4) Transpose() Mat4x2 {
	return Mat4x2{m1[0], m1[4], m1[1], m1[5], m1[2], m1[6], m1[3], m1[7]}
}

// Transpose produces the transpose of this matrix. For any MxN matrix
// the transpose is an NxM matrix with the rows swapped with the columns. For instance
// the transpose of the Mat3x2 is a Mat2x3 like so:
//
//    [[a b]]    [[a c e]]
//    [[c d]] =  [[b d f]]
//    [[e f]]
func (m1 Mat3x2) Transpose() Mat2x3 {
	return Mat2x3{m1[0], m1[2], m1[4], m1[1], m1[3], m1[5]}
}

// Transpose produces the transpose of this matrix. For any MxN matrix
// the transpose is an NxM matrix with the rows swapped with the columns. For instance
// the transpose of the Mat3x2 is a Mat2x3 like so:
//
//    [[a b]]    [[a c e]]
//    [[c d]] =  [[b d f]]
//    [[e f]]
func (m1 Mat3) Transpose() Mat3 {
	return Mat3{m1[0], m1[3], m1[6], m1[1], m1[4], m1[7], m1[2], m1[5], m1[8]}
}

// Transpose produces the transpose of this matrix. For any MxN matrix
// the transpose is an NxM matrix with the rows swapped with the columns. For instance
// the transpose of the Mat3x2 is a Mat2x3 like so:
//
//    [[a b]]    [[a c e]]
//    [[c d]] =  [[b d f]]
//    [[e f]]
func (m1 Mat3x4) Transpose() Mat4x3 {
	return Mat4x3{m1[0], m1[4], m1[8], m1[1], m1[5], m1[9], m1[2], m1[6], m1[10], m1[3], m1[7], m1[11]}
}

// Transpose produces the transpose of this matrix. For any MxN matrix
// the transpose is an NxM matrix with the rows swapped with the columns. For instance
// the transpose of the Mat3x2 is a Mat2x3 like so:
//
//    [[a b]]    [[a c e]]
//    [[c d]] =  [[b d f]]
//    [[e f]]
func (m1 Mat4x2) Transpose() Mat2x4 {
	return Mat2x4{m1[0], m1[2], m1[4], m1[6], m1[1], m1[3], m1[5], m1[7]}
}

// Transpose produces the transpose of this matrix. For any MxN matrix
// the transpose is an NxM matrix with the rows swapped with the columns. For instance
// the transpose of the Mat3x2 is a Mat2x3 like so:
//
//    [[a b]]    [[a c e]]
//    [[c d]] =  [[b d f]]
//    [[e f]]
func (m1 Mat4x3) Transpose() Mat3x4 {
	return Mat3x4{m1[0], m1[3], m1[6], m1[9], m1[1], m1[4], m1[7], m1[10], m1[2], m1[5], m1[8], m1[11]}
}

// Transpose produces the transpose of this matrix. For any MxN matrix
// the transpose is an NxM matrix with the rows swapped with the columns. For instance
// the transpose of the Mat3x2 is a Mat2x3 like so:
//
//    [[a b]]    [[a c e]]
//    [[c d]] =  [[b d f]]
//    [[e f]]
func (m1 Mat4) Transpose() Mat4 {
	return Mat4{m1[0], m1[4], m1[8], m1[12], m1[1], m1[5], m1[9], m1[13], m1[2], m1[6], m1[10], m1[14], m1[3], m1[7], m1[11], m1[15]}
}

// The determinant of a matrix is a measure of a square matrix's
// singularity and invertability, among other things. In this library, the
// determinant is hard coded based on pre-computed cofactor expansion, and uses
// no loops. Of course, the addition and multiplication must still be done.
func (m Mat2) Det() float64 {
	return m[0]*m[2] - m[1]*m[3]
}

// The determinant of a matrix is a measure of a square matrix's
// singularity and invertability, among other things. In this library, the
// determinant is hard coded based on pre-computed cofactor expansion, and uses
// no loops. Of course, the addition and multiplication must still be done.
func (m Mat3) Det() float64 {
	return m[0]*m[4]*m[8] + m[3]*m[7]*m[2] + m[6]*m[1]*m[5] - m[6]*m[4]*m[2] - m[3]*m[1]*m[8] - m[0]*m[7]*m[5]
}

// The determinant of a matrix is a measure of a square matrix's
// singularity and invertability, among other things. In this library, the
// determinant is hard coded based on pre-computed cofactor expansion, and uses
// no loops. Of course, the addition and multiplication must still be done.
func (m Mat4) Det() float64 {
	return m[0]*m[5]*m[10]*m[15] - m[0]*m[5]*m[11]*m[14] - m[0]*m[6]*m[9]*m[15] + m[0]*m[6]*m[11]*m[13] + m[0]*m[7]*m[9]*m[14] - m[0]*m[7]*m[10]*m[13] - m[1]*m[4]*m[10]*m[15] + m[1]*m[4]*m[11]*m[14] + m[1]*m[6]*m[8]*m[15] - m[1]*m[6]*m[11]*m[12] - m[1]*m[7]*m[8]*m[14] + m[1]*m[7]*m[10]*m[12] + m[2]*m[4]*m[9]*m[15] - m[2]*m[4]*m[11]*m[13] - m[2]*m[5]*m[8]*m[15] + m[2]*m[5]*m[11]*m[12] + m[2]*m[7]*m[8]*m[13] - m[2]*m[7]*m[9]*m[12] - m[3]*m[4]*m[9]*m[14] + m[3]*m[4]*m[10]*m[13] + m[3]*m[5]*m[8]*m[14] - m[3]*m[5]*m[10]*m[12] - m[3]*m[6]*m[8]*m[13] + m[3]*m[6]*m[9]*m[12]
}

// Inv computes the inverse of a square matrix. An inverse is a square matrix such that when multiplied by the
// original, yields the identity.
//
// M_inv * M = M * M_inv = I
//
// In this library, the math is precomputed, and uses no loops, though the multiplications, additions, determinant calculation, and scaling
// are still done. This can still be (relatively) expensive for a 4x4.
//
// This function does not check the determinant to see if the matrix is invertible.
// If the determinant is 0.0, the value of all elements will be
// infinite. (See here for why: http://play.golang.org/p/Aaj7SnbqIp )
// Therefore, if the program really cares, it should check the determinant first.
// In the future, an alternate function may be written which takes in a pre-computed determinant.
func (m Mat2) Inv() Mat2 {
	det := m.Det()
	if FloatEqual(det, float64(0.0)) {
		return Mat2{}
	}
	retMat := Mat2{m[3], -m[1], -m[2], m[0]}
	return retMat.Mul(1 / det)
}

// Inv computes the inverse of a square matrix. An inverse is a square matrix such that when multiplied by the
// original, yields the identity.
//
// M_inv * M = M * M_inv = I
//
// In this library, the math is precomputed, and uses no loops, though the multiplications, additions, determinant calculation, and scaling
// are still done. This can still be (relatively) expensive for a 4x4.
//
// This function does not check the determinant to see if the matrix is invertible.
// If the determinant is 0.0, the value of all elements will be
// infinite. (See here for why: http://play.golang.org/p/Aaj7SnbqIp )
// Therefore, if the program really cares, it should check the determinant first.
// In the future, an alternate function may be written which takes in a pre-computed determinant.
func (m Mat3) Inv() Mat3 {
	det := m.Det()
	if FloatEqual(det, float64(0.0)) {
		return Mat3{}
	}
	retMat := Mat3{m[4]*m[8] - m[5]*m[7], m[2]*m[7] - m[1]*m[8], m[1]*m[5] - m[2]*m[4], m[5]*m[6] - m[3]*m[8], m[0]*m[8] - m[2]*m[6], m[2]*m[3] - m[0]*m[5], m[3]*m[7] - m[4]*m[6], m[1]*m[6] - m[0]*m[7], m[0]*m[4] - m[1]*m[3]}
	return retMat.Mul(1 / det)
}

// Inv computes the inverse of a square matrix. An inverse is a square matrix such that when multiplied by the
// original, yields the identity.
//
// M_inv * M = M * M_inv = I
//
// In this library, the math is precomputed, and uses no loops, though the multiplications, additions, determinant calculation, and scaling
// are still done. This can still be (relatively) expensive for a 4x4.
//
// This function does not check the determinant to see if the matrix is invertible.
// If the determinant is 0.0, the value of all elements will be
// infinite. (See here for why: http://play.golang.org/p/Aaj7SnbqIp )
// Therefore, if the program really cares, it should check the determinant first.
// In the future, an alternate function may be written which takes in a pre-computed determinant.
func (m Mat4) Inv() Mat4 {
	det := m.Det()
	if FloatEqual(det, float64(0.0)) {
		return Mat4{}
	}
	retMat := Mat4{-m[7]*m[10]*m[13] + m[6]*m[11]*m[13] + m[7]*m[9]*m[14] - m[5]*m[11]*m[14] - m[6]*m[9]*m[15] + m[5]*m[10]*m[15], m[3]*m[10]*m[13] - m[2]*m[11]*m[13] - m[3]*m[9]*m[14] + m[1]*m[11]*m[14] + m[2]*m[9]*m[15] - m[1]*m[10]*m[15], -m[3]*m[6]*m[13] + m[2]*m[7]*m[13] + m[3]*m[5]*m[14] - m[1]*m[7]*m[14] - m[2]*m[5]*m[15] + m[1]*m[6]*m[15], m[3]*m[6]*m[9] - m[2]*m[7]*m[9] - m[3]*m[5]*m[10] + m[1]*m[7]*m[10] + m[2]*m[5]*m[11] - m[1]*m[6]*m[11], m[7]*m[10]*m[12] - m[6]*m[11]*m[12] - m[7]*m[8]*m[14] + m[4]*m[11]*m[14] + m[6]*m[8]*m[15] - m[4]*m[10]*m[15], -m[3]*m[10]*m[12] + m[2]*m[11]*m[12] + m[3]*m[8]*m[14] - m[0]*m[11]*m[14] - m[2]*m[8]*m[15] + m[0]*m[10]*m[15], m[3]*m[6]*m[12] - m[2]*m[7]*m[12] - m[3]*m[4]*m[14] + m[0]*m[7]*m[14] + m[2]*m[4]*m[15] - m[0]*m[6]*m[15], -m[3]*m[6]*m[8] + m[2]*m[7]*m[8] + m[3]*m[4]*m[10] - m[0]*m[7]*m[10] - m[2]*m[4]*m[11] + m[0]*m[6]*m[11], -m[7]*m[9]*m[12] + m[5]*m[11]*m[12] + m[7]*m[8]*m[13] - m[4]*m[11]*m[13] - m[5]*m[8]*m[15] + m[4]*m[9]*m[15], m[3]*m[9]*m[12] - m[1]*m[11]*m[12] - m[3]*m[8]*m[13] + m[0]*m[11]*m[13] + m[1]*m[8]*m[15] - m[0]*m[9]*m[15], -m[3]*m[5]*m[12] + m[1]*m[7]*m[12] + m[3]*m[4]*m[13] - m[0]*m[7]*m[13] - m[1]*m[4]*m[15] + m[0]*m[5]*m[15], m[3]*m[5]*m[8] - m[1]*m[7]*m[8] - m[3]*m[4]*m[9] + m[0]*m[7]*m[9] + m[1]*m[4]*m[11] - m[0]*m[5]*m[11], m[6]*m[9]*m[12] - m[5]*m[10]*m[12] - m[6]*m[8]*m[13] + m[4]*m[10]*m[13] + m[5]*m[8]*m[14] - m[4]*m[9]*m[14], -m[2]*m[9]*m[12] + m[1]*m[10]*m[12] + m[2]*m[8]*m[13] - m[0]*m[10]*m[13] - m[1]*m[8]*m[14] + m[0]*m[9]*m[14], m[2]*m[5]*m[12] - m[1]*m[6]*m[12] - m[2]*m[4]*m[13] + m[0]*m[6]*m[13] + m[1]*m[4]*m[14] - m[0]*m[5]*m[14], -m[2]*m[5]*m[8] + m[1]*m[6]*m[8] + m[2]*m[4]*m[9] - m[0]*m[6]*m[9] - m[1]*m[4]*m[10] + m[0]*m[5]*m[10]}
	return retMat.Mul(1 / det)
}

// ApproxEqual performs an element-wise approximate equality test between two matrices,
// as if FloatEqual had been used.
func (m1 Mat2) ApproxEqual(m2 Mat2) bool {
	for i := range m1 {
		if !FloatEqual(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqual performs an element-wise approximate equality test between two matrices,
// as if FloatEqual had been used.
func (m1 Mat2x3) ApproxEqual(m2 Mat2x3) bool {
	for i := range m1 {
		if !FloatEqual(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqual performs an element-wise approximate equality test between two matrices,
// as if FloatEqual had been used.
func (m1 Mat2x4) ApproxEqual(m2 Mat2x4) bool {
	for i := range m1 {
		if !FloatEqual(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqual performs an element-wise approximate equality test between two matrices,
// as if FloatEqual had been used.
func (m1 Mat3x2) ApproxEqual(m2 Mat3x2) bool {
	for i := range m1 {
		if !FloatEqual(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqual performs an element-wise approximate equality test between two matrices,
// as if FloatEqual had been used.
func (m1 Mat3) ApproxEqual(m2 Mat3) bool {
	for i := range m1 {
		if !FloatEqual(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqual performs an element-wise approximate equality test between two matrices,
// as if FloatEqual had been used.
func (m1 Mat3x4) ApproxEqual(m2 Mat3x4) bool {
	for i := range m1 {
		if !FloatEqual(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqual performs an element-wise approximate equality test between two matrices,
// as if FloatEqual had been used.
func (m1 Mat4x2) ApproxEqual(m2 Mat4x2) bool {
	for i := range m1 {
		if !FloatEqual(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqual performs an element-wise approximate equality test between two matrices,
// as if FloatEqual had been used.
func (m1 Mat4x3) ApproxEqual(m2 Mat4x3) bool {
	for i := range m1 {
		if !FloatEqual(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqual performs an element-wise approximate equality test between two matrices,
// as if FloatEqual had been used.
func (m1 Mat4) ApproxEqual(m2 Mat4) bool {
	for i := range m1 {
		if !FloatEqual(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqualThreshold performs an element-wise approximate equality test between two matrices
// with a given epsilon threshold, as if FloatEqualThreshold had been used.
func (m1 Mat2) ApproxEqualThreshold(m2 Mat2, threshold float64) bool {
	for i := range m1 {
		if !FloatEqualThreshold(m1[i], m2[i], threshold) {
			return false
		}
	}
	return true
}

// ApproxEqualThreshold performs an element-wise approximate equality test between two matrices
// with a given epsilon threshold, as if FloatEqualThreshold had been used.
func (m1 Mat2x3) ApproxEqualThreshold(m2 Mat2x3, threshold float64) bool {
	for i := range m1 {
		if !FloatEqualThreshold(m1[i], m2[i], threshold) {
			return false
		}
	}
	return true
}

// ApproxEqualThreshold performs an element-wise approximate equality test between two matrices
// with a given epsilon threshold, as if FloatEqualThreshold had been used.
func (m1 Mat2x4) ApproxEqualThreshold(m2 Mat2x4, threshold float64) bool {
	for i := range m1 {
		if !FloatEqualThreshold(m1[i], m2[i], threshold) {
			return false
		}
	}
	return true
}

// ApproxEqualThreshold performs an element-wise approximate equality test between two matrices
// with a given epsilon threshold, as if FloatEqualThreshold had been used.
func (m1 Mat3x2) ApproxEqualThreshold(m2 Mat3x2, threshold float64) bool {
	for i := range m1 {
		if !FloatEqualThreshold(m1[i], m2[i], threshold) {
			return false
		}
	}
	return true
}

// ApproxEqualThreshold performs an element-wise approximate equality test between two matrices
// with a given epsilon threshold, as if FloatEqualThreshold had been used.
func (m1 Mat3) ApproxEqualThreshold(m2 Mat3, threshold float64) bool {
	for i := range m1 {
		if !FloatEqualThreshold(m1[i], m2[i], threshold) {
			return false
		}
	}
	return true
}

// ApproxEqualThreshold performs an element-wise approximate equality test between two matrices
// with a given epsilon threshold, as if FloatEqualThreshold had been used.
func (m1 Mat3x4) ApproxEqualThreshold(m2 Mat3x4, threshold float64) bool {
	for i := range m1 {
		if !FloatEqualThreshold(m1[i], m2[i], threshold) {
			return false
		}
	}
	return true
}

// ApproxEqualThreshold performs an element-wise approximate equality test between two matrices
// with a given epsilon threshold, as if FloatEqualThreshold had been used.
func (m1 Mat4x2) ApproxEqualThreshold(m2 Mat4x2, threshold float64) bool {
	for i := range m1 {
		if !FloatEqualThreshold(m1[i], m2[i], threshold) {
			return false
		}
	}
	return true
}

// ApproxEqualThreshold performs an element-wise approximate equality test between two matrices
// with a given epsilon threshold, as if FloatEqualThreshold had been used.
func (m1 Mat4x3) ApproxEqualThreshold(m2 Mat4x3, threshold float64) bool {
	for i := range m1 {
		if !FloatEqualThreshold(m1[i], m2[i], threshold) {
			return false
		}
	}
	return true
}

// ApproxEqualThreshold performs an element-wise approximate equality test between two matrices
// with a given epsilon threshold, as if FloatEqualThreshold had been used.
func (m1 Mat4) ApproxEqualThreshold(m2 Mat4, threshold float64) bool {
	for i := range m1 {
		if !FloatEqualThreshold(m1[i], m2[i], threshold) {
			return false
		}
	}
	return true
}

// ApproxEqualFunc performs an element-wise approximate equality test between two matrices
// with a given equality functions, intended to be used with FloatEqualFunc; although and comparison
// function may be used in practice.
func (m1 Mat2) ApproxFuncEqual(m2 Mat2, eq func(float64, float64) bool) bool {
	for i := range m1 {
		if !eq(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqualFunc performs an element-wise approximate equality test between two matrices
// with a given equality functions, intended to be used with FloatEqualFunc; although and comparison
// function may be used in practice.
func (m1 Mat2x3) ApproxFuncEqual(m2 Mat2x3, eq func(float64, float64) bool) bool {
	for i := range m1 {
		if !eq(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqualFunc performs an element-wise approximate equality test between two matrices
// with a given equality functions, intended to be used with FloatEqualFunc; although and comparison
// function may be used in practice.
func (m1 Mat2x4) ApproxFuncEqual(m2 Mat2x4, eq func(float64, float64) bool) bool {
	for i := range m1 {
		if !eq(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqualFunc performs an element-wise approximate equality test between two matrices
// with a given equality functions, intended to be used with FloatEqualFunc; although and comparison
// function may be used in practice.
func (m1 Mat3x2) ApproxFuncEqual(m2 Mat3x2, eq func(float64, float64) bool) bool {
	for i := range m1 {
		if !eq(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqualFunc performs an element-wise approximate equality test between two matrices
// with a given equality functions, intended to be used with FloatEqualFunc; although and comparison
// function may be used in practice.
func (m1 Mat3) ApproxFuncEqual(m2 Mat3, eq func(float64, float64) bool) bool {
	for i := range m1 {
		if !eq(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqualFunc performs an element-wise approximate equality test between two matrices
// with a given equality functions, intended to be used with FloatEqualFunc; although and comparison
// function may be used in practice.
func (m1 Mat3x4) ApproxFuncEqual(m2 Mat3x4, eq func(float64, float64) bool) bool {
	for i := range m1 {
		if !eq(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqualFunc performs an element-wise approximate equality test between two matrices
// with a given equality functions, intended to be used with FloatEqualFunc; although and comparison
// function may be used in practice.
func (m1 Mat4x2) ApproxFuncEqual(m2 Mat4x2, eq func(float64, float64) bool) bool {
	for i := range m1 {
		if !eq(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqualFunc performs an element-wise approximate equality test between two matrices
// with a given equality functions, intended to be used with FloatEqualFunc; although and comparison
// function may be used in practice.
func (m1 Mat4x3) ApproxFuncEqual(m2 Mat4x3, eq func(float64, float64) bool) bool {
	for i := range m1 {
		if !eq(m1[i], m2[i]) {
			return false
		}
	}
	return true
}

// ApproxEqualFunc performs an element-wise approximate equality test between two matrices
// with a given equality functions, intended to be used with FloatEqualFunc; although and comparison
// function may be used in practice.
func (m1 Mat4) ApproxFuncEqual(m2 Mat4, eq func(float64, float64) bool) bool {
	for i := range m1 {
		if !eq(m1[i], m2[i]) {
			return false
		}
	}
	return true
}
