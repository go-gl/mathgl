package mgl32

import (
	"math"
)

func CartesianToSpherical(coord Vec3) (r, theta, phi float32) {
	r = coord.Len()
	theta = float32(math.Acos(float64(coord[2] / r)))
	phi = float32(math.Atan2(float64(coord[1]), float64(coord[0])))

	return
}

func CartesianToCylindical(coord Vec3) (rho, phi, z float32) {
	rho = float32(math.Hypot(float64(coord[0]), float64(coord[1])))

	phi = float32(math.Atan2(float64(coord[1]), float64(coord[0])))

	z = coord[2]

	return
}

func SphericalToCartesian(r, theta, phi float32) Vec3 {
	st, ct := math.Sincos(float64(theta))
	sp, cp := math.Sincos(float64(phi))

	return Vec3{r * float32(st*cp), r * float32(st*sp), r * float32(ct)}
}

func SphericalToCylindrical(r, theta, phi float32) (rho, phi2, z float32) {
	s, c := math.Sincos(float64(theta))

	rho = r * float32(s)
	z = r * float32(c)
	phi2 = phi

	return
}

func CylindircalToSpherical(rho, phi, z float32) (r, theta, phi2 float32) {
	r = float32(math.Hypot(float64(rho), float64(z)))
	phi2 = phi
	theta = float32(math.Atan2(float64(rho), float64(z)))

	return
}

func CylindricalToCartesian(rho, phi, z float32) Vec3 {
	s, c := math.Sincos(float64(phi))

	return Vec3{rho * float32(c), rho * float32(s), z}
}
