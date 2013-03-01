package mathgl_test

import (
	"github.com/Jragonmiris/mathgl"
	"testing"
)

func TestMulIdent(t *testing.T){
	i1 := [...]float32{1,0,0,0, 0,1,0,0, 0,0,1,0, 0,0,0,1}
	i2 := mathgl.Ident4f()
	i3 := mathgl.Ident4f()
	
	mul := i2.Mul4(i3)
	
	for i := range mul {
		if mul[i] != i1[i] {
			t.Fatalf("Multiplication of identities does not yield identity")
		}
	}
}