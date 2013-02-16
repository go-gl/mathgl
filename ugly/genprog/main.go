package main

import "fmt"

func main() {
	s := ""
	for m := 2; m <= 4; m++ {
		for n := 2; n <= 4; n++ {
			if m == n {
				continue
			}
			for o := 1; o <= 4; o++ {
				s += fmt.Sprintf(GenMul(m,n,o))
			}
		}
	}
	
	fmt.Println(s)
}

func GenMul(m, n, o int) (s string) {
	s = "func " + "(m1 " + GenName(m,n) + "f) Mul" + fmt.Sprintf("%d",n)
	
	if n != o {
		s += fmt.Sprintf("x%d",o)
	}
	
	s += "f(m2 " + GenName(n,o) + "f) " + GenName(m,o) + "f {\n\treturn " + GenName(m,o) + "f{"
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
	s += "}\n}\n"
	return s
}

func GenName(m, n int) string {
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