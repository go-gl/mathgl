package mgl32

import (
	"testing"
)

func TestMxNTransposeWide(t *testing.T) {
	m := Mat2x3FromCols([3]Vec2{
		Vec2{1, 2},
		Vec2{3, 4},
		Vec2{5, 6},
	})

	mn := NewBackedMatrix(m[:], 2, 3)

	transpose := m.Transpose()

	transposeMN := mn.Transpose(nil)

	correct := NewBackedMatrix(transpose[:], 3, 2)

	if !correct.ApproxEqualThreshold(transposeMN, 1e-4) {
		t.Errorf("Transpose gives incorrect result; got: %v, expected: %v", transposeMN, correct)
	}
}

func TestMxNTransposeTall(t *testing.T) {
	m := Mat3x2FromCols([2]Vec3{
		Vec3{1, 2, 3},
		Vec3{4, 5, 6},
	})

	mn := NewBackedMatrix(m[:], 3, 2)

	transpose := m.Transpose()

	transposeMN := mn.Transpose(nil)

	correct := NewBackedMatrix(transpose[:], 2, 3)

	if !correct.ApproxEqualThreshold(transposeMN, 1e-4) {
		t.Errorf("Transpose gives incorrect result; got: %v, expected: %v", transposeMN, correct)
	}
}

func TestMxNTransposeSquare(t *testing.T) {
	m := Mat3FromCols([3]Vec3{
		Vec3{1, 2, 3},
		Vec3{4, 5, 6},
		Vec3{7, 8, 9},
	})

	mn := NewBackedMatrix(m[:], 3, 3)

	transpose := m.Transpose()

	transposeMN := mn.Transpose(nil)

	correct := NewBackedMatrix(transpose[:], 3, 3)

	if !correct.ApproxEqualThreshold(transposeMN, 1e-4) {
		t.Errorf("Transpose gives incorrect result; got: %v, expected: %v", transposeMN, correct)
	}
}

func TestMxNAtSet(t *testing.T) {
	m := Mat3{1, 2, 3, 4, 5, 6, 7, 8, 9}

	mn := NewBackedMatrix(m[:], 3, 3)

	v := mn.At(0, 2)

	if !FloatEqualThreshold(v, 7, 1e-4) {
		t.Errorf("Incorrect value gotten by At: %v, expected %v", v, 7)
	}

	mn.Set(0, 2, 9001)

	v = mn.At(0, 2)

	if !FloatEqualThreshold(v, 9001, 1e-4) {
		t.Errorf("Incorrect value set by Set: %v, expected %v", v, 9001)
	}

	correct := Mat3{1, 2, 3, 4, 5, 6, 9001, 8, 9}
	correctMN := NewBackedMatrix(correct[:], 3, 3)

	if !correctMN.ApproxEqualThreshold(mn, 1e-4) {
		t.Errorf("Set matrix does not equal correct matrix. Got: %v, expected: %v", mn, correctMN)
	}
}
