package mathgl

import (
	"math"
	"testing"
)

func TestQuatMulIdentity(t *testing.T) {
	i1 := Quatd{1.0, Vec3d{0, 0, 0}}
	i2 := QuatIdentd()
	i3 := QuatIdentd()

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
	angleDegrees := 30.0
	axis := Vec3d{1, 0, 0}

	i1 := QuatRotated(angleDegrees, axis)

	rotatedAxis := i1.Rotate(axis)

	for i := range rotatedAxis {
		if rotatedAxis[i] != axis[i] {
			t.Errorf("Rotation of axis does not yield identity")
		}
	}
}

func TestQuatRotateOffAxis(t *testing.T) {
	angleDegrees := 30.0
	angleRads := angleDegrees * math.Pi / 180.0
	axis := Vec3d{1, 0, 0}

	i1 := QuatRotated(angleDegrees, axis)

	vector := Vec3d{0, 1, 0}
	rotatedVector := i1.Rotate(vector)
	answer := Vec3d{0, math.Cos(angleRads), math.Sin(angleRads)}

	for i := range rotatedVector {
		if rotatedVector[i] != answer[i] {
			t.Errorf("Rotation of vector does not yield answer")
		}
	}
}
