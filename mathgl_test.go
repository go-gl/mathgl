package mathgl_test

import (
	"github.com/Jragonmiris/mathgl"
	"testing"
	"math"
)

func TestVecSetGet(t *testing.T) {
	v1 := mathgl.NewVector(mathgl.FLOAT32)
	if v1 == nil {
		t.Fatalf("Failed to create new vector")
	}
	
	input := []mathgl.VecNum{mathgl.VecInt32(1),mathgl.VecFloat32(2),mathgl.VecFloat32(3),mathgl.VecFloat32(4),mathgl.VecFloat32(5)}
	
	if err := v1.AddElements(input); err == nil {
		t.Errorf("Added list with bad element to vector")
	}
	
	input[0] = mathgl.VecFloat32(1)
	
	if err := v1.AddElements(input); err != nil {
		t.Fatalf("Failed to set vector with correct list")
	}
	
	if a := v1.GetElement(0); a == nil || float32(a.(mathgl.VecFloat32)) - float32(1.) > 0.000001 {
		t.Errorf("Didn't get/set correct element of vector")
	}
	
	v2,e := mathgl.VectorOf(mathgl.FLOAT32,input)
	if v2 == nil || e != nil {
		t.Fatalf("VectorOf failed on good input")
	}
	
	if a,b := v1.GetElement(2), v2.GetElement(2); a == nil || b == nil || float32(a.(mathgl.VecFloat32)) - float32(b.(mathgl.VecFloat32)) > .000000001 {
		t.Errorf("Two vectors not the same despite being made from same list")
	}
	
	
	if err := v1.SetElement(25, mathgl.VecFloat32(1)); err == nil {
		t.Errorf("Set out of bounds vector element")
	}
	
	if err := v1.SetElement(-3, mathgl.VecFloat32(1)); err == nil {
		t.Errorf("Set out of bounds vector element")
	}
	
	if err := v1.SetElement(4, mathgl.VecFloat32(42)); err != nil {
		t.Fatalf("Didn't set in-bounds vector element")
	}
	
	if a := v1.GetElement(4); math.Abs(float64(float32(a.(mathgl.VecFloat32)) - float32(42))) > .00001 {
		t.Errorf("Did not correctly set single-in bounds vector element")
	}
	
	if a := v2.GetElement(4); math.Abs(float64(float32(a.(mathgl.VecFloat32)) - float32(42))) < .00001 {
		t.Errorf("Changing one vector changed another %d")
	}
	
}