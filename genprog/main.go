// Copyright 2012 The go-gl/mathgl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Println("Making vectorf.go")
	vecs := GenVec()
	//fmt.Println(vecs)
	vecf, err := os.Create("../mgl32/vector.go")
	if err != nil {
		panic(err)
	}
	defer vecf.Close()

	_, err = vecf.Write([]byte(vecs))
	if err != nil {
		panic(err)
	}

	//fmt.Println("Making matrixf.go")
	mats := GenMat()
	//fmt.Println(mats)
	matf, err := os.Create("../mgl32/matrix.go")
	if err != nil {
		panic(err)
	}
	defer matf.Close()

	_, err = matf.Write([]byte(mats))
	if err != nil {
		panic(err)
	}
	//fmt.Println(mats)
	//fmt.Println("Done")
}

func GenVec() (s string) {
	vecs := `// Copyright 2012 The go-gl/mathgl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mgl32

import(
	"math"
)

`

	for m := 2; m <= 4; m++ {
		vecs += GenVecDef(m)
	}

	vecs += "\n"
	for m := 2; m <= 4; m++ {
		vecs += GenVecAdd(m)
	}

	for m := 2; m <= 4; m++ {
		vecs += GenVecSub(m)
	}

	for m := 2; m <= 4; m++ {
		vecs += GenVecMul(m)
	}

	for m := 2; m <= 4; m++ {
		vecs += GenVecDot(m)
	}

	for m := 2; m <= 4; m++ {
		vecs += GenVecLen(m)
	}

	for m := 2; m <= 4; m++ {
		vecs += GenVecNormalize(m)
	}

	vecs += GenVecCross()

	for m := 2; m <= 4; m++ {
		vecs += GenVecEq(m)
	}

	for m := 2; m <= 4; m++ {
		vecs += GenVecThresholdEq(m)
	}

	for m := 2; m <= 4; m++ {
		vecs += GenVecFuncEq(m)
	}

	/*for m := 2; m <= 4; m++ {
		for o := 2; o <= 4; o++ {
			vecs += GenVecMatMul(m,o)
		}
	}*/

	return vecs
}

func GenVecCross() string {
	header := `// The vector cross product is an operation only defined on 3D vectors. It is equivalent to
// Vec3{v1[1]*v2[2]-v1[2]*v2[1], v1[2]*v2[0]-v1[0]*v2[2], v1[0]*v2[1] - v1[1]*v2[0]}. 
// Another interpretation is |v1||v2|sin(theta) where there is the angle between v1 and v2.
//
// Technically, a generalized cross product exists as an "(N-1)ary" operation 
// (that is, the 4D cross product requires 3 4D vectors). But the binary
// 3D (and 7D) cross product is the most important. It can be considered 
// the area of a parallelograph with sides v1 and v2.
//
// Like the dot product, the cross product is roughly a measure of directionality. 
// Two normalized perpendicular vectors will return a value of
// 1.0 or 0.0 and two parallel vectors will return a value of 0. 
// The cross product is "anticommutative" meaning v1.Cross(v2) = -v2.Cross(v1),
// this property can be useful to know when finding normals, 
// as taking the wrong cross product can lead to the opposite normal of the one you want.
`
	return header + "func (v1 Vec3) Cross(v2 Vec3) Vec3 {\n\treturn Vec3{v1[1]*v2[2]-v1[2]*v2[1], v1[2]*v2[0]-v1[0]*v2[2], v1[0]*v2[1] - v1[1]*v2[0]}\n}\n\n"
}

func GenVecDef(m int) (s string) {
	return fmt.Sprintf("type Vec%d [%d]float32\n", m, m)
}

/*func GenVecMatMul(m, n, o int) (s string) {
	s = fmt.Sprintf("func (v1 %s) Mul%s(v2 %s) %s {\n\treturn %s{", GenMatName(m,n), GenMatName(n,o), GenMatName(m,o), GenMatName(m,o))
	for j := 0; j < o; j++ {
		for i := 0; i < m; i++ {
			for k := 0; k < n; k++ {
				s += fmt.Sprintf("m1[%d] * m2[%d]", i+k*m, j*n+k)
				s += " + "
			}
			s = s[:len(s)-3]
			s += ", "
		}
	}
	s = s[:len(s)-2]
	s += "}\n}\n\n"
	return s
}*/

func GenVecAdd(m int) (s string) {
	s = `// Add performs element-wise addition between two vectors. It is equivalent to iterating
// over every element of v1 and adding the corresponding element of v2 to it.
`
	s += fmt.Sprintf("func (v1 %s) Add(v2 %s) %s {\n\treturn %s{", VecName(m), VecName(m), VecName(m), VecName(m))
	for i := 0; i < m; i++ {
		s += fmt.Sprintf("v1[%d] + v2[%d]", i, i)
		if i != m-1 {
			s += ","
		}
	}
	s += "}\n}\n\n"
	return s
}

func GenVecSub(m int) (s string) {
	s = `// Sub performs element-wise subtraction between two vectors. It is equivalent to iterating
// over every element of v1 and subtracting the corresponding element of v2 from it.
`
	s += fmt.Sprintf("func (v1 %s) Sub(v2 %s) %s {\n\treturn %s{", VecName(m), VecName(m), VecName(m), VecName(m))
	for i := 0; i < m; i++ {
		s += fmt.Sprintf("v1[%d] - v2[%d]", i, i)
		if i != m-1 {
			s += ","
		}
	}
	s += "}\n}\n\n"
	return s
}

func GenVecMul(m int) (s string) {
	s = `// Mul performs a scalar multiplication between the vector and some constant value
// c. This is equivalent to iterating over every vector element and multiplying by c.
`
	s += fmt.Sprintf("func (v1 %s) Mul(c float32) %s {\n\treturn %s{", VecName(m), VecName(m), VecName(m))
	for i := 0; i < m; i++ {
		s += fmt.Sprintf("v1[%d] * c", i)
		if i != m-1 {
			s += ","
		}
	}
	s += "}\n}\n\n"

	return s
}

func GenVecDot(m int) (s string) {
	s = `// Dot returns the dot product of this vector with another. There are multiple ways
// to describe this value. One is the multiplication of their lengths and cos(theta) where
// theta is the angle between the vectors: v1.v2 = |v1||v2|cos(theta).
//
// The other (and what is actually done) is the sum of the element-wise multiplication of all
// elements. So for instance, two Vec3s would yield v1.x * v2.x + v1.y * v2.y + v1.z * v2.z.
//
// This means that the dot product of a vector and itself is the square of its Len (within
// the bounds of floating points error).
//
// The dot product is roughly a measure of how closely two vectors are to pointing in the same
// direction. If both vectors are normalized, the value will be -1 for opposite pointing,
// one for same pointing, and 0 for perpendicular vectors.
`
	s += fmt.Sprintf("func (v1 %s) Dot(v2 %s) float32 {\n\treturn ", VecName(m), VecName(m))
	for i := 0; i < m; i++ {
		s += fmt.Sprintf("v1[%d] * v2[%d]", i, i)
		if i != m-1 {
			s += "+"
		}
	}
	s += "\n}\n\n"

	return s
}

func GenVecLen(m int) (s string) {
	s = `// Len returns the vector's length. Note that this is NOT the dimension of
// the vector (len(v)), but the mathematical length. This is equivalent to the square
// root of the sum of the squares of all elements. E.G. for a Vec2 it's 
// math.Hypot(v[0], v[1]).
`
	if m != 2 {
		s += fmt.Sprintf("func (v1 %s) Len() float32 {\n\treturn float32(math.Sqrt(float64(", VecName(m))
		for i := 0; i < m; i++ {
			s += fmt.Sprintf("v1[%d] * v1[%d]", i, i)
			if i != m-1 {
				s += "+"
			}
		}
		s += ")))\n}\n\n"
	} else {
		s += fmt.Sprintf("func (v1 %s) Len() float32 {\n\treturn float32(math.Hypot(float64(v1[0]), float64(v1[1])))\n}\n\n", VecName())
	}
	return s
}

func GenVecNormalize(m int) (s string) {
	s = `// Normalize normalizes the vector. Normalization is (1/|v|)*v,
// making this equivalent to v.Scale(1/v.Len()). If the len is 0.0,
// this function will return an infinite value for all elements due
// to how floating point division works in Go (n/0.0 = math.Inf(Sign(n))).
//
// Normalization makes a vector's Len become 1.0 (within the margin of floating point error),
// while maintaining its directionality.
//
// (Can be seen here: http://play.golang.org/p/Aaj7SnbqIp )
`
	s += fmt.Sprintf("func (v1 %s) Normalize() %s {\n\tl := 1.0/v1.Len()", VecName(m), VecName(m))
	s += fmt.Sprintf("))\n\treturn %s{", VecName(m))
	for i := 0; i < m; i++ {
		s += fmt.Sprintf("float32(float64(v1[%d]) * l)", i)
		if i != m-1 {
			s += ","
		}
	}
	s += "}\n}\n\n"
	return s
}

func GenVecEq(m int) (s string) {
	s = `// ApproxEqual takes in a vector and does an element-wise
// approximate float comparison as if FloatEqual had been used
`
	s += fmt.Sprintf("func (v1 %s) ApproxEqual(v2 %s) bool {\n\t", VecName(m), VecName(m))

	s += "for i := range v1 {\n\t\t"
	s += "if !FloatEqual(v1[i],v2[i]) {\n\t\t\t"
	s += "return false\n\t\t"
	s += "}\n\t}\n\t"
	s += "return true\n}\n\n"

	return s
}

func GenVecThresholdEq(m int) (s string) {
	s = `// ApproxThresholdEq takes in a threshold for comparing two floats, and uses it to do an
// element-wise comparison of the vector to another.
`
	s += fmt.Sprintf("func (v1 %s) ApproxEqualThreshold(v2 %s, threshold float32) bool {\n\t", VecName(m), VecName(m))

	s += "for i := range v1 {\n\t\t"
	s += "if !FloatEqualThreshold(v1[i],v2[i], threshold) {\n\t\t\t"
	s += "return false\n\t\t"
	s += "}\n\t}\n\t"
	s += "return true\n}\n\n"

	return s
}

func GenVecFuncEq(m int) (s string) {
	s = `// ApproxFuncEq takes in a func that compares two floats, and uses it to do an element-wise
// comparison of the vector to another. This is intended to be used with FloatEqualFunc
`
	s += fmt.Sprintf("func (v1 %s) ApproxFuncEqual(v2 %s, eq func(float32,float32) bool) bool {\n\t", VecName(m), VecName(m))

	s += "for i := range v1 {\n\t\t"
	s += "if !eq(v1[i],v2[i]) {\n\t\t\t"
	s += "return false\n\t\t"
	s += "}\n\t}\n\t"
	s += "return true\n}\n\n"

	return s
}

func VecName(m int) (s string) {
	return fmt.Sprintf("Vec%d", m)
}

func GenMat() string {
	mats := `// Copyright 2012 The go-gl/mathgl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mgl32

import(
	"math"
)

`

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			mats += GenMatDef(m, n)
		}
	}
	mats += "\n"

	for m := 2; m <= 4; m++ {
		mats += GenMatIden(m)
	}

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			mats += GenMatAdd(m, n)
		}
	}

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			mats += GenMatSub(m, n)
		}
	}

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			mats += GenScalarMul(m, n)
		}
	}

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			for o := 1; o <= 4; o++ {
				mats += GenMatMul(m, n, o)
			}
		}
	}

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			mats += GenTranspose(m, n)
		}
	}

	for m := 2; m <= 4; m++ {
		mats += GenDet(m)
	}

	for m := 2; m <= 4; m++ {
		mats += GenInv(m)
	}

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			mats += GenMatEq(m, n)
		}
	}

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			mats += GenMatThresholdEq(m, n)
		}
	}

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			mats += GenMatFuncEq(m, n)
		}
	}

	//fmt.Println(mats)
	return mats
}

func GenMatIden(m int) (s string) {
	s = `// Ident<N> returns the NxN identity matrix.
// The identity matrix is a square matrix with the value 1 on its
// diagonals. The characteristic property of the identity matrix is that
// any matrix multiplied by it is itself. (MI = M; IN = N)
`
	s += fmt.Sprintf("func Ident%d() %s {\n\treturn %s{", m, GenMatName(m, m), GenMatName(m, m))

	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			if i != j {
				s += "0"
			} else {
				s += "1"
			}

			if i != m-1 || j != m-1 {
				s += ","
			}
		}
	}

	s += "}\n}\n\n"

	return s
}

func GenMatDef(m, n int) (s string) {
	return fmt.Sprintf("type %s [%d]float32\n", GenMatName(m, n), m*n)
}

func GenMatAdd(m, n int) (s string) {
	s = `// Add performs an element-wise addition of two matrices, this is
// equivalent to iterating over every element of m1 and adding the corresponding value of m2.
`
	s += fmt.Sprintf("func (m1 %s) Add(m2 %s) %s {\n\treturn %s {", GenMatName(m, n), GenMatName(m, n), GenMatName(m, n), GenMatName(m, n))
	for i := 0; i < m*n; i++ {
		s += fmt.Sprintf("m1[%d] + m2[%d]", i, i)
		if i != (m*n)-1 {
			s += ","
		}
	}
	s += "}\n}\n\n"
	return s
}

func GenMatSub(m, n int) (s string) {
	s = `// Sub performs an element-wise subtraction of two matrices, this is
// equivalent to iterating over every element of m1 and subtracting the corresponding value of m2.
`
	s += fmt.Sprintf("func (m1 %s) Sub(m2 %s) %s {\n\treturn %s {", GenMatName(m, n), GenMatName(m, n), GenMatName(m, n), GenMatName(m, n))
	for i := 0; i < m*n; i++ {
		s += fmt.Sprintf("m1[%d] - m2[%d]", i, i)
		if i != (m*n)-1 {
			s += ","
		}
	}
	s += "}\n}\n\n"
	return s
}

func GenScalarMul(m, n int) (s string) {
	s = `// Mul performs a scalar multiplcation of the matrix. This is equivalent to iterating
// over every element of the matrix and multiply it by c.
`
	s += fmt.Sprintf("func (m1 %s) Mul(c float32) %s {\n\treturn %s{", GenMatName(m, n), GenMatName(m, n), GenMatName(m, n))
	for i := 0; i < m*n; i++ {
		s += fmt.Sprintf("m1[%d] *c", i)
		if i != (m*n)-1 {
			s += ","
		}
	}
	s += "}\n}\n\n"
	return s
}

func GenMatMul(m, n, o int) (s string) {
	s = `// Mul<Dim> performs a "matrix product" between this matrix
// and another of the given dimension. For any two matrices of dimensionality
// MxN and NxO, the result will be MxO. For instance, Mat4 multiplied using
// Mul4x2 will result in a Mat4x2.
`
	s += "func " + "(m1 " + GenMatName(m, n) + ") Mul" + fmt.Sprintf("%d", n)

	if n != o {
		s += fmt.Sprintf("x%d", o)
	}

	s += "(m2 " + GenMatName(n, o) + ") " + GenMatName(m, o) + " {\n\treturn " + GenMatName(m, o) + "{"
	for j := 0; j < o; j++ { // For each element of the output array
		for i := 0; i < m; i++ {
			for k := 0; k < n; k++ { // For each element of the vector we're multiplying
				s += fmt.Sprintf("m1[%d] * m2[%d]", i+k*m, j*n+k)
				s += " + "
			}
			s = s[:len(s)-3]
			s += ", "
		}
	}
	s = s[:len(s)-2]
	s += "}\n}\n\n"
	return s
}

func GenTranspose(m, n int) (s string) {
	s = `// Transpose produces the transpose of this matrix. For any MxN matrix
// the transpose is an NxM matrix with the rows swapped with the columns. For instance
// the transpose of the Mat3x2 is a Mat2x3 like so:
//
//    [[a b]]    [[a c e]]
//    [[c d]] =  [[b d f]]
//    [[e f]]    
`
	s += fmt.Sprintf("func (m1 %s) Transpose() %s {\n\treturn %s{", GenMatName(m, n), GenMatName(n, m), GenMatName(n, m))

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			s += fmt.Sprintf("m1[%d]", i+j*n)
			if i != m-1 || j != n-1 {
				s += ","
			}
		}
	}
	s += "}\n}\n\n"

	return s
}

/*func GenDet(m,n int) (s string) {
	s = fmt.Sprintf("func (m1 %s) Mul() float32 {\n\treturn %s{", GenMatName(m,n), GenMatName(n,o), GenMatName(m,o), GenMatName(m,o))
}*/

func GenMatName(m, n int) string {
	if m == 1 {
		return fmt.Sprintf("Vec%d", n)
	}

	if n == 1 {
		return fmt.Sprintf("Vec%d", m)
	}

	s := fmt.Sprintf("Mat%d", m)

	if m != n {
		s += fmt.Sprintf("x%d", n)
	}

	return s
}

func GenDet(m int) string {
	s := `// The determinant of a matrix is a measure of a square matrix's
// singularity and invertability, among other things. In this library, the
// determinant is hard coded based on pre-computed cofactor expansion, and uses
// no loops. Of course, the addition and multiplication must still be done.
`
	s += fmt.Sprintf("func (m %s) Det() float32 {\n\treturn ", GenMatName(m, m))

	switch m {
	case 2:
		s += "m[0] * m[2] - m[1] * m[3]"
	case 3:
		s += "m[0]*m[4]*m[8] + m[3] * m[7] * m[2] + m[6] * m[1] * m[5] - m[6] * m[4] * m[2] - m[3] * m[1] * m[8] - m[0] * m[7] * m[5]"
	case 4:
		s += "m[0]*m[5]*m[10]*m[15] - m[0]*m[5]*m[11]*m[14] - m[0]*m[6]*m[9]*m[15] + m[0]*m[6]*m[11]*m[13] + m[0]*m[7]*m[9]*m[14] - m[0]*m[7]*m[10]*m[13]" +
			" - m[1]*m[4]*m[10]*m[15] + m[1]*m[4]*m[11]*m[14] + m[1]*m[6]*m[8]*m[15] - m[1]*m[6]*m[11]*m[12] - m[1]*m[7]*m[8]*m[14] + m[1]*m[7]*m[10]*m[12]" +
			" + m[2]*m[4]*m[9]*m[15] - m[2]*m[4]*m[11]*m[13] - m[2]*m[5]*m[8]*m[15] + m[2]*m[5]*m[11]*m[12] + m[2]*m[7]*m[8]*m[13] - m[2]*m[7]*m[9]*m[12]" +
			" - m[3]*m[4]*m[9]*m[14] + m[3]*m[4]*m[10]*m[13] + m[3]*m[5]*m[8]*m[14] - m[3]*m[5]*m[10]*m[12] - m[3]*m[6]*m[8]*m[13] + m[3]*m[6]*m[9]*m[12]"
	}
	s += "\n}\n\n"

	return s
}

func GenInv(m int) string {
	s := `// Inv computes the inverse of a square matrix. An inverse is a square matrix such that when multiplied by the
// original, yields the identity.
//
// M_inv * M = M * M_inv = I
//
// In this library, the math is precomputed, and uses no loops, though the multiplications, additions, determinant calculation, and scaling
// are still done. This can still be (relatively) expensive for a 4x4.
//
// This function does not check the determinant to see if the matrix is invertible. 
// If the determinant is 0.0, the value of all elements will be
// infinite. (See here for why: http://play.golang.org/p/Aaj7SnbqIp ) 
// Therefore, if the program really cares, it should check the determinant first.
// In the future, an alternate function may be written which takes in a pre-computed determinant. 
`
	s += fmt.Sprintf("func (m %s) Inv() %s {\n\t", GenMatName(m, m), GenMatName(m, m))
	s += "det := m.Det()\n\t if FloatEqual(det,float32(0.0)) { \n\t\t return " + GenMatName(m, m) + "{}\n\t}\n\t"
	s += "retMat := " + GenMatName(m, m) + "{"

	switch m {
	case 2:
		s += "m[3], -m[1], -m[2], m[0]"
	case 3:
		s += "m[4] * m[8] -m[5] * m[7] , m[2] * m[7] -m[1] * m[8] ,m[1] * m[5] -m[2] * m[4] ,m[5] * m[6] -m[3] * m[8] ,m[0] * m[8] -m[2] * m[6] ,m[2] * m[3] -m[0] * m[5] ,m[3] * m[7] -m[4] * m[6] ,m[1] * m[6] -m[0] * m[7] ,m[0] * m[4] -m[1] * m[3]"
	case 4:
		s += "-m[7] * m[10] * m[13] +m[6] * m[11] * m[13] +m[7] * m[9] * m[14] -m[5] * m[11] * m[14] -m[6] * m[9] * m[15] +m[5] * m[10] * m[15] ," +
			"m[3] * m[10] * m[13] -m[2] * m[11] * m[13] -m[3] * m[9] * m[14] +m[1] * m[11] * m[14] +m[2] * m[9] * m[15] -m[1] * m[10] * m[15] ," +
			"-m[3] * m[6] * m[13] +m[2] * m[7] * m[13] +m[3] * m[5] * m[14] -m[1] * m[7] * m[14] -m[2] * m[5] * m[15] +m[1] * m[6] * m[15] ," +
			"m[3] * m[6] * m[9] -m[2] * m[7] * m[9] -m[3] * m[5] * m[10] +m[1] * m[7] * m[10] +m[2] * m[5] * m[11] -m[1] * m[6] * m[11] ," +
			"m[7] * m[10] * m[12] -m[6] * m[11] * m[12] -m[7] * m[8] * m[14] +m[4] * m[11] * m[14] +m[6] * m[8] * m[15] -m[4] * m[10] * m[15] ," +
			"-m[3] * m[10] * m[12] +m[2] * m[11] * m[12] +m[3] * m[8] * m[14] -m[0] * m[11] * m[14] -m[2] * m[8] * m[15] +m[0] * m[10] * m[15] ," +
			" m[3] * m[6] * m[12] -m[2] * m[7] * m[12] -m[3] * m[4] * m[14] +m[0] * m[7] * m[14] +m[2] * m[4] * m[15] -m[0] * m[6] * m[15] ," +
			"-m[3] * m[6] * m[8] +m[2] * m[7] * m[8] +m[3] * m[4] * m[10] -m[0] * m[7] * m[10] -m[2] * m[4] * m[11] +m[0] * m[6] * m[11] ," +
			"-m[7] * m[9] * m[12] +m[5] * m[11] * m[12] +m[7] * m[8] * m[13] -m[4] * m[11] * m[13] -m[5] * m[8] * m[15] +m[4] * m[9] * m[15] ," +
			"m[3] * m[9] * m[12] -m[1] * m[11] * m[12] -m[3] * m[8] * m[13] +m[0] * m[11] * m[13] +m[1] * m[8] * m[15] -m[0] * m[9] * m[15] ," +
			"-m[3] * m[5] * m[12] +m[1] * m[7] * m[12] +m[3] * m[4] * m[13] -m[0] * m[7] * m[13] -m[1] * m[4] * m[15] +m[0] * m[5] * m[15] ," +
			"m[3] * m[5] * m[8] -m[1] * m[7] * m[8] -m[3] * m[4] * m[9] +m[0] * m[7] * m[9] +m[1] * m[4] * m[11] -m[0] * m[5] * m[11] ," +
			"m[6] * m[9] * m[12] -m[5] * m[10] * m[12] -m[6] * m[8] * m[13] +m[4] * m[10] * m[13] +m[5] * m[8] * m[14] -m[4] * m[9] * m[14] ," +
			"-m[2] * m[9] * m[12] +m[1] * m[10] * m[12] +m[2] * m[8] * m[13] -m[0] * m[10] * m[13] -m[1] * m[8] * m[14] +m[0] * m[9] * m[14] ," +
			"m[2] * m[5] * m[12] -m[1] * m[6] * m[12] -m[2] * m[4] * m[13] +m[0] * m[6] * m[13] +m[1] * m[4] * m[14] -m[0] * m[5] * m[14] ," +
			"-m[2] * m[5] * m[8] +m[1] * m[6] * m[8] +m[2] * m[4] * m[9] -m[0] * m[6] * m[9] -m[1] * m[4] * m[10] +m[0] * m[5] * m[10]"
	}

	s += "}\n\t return retMat.Mul(1/det)\n}\n\n"

	return s
}

func GenMatEq(m, n int) (s string) {
	s = `// ApproxEqual performs an element-wise approximate equality test between two matrices,
// as if FloatEqual had been used.
`
	s += fmt.Sprintf("func (m1 %s) ApproxEqual(m2 %s) bool {\n\t", GenMatName(m, n), GenMatName(m, n))

	s += "for i := range m1 {\n\t\t"
	s += "if !FloatEqual(m1[i],m2[i]) {\n\t\t\t"
	s += "return false\n\t\t"
	s += "}\n\t}\n\t"
	s += "return true\n}\n\n"

	return s
}

func GenMatThresholdEq(m, n int) (s string) {
	s = `// ApproxEqualThreshold performs an element-wise approximate equality test between two matrices
// with a given epsilon threshold, as if FloatEqualThreshold had been used.
`
	s += fmt.Sprintf("func (m1 %s) ApproxEqualThreshold(m2 %s, threshold float32) bool {\n\t", GenMatName(m, n), GenMatName(m, n))

	s += "for i := range m1 {\n\t\t"
	s += "if !FloatEqualThreshold(m1[i],m2[i], threshold) {\n\t\t\t"
	s += "return false\n\t\t"
	s += "}\n\t}\n\t"
	s += "return true\n}\n\n"

	return s
}

func GenMatFuncEq(m, n int) (s string) {
	s = `// ApproxEqualFunc performs an element-wise approximate equality test between two matrices
// with a given equality functions, intended to be used with FloatEqualFunc; although and comparison
// function may be used in practice.
`
	s += fmt.Sprintf("func (m1 %s) ApproxFuncEqual(m2 %s, eq func(float32,float32) bool) bool {\n\t", GenMatName(m, n), GenMatName(m, n))

	s += "for i := range m1 {\n\t\t"
	s += "if !eq(m1[i],m2[i]) {\n\t\t\t"
	s += "return false\n\t\t"
	s += "}\n\t}\n\t"
	s += "return true\n}\n\n"

	return s
}
