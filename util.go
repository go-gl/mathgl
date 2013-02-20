package mathgl

import "math"

const epsilon float64 = 1e-15

// A safe utility function to compare floats
// Taken from http://floating-point-gui.de/errors/comparison/
// Slightly altered to not call Abs when not needed
func FloatEqual(a, b float64) bool {
	if a == b { // Handles the case of inf
		return true
	} else if a*b == 0 { // If a or b are 0
		return math.Abs(a-b) < epsilon*epsilon
	}

	// Else compare difference
	return math.Abs(a-b)/(math.Abs(a)+math.Abs(b)) < epsilon
}
