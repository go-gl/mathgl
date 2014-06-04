// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mgl64

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestQuatMulIdentity(t *testing.T) {
	i1 := Quat{1.0, Vec3{0, 0, 0}}
	i2 := QuatIdent()
	i3 := QuatIdent()

	mul := i2.Mul(i3)

	if !FloatEqual(mul.W, 1.0) {
		t.Errorf("Multiplication of identities does not yield identity")
	}

	for i := range mul.V {
		if mul.V[i] != i1.V[i] {
			t.Errorf("Multiplication of identities does not yield identity")
		}
	}
}

func TestQuatRotateOnAxis(t *testing.T) {
	var angleDegrees float64 = 30.0
	axis := Vec3{1, 0, 0}

	i1 := QuatRotate(DegToRad(angleDegrees), axis)

	rotatedAxis := i1.Rotate(axis)

	for i := range rotatedAxis {
		if !FloatEqualThreshold(rotatedAxis[i], axis[i], 1e-4) {
			t.Errorf("Rotation of axis does not yield identity")
		}
	}
}

func TestQuatRotateOffAxis(t *testing.T) {
	var angleRads float64 = DegToRad(30.0)
	axis := Vec3{1, 0, 0}

	i1 := QuatRotate(angleRads, axis)

	vector := Vec3{0, 1, 0}
	rotatedVector := i1.Rotate(vector)

	s, c := math.Sincos(float64(angleRads))
	answer := Vec3{0, float64(c), float64(s)}

	for i := range rotatedVector {
		if !FloatEqualThreshold(rotatedVector[i], answer[i], 1e-4) {
			t.Errorf("Rotation of vector does not yield answer")
		}
	}
}

func TestQuatIdentityToMatrix(t *testing.T) {
	quat := QuatIdent()
	matrix := quat.Mat4()
	answer := Ident4()

	if !matrix.ApproxEqual(answer) {
		t.Errorf("Identity quaternion does not yield identity matrix")
	}
}

func TestQuatRotationToMatrix(t *testing.T) {
	var angle float64 = DegToRad(45.0)

	axis := Vec3{1, 2, 3}.Normalize()
	quat := QuatRotate(angle, axis)
	matrix := quat.Mat4()
	answer := HomogRotate3D(angle, axis)

	if !matrix.ApproxEqualThreshold(answer, 1e-4) {
		t.Errorf("Rotation quaternion does not yield correct rotation matrix; got: %v expected: %v", matrix, answer)
	}
}

// Taken from the Matlab AnglesToQuat documentation example
func TestAnglesToQuatZYX(t *testing.T) {
	q := AnglesToQuat(.7854, 0.1, 0, ZYX)

	t.Log("Calculated quaternion: ", q, "\n")

	if !FloatEqualThreshold(q.W, .9227, 1e-3) {
		t.Errorf("Quaternion W incorrect. Got: %f Expected: %f", q.W, .9227)
	}

	if !q.V.ApproxEqualThreshold(Vec3{-0.0191, 0.0462, 0.3822}, 1e-3) {
		t.Errorf("Quaternion V incorrect. Got: %v, Expected: %v", q.V, Vec3{-0.0191, 0.0462, 0.3822})
	}
}

func TestQuatMatRotateY(t *testing.T) {
	q := QuatRotate(float64(math.Pi), Vec3{0, 1, 0})
	q = q.Normalize()
	v := Vec3{1, 0, 0}

	result := q.Rotate(v)

	expected := Rotate3DY(float64(math.Pi)).Mul3x1(v)
	t.Logf("Computed from rotation matrix: %v", expected)
	if !result.ApproxEqualThreshold(expected, 1e-4) {
		t.Errorf("Quaternion rotating vector doesn't match 3D matrix method. Got: %v, Expected: %v", result, expected)
	}

	expected = q.Mul(Quat{0, v}).Mul(q.Conjugate()).V
	t.Logf("Computed from conjugate method: %v", expected)
	if !result.ApproxEqualThreshold(expected, 1e-4) {
		t.Errorf("Quaternion rotating vector doesn't match slower conjugate method. Got: %v, Expected: %v", result, expected)
	}

	expected = Vec3{-1, 0, 0}
	if !result.ApproxEqualThreshold(expected, 4e-4) { // The result we get for z is like 8e-8, but a 1e-4 threshold juuuuuust causes it to freak out when compared to 0.0
		t.Errorf("Quaternion rotating vector doesn't match hand-computed result. Got: %v, Expected: %v", result, expected)
	}
}

func BenchmarkQuatRotateOptimized(b *testing.B) {
	b.StopTimer()
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		q := QuatRotate(rand.Float64(), Vec3{rand.Float64(), rand.Float64(), rand.Float64()})
		v := Vec3{rand.Float64(), rand.Float64(), rand.Float64()}
		q = q.Normalize()
		b.StartTimer()

		v = q.Rotate(v)
	}
}

func BenchmarkQuatRotateConjugate(b *testing.B) {
	b.StopTimer()
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		q := QuatRotate(rand.Float64(), Vec3{rand.Float64(), rand.Float64(), rand.Float64()})
		v := Vec3{rand.Float64(), rand.Float64(), rand.Float64()}
		q = q.Normalize()
		b.StartTimer()

		v = q.Mul(Quat{0, v}).Mul(q.Conjugate()).V
	}
}

func BenchmarkQuatArrayAccess(b *testing.B) {
	b.StopTimer()
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		q := QuatRotate(rand.Float64(), Vec3{rand.Float64(), rand.Float64(), rand.Float64()})
		b.StartTimer()

		_ = q.V[0]
	}
}

func BenchmarkQuatFuncElementAccess(b *testing.B) {
	b.StopTimer()
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		q := QuatRotate(rand.Float64(), Vec3{rand.Float64(), rand.Float64(), rand.Float64()})
		b.StartTimer()

		_ = q.X()
	}
}
