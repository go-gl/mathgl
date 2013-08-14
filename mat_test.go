package mathgl

import (
	"math/rand"
	"testing"
	"time"
)

func TestMulIdent(t *testing.T) {
	i1 := [...]float32{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1}
	i2 := Ident4f()
	i3 := Ident4f()

	mul := i2.Mul4(i3)

	for i := range mul {
		if mul[i] != i1[i] {
			t.Errorf("Multiplication of identities does not yield identity")
		}
	}
}

func BenchmarkMatAdd(b *testing.B) {
	b.StopTimer()
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m1 := Mat4f{}
		m2 := Mat4f{}

		for j := 0; j < len(m1); j++ {
			m1[j], m2[j] = rand.Float32(), rand.Float32()
		}
		b.StartTimer()

		m1.Add(m2)
	}
}

func BenchmarkMatScale(b *testing.B) {
	b.StopTimer()
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m1 := Mat4f{}

		for j := 0; j < len(m1); j++ {
			m1[j] = rand.Float32()
		}
		c := rand.Float32()
		b.StartTimer()

		m1.Mul(c)
	}
}

func BenchmarkMatMul(b *testing.B) {
	b.StopTimer()
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m1 := Mat4f{}
		m2 := Mat4f{}

		for j := 0; j < len(m1); j++ {
			m1[j], m2[j] = rand.Float32(), rand.Float32()
		}
		b.StartTimer()

		m1.Mul4(m2)
	}
}

func BenchmarkMatTranspose(b *testing.B) {
	b.StopTimer()
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m1 := Mat4f{}

		for j := 0; j < len(m1); j++ {
			m1[j] = rand.Float32()
		}
		b.StartTimer()

		m1.Transpose()
	}
}

func BenchmarkMatDet(b *testing.B) {
	b.StopTimer()
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m1 := Mat4f{}

		for j := 0; j < len(m1); j++ {
			m1[j] = rand.Float32()
		}
		b.StartTimer()

		m1.Det()
	}
}

func BenchmarkMatInv(b *testing.B) {
	b.StopTimer()
	rand := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		m1 := Mat4f{}

		for j := 0; j < len(m1); j++ {
			m1[j] = rand.Float32()
		}
		b.StartTimer()

		m1.Inv()
	}
}
