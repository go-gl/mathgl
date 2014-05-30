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
	vecs := "package mgl32\n\nimport(\n\t \"math\"\n)\n\n"

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
	return "func (v1 Vec3) Cross(v2 Vec3) Vec3 {\n\treturn Vec3{v1[1]*v2[2]-v1[2]*v2[1], v1[2]*v2[0]-v1[0]*v2[2], v1[0]*v2[1] - v1[1]*v2[0]}\n}\n\n"
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
	s = fmt.Sprintf("func (v1 %s) Add(v2 %s) %s {\n\treturn %s{", VecName(m), VecName(m), VecName(m), VecName(m))
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
	s = fmt.Sprintf("func (v1 %s) Sub(v2 %s) %s {\n\treturn %s{", VecName(m), VecName(m), VecName(m), VecName(m))
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
	s = fmt.Sprintf("func (v1 %s) Mul(c float32) %s {\n\treturn %s{", VecName(m), VecName(m), VecName(m))
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
	s = fmt.Sprintf("func (v1 %s) Dot(v2 %s) float32 {\n\treturn ", VecName(m), VecName(m))
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
	s = fmt.Sprintf("func (v1 %s) Len() float32 {\n\treturn float32(math.Sqrt(float64(", VecName(m))
	for i := 0; i < m; i++ {
		s += fmt.Sprintf("v1[%d] * v1[%d]", i, i)
		if i != m-1 {
			s += "+"
		}
	}
	s += ")))\n}\n\n"
	return s
}

func GenVecNormalize(m int) (s string) {
	s = fmt.Sprintf("func (v1 %s) Normalize() %s {\n\tl := 1.0/math.Sqrt(float64(", VecName(m), VecName(m))
	for i := 0; i < m; i++ {
		s += fmt.Sprintf("v1[%d] * v1[%d]", i, i)
		if i != m-1 {
			s += "+"
		}
	}
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
	s = fmt.Sprintf("func (v1 %s) ApproxEqual(v2 %s) bool {\n\t", VecName(m), VecName(m))

	s += "for i := range v1 {\n\t\t"
	s += "if !FloatEqual(v1[i],v2[i]) {\n\t\t\t"
	s += "return false\n\t\t"
	s += "}\n\t}\n\t"
	s += "return true\n}\n\n"

	return s
}

func GenVecThresholdEq(m int) (s string) {
	s = fmt.Sprintf("func (v1 %s) ApproxEqualThreshold(v2 %s, threshold float32) bool {\n\t", VecName(m), VecName(m))

	s += "for i := range v1 {\n\t\t"
	s += "if !FloatEqualThreshold(v1[i],v2[i], threshold) {\n\t\t\t"
	s += "return false\n\t\t"
	s += "}\n\t}\n\t"
	s += "return true\n}\n\n"

	return s
}

func GenVecFuncEq(m int) (s string) {
	s = fmt.Sprintf("func (v1 %s) ApproxFuncEqual(v2 %s, eq func(float32,float32) bool) bool {\n\t", VecName(m), VecName(m))

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
	mats := "package mgl32\n\nimport(\n\t //\"math\"\n)\n\n"

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
	s = fmt.Sprintf("func Ident%d() %s {\n\treturn %s{", m, GenMatName(m, m), GenMatName(m, m))

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
	s = fmt.Sprintf("func (m1 %s) Add(m2 %s) %s {\n\treturn %s {", GenMatName(m, n), GenMatName(m, n), GenMatName(m, n), GenMatName(m, n))
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
	s = fmt.Sprintf("func (m1 %s) Sub(m2 %s) %s {\n\treturn %s {", GenMatName(m, n), GenMatName(m, n), GenMatName(m, n), GenMatName(m, n))
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
	s = fmt.Sprintf("func (m1 %s) Mul(c float32) %s {\n\treturn %s{", GenMatName(m, n), GenMatName(m, n), GenMatName(m, n))
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
	s = "func " + "(m1 " + GenMatName(m, n) + ") Mul" + fmt.Sprintf("%d", n)

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
	s = fmt.Sprintf("func (m1 %s) Transpose() %s {\n\treturn %s{", GenMatName(m, n), GenMatName(n, m), GenMatName(n, m))

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
	s := fmt.Sprintf("func (m %s) Det() float32 {\n\treturn ", GenMatName(m, m))

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
	s := fmt.Sprintf("func (m %s) Inv() %s {\n\t", GenMatName(m, m), GenMatName(m, m))
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
	s = fmt.Sprintf("func (m1 %s) ApproxEqual(m2 %s) bool {\n\t", GenMatName(m, n), GenMatName(m, n))

	s += "for i := range m1 {\n\t\t"
	s += "if !FloatEqual(m1[i],m2[i]) {\n\t\t\t"
	s += "return false\n\t\t"
	s += "}\n\t}\n\t"
	s += "return true\n}\n\n"

	return s
}

func GenMatThresholdEq(m, n int) (s string) {
	s = fmt.Sprintf("func (m1 %s) ApproxEqualThreshold(m2 %s, threshold float32) bool {\n\t", GenMatName(m, n), GenMatName(m, n))

	s += "for i := range m1 {\n\t\t"
	s += "if !FloatEqualThreshold(m1[i],m2[i], threshold) {\n\t\t\t"
	s += "return false\n\t\t"
	s += "}\n\t}\n\t"
	s += "return true\n}\n\n"

	return s
}

func GenMatFuncEq(m, n int) (s string) {
	s = fmt.Sprintf("func (m1 %s) ApproxFuncEqual(m2 %s, eq func(float32,float32) bool) bool {\n\t", GenMatName(m, n), GenMatName(m, n))

	s += "for i := range m1 {\n\t\t"
	s += "if !eq(m1[i],m2[i]) {\n\t\t\t"
	s += "return false\n\t\t"
	s += "}\n\t}\n\t"
	s += "return true\n}\n\n"

	return s
}
