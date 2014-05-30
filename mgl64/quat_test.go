package mgl64

import (
	"math"
	"testing"
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

	i1 := QuatRotate(angleDegrees, axis)

	rotatedAxis := i1.Rotate(axis)

	for i := range rotatedAxis {
		if !FloatEqualThreshold(rotatedAxis[i], axis[i], 1e-4) {
			t.Errorf("Rotation of axis does not yield identity")
		}
	}
}

func TestQuatRotateOffAxis(t *testing.T) {
	var angleDegrees float64 = 30.0
	var angleRads float64 = angleDegrees * math.Pi / 180.0
	axis := Vec3{1, 0, 0}

	i1 := QuatRotate(angleDegrees, axis)

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
	var angle float64 = 45.0
	axis := Vec3{1, 2, 3}.Normalize()
	quat := QuatRotate(angle, axis)
	matrix := quat.Mat4()
	answer := HomogRotate3D(angle*math.Pi/180, axis)

	if !matrix.ApproxEqualThreshold(answer, 1e-4) {
		t.Errorf("Rotation quaternion does not yield correct rotation matrix; got: %v expected: %v", matrix, answer)
	}
}
