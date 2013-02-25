package mathgl_test

import (
	"github.com/Jragonmiris/mathgl"
	"testing"
)

// Benchmarks a standard model-view-projection multiplication
func BenchmarkMVPMul(b *testing.B) {
	b.StopTimer()
	
	M := mathgl.Identity(4, mathgl.FLOAT64)
	V := mathgl.LookAt(4., 3., 3., 0., 0., 0., 0, 1., 0)
	P := mathgl.Perspective(45.0, 4./3., .1, 100.0)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		P.Mul(V).Mul(M)
		//MVP.AsCMOArray(mathgl.FLOAT32)
	}
}

// Benchmarks creation AND multiplication of a standard MVP matrix
func BenchmarkMVPMake(b *testing.B) {

	for i := 0; i < b.N; i++ {
	
		M := mathgl.Identity(4, mathgl.FLOAT64)
		V := mathgl.LookAt(4., 3., 3., 0., 0., 0., 0, 1., 0)
		P := mathgl.Perspective(45.0, 4./3., .1, 100.0)
		P.Mul(V).Mul(M)
	}
}

func BenchmarkPerspective(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mathgl.Perspective(45.0, 4./3., .1, 100.0)
	}
}

func BenchmarkIdentity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mathgl.Identity(4, mathgl.FLOAT64)
	}
}

// This should be the most expensive because of all the cross products and multiplications it does
func BenchmarkLookAt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mathgl.LookAt(4., 3., 3., 0., 0., 0., 0, 1., 0)
	}
}

func BenchmarkLookAtV(b *testing.B) {
	b.StopTimer()
	eye,_ := mathgl.InferVectorOf(4., 3., 3.)
	center,_ := mathgl.InferVectorOf(0.,0.,0.)
	up,_ := mathgl.InferVectorOf(0.,1.,0.)

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mathgl.LookAtV(*eye, *center, *up)
	}
}

func BenchmarkVectorOf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mathgl.VectorOf([]mathgl.Scalar {
			mathgl.MakeScalar(4.0, mathgl.FLOAT64),
			mathgl.MakeScalar(3.1, mathgl.FLOAT64),
			mathgl.MakeScalar(9.2, mathgl.FLOAT64),
			mathgl.MakeScalar(1.0, mathgl.FLOAT64)}, mathgl.FLOAT64)
	}
}

func BenchmarkInferVectorOf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mathgl.InferVectorOf(
			4.0,
			3.1,
			9.2,
			1.0)
	}
}
