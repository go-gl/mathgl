package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Making vectorf.go")
	vecs := GenVec()
	vecf,err := os.Create("vectorf.go")
	if err != nil {
		panic(err)
	}
	defer vecf.Close()

	_,err = vecf.Write([]byte(vecs))
	if err != nil {
		panic(err)
	}

	fmt.Println("Making matrixf.go")
	mats := GenMat()
	//fmt.Println(mats)
	matf,err := os.Create("matrixf.go")
	if err != nil {
		panic(err)
	}
	defer matf.Close()

	_,err = matf.Write([]byte(mats))
	if err != nil {
		panic(err)
	}

	
	fmt.Println("Done")
}

func GenVec() (s string) {
	vecs := "package mathgl\n\nimport(\n\t \"math\"\n)\n\n"
	
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

	/*for m := 2; m <= 4; m++ {
		for o := 2; o <= 4; o++ {
			vecs += GenVecMatMul(m,o)
		}
	}*/

	return vecs
}

func GenVecDef(m int) (s string) {
	return fmt.Sprintf("type Vec%df [%d]float32\n", m, m)
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

func VecName(m int) (s string) {
	return fmt.Sprintf("Vec%df", m)
}

func GenMat() string {
	mats := "package mathgl\n\nimport(\n\t //\"math\"\n)\n\n"

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			mats += GenMatDef(m,n)
		}
	}
	mats += "\n"

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			mats += GenMatAdd(m,n)
		}
	}

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			mats += GenMatSub(m,n)
		}
	}

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			mats += GenScalarMul(m,n)
		}
	}

	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			for o := 1; o <= 4; o++ {
				mats += GenMatMul(m,n,o)
			}
		}
	}

	return mats
}

func GenMatDef(m,n int) (s string) {
	return fmt.Sprintf("type %sf [%d]float32\n", GenMatName(m,n), m*n)
}

func GenMatAdd(m,n int) (s string) {
	s = fmt.Sprintf("func (m1 %sf) Add(m2 %sf) %sf {\n\treturn %sf {", GenMatName(m,n), GenMatName(m,n), GenMatName(m,n), GenMatName(m,n))
	for i := 0; i < m*n; i++ {
		s += fmt.Sprintf("m1[%d] + m2[%d]", i, i)
		if i != (m*n)-1 {
			s += ","
		}
	}
	s += "}\n}\n\n"
	return s
}

func GenMatSub(m,n int) (s string) {
	s = fmt.Sprintf("func (m1 %sf) Sub(m2 %sf) %sf {\n\treturn %sf {", GenMatName(m,n), GenMatName(m,n), GenMatName(m,n), GenMatName(m,n))
	for i := 0; i < m*n; i++ {
		s += fmt.Sprintf("m1[%d] - m2[%d]", i, i)
		if i != (m*n)-1 {
			s += ","
		}
	}
	s += "}\n}\n\n"
	return s
}

func GenScalarMul(m,n int) (s string) {
	s = fmt.Sprintf("func (m1 %sf) Mul(c float32) %sf {\n\treturn %sf{", GenMatName(m,n), GenMatName(m,n), GenMatName(m,n))
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
	s = "func " + "(m1 " + GenMatName(m,n) + "f) Mul" + fmt.Sprintf("%d",n)

	if n != o {
		s += fmt.Sprintf("x%d",o)
	}

	s += "f(m2 " + GenMatName(n,o) + "f) " + GenMatName(m,o) + "f {\n\treturn " + GenMatName(m,o) + "f{"
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

/*func GenDet(m,n int) (s string) {
	s = fmt.Sprintf("func (m1 %s) Mul() float32 {\n\treturn %s{", GenMatName(m,n), GenMatName(n,o), GenMatName(m,o), GenMatName(m,o))
}*/

func GenMatName(m, n int) string {
	if m == 1 {
		return fmt.Sprintf("Vec%d",n)
	}
	
	if n == 1 {
		return fmt.Sprintf("Vec%d",m)
	}
	
	s :=  fmt.Sprintf("Matrix%d",m)
	
	if m != n {
		s +=  fmt.Sprintf("x%d",n)
	}
	
	return s
}
