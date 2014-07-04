package mgl32

import (
	"testing"
)

func TestPointCollision(t *testing.T) {
	var pointTests = [...]struct {
		aabb  AABB
		point Vec3

		out bool
	}{
		{AABB{Min: Vec3{0, 0, 0}, Max: Vec3{1, 1, 1}}, Vec3{.5, .5, .5}, true},
		{AABB{Min: Vec3{0, 0, 0}, Max: Vec3{1, 1, 1}}, Vec3{1, 1, 1.5}, false},
	}

	for _, test := range pointTests {
		if test.aabb.IsPointInside(test.point) != test.out {
			t.Errorf("Test of AABB %v and Point %v does not return expected result %v", test.aabb, test.point, test.out)
		}
	}
}

func TestAABBCollision(t *testing.T) {
	var aabbTests = [...]struct {
		aabb1 AABB
		aabb2 AABB

		out bool
	}{
		{AABB{Min: Vec3{0, 0, 0}, Max: Vec3{1, 1, 1}}, AABB{Min: Vec3{-.5, -.5, -.5}, Max: Vec3{.5, .5, .5}}, true},
		{AABB{Min: Vec3{0, 0, 0}, Max: Vec3{1, 1, 1}}, AABB{Min: Vec3{-.5, -.5, -.5}, Max: Vec3{-.1, -.1, -.1}}, false},
	}

	for _, test := range aabbTests {
		if test.aabb1.Collides(test.aabb2) != test.out {
			t.Errorf("Test of AABB %v and AABB %v does not return expected result %v", test.aabb1, test.aabb2, test.out)
		}
	}
}

func TestTriangleCollision(t *testing.T) {
	var triangleTests = [...]struct {
		aabb       AABB
		p1, p2, p3 Vec3

		out bool
	}{
		{AABB{Min: Vec3{0, 0, 0}, Max: Vec3{1, 1, 1}}, Vec3{.5, .5, .5}, Vec3{0, 0, 0}, Vec3{1, 1, 1}, true},
		{AABB{Min: Vec3{0, 0, 0}, Max: Vec3{1, 1, 1}}, Vec3{1.1, 1.1, 1.1}, Vec3{1.5, 1.5, 1.5}, Vec3{2.4, 3.7, 9.1}, false},
	}

	for _, test := range triangleTests {
		if test.aabb.CollidesTriangle(test.p1, test.p2, test.p3) != test.out {
			t.Errorf("Test of AABB %v and Triangle {%v,%v,%v} does not return expected result %v", test.aabb, test.p1, test.p2, test.p3, test.out)
		}
	}
}

func TestRayIntersection(t *testing.T) {
	var rayTests = [...]struct {
		aabb         AABB
		rayFrom, dir Vec3

		out          bool
		correctedOut bool // If the intersection is behind the ray
	}{
		{AABB{Min: Vec3{0, 0, 0}, Max: Vec3{1, 1, 1}}, Vec3{-1, .5, .5}, Vec3{1, 0, 0}, true, true},
		{AABB{Min: Vec3{0, 0, 0}, Max: Vec3{1, 1, 1}}, Vec3{-1.5, .5, .5}, Vec3{-1, 0, 0}, true, false},
	}

	for _, test := range rayTests {
		ok, tmin, tmax := test.aabb.IntersectsRay(test.rayFrom, test.dir)
		if ok != test.out {
			t.Errorf("Test of AABB %v and ray starting at %v with direction %v does not return expected result %v.", test.aabb, test.rayFrom, test.dir, test.out)
		}
		if tmin < 0 && tmax < 0 {
			ok = false
		}
		t.Logf("Got tmin, tmax %v; %v", tmin, tmax)
		if ok != test.correctedOut {
			t.Errorf("Test of AABB %v and ray starting at %v with direction %v does not return expected result %v after correcting for direction.", test.aabb, test.rayFrom, test.dir, test.correctedOut)
		}
	}
}

func TestAABBTransform(t *testing.T) {
	var transformTests = [...]struct {
		aabb      AABB
		transform Mat4

		out AABB
	}{
		{AABB{Min: Vec3{0, 0, 0}, Max: Vec3{1, 1, 1}}, Translate3D(0, 0, .5), AABB{Min: Vec3{0, 0, .5}, Max: Vec3{1, 1, 1.5}}},
		{AABB{Min: Vec3{0, 0, 0}, Max: Vec3{1, 1, 1}}, Ident4(), AABB{Min: Vec3{0, 0, 0}, Max: Vec3{1, 1, 1}}},
		{AABB{Min: Vec3{0, 0, 0}, Max: Vec3{1, 1, 1}}, Translate3D(0, 0, .5).Mul4(Scale3D(2, 1, 1)), AABB{Min: Vec3{0, 0, .5}, Max: Vec3{2, 1, 1.5}}},
	}

	for _, test := range transformTests {
		out := test.aabb.Transform(test.transform)
		if !out.Min.ApproxEqualThreshold(test.out.Min, 1e-4) || !out.Max.ApproxEqualThreshold(test.out.Max, 1e-4) {
			t.Errorf("AABB %v transformed by matrix %v results in AABB %v and not AABB %v", test.aabb, test.transform, out, test.out)
		}
	}
}
