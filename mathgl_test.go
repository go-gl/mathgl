package mathgl_test

import (
	"github.com/Jragonmiris/mathgl"
	"math"
	"testing"
)

func TestScalar(t *testing.T) {
	one := mathgl.MakeScalar(1, mathgl.FLOAT32)
	if one == nil {
		t.Fatalf("Couldn't make a scalar")
	}

	if math.Abs(float64(one.(mathgl.ScalarFloat32))-float64(1)) > .0000001 {
		t.Errorf("Scalar not set correctly")
	}

	if one.Type() != mathgl.FLOAT32 {
		t.Errorf("Scalar not of correct type after being made")
	}

	alsoOne := mathgl.MakeScalar(1, mathgl.FLOAT32)
	if !one.Equal(alsoOne) {
		t.Fatalf("One doesn't equal one, or equal method failed") // If equal isn't working, that's REALLY bad for the rest of the tests, hence fatal
	}

	pointZeroOne := mathgl.MakeScalar(.01, mathgl.FLOAT32)
	onePointZeroOne := one.Add(pointZeroOne)
	if one.Equal(onePointZeroOne) || !onePointZeroOne.Equal(mathgl.MakeScalar(1.01, mathgl.FLOAT32)) {
		t.Errorf("Addition failed")
	}

	oneAgain := onePointZeroOne.Sub(pointZeroOne)
	if !oneAgain.Equal(one) {
		t.Errorf("Sub failed")
	}

	if five := mathgl.MakeScalar(5, mathgl.FLOAT32); !five.Equal(one.Mul(five)) {
		t.Errorf("Multiplcation failed")
	}

	two := mathgl.MakeScalar(2, mathgl.FLOAT32)

	if !one.Equal(two.Div(two)) {
		t.Errorf("Division failed")
	}

	inSlice := []interface{}{int(1), float64(1.0), float32(1.01), uint32(2)}
	slice := mathgl.ScalarSlice(inSlice, mathgl.FLOAT32)

	if slice == nil || len(slice) < 4 {
		t.Fatalf("Slice not the correct length or does not exist after conversion %v", slice)
	}

	for _, el := range slice {
		if el == nil || el.Type() != mathgl.FLOAT32 {
			t.Fatalf("Making a scalar slice failed")
		}
	}

	if !slice[0].Equal(one) || !one.Equal(slice[1]) || !onePointZeroOne.Equal(slice[2]) || !two.Equal(slice[3]) {
		t.Errorf("Making a slice returned incorrect values")
	}

	oneInt := mathgl.MakeScalar(1.0, mathgl.INT32)
	anotherOneInt := mathgl.MakeScalar(int32(1), mathgl.INT32)
	if !oneInt.Equal(anotherOneInt) {
		t.Errorf("Equality fails for integers")
	}

	defer func() { recover() }() // Ignore the panic we're about to cause
	one.Equal(oneInt)

	t.Errorf("Did not panic on attempt to perform equality on differently typed scalars")
}

func TestVecSetGet(t *testing.T) {
	v1 := mathgl.NewVector(mathgl.FLOAT32)
	if v1 == nil {
		t.Fatalf("Failed to create new vector")
	}

	input := []mathgl.Scalar{mathgl.ScalarInt32(1), mathgl.ScalarFloat32(2), mathgl.ScalarFloat32(3), mathgl.ScalarFloat32(4), mathgl.ScalarFloat32(5)}

	if err := v1.AddElements(input); err == nil {
		t.Errorf("Added list with bad element to vector")
	}

	input[0] = mathgl.ScalarFloat32(1)

	if err := v1.AddElements(input); err != nil {
		t.Fatalf("Failed to set vector with correct list")
	}

	if a := v1.GetElement(0); a == nil || float32(a.(mathgl.ScalarFloat32))-float32(1.) > 0.000001 {
		t.Errorf("Didn't get/set correct element of vector")
	}

	v2, e := mathgl.VectorOf(input, mathgl.FLOAT32)
	if v2 == nil || e != nil {
		t.Fatalf("VectorOf failed on good input")
	}

	if a, b := v1.GetElement(2), v2.GetElement(2); a == nil || b == nil || float32(a.(mathgl.ScalarFloat32))-float32(b.(mathgl.ScalarFloat32)) > .000000001 {
		t.Errorf("Two vectors not the same despite being made from same list")
	}

	if !v1.Equal(*v2) { // We should have checked if this was equal in the last step. So if this fails equal is PROBABLY bad
		t.Errorf("Vectors are not equal or equal function failed, v1: %v v2: %v", v1, v2)
	}

	if err := v1.SetElement(25, mathgl.ScalarFloat32(1)); err == nil {
		t.Errorf("Set out of bounds vector element")
	}

	if err := v1.SetElement(-3, mathgl.ScalarFloat32(1)); err == nil {
		t.Errorf("Set out of bounds vector element")
	}

	if err := v1.SetElement(4, mathgl.ScalarFloat32(42)); err != nil {
		t.Fatalf("Didn't set in-bounds vector element")
	}

	if a := v1.GetElement(4); math.Abs(float64(float32(a.(mathgl.ScalarFloat32))-float32(42))) > .0000001 {
		t.Errorf("Did not correctly set single-in bounds vector element")
	}

	if a := v2.GetElement(4); math.Abs(float64(float32(a.(mathgl.ScalarFloat32))-float32(42))) < .0000001 {
		t.Errorf("Changing one vector changed another")
	}

	if v1.Equal(*v2) {
		t.Errorf("Vectors are equal despite changing v1, or equal is wrong, v1: %v v2: %v", v1, v2)
	}

}

func TestMatrixCreation(t *testing.T) {
	iden2 := mathgl.Identity(2, mathgl.FLOAT64)
	for i, el := range iden2.AsSlice() {
		if (i == 0 || i == 3) && math.Abs(float64(el.(mathgl.ScalarFloat64))-float64(1)) > .000001 {
			t.Errorf("Diagonals not 1 in 2x2 identity el: %v", el)
		} else if (i == 1 || i == 2) && math.Abs(float64(el.(mathgl.ScalarFloat64))-float64(0)) > .000001 {
			t.Errorf("Off-diagonals not 0 in 2x2 identity el: %v", el)
		}
	}

	lopsided := [][]mathgl.Scalar{mathgl.ScalarSlice([]interface{}{1, 2}, mathgl.INT32), mathgl.ScalarSlice([]interface{}{3, 4}, mathgl.INT32)}
	rowMat, err := mathgl.MatrixFromRows(lopsided, mathgl.INT32)
	colMat, err2 := mathgl.MatrixFromCols(lopsided, mathgl.INT32)

	if rowMat == nil || err != nil {
		t.Fatalf("MatrixFromRows failed")
	}

	if colMat == nil || err2 != nil {
		t.Fatalf("MatrixFromCols failed")
	}

	rowSlice := rowMat.AsSlice()
	colSlice := colMat.AsSlice()

	if 1 != rowSlice[0].Int32() || 2 != rowSlice[1].Int32() || 3 != rowSlice[2].Int32() || 4 != rowSlice[3].Int32() {
		t.Errorf("Matrix from rows did not order elements correctly %v", rowSlice)
	}

	if 1 != colSlice[0].Int32() || 2 != colSlice[2].Int32() || 3 != colSlice[1].Int32() || 4 != colSlice[3].Int32() {
		t.Errorf("Matrix from cols did not order elements correctly %v", colSlice)
	}

	askew := [][]mathgl.Scalar{mathgl.ScalarSlice([]interface{}{1, 2, 3}, mathgl.INT32), mathgl.ScalarSlice([]interface{}{4, 5, 6}, mathgl.INT32)}
	row2, err := mathgl.MatrixFromRows(askew, mathgl.INT32)
	col2, err2 := mathgl.MatrixFromCols(askew, mathgl.INT32)

	if row2 == nil || err != nil {
		t.Fatalf("MatrixFromRows failed on non-square matrix")
	}

	if col2 == nil || err2 != nil {
		t.Fatalf("MatrixFromCols failed on non-square matrix")
	}

	rowSlice = row2.AsSlice()
	colSlice = col2.AsSlice()

	if 1 != rowSlice[0].Int32() || 2 != rowSlice[1].Int32() || 3 != rowSlice[2].Int32() || 4 != rowSlice[3].Int32() || 5 != rowSlice[4].Int32() || 6 != rowSlice[5].Int32() {
		t.Errorf("Matrix from rows did not order elements correctly %v", rowSlice)
	}

	if 1 != colSlice[0].Int32() || 2 != colSlice[2].Int32() || 3 != colSlice[4].Int32() || 4 != colSlice[1].Int32() || 5 != colSlice[3].Int32() || 6 != colSlice[5].Int32() {
		t.Errorf("Matrix from cols did not order elements correctly %v", colSlice)
	}

	fromSlice, err := mathgl.MatrixFromSlice(mathgl.ScalarSlice([]interface{}{1, 2, 3, 4, 5, 6}, mathgl.INT32), mathgl.INT32, 2, 3)
	if fromSlice == nil || err != nil {
		t.Fatalf("Making a matrix from a slice failed")
	}
	testSlice := fromSlice.AsSlice()

	for i := range testSlice {
		if testSlice[i].Int32() != rowSlice[i].Int32() {
			t.Errorf("Matrix from slice not equal to one from rows, though it should be")
		}
	}

	if !fromSlice.Equal(*row2) {
		t.Errorf("Equal gives false negative")
	}

	if !row2.Equal(*fromSlice) {
		t.Errorf("Equal not transitive (or gives false negative transitively)")
	}

	if col2.Equal(*fromSlice) {
		t.Errorf("Equal gives false positive")
	}
}

func TestVecMath(t *testing.T) {
	vec1, _ := mathgl.VectorOf(mathgl.ScalarSlice([]interface{}{1, 2, 0, 1}, mathgl.FLOAT64), mathgl.FLOAT64)
	vec2, _ := mathgl.VectorOf(mathgl.ScalarSlice([]interface{}{0, 9, 1, 2.35}, mathgl.FLOAT64), mathgl.FLOAT64)
	eq, _ := mathgl.VectorOf(mathgl.ScalarSlice([]interface{}{1, 11, 1, 3.35}, mathgl.FLOAT64), mathgl.FLOAT64)

	if sum := vec1.Add(*vec2); !eq.Equal(sum) {
		t.Fatalf("Addition not working properly %v", sum)
	}

	if sum := vec2.Add(*vec1); !eq.Equal(sum) {
		t.Errorf("Addition not transitive %v", sum)
	}

	if diff := eq.Sub(*vec1); !vec2.Equal(diff) {
		t.Errorf("Subtraction fails %v", diff)
	}

	if diff := eq.Sub(*vec2); !vec1.Equal(diff) {
		t.Errorf("Subtraction fails %v", diff)
	}

	if dot := vec1.Dot(*vec2); math.Abs(dot.Fl64()-20.35) > .00000001 {
		t.Fatalf("Dot product produces incorrect answer %v", dot)
	}

	if dot := vec2.Dot(*vec1); math.Abs(dot.Fl64()-20.35) > .00000001 {
		t.Fatalf("Dot product intransitive %v", dot)
	}

	if dot := vec2.Dot(*vec2); math.Abs(dot.Fl64()-87.5225) > .00000001 {
		t.Fatalf("Dot product failed %v", dot)
	}

	length := vec2.Len()
	if math.Abs(length-9.35534606522) > 1e-10 {
		t.Fatalf("Length is incorrect %f", length)
	}

	norm := vec2.Normalize()

	if math.Abs(norm.GetElement(0).Fl64()) > 1e-10 || math.Abs(norm.GetElement(1).Fl64()-0.96201679096) > 1e-10 || math.Abs(norm.GetElement(2).Fl64()-0.10689075455) > 1e-10 ||
		math.Abs(norm.GetElement(3).Fl64()-0.25119327319) > .00001 {
		t.Errorf("Normalization of vector failed %v", norm)
	}

	zero, _ := mathgl.VectorOf(mathgl.ScalarSlice([]interface{}{0, 0, 0, 0, 0, 0, 0}, mathgl.FLOAT64), mathgl.FLOAT64)

	if !zero.Equal(zero.Normalize()) {
		t.Errorf("Normalization of zero vector changes vector")
	}

	cr1, _ := mathgl.VectorOf(mathgl.ScalarSlice([]interface{}{5, 15.7, 2}, mathgl.FLOAT64), mathgl.FLOAT64)
	cr2, _ := mathgl.VectorOf(mathgl.ScalarSlice([]interface{}{3, 0, -.2}, mathgl.FLOAT64), mathgl.FLOAT64)

	ver, _ := mathgl.VectorOf(mathgl.ScalarSlice([]interface{}{-3.14, 7, -47.1}, mathgl.FLOAT64), mathgl.FLOAT64)

	if cross := cr1.Cross(*cr2); !ver.Equal(cross) {
		t.Errorf("Cross product is wrong %v", cross)
	}

	ver2 := ver.ScalarMul(mathgl.MakeScalar(-1, mathgl.FLOAT64))
	if math.Abs(ver2.GetElement(0).Fl64()-3.14) > 1e-10 || math.Abs(ver2.GetElement(1).Fl64()-(-7)) > 1e-10 || math.Abs(ver2.GetElement(2).Fl64()-47.1) > 1e-10 {
		t.Errorf("Scalar multiply failed to work %v", ver2)
	}

	// u x v = -(v x u)
	if cross := cr2.Cross(*cr1); !ver2.Equal(cross) {
		t.Errorf("Cross product failed %v", cross)
	}
}

// Tests conversion between vec/mat/scalar (and later, AsArray)
func TestConversion(t *testing.T) {
	conv,_ := mathgl.VectorOf([]mathgl.Scalar{mathgl.MakeScalar(1, mathgl.INT32)}, mathgl.INT32)
	if conv.ToScalar().Int32() != 1 {
		t.Errorf("Vector's ToScalar doesn't work")
	}
	
	mat,_ := conv.AsMatrix(false)
	comp,_ := mathgl.MatrixFromSlice([]mathgl.Scalar{mathgl.MakeScalar(1, mathgl.INT32)}, mathgl.INT32,1,1)
	
	if !mat.Equal(*comp) {
		t.Errorf("Matrix conversion failed")
	}
	
	if !mat.AsVector().Equal(*conv) {
		t.Errorf("Matrix conversion to vector failed")
	}
	
	if mat.ToScalar().Int32() != 1 {
		t.Errorf("Matrix ToScalar doesn't work")
	}
	
	// Needs more tests, but should suffice minimally for now
}

func TestMatrixMath(t *testing.T) {
	i3 := mathgl.Identity(3, mathgl.FLOAT64)
	
	mul3,_ := mathgl.MatrixFromRows([][]mathgl.Scalar{
		mathgl.ScalarSlice([]interface{}{2,0,0}, mathgl.FLOAT64),
		mathgl.ScalarSlice([]interface{}{0,2,0}, mathgl.FLOAT64),
		mathgl.ScalarSlice([]interface{}{0,0,2}, mathgl.FLOAT64)}, mathgl.FLOAT64)
		
	if sum := i3.Add(i3); !mul3.Equal(sum) {
		t.Errorf("Adding fails %v %v", sum, mul3)
	}
	
	if diff := mul3.Sub(i3); !i3.Equal(diff) {
		t.Errorf("Subtracting fails %v %v", diff, i3)
	}
	
	if prod := i3.ScalarMul(mathgl.MakeScalar(2,mathgl.FLOAT64)); !prod.Equal(*mul3) {
		t.Errorf("Scaling fails %v %v", prod, mul3)
	}
	
	if prod := i3.Mul(*mul3); !mul3.Equal(prod) {
		t.Fatalf("Multiplication by identity failed %v", prod)
	}
	
	mat3x2,_ := mathgl.MatrixFromRows([][]mathgl.Scalar{
		mathgl.ScalarSlice([]interface{}{1,0}, mathgl.FLOAT64),
		mathgl.ScalarSlice([]interface{}{1,2}, mathgl.FLOAT64),
		mathgl.ScalarSlice([]interface{}{0,1}, mathgl.FLOAT64)}, mathgl.FLOAT64)
		
	mat2x4,_ := mathgl.MatrixFromRows([][]mathgl.Scalar{
		mathgl.ScalarSlice([]interface{}{1,0,3,1}, mathgl.FLOAT64),
		mathgl.ScalarSlice([]interface{}{1,2,1,4}, mathgl.FLOAT64)}, mathgl.FLOAT64)
		
	out3x4,_ := mathgl.MatrixFromRows([][]mathgl.Scalar{
		mathgl.ScalarSlice([]interface{}{1,0,3,1}, mathgl.FLOAT64),
		mathgl.ScalarSlice([]interface{}{3,4,5,9}, mathgl.FLOAT64),
		mathgl.ScalarSlice([]interface{}{1,2,1,4}, mathgl.FLOAT64)}, mathgl.FLOAT64)
		
	if prod := mat3x2.Mul(*mat2x4); !out3x4.Equal(prod) {
		t.Fatalf("Non-square product failed %v", prod)
	}
	
	vec,_ := mathgl.VectorOf(mathgl.ScalarSlice([]interface{}{1,5,1.5}, mathgl.FLOAT64), mathgl.FLOAT64)
	result,_ := mathgl.MatrixFromRows([][]mathgl.Scalar{
		mathgl.ScalarSlice([]interface{}{6,11.5}, mathgl.FLOAT64)}, mathgl.FLOAT64)
		
	if prod := vec.Mul(*mat3x2); !result.Equal(prod) {
		t.Fatalf("Vector/Matrix Multiplication failed")
	}
	
	leftMul,_ := mathgl.MatrixFromRows([][]mathgl.Scalar{
		mathgl.ScalarSlice([]interface{}{2,1.5,0}, mathgl.FLOAT64),
		mathgl.ScalarSlice([]interface{}{1,2,0}, mathgl.FLOAT64),
		mathgl.ScalarSlice([]interface{}{7,0,2}, mathgl.FLOAT64)}, mathgl.FLOAT64)
	
	result,_ = mathgl.MatrixFromRows([][]mathgl.Scalar{
		[]mathgl.Scalar{mathgl.MakeScalar(9.5,mathgl.FLOAT64)}, 
		[]mathgl.Scalar{mathgl.MakeScalar(11,mathgl.FLOAT64)}, 
		[]mathgl.Scalar{mathgl.MakeScalar(10,mathgl.FLOAT64)}}, mathgl.FLOAT64)
		
	if prod := leftMul.Mul(*vec); !result.Equal(prod) {
		t.Fatalf("Matrix/Vector Multiplication failed")
	}
	
	// TODO: Test transpose
	
}

// My determinant method is slow. Disable this test for quick testing
func TestInvertible(t *testing.T) {
	invmat,_ := mathgl.MatrixFromRows([][]mathgl.Scalar {
		mathgl.ScalarSlice([]interface{}{1, 3, 2, 5},mathgl.FLOAT64),
		mathgl.ScalarSlice([]interface{}{1, 5, 6, 2},mathgl.FLOAT64),
		mathgl.ScalarSlice([]interface{}{1.5, 2.5, 1, 2},mathgl.FLOAT64),
		mathgl.ScalarSlice([]interface{}{2, 4, 7, 1},mathgl.FLOAT64)}, mathgl.FLOAT64)
		
	if math.Abs(invmat.Det() - 71) > 1e-7 {
		t.Fatalf("Determinant incorrect")
	}
	
	// This is just too hard in general to hard code the answer (precision needed is too great), so I'm using a for loop
	slice := invmat.Transpose().AsSlice()
	for i,el := range slice {
		slice[i] = mathgl.MakeScalar(float64(1)/float64(71) * el.Fl64(), mathgl.FLOAT64)
	}
	
	realInv,_ := mathgl.MatrixFromSlice(slice, mathgl.FLOAT64, 4, 4)
	
	if genInv := invmat.Inverse(); !realInv.Equal(genInv) {
		t.Errorf("Matrix did not generate correct inverse\n %v,\n %v", realInv, genInv )
	}
	
	noninvmat,_ := mathgl.MatrixFromRows([][]mathgl.Scalar {
		mathgl.ScalarSlice([]interface{}{1, 2},mathgl.FLOAT64),
		mathgl.ScalarSlice([]interface{}{2, 4},mathgl.FLOAT64)}, mathgl.FLOAT64)
	
	if math.Abs(noninvmat.Det()) > 1e-7 {
		t.Errorf("Non-invertible matrix with det > 0")
	}
	
	if m := mathgl.NewMatrix(0,0,mathgl.FLOAT64); m.Equal(noninvmat.Inverse()) {
		t.Errorf("Non-invertible matrix calling Inverse does not return matrix zero-type")
	}
}