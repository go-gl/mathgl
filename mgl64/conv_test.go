package mgl32

import (
	"testing"
)

func TestCartesianToSphere(t *testing.T) {
	v := Vec3{5, 12, 9}

	r, theta, phi := CartesianToSpherical(v)

	if !FloatEqualThreshold(r, 15.8114, 1e-4) {
		t.Errorf("Got incorrect value for radius. Got: %f, expected: %f", r, 15.8114)
	}

	if !FloatEqualThreshold(theta, 0.965250852, 1e-4) {
		t.Errorf("Got incorrect value for theta. Got: %f, expected: %f", theta, 0.965250852)
	}

	if !FloatEqualThreshold(phi, 1.1760046, 1e-4) {
		t.Errorf("Got incorrect value for phi. Got: %f, expected: %f", phi, 1.1760046)
	}
}

func TestSphereToCartesian(t *testing.T) {
	v := Vec3{5, 12, 9}

	result := SphericalToCartesian(15.8114, 0.965250852, 1.1760046)

	if !v.ApproxEqualThreshold(result, 1e-4) {
		t.Errorf("Got incorrect vector. Got: %v, Expected: %v", result, v)
	}
}

func TestCartesianToCylinder(t *testing.T) {
	v := Vec3{5, 12, 9}

	rho, phi, z := CartesianToCylindical(v)

	if !FloatEqualThreshold(rho, 13, 1e-4) {
		t.Errorf("Got incorrect value for radius. Got: %f, expected: %f", rho, 13)
	}

	if !FloatEqualThreshold(phi, 1.17601, 1e-4) {
		t.Errorf("Got incorrect value for theta. Got: %f, expected: %f", phi, 1.17601)
	}

	if !FloatEqualThreshold(z, 9, 1e-4) {
		t.Errorf("Got incorrect value for phi. Got: %f, expected: %f", z, 9)
	}
}

func TestCylinderToCartesian(t *testing.T) {
	v := Vec3{5, 12, 9}

	result := CylindricalToCartesian(13, 1.17601, 9)

	if !v.ApproxEqualThreshold(result, 1e-4) {
		t.Errorf("Got incorrect vector. Got: %v, expected: %v", result, v)
	}
}

func TestCylinderToSphere(t *testing.T) {
	r, theta, phi := CylindircalToSpherical(13, 1.17601, 9)

	if !FloatEqualThreshold(r, 15.8114, 1e-4) {
		t.Errorf("Got incorrect value for radius. Got: %f, expected: %f", r, 15.8114)
	}

	if !FloatEqualThreshold(theta, 0.965250852, 1e-4) {
		t.Errorf("Got incorrect value for theta. Got: %f, expected: %f", theta, 0.965250852)
	}

	if !FloatEqualThreshold(phi, 1.1760046, 1e-4) {
		t.Errorf("Got incorrect value for phi. Got: %f, expected: %f", phi, 1.1760046)
	}
}

func TestSphereToCylinder(t *testing.T) {
	rho, phi, z := SphericalToCylindrical(15.8114, 0.965250852, 1.1760046)

	if !FloatEqualThreshold(rho, 13, 1e-4) {
		t.Errorf("Got incorrect value for radius. Got: %f, expected: %f", rho, 13)
	}

	if !FloatEqualThreshold(phi, 1.17601, 1e-4) {
		t.Errorf("Got incorrect value for theta. Got: %f, expected: %f", phi, 1.17601)
	}

	if !FloatEqualThreshold(z, 9, 1e-4) {
		t.Errorf("Got incorrect value for phi. Got: %f, expected: %f", z, 9)
	}
}
