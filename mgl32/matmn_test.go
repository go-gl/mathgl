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

func TestMxNReallocCallback(t *testing.T) {
	var retVal []float32 = nil
	reallocCallback = func(buf []float32) {
		t.Log("In realloc callback")
		retVal = buf
	}

	// To prevent affecting other tests
	defer func() {
		reallocCallback = nil
	}()

	a := NewMatrix(3, 3)
	a.Reshape(4, 4)

	if retVal == nil {
		t.Errorf("Realloc callback not set or called correctly")
	}

	a = nil
	a.Reshape(4, 4)

	if retVal == nil {
		t.Errorf("Realloc callback is being called to realloc over a nil slice")
	}

}

func TestMxNMulMxN(t *testing.T) {
	m := Ident4()
	r := HomogRotate3DX(DegToRad(45))
	tr := Translate3D(1, 0, 0)
	s := Scale3D(2, 2, 2)

	correct := tr.Mul4(r.Mul4(s.Mul4(m))) // tr*r*s
	correctMN := NewBackedMatrix(correct[:], 4, 4)

	mn := NewBackedMatrix(m[:], 4, 4)
	rmn := NewBackedMatrix(r[:], 4, 4)
	trmn := NewBackedMatrix(tr[:], 4, 4)
	smn := NewBackedMatrix(s[:], 4, 4)

	result := trmn.MulMxN(nil, rmn.MulMxN(nil, smn.MulMxN(nil, mn)))

	if !result.ApproxEqualThreshold(correctMN, 1e-4) {
		t.Errorf("Multiplication of MxN matrix and 4x4 matrix not the same. Got: %v expected: %v", result, correctMN)
	}
}

func TestMxNMulMxNErrorHandling(t *testing.T) {
	mn := NewMatrix(4, 12)
	mn2 := NewMatrix(9, 3)

	result := mn2.MulMxN(nil, mn)

	if result != nil {
		t.Errorf("Nil not returned for bad matrix multiplication, got %v instead", result)
	}
}

func TestMxNMul(t *testing.T) {
	m := Mat3{2, 4, 6, 1, 9, 12, 7, 4, 3}
	mn := NewBackedMatrix(m[:], 3, 3)

	correct := m.Mul(15)
	correctMN := NewBackedMatrix(correct[:], 3, 3)

	result := mn.Mul(nil, 15)

	if !correctMN.ApproxEqualThreshold(result, 1e-4) {
		t.Errorf("Scaling a matrix produces weird results got: %v, expected: %v", result, correct)
	}
}
