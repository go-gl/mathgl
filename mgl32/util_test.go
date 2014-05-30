package mgl32

import (
	"math/rand"
	"testing"
	"time"
)

func TestEqual(t *testing.T) {
	var a float32 = 1.5
	var b float32 = 1.0 + .5

	if !FloatEqual(a, a) {
		t.Errorf("Float Equal fails on comparing a number with itself")
	}

	if !FloatEqual(a, b) {
		t.Errorf("Float Equal fails to compare two equivalent numbers with minimal drift")
	} else if !FloatEqual(b, a) {
		t.Errorf("Float Equal is not symmetric for some reason")
	}

	if !FloatEqual(0.0, 0.0) {
		t.Errorf("Float Equal fails to compare zero values correctly")
	}

	if FloatEqual(1.5, 1.51) {
		t.Errorf("Float Equal gives false positive on large difference")
	}

	if FloatEqual(1.5, 1.5000001) {
		t.Errorf("Float Equal gives false positive on small difference")
	}

	if FloatEqual(1.5, 0.0) {
		t.Errorf("Float Equal gives false positive comparing with zero")
	}
}

func TestEqualThreshold(t *testing.T) {
	// |1.0 - 1.01| < .1
	if !FloatEqualThreshold(1.0, 1.01, 1e-1) {
		t.Errorf("Thresholded equal returns negative on threshold")
	}

	// Comes out to |1.0 - 1.01| < .0001
	if FloatEqualThreshold(1.0, 1.01, 1e-3) {
		t.Errorf("Thresholded equal returns false positive on tolerant threshold")
	}
}

func TestEqual32(t *testing.T) {
	a := float32(1.5)
	b := float32(1.0 + .5)

	if !FloatEqual(a, a) {
		t.Errorf("Float Equal fails on comparing a number with itself")
	}

	if !FloatEqual(a, b) {
		t.Errorf("Float Equal fails to compare two equivalent numbers with minimal drift")
	} else if !FloatEqual(b, a) {
		t.Errorf("Float Equal is not symmetric for some reason")
	}

	if !FloatEqual(0.0, 0.0) {
		t.Errorf("Float Equal fails to compare zero values correctly")
	}

	if FloatEqual(1.5, 1.51) {
		t.Errorf("Float Equal gives false positive on large difference")
	}

	if FloatEqual(1.5, 0.0) {
		t.Errorf("Float Equal gives false positive comparing with zero")
	}
}

func TestClampf(t *testing.T) {
	if !FloatEqual(Clampf(-1.0, 0.0, 1.0), 0.0) {
		t.Errorf("Clamp returns incorrect value for below threshold")
	}

	if !FloatEqual(Clampf(0.0, 0.0, 1.0), 0.0) {
		t.Errorf("Clamp does something weird when value is at threshold")
	}

	if !FloatEqual(Clampf(.14, 0.0, 1.0), .14) {
		t.Errorf("Clamp fails to return correct value when value is within threshold")
	}

	if !FloatEqual(Clampf(1.1, 0.0, 1.0), 1.0) {
		t.Errorf("Clamp fails to return max threshold when appropriate")
	}
}

func TestClampd(t *testing.T) {
	if !FloatEqual(Clampf(-1.0, 0.0, 1.0), 0.0) {
		t.Errorf("Clamp returns incorrect value for below threshold")
	}

	if !FloatEqual(Clampf(0.0, 0.0, 1.0), 0.0) {
		t.Errorf("Clamp does something weird when value is at threshold")
	}

	if !FloatEqual(Clampf(.14, 0.0, 1.0), .14) {
		t.Errorf("Clamp fails to return correct value when value is within threshold")
	}

	if !FloatEqual(Clampf(1.1, 0.0, 1.0), 1.0) {
		t.Errorf("Clamp fails to return max threshold when appropriate")
	}
}

func TestIsClamped(t *testing.T) {
	if IsClampedf(-1.0, 0.0, 1.0) {
		t.Errorf("Test below min is considered clamped")
	}

	if !IsClampedf(.15, 0.0, 1.0) {
		t.Errorf("Test in threshold returns false")
	}

	if IsClampedf(1.5, 0.0, 1.0) {
		t.Errorf("Test above max threshold returns false positive")
	}
}

/* These benchmarks probably aren't very interesting, there's not really many ways to optimize the functions they're benchmarking */

func BenchmarkEqual(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		f1 := r.Float32()
		f2 := r.Float32()
		b.StartTimer()

		FloatEqual(f1, f2)
	}
}

// Here just to get a baseline of how much worse the safer equal is
func BenchmarkBuiltinEqual(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		f1 := r.Float32()
		f2 := r.Float32()
		b.StartTimer()

		_ = f1 == f2
	}
}

func BenchmarkClampf(b *testing.B) {
	b.StopTimer()
	r := rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		a := r.Float32()
		t1 := r.Float32()
		t2 := r.Float32()
		b.StartTimer()

		Clampf(a, t1, t2)
	}
}
