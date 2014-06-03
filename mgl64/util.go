// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mgl64

// Epsilon is some tiny value that determines how precisely equal we want our floats to be
// This is exported and left as a variable in case you want to change the default threshold for the
// purposes of certain methods (e.g. Unproject uses the default epsilon when determining
// if the determinant is "close enough" to zero to mean there's no inverse).
//
// This is, obviously, not mutex protected so be **absolutely sure** that no functions using Epsilon
// are being executed when you change this.
//
// This is left as a float64 for more precise comparisons
var Epsilon float64 = 1e-10

// A direct copy of the math package's Abs. This is here for the mgl32
// package, to prevent rampant type conversions during equality tests.
func Abs(a float64) float64 {
	if a < 0 {
		return -a
	} else if a == 0 {
		return 0
	}

	return a
}

// FloatEqual is a safe utility function to compare floats.
// It's Taken from http://floating-point-gui.de/errors/comparison/
//
// It is slightly altered to not call Abs when not needed.
// Keep in mind that it expects float32s to be converted to float64s before being passed in, because they have to be converted for Abs anyway
func FloatEqual(a, b float64) bool {

	if a == b { // Handles the case of inf or shortcuts the loop when no significant error has accumulated
		return true
	} else if a*b == 0 { // If a or b are 0
		return Abs(a-b) < Epsilon*Epsilon
	}

	// Else compare difference
	return Abs(a-b)/(Abs(a)+Abs(b)) < Epsilon
}

// FloatEqualFunc is a utility closure that will generate a function that
// always approximately compares floats like FloatEqualThreshold with a different
// threshold.
func FloatEqualFunc(epsilon float64) func(float64, float64) bool {
	return func(a, b float64) bool {
		return FloatEqualThreshold(a, b, epsilon)
	}
}

// FloatEqualThreshold is a utility function to compare floats.
// It's Taken from http://floating-point-gui.de/errors/comparison/
//
// It is slightly altered to not call Abs when not needed.
// Keep in mind that it expects float32s to be converted to float64s before being passed in, because they have to be converted for Abs anyway
//
// This differs from FloatEqual in that it lets you pass in your comparison threshold, so that you can adjust the comparison value to your specific needs
func FloatEqualThreshold(a, b, epsilon float64) bool {

	if a == b { // Handles the case of inf or shortcuts the loop when no significant error has accumulated
		return true
	} else if a*b == 0 { // If a or b is 0
		return Abs(a-b) < epsilon*epsilon
	}

	// Else compare difference
	return Abs(a-b)/(Abs(a)+Abs(b)) < epsilon
}

func Clamp(a, t1, t2 float64) float64 {
	if a < t1 {
		return t1
	} else if a > t2 {
		return t2
	}

	return a
}

func ClampFunc(t1, t2 float64) func(float64) {
	return func(a float64) {
		Clamp(a, t1, t2)
	}
}

/* The IsClamped functions use strict equality (meaning: not the FloatEqual function)
there shouldn't be any major issues with this since clamp is often used to fix minor errors*/

func IsClamped(a, t1, t2 float64) bool {
	return a >= t1 && a <= t2
}

func SetMin(a, b *float64) {
	if *b < *a {
		*a = *b
	}
}

func SetMax(a, b *float64) {
	if *a < *b {
		*a = *b
	}
}
