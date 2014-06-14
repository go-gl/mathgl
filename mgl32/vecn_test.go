package mgl32

import (
	"testing"
)

func TestVecNCross(t *testing.T) {
	v1 := Vec3{1, 3, 5}
	v2 := Vec3{2, 4, 6}

	correct := v1.Cross(v2)
	correctN := NewBackedVecN(correct[:])

	v1n := NewBackedVecN(v1[:])
	v2n := NewBackedVecN(v2[:])

	result := v1n.Cross(nil, v2n)

	if !correctN.ApproxEqualThreshold(result, 1e-4) {
		t.Errorf("VecN cross product is incorrect. Got: %v; Expected: %v", result, correct)
	}
}

func TestVecNDot(t *testing.T) {
	v1 := Vec3{1, 3, 5}
	v2 := Vec3{2, 4, 6}

	correct := v1.Dot(v2)

	v1n := NewBackedVecN(v1[:])
	v2n := NewBackedVecN(v2[:])

	result := v1n.Dot(v2n)

	if !FloatEqualThreshold(correct, result, 1e-4) {
		t.Errorf("Dot product doesn't work for VecN. Got: %v, Expected: %v", result, correct)
	}
}

func TestVecNMul(t *testing.T) {
	v1 := Vec3{1, 3, 5}

	correct := v1.Mul(3)
	correctN := NewBackedVecN(correct[:])

	v1n := NewBackedVecN(v1[:])

	result := v1n.Mul(nil, 3)

	if !correctN.ApproxEqualThreshold(result, 1e-4) {
		t.Errorf("VecN scalar multiplication is incorrect. Got: %v; Expected: %v", result, correct)
	}
}

func TestVecNNormalize(t *testing.T) {
	v1 := Vec3{1, 3, 5}

	correct := v1.Normalize()
	correctN := NewBackedVecN(correct[:])

	v1n := NewBackedVecN(v1[:])

	result := v1n.Normalize(nil)

	if !correctN.ApproxEqualThreshold(result, 1e-4) {
		t.Errorf("VecN normalization is incorrect. Got: %v; Expected: %v", result, correct)
	}
}

func TestVecNAdd(t *testing.T) {
	v1 := Vec3{1, 3, 5}
	v2 := Vec3{2, 4, 6}

	correct := v1.Add(v2)
	correctN := NewBackedVecN(correct[:])

	v1n := NewBackedVecN(v1[:])
	v2n := NewBackedVecN(v2[:])

	result := v1n.Add(nil, v2n)

	if !correctN.ApproxEqualThreshold(result, 1e-4) {
		t.Errorf("VecN addition is incorrect. Got: %v; Expected: %v", result, correct)
	}
}

func TestVecNSub(t *testing.T) {
	v1 := Vec3{1, 3, 5}
	v2 := Vec3{2, 4, 6}

	correct := v1.Sub(v2)
	correctN := NewBackedVecN(correct[:])

	v1n := NewBackedVecN(v1[:])
	v2n := NewBackedVecN(v2[:])

	result := v1n.Sub(nil, v2n)

	if !correctN.ApproxEqualThreshold(result, 1e-4) {
		t.Errorf("VecN subtraction is incorrect. Got: %v; Expected: %v", result, correct)
	}
}
