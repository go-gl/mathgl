package mgl32

import (
	"testing"
)

func TestMNTransposeWide(t *testing.T) {
	m := Mat2x3{1, 2, 3, 4, 5, 6}

	mn := NewBackedMatrix(m[:], 2, 3)

	transpose := m.Transpose()

	transposeMN := mn.Transpose(nil)

	correct := NewBackedMatrix(transpose[:], 3, 2)

	if !correct.ApproxEqualThreshold(transposeMN, 1e-4) {
		t.Errorf("Transpose gives incorrect result; got: %v, expected: %v", transposeMN, correct)
	}
}

func TestMNTransposeTall(t *testing.T) {
	m := Mat3x2{1, 2, 3, 4, 5, 6}

	mn := NewBackedMatrix(m[:], 3, 2)

	transpose := m.Transpose()

	transposeMN := mn.Transpose(nil)

	correct := NewBackedMatrix(transpose[:], 2, 3)

	if !correct.ApproxEqualThreshold(transposeMN, 1e-4) {
		t.Errorf("Transpose gives incorrect result; got: %v, expected: %v", transposeMN, correct)
	}
}

func TestMNTransposeSquare(t *testing.T) {
	m := Mat3{1, 2, 3, 4, 5, 6, 7, 8, 9}

	mn := NewBackedMatrix(m[:], 3, 3)

	transpose := m.Transpose()

	transposeMN := mn.Transpose(nil)

	correct := NewBackedMatrix(transpose[:], 3, 3)

	if !correct.ApproxEqualThreshold(transposeMN, 1e-4) {
		t.Errorf("Transpose gives incorrect result; got: %v, expected: %v", transposeMN, correct)
	}
}

func TestMNAtSet(t *testing.T) {
	m := Mat3{1, 2, 3, 4, 5, 6, 7, 8, 9}

	mn := NewBackedMatrix(m[:], 3, 3)

	v := mn.At(0, 2)

	if !FloatEqualThreshold(v, 3, 1e-4) {
		t.Errorf("Incorrect value gotten by At: %v, expected %v", v, 3)
	}

	mn.Set(0, 2, 9001)

	v = mn.At(0, 2)

	if !FloatEqualThreshold(v, 9001, 1e-4) {
		t.Errorf("Incorrect value set by Set: %v, expected %v", v, 9001)
	}
}
