// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mgl32

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
	var angleDegrees float32 = 30.0
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
	var angleRads float32 = DegToRad(30.0)
	axis := Vec3{1, 0, 0}

	i1 := QuatRotate(angleRads, axis)

	vector := Vec3{0, 1, 0}
	rotatedVector := i1.Rotate(vector)

	s, c := math.Sincos(float64(angleRads))
	answer := Vec3{0, float32(c), float32(s)}

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
	var angle float32 = DegToRad(45.0)

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
	q := QuatRotate(float32(math.Pi), Vec3{0, 1, 0})
	q = q.Normalize()
	v := Vec3{1, 0, 0}

	result := q.Rotate(v)

	expected := Rotate3DY(float32(math.Pi)).Mul3x1(v)
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
		q := QuatRotate(rand.Float32(), Vec3{rand.Float32(), rand.Float32(), rand.Float32()})
		v := Vec3{rand.Float32(), rand.Float32(), rand.Float32()}
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
		q := QuatRotate(rand.Float32(), Vec3{rand.Float32(), rand.Float32(), rand.Float32()})
		v := Vec3{rand.Float32(), rand.Float32(), rand.Float32()}
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
		q := QuatRotate(rand.Float32(), Vec3{rand.Float32(), rand.Float32(), rand.Float32()})
		b.StartTimer()

		_ = q.V[0]
	}
}

func BenchmarkQuatFuncElementAccess(b *testing.B) {
	b.StopTimer()
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		q := QuatRotate(rand.Float32(), Vec3{rand.Float32(), rand.Float32(), rand.Float32()})
		b.StartTimer()

		_ = q.X()
	}
}

func TestMat4ToQuat(t *testing.T) {
	// http://www.euclideanspace.com/maths/geometry/rotations/conversions/matrixToQuaternion/examples/index.htm

	tests := []struct {
		Description string
		Rotation    Mat4
		Expected    Quat
	}{
		{
			"forward",
			Ident4(),
			QuatIdent(),
		},
		{
			"heading 90 degree",
			Mat4{
				0, 0, -1, 0,
				0, 1, 0, 0,
				1, 0, 0, 0,
				0, 0, 0, 1,
			},
			Quat{0.7071, Vec3{0, 0.7071, 0}},
		},
		{
			"heading 180 degree",
			Mat4{
				-1, 0, 0, 0,
				0, 1, 0, 0,
				0, 0, -1, 0,
				0, 0, 0, 1,
			},
			Quat{0, Vec3{0, 1, 0}},
		},
		{
			"attitude 90 degree",
			Mat4{
				0, 1, 0, 0,
				-1, 0, 0, 0,
				0, 0, 1, 0,
				0, 0, 0, 1,
			},
			Quat{0.7071, Vec3{0, 0, 0.7071}},
		},
		{
			"bank 90 degree",
			Mat4{
				1, 0, 0, 0,
				0, 0, 1, 0,
				0, -1, 0, 0,
				0, 0, 0, 1,
			},
			Quat{0.7071, Vec3{0.7071, 0, 0}},
		},
	}

	threshold := float32(math.Pow(10, -2))
	for _, c := range tests {
		if r := Mat4ToQuat(c.Rotation); !r.ApproxEqualThreshold(c.Expected, threshold) {
			t.Errorf("%v failed: Mat4ToQuat(%v) != %v (got %v)", c.Description, c.Rotation, c.Expected, r)
		}
	}
}

func TestQuatRotate(t *testing.T) {
	tests := []struct {
		Description string
		Angle       float32
		Axis        Vec3
		Expected    Quat
	}{
		{
			"forward",
			0, Vec3{0, 0, 0},
			QuatIdent(),
		},
		{
			"heading 90 degree",
			DegToRad(90), Vec3{0, 1, 0},
			Quat{0.7071, Vec3{0, 0.7071, 0}},
		},
		{
			"heading 180 degree",
			DegToRad(180), Vec3{0, 1, 0},
			Quat{0, Vec3{0, 1, 0}},
		},
		{
			"attitude 90 degree",
			DegToRad(90), Vec3{0, 0, 1},
			Quat{0.7071, Vec3{0, 0, 0.7071}},
		},
		{
			"bank 90 degree",
			DegToRad(90), Vec3{1, 0, 0},
			Quat{0.7071, Vec3{0.7071, 0, 0}},
		},
	}

	threshold := float32(math.Pow(10, -2))
	for _, c := range tests {
		if r := QuatRotate(c.Angle, c.Axis); !r.ApproxEqualThreshold(c.Expected, threshold) {
			t.Errorf("%v failed: QuatRotate(%v, %v) != %v (got %v)", c.Description, c.Angle, c.Axis, c.Expected, r)
		}
	}
}

func TestQuatLookAtV(t *testing.T) {
	// http://www.euclideanspace.com/maths/algebra/realNormedAlgebra/quaternions/transforms/examples/index.htm

	tests := []struct {
		Description     string
		Eye, Center, Up Vec3
		Expected        Quat
	}{
		{
			"forward",
			Vec3{0, 0, 0},
			Vec3{0, 0, -1},
			Vec3{0, 1, 0},
			QuatIdent(),
		},
		{
			"heading 90 degree",
			Vec3{0, 0, 0},
			Vec3{1, 0, 0},
			Vec3{0, 1, 0},
			Quat{0.7071, Vec3{0, 0.7071, 0}},
		},
		{
			"heading 180 degree",
			Vec3{0, 0, 0},
			Vec3{0, 0, 1},
			Vec3{0, 1, 0},
			Quat{0, Vec3{0, 1, 0}},
		},
		{
			"attitude 90 degree",
			Vec3{0, 0, 0},
			Vec3{0, 0, -1},
			Vec3{1, 0, 0},
			Quat{0.7071, Vec3{0, 0, 0.7071}},
		},
		{
			"bank 90 degree",
			Vec3{0, 0, 0},
			Vec3{0, -1, 0},
			Vec3{0, 0, -1},
			Quat{0.7071, Vec3{0.7071, 0, 0}},
		},
	}

	threshold := float32(math.Pow(10, -2))
	for _, c := range tests {
		if r := QuatLookAtV(c.Eye, c.Center, c.Up); !r.ApproxEqualThreshold(c.Expected, threshold) {
			t.Errorf("%v failed: QuatLookAtV(%v, %v, %v) != %v (got %v)", c.Description, c.Eye, c.Center, c.Up, c.Expected, r)
		}
	}
}

func TestQuatMatConversion(t *testing.T) {
	tests := []struct {
		Angle float32
		Axis  Vec3
	}{}

	for a := 0.0; a <= math.Pi*2; a += math.Pi / 4.0 {
		af := float32(a)
		tests = append(tests, []struct {
			Angle float32
			Axis  Vec3
		}{
			{af, Vec3{1, 0, 0}},
			{af, Vec3{0, 1, 0}},
			{af, Vec3{0, 0, 1}},
		}...)
	}

	threshold := float32(1e-4 /*math.Pow(10, -2)*/)
	for _, c := range tests {
		m1 := HomogRotate3D(c.Angle, c.Axis)
		q1 := Mat4ToQuat(m1)

		q2 := QuatRotate(c.Angle, c.Axis)
		m2 := q2.Mat4()

		if !m1.ApproxEqualThreshold(m2, threshold) {
			t.Errorf("Rotation matrices for %v %v do not match:\n%v\n%v", RadToDeg(c.Angle), c.Axis, m1, m2)
		}

		if !q1.ApproxEqualThreshold(q2, threshold) {
			t.Errorf("Quaternions for %v %v do not match:\n%v\n%v", RadToDeg(c.Angle), c.Axis, q1, q2)
		}
	}
}
