// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mgl32

import (
	"github.com/luxengine/math"
)

// Converts 3-dimensional cartesian coordinates (x,y,z) to spherical
// coordinates with radius r, inclination theta, and azimuth phi.
//
// All angles are in radians.
func CartesianToSpherical(coord Vec3) (r, theta, phi float32) {
	r = coord.Len()
	theta = math.Acos(coord[2] / r)
	phi = math.Atan2(coord[1], coord[0])

	return
}

// Converts 3-dimensional cartesian coordinates (x,y,z) to cylindrical
// coordinates with radial distance r, azimuth phi, and height z.
//
// All angles are in radians.
func CartesianToCylindical(coord Vec3) (rho, phi, z float32) {
	rho = math.Hypot(coord[0], coord[1])

	phi = math.Atan2(coord[1], coord[0])

	z = coord[2]

	return
}

// Converts spherical coordinates with radius r, inclination theta,
// and azimuth phi to cartesian coordinates (x,y,z).
//
// Angles are in radians.
func SphericalToCartesian(r, theta, phi float32) Vec3 {
	st, ct := math.Sincos(theta)
	sp, cp := math.Sincos(phi)

	return Vec3{r * st * cp, r * st * sp, r * ct}
}

// Converts spherical coordinates with radius r, inclination theta,
// and azimuth phi to cylindrical coordinates with radial distance r,
// azimuth phi, and height z.
//
// Angles are in radians
func SphericalToCylindrical(r, theta, phi float32) (rho, phi2, z float32) {
	s, c := math.Sincos(theta)

	rho = r * float32(s)
	z = r * float32(c)
	phi2 = phi

	return
}

// Converts cylindrical coordinates with radial distance r,
// azimuth phi, and height z to spherical coordinates with radius r,
// inclination theta, and azimuth phi.
//
// Angles are in radians
func CylindircalToSpherical(rho, phi, z float32) (r, theta, phi2 float32) {
	r = math.Hypot(rho, z)
	phi2 = phi
	theta = math.Atan2(rho, z)

	return
}

// Converts cylindrical coordinates with radial distance r,
// azimuth phi, and height z to cartesian coordinates (x,y,z)
//
// Angles are in radians.
func CylindricalToCartesian(rho, phi, z float32) Vec3 {
	s, c := math.Sincos(phi)

	return Vec3{rho * c, rho * s, z}
}

// Converts degrees to radians
func DegToRad(angle float32) float32 {
	return angle * math.Pi / 180
}

// Converts radians to degrees
func RadToDeg(angle float32) float32 {
	return angle * 180 / math.Pi
}
