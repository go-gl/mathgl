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

		m1.Add(m2)
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

		m1.Mul(c)
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

		m1.Mul4(m2)
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

		m1.Transpose()
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

		m1.Det()
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

		m1.Inv()
	}
}
