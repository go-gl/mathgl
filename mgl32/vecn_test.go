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
