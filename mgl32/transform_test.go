// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mgl32

import (
	"math"
	"testing"
)

func TestHomogRotate3D(t *testing.T) {
	tests := []struct {
		Angle    float32
		Axis     Vec3
		Expected Mat4
	}{
		{0, Vec3{0, 0, 0}, Ident4()},
		{
			DegToRad(-90), Vec3{0, 1, 0}, // heading 90 degree
			Mat4{
				0, 0, 1, 0,
				0, 1, 0, 0,
				-1, 0, 0, 0,
				0, 0, 0, 1,
			},
		},
		{
			DegToRad(180), Vec3{0, 1, 0}, // heading 180 degree
			Mat4{
				-1, 0, 0, 0,
				0, 1, 0, 0,
				0, 0, -1, 0,
				0, 0, 0, 1,
			},
		},
		{
			DegToRad(-90), Vec3{0, 0, 1}, // attitude 90 degree
			Mat4{
				0, -1, 0, 0,
				1, 0, 0, 0,
				0, 0, 1, 0,
				0, 0, 0, 1,
			},
		},
		{
			DegToRad(-90), Vec3{1, 0, 0}, // bank 90 degree
			Mat4{
				1, 0, 0, 0,
				0, 0, -1, 0,
				0, 1, 0, 0,
				0, 0, 0, 1,
			},
		},
	}

	threshold := float32(math.Pow(10, -2))
	for _, c := range tests {
		if r := HomogRotate3D(c.Angle, c.Axis); !r.ApproxEqualThreshold(c.Expected, threshold) {
			t.Errorf("HomogRotate3D(%v, %v) != %v (got %v)", c.Angle, c.Axis, c.Expected, r)
		}
	}
}
