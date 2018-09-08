package mgl32

import (
	"math"
)

func Atan2(y, x float32) float32 {
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

	angle := math.Atan(float64(y) / float64(x))

	if x < 0.0 && y > 0.0 {
		return float32(angle + math.Pi)
	} else if x < 0.0 && y < 0.0 {
		return float32(angle + math.Pi)
	} else if x > 0.0 && y < 0.0 {
		return float32(angle + 2.0*math.Pi)
	}

	return float32(angle)
}
