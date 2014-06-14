package mgl32

import (
	"testing"
)

func TestTranspose(t *testing.T) {
	m := Mat2x3{1, 2, 3, 4, 5, 6}

	mn := NewBackedMatrix(m[:], 2, 3)

	transpose := m.Transpose()
	t.Log(transpose)

	transposeMN := mn.Transpose(nil)

	correct := NewBackedMatrix(transpose[:], 3, 2)

	if !correct.ApproxEqualThreshold(transposeMN, 1e-4) {
		t.Errorf("Transpose gives incorrect result; got: %v, expected: %v", transposeMN, correct)
	}
}
