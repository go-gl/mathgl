// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mgl64

import (
	"math/rand"
	"testing"
	"time"
)

func TestMulIdent(t *testing.T) {
	i1 := [...]float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
	i2 := Ident4()
	i3 := Ident4()

	mul := i2.Mul4(i3)

	for i := range mul {
		if mul[i] != i1[i] {
			t.Errorf("Multiplication of identities does not yield identity")
		}
	}
}

// M>N
func TestMatRowsMGTN(t *testing.T) {
	rows := [2]Vec3{
		Vec3{1, 2, 3},
		Vec3{4, 5, 6},
	}
	m1 := Mat2x3FromRows(rows)

	t.Logf("2x3 matrix as built from rows: %v", m1)
	for r := 0; r < 2; r++ {
		for c := 0; c < 3; c++ {
			if !FloatEqualThreshold(m1.At(r, c), rows[r][c], 1e-5) {
				t.Errorf("Matrix element at (%d,%d) wrong when built from rows. Got: %f, Expected: %f", r, c, m1.At(r, c), rows[r][c])
			}
		}
	}

	r2 := m1.Rows()

	t.Logf("2x3 matrix returned rows: %v", r2)
	for r := 0; r < 2; r++ {
		for c := 0; c < 3; c++ {
			if !FloatEqualThreshold(r2[r][c], rows[r][c], 1e-5) {
				t.Errorf("Matrix element at (%d,%d) wrong when rows are gotten. Got: %f, Expected: %f", r, c, r2[r][c], rows[r][c])
			}
		}
	}
}

// M<N
func TestMatRowsMLTN(t *testing.T) {
	rows := [4]Vec3{
		Vec3{1, 2, 3},
		Vec3{4, 5, 6},
		Vec3{7, 8, 9},
		Vec3{10, 11, 12},
	}
	m1 := Mat4x3FromRows(rows)

	t.Logf("4x3 matrix as built from rows: %v", m1)
	for r := 0; r < 4; r++ {
		for c := 0; c < 3; c++ {
			if !FloatEqualThreshold(m1.At(r, c), rows[r][c], 1e-5) {
				t.Errorf("Matrix element at (%d,%d) wrong when built from rows. Got: %f, Expected: %f", r, c, m1.At(r, c), rows[r][c])
			}
		}
	}

	r2 := m1.Rows()

	t.Logf("4x3 matrix returned rows: %v", r2)
	for r := 0; r < 4; r++ {
		for c := 0; c < 3; c++ {
			if !FloatEqualThreshold(r2[r][c], rows[r][c], 1e-5) {
				t.Errorf("Matrix element at (%d,%d) wrong when rows are gotten. Got: %f, Expected: %f", r, c, r2[r][c], rows[r][c])
			}
		}
	}
}

// Square matrix
func TestMatRowsM(t *testing.T) {
	rows := [4]Vec4{
		Vec4{1, 2, 3, 4},
		Vec4{5, 6, 7, 8},
		Vec4{9, 10, 11, 12},
		Vec4{13, 14, 15, 16},
	}
	m1 := Mat4FromRows(rows)

	t.Logf("4x4 matrix as built from rows: %v", m1)
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			if !FloatEqualThreshold(m1.At(r, c), rows[r][c], 1e-5) {
				t.Errorf("Matrix element at (%d,%d) wrong when built from rows. Got: %f, Expected: %f", r, c, m1.At(r, c), rows[r][c])
			}
		}
	}

	r2 := m1.Rows()

	t.Logf("4x4 matrix returned rows: %v", r2)
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			if !FloatEqualThreshold(r2[r][c], rows[r][c], 1e-5) {
				t.Errorf("Matrix element at (%d,%d) wrong when rows are gotten. Got: %f, Expected: %f", r, c, r2[r][c], rows[r][c])
			}
		}
	}
}

// M<N
func TestMatColsMLTN(t *testing.T) {
	cols := [2]Vec3{
		Vec3{1, 2, 3},
		Vec3{4, 5, 6},
	}
	m1 := Mat3x2FromCols(cols)

	t.Logf("3x2 matrix as built from cols: %v", m1)
	for r := 0; r < 3; r++ {
		for c := 0; c < 2; c++ {
			if !FloatEqualThreshold(m1.At(r, c), cols[c][r], 1e-5) {
				t.Errorf("Matrix element at (%d,%d) wrong when built from rows. Got: %f, Expected: %f", r, c, m1.At(r, c), cols[c][r])
			}
		}
	}

	r2 := m1.Cols()

	t.Logf("3x2 matrix returned cols: %v", r2)
	for r := 0; r < 3; r++ {
		for c := 0; c < 2; c++ {
			if !FloatEqualThreshold(r2[c][r], cols[c][r], 1e-5) {
				t.Errorf("Matrix element at (%d,%d) wrong when rows are gotten. Got: %f, Expected: %f", r, c, r2[c][r], cols[c][r])
			}
		}
	}
}

// M>N
func TestMatColsMGTN(t *testing.T) {
	cols := [4]Vec3{
		Vec3{1, 2, 3},
		Vec3{4, 5, 6},
		Vec3{7, 8, 9},
		Vec3{10, 11, 12},
	}
	m1 := Mat3x4FromCols(cols)

	t.Logf("3x4 matrix as built from cols: %v", m1)
	for r := 0; r < 3; r++ {
		for c := 0; c < 4; c++ {
			if !FloatEqualThreshold(m1.At(r, c), cols[c][r], 1e-5) {
				t.Errorf("Matrix element at (%d,%d) wrong when built from cols. Got: %f, Expected: %f", r, c, m1.At(r, c), cols[c][r])
			}
		}
	}

	r2 := m1.Cols()

	t.Logf("3x4 matrix returned cols: %v", r2)
	for r := 0; r < 3; r++ {
		for c := 0; c < 4; c++ {
			if !FloatEqualThreshold(r2[c][r], cols[c][r], 1e-5) {
				t.Errorf("Matrix element at (%d,%d) wrong when rows are gotten. Got: %f, Expected: %f", r, c, r2[c][r], cols[c][r])
			}
		}
	}
}

// Square matrix
func TestMatColsM(t *testing.T) {
	cols := [4]Vec4{
		Vec4{1, 2, 3, 4},
		Vec4{5, 6, 7, 8},
		Vec4{9, 10, 11, 12},
		Vec4{13, 14, 15, 16},
	}
	m1 := Mat4FromCols(cols)

	t.Logf("4x4 matrix as built from cols: %v", m1)
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			if !FloatEqualThreshold(m1.At(r, c), cols[c][r], 1e-5) {
				t.Errorf("Matrix element at (%d,%d) wrong when built from rows. Got: %f, Expected: %f", r, c, m1.At(r, c), cols[c][r])
			}
		}
	}

	r2 := m1.Cols()

	t.Logf("4x4 matrix returned cols: %v", r2)
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			if !FloatEqualThreshold(r2[c][r], cols[c][r], 1e-5) {
				t.Errorf("Matrix element at (%d,%d) wrong when rows are gotten. Got: %f, Expected: %f", r, c, r2[c][r], cols[c][r])
			}
		}
	}
}

func BenchmarkMatAdd(b *testing.B) {
	b.StopTimer()
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m1 := Mat4{}
		m2 := Mat4{}

		for j := 0; j < len(m1); j++ {
			m1[j], m2[j] = rand.Float64(), rand.Float64()
		}
		b.StartTimer()

		m1 = m1.Add(m2)
	}
}

func BenchmarkMatScale(b *testing.B) {
	b.StopTimer()
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m1 := Mat4{}

		for j := 0; j < len(m1); j++ {
			m1[j] = rand.Float64()
		}
		c := rand.Float64()
		b.StartTimer()

		m1 = m1.Mul(c)
	}
}

func BenchmarkMatMul(b *testing.B) {
	b.StopTimer()
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m1 := Mat4{}
		m2 := Mat4{}

		for j := 0; j < len(m1); j++ {
			m1[j], m2[j] = rand.Float64(), rand.Float64()
		}
		b.StartTimer()

		m1 = m1.Mul4(m2)
	}
}

func BenchmarkMatTranspose(b *testing.B) {
	b.StopTimer()
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m1 := Mat4{}

		for j := 0; j < len(m1); j++ {
			m1[j] = rand.Float64()
		}
		b.StartTimer()

		_ = m1.Transpose()
	}
}

func BenchmarkMatDet(b *testing.B) {
	b.StopTimer()
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m1 := Mat4{}

		for j := 0; j < len(m1); j++ {
			m1[j] = rand.Float64()
		}
		b.StartTimer()

		_ = m1.Det()
	}
}

func BenchmarkMatInv(b *testing.B) {
	b.StopTimer()
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m1 := Mat4{}

		for j := 0; j < len(m1); j++ {
			m1[j] = rand.Float64()
		}
		b.StartTimer()

		m1 = m1.Inv()
	}
}
