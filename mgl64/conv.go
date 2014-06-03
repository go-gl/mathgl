package mgl64

import (
	"math"
)

func CartesianToSpherical(coord Vec3) (r, theta, phi float64) {
	r = coord.Len()
	theta = float64(math.Acos(float64(coord[2] / r)))
	phi = float64(math.Atan2(float64(coord[1]), float64(coord[0])))

	return
}

func CartesianToCylindical(coord Vec3) (rho, phi, z float64) {
	rho = float64(math.Hypot(float64(coord[0]), float64(coord[1])))

	if xeq, yeq := FloatEqual(coord[0], 0), FloatEqual(coord[1], 0); xeq || yeq {
		phi = 0
	} else if coord[0] > 0 || xeq {
		phi = float64(math.Atan2(float64(coord[1]), float64(coord[0])))
	} else {
		phi = float64(-math.Asin(float64(coord[1]/rho)) + math.Pi)
	}

	z = coord[2]

	return
}

func SphericalToCartesian(r, theta, phi float64) Vec3 {
	st, ct := math.Sincos(float64(theta))
	sp, cp := math.Sincos(float64(phi))

	return Vec3{r * float64(st*cp), r * float64(st*sp), r * float64(ct)}
}

func SpericalToCylindrical(r, theta, phi float64) (rho, phi2, z float64) {
	s, c := math.Sincos(float64(theta))

	rho = r * float64(s)
	z = r * float64(c)
	phi2 = phi

	return
}

func CylindircalToSpherical(rho, phi, z float64) (r, theta, phi2 float64) {
	r = float64(math.Hypot(float64(rho), float64(z)))
	phi2 = phi
	theta = float64(math.Atan2(float64(rho), float64(z)))

	return
}

func CylindricalToCartesian(rho, phi, z float64) Vec3 {
	s, c := math.Sincos(float64(phi))

	return Vec3{rho * float64(c), rho * float64(s), z}
}
