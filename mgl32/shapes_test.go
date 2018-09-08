// Copyright 2018 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.package mgl64_test

package mgl32

import (
	"testing"
)

func TestScreenToGLCoords(t *testing.T) {
	// use a small screen size in order to minimize errors due to fp rounding.
	const (
		sw = 100
		sh = 100
	)
	x, y := ScreenToGLCoords(0, sh-1, sw, sh)
	if x != -1 {
		t.Errorf("x = %f, expected -1.0", x)
	}
	if y != -1 {
		t.Errorf("y = %f, expected -1.0", y)
	}

	x, y = ScreenToGLCoords(sw-1, 0, sw, sh)
	if x != 1 {
		t.Errorf("x = %f, expected 1.0", x)
	}
	if y != 1 {
		t.Errorf("y = %f, expected 1.0", y)
	}
}

func TestGLToScreenCoords(t *testing.T) {
	const (
		sw = 100
		sh = 100
	)
	x, y := GLToScreenCoords(-1, -1, sw, sh)
	if x != 0 {
		t.Errorf("x = %d, expected 0", x)
	}
	if y != sh-1 {
		t.Errorf("y = %d, expected %d", y, sh-1)
	}

	x, y = GLToScreenCoords(1, 1, sw, sh)
	if x != sw-1 {
		t.Errorf("x = %d, expected %d", x, sw-1)
	}
	if y != 0 {
		t.Errorf("y = %d, expected 0", y)
	}
}

func Test_choose(t *testing.T) {

	tests := []struct {
		name string
		n    int
		want []int
	}{
		// test cases.
		{
			name: "C(2,k)",
			n:    2,
			want: []int{1, 2, 1},
		},
		{
			name: "C(3,k)",
			n:    3,
			want: []int{1, 3, 3, 1},
		},
		{
			name: "C(4,k)",
			n:    4,
			want: []int{1, 4, 6, 4, 1},
		},
		{
			name: "C(5,k)",
			n:    5,
			want: []int{1, 5, 10, 10, 5, 1},
		},
		{
			name: "C(6,k)",
			n:    6,
			want: []int{1, 6, 15, 20, 25, 6, 1},
		},
	}
	for _, tt := range tests {
		// t.Run(tt.name, func(t *testing.T) {
		got := []int{}
		fail := false
		for k := 0; k < tt.n; k++ {
			got = append(got, choose(tt.n, k))
			fail = got[k] != tt.want[k]
		}
		if fail {
			t.Errorf("choose() = %v, want %v", got, tt.want)
		}
		// })
	}
}
