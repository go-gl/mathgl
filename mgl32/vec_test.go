// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mgl32

import (
	"math/rand"
	"testing"
	"time"
)

/* Only floats are tested because the double versions are simply a find->replace on floats */

func Test2DVecAdd(t *testing.T) {
	v1 := Vec2{1.0, 2.5}
	v2 := Vec2{0.0, 1.0}

	v3 := v1.Add(v2)

	if !FloatEqual(v3[0], 1.0) || !FloatEqual(v3[1], 3.5) {
		t.Errorf("Add not adding properly")
	}

	v4 := v2.Add(v1)

	if !FloatEqual(v3[0], v4[0]) || !FloatEqual(v3[0], v4[0]) {
		t.Errorf("Addition is somehow not commutative")
	}

}

func Test3DVecAdd(t *testing.T) {
	v1 := Vec3{1.0, 2.5, 1.1}
	v2 := Vec3{0.0, 1.0, 9.9}

	v3 := v1.Add(v2)

	if !FloatEqual(v3[0], 1.0) || !FloatEqual(v3[1], 3.5) || !FloatEqual(v3[2], 11.0) {
		t.Errorf("Add not adding properly")
	}

	v4 := v2.Add(v1)

	if !FloatEqual(v3[0], v4[0]) || !FloatEqual(v3[0], v4[0]) || !FloatEqual(v3[2], v4[2]) {
		t.Errorf("Addition is somehow not commutative")
	}

}

func Test4DVecAdd(t *testing.T) {
	v1 := Vec4{1.0, 2.5, 1.1, 2.0}
	v2 := Vec4{0.0, 1.0, 9.9, 100.0}

	v3 := v1.Add(v2)

	if !FloatEqual(v3[0], 1.0) || !FloatEqual(v3[1], 3.5) || !FloatEqual(v3[2], 11.0) || !FloatEqual(v3[3], 102.0) {
		t.Errorf("Add not adding properly")
	}

	v4 := v2.Add(v1)

	if !FloatEqual(v3[0], v4[0]) || !FloatEqual(v3[0], v4[0]) || !FloatEqual(v3[2], v4[2]) || !FloatEqual(v3[3], v4[3]) {
		t.Errorf("Addition is somehow not commutative")
	}

}

func Test2DVecSub(t *testing.T) {
	v1 := Vec2{1.0, 2.5}
	v2 := Vec2{0.0, 1.0}

	v3 := v1.Sub(v2)

	if !FloatEqual(v3[0], 1.0) || !FloatEqual(v3[1], 1.5) {
		t.Errorf("Sub not subtracting properly [%f, %f]", v3[0], v3[1])
	}

}

func Test3DVecSub(t *testing.T) {
	v1 := Vec3{1.0, 2.5, 1.1}
	v2 := Vec3{0.0, 1.0, 9.9}

	v3 := v1.Sub(v2)

	// 1.1-9.9 does stupid things to floats, so we need threshold
	if !FloatEqual(v3[0], 1.0) || !FloatEqual(v3[1], 1.5) || !FloatEqualThreshold(v3[2], -8.8, 1e-5) {
		t.Errorf("Sub not subtracting properly [%f, %f, %f]", v3[0], v3[1], v3[2])
	}

}

func Test4DVecSub(t *testing.T) {
	v1 := Vec4{1.0, 2.5, 1.1, 2.0}
	v2 := Vec4{0.0, 1.0, 9.9, 100.0}

	v3 := v1.Sub(v2)

	// 1.1-9.9 does stupid things to floats, so we need a more tolerant threshold
	if !FloatEqual(v3[0], 1.0) || !FloatEqual(v3[1], 1.5) || !FloatEqualThreshold(v3[2], -8.8, 1e-5) || !FloatEqual(v3[3], -98.0) {
		t.Errorf("Sub not subtracting properly [%f, %f, %f, %f]", v3[0], v3[1], v3[2], v3[3])
	}

}

func TestVecScale(t *testing.T) {
	v := Vec2{1.0, 0.0}
	v = v.Mul(15.0)

	if !FloatEqual(v[0], 15.0) || !FloatEqual(v[1], 0.0) {
		t.Errorf("Vec mul does something weird [%f, %f]", v[0], v[1])
	}

	v2 := Vec3{1.0, 0.0, 100.1}
	v2 = v2.Mul(15.0)

	if !FloatEqual(v2[0], 15.0) || !FloatEqual(v2[1], 0.0) || !FloatEqual(v2[2], 1501.5) {
		t.Errorf("Vec mul does something weird [%f, %f, %f]", v2[0], v2[1], v2[2])
	}

	v3 := Vec4{1.0, 0.0, 100.1, -1.0}
	v3 = v3.Mul(15.0)

	if !FloatEqual(v3[0], 15.0) || !FloatEqual(v3[1], 0.0) || !FloatEqual(v3[2], 1501.5) || !FloatEqual(v3[3], -15.0) {
		t.Errorf("Vec mul does something weird [%f, %f, %f, %f]", v3[0], v3[1], v3[2], v3[3])
	}
}

func BenchmarkVec4Add(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v1 := Vec4{r.Float32(), r.Float32(), r.Float32(), r.Float32()}
		v2 := Vec4{r.Float32(), r.Float32(), r.Float32(), r.Float32()}
		b.StartTimer()

		v1.Add(v2)
	}
}

func BenchmarkVec4Sub(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v1 := Vec4{r.Float32(), r.Float32(), r.Float32(), r.Float32()}
		v2 := Vec4{r.Float32(), r.Float32(), r.Float32(), r.Float32()}
		b.StartTimer()

		v1.Sub(v2)
	}
}

func BenchmarkVec4Scale(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v1 := Vec4{r.Float32(), r.Float32(), r.Float32(), r.Float32()}
		c := r.Float32()
		b.StartTimer()

		v1.Mul(c)
	}
}

func BenchmarkVec4Dot(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v1 := Vec4{r.Float32(), r.Float32(), r.Float32(), r.Float32()}
		v2 := Vec4{r.Float32(), r.Float32(), r.Float32(), r.Float32()}
		b.StartTimer()

		v1.Dot(v2)
	}
}

func BenchmarkVec4Len(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v1 := Vec4{r.Float32(), r.Float32(), r.Float32(), r.Float32()}
		b.StartTimer()

		v1.Len()
	}
}

func BenchmarkVec4Normalize(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v1 := Vec4{r.Float32(), r.Float32(), r.Float32(), r.Float32()}
		b.StartTimer()

		v1.Normalize()
	}
}

func BenchmarkVecCross(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		v1 := Vec3{r.Float32(), r.Float32(), r.Float32()}
		v2 := Vec3{r.Float32(), r.Float32(), r.Float32()}
		b.StartTimer()

		v1.Cross(v2)
	}
}
