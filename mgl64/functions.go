package mgl64

import (
	"math"
)

func Atan2(y, x float64) float64 {
	if x == 0.0 {
		if y >= 0.0 {
			return math.Pi / 2.0
		} else {
			return math.Pi * 1.5
		}
	} else if y == 0.0 {
		if x > 0.0 {
			return 0.0
		} else {
			return math.Pi
		}
	}

	angle := math.Atan(y / x)

	if x < 0.0 && y > 0.0 {
		return angle + math.Pi
	} else if x < 0.0 && y < 0.0 {
		return angle + math.Pi
	} else if x > 0.0 && y < 0.0 {
		return angle + 2.0*math.Pi
	}

	return angle
}
