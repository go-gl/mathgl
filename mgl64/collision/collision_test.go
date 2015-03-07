package collision

import (
	"testing"

	"github.com/go-gl/mathgl/mgl64"
)

func TestPointCollision(t *testing.T) {
	var pointTests = [...]struct {
		aabb  AABB
		point mgl64.Vec3

		out bool
	}{
		{AABB{Min: mgl64.Vec3{0, 0, 0}, Max: mgl64.Vec3{1, 1, 1}}, mgl64.Vec3{.5, .5, .5}, true},
		{AABB{Min: mgl64.Vec3{0, 0, 0}, Max: mgl64.Vec3{1, 1, 1}}, mgl64.Vec3{1, 1, 1.5}, false},
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
		{AABB{Min: mgl64.Vec3{0, 0, 0}, Max: mgl64.Vec3{1, 1, 1}}, AABB{Min: mgl64.Vec3{-.5, -.5, -.5}, Max: mgl64.Vec3{.5, .5, .5}}, true},
		{AABB{Min: mgl64.Vec3{0, 0, 0}, Max: mgl64.Vec3{1, 1, 1}}, AABB{Min: mgl64.Vec3{-.5, -.5, -.5}, Max: mgl64.Vec3{-.1, -.1, -.1}}, false},
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
		p1, p2, p3 mgl64.Vec3

		out bool
	}{
		{AABB{Min: mgl64.Vec3{0, 0, 0}, Max: mgl64.Vec3{1, 1, 1}}, mgl64.Vec3{.5, .5, .5}, mgl64.Vec3{0, 0, 0}, mgl64.Vec3{1, 1, 1}, true},
		{AABB{Min: mgl64.Vec3{0, 0, 0}, Max: mgl64.Vec3{1, 1, 1}}, mgl64.Vec3{1.1, 1.1, 1.1}, mgl64.Vec3{1.5, 1.5, 1.5}, mgl64.Vec3{2.4, 3.7, 9.1}, false},
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
		rayFrom, dir mgl64.Vec3

		out          bool
		correctedOut bool // If the intersection is behind the ray
	}{
		{AABB{Min: mgl64.Vec3{0, 0, 0}, Max: mgl64.Vec3{1, 1, 1}}, mgl64.Vec3{-1, .5, .5}, mgl64.Vec3{1, 0, 0}, true, true},
		{AABB{Min: mgl64.Vec3{0, 0, 0}, Max: mgl64.Vec3{1, 1, 1}}, mgl64.Vec3{-1.5, .5, .5}, mgl64.Vec3{-1, 0, 0}, true, false},
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
		transform mgl64.Mat4

		out AABB
	}{
		{AABB{Min: mgl64.Vec3{0, 0, 0}, Max: mgl64.Vec3{1, 1, 1}}, mgl64.Translate3D(0, 0, .5), AABB{Min: mgl64.Vec3{0, 0, .5}, Max: mgl64.Vec3{1, 1, 1.5}}},
		{AABB{Min: mgl64.Vec3{0, 0, 0}, Max: mgl64.Vec3{1, 1, 1}}, mgl64.Ident4(), AABB{Min: mgl64.Vec3{0, 0, 0}, Max: mgl64.Vec3{1, 1, 1}}},
		{AABB{Min: mgl64.Vec3{0, 0, 0}, Max: mgl64.Vec3{1, 1, 1}}, mgl64.Translate3D(0, 0, .5).Mul4(mgl64.Scale3D(2, 1, 1)), AABB{Min: mgl64.Vec3{0, 0, .5}, Max: mgl64.Vec3{2, 1, 1.5}}},
	}

	for _, test := range transformTests {
		out := test.aabb.Transform(test.transform)
		if !out.Min.ApproxEqualThreshold(test.out.Min, 1e-4) || !out.Max.ApproxEqualThreshold(test.out.Max, 1e-4) {
			t.Errorf("AABB %v transformed by matrix %v results in AABB %v and not AABB %v", test.aabb, test.transform, out, test.out)
		}
	}
}

func TestRay(t *testing.T) {
	var tests = [...]struct {
		rayFrom, dir mgl64.Vec3
		t            float64

		out mgl64.Vec3
	}{
		{mgl64.Vec3{0, 0, 0}, mgl64.Vec3{1, 0, 0}, .5, mgl64.Vec3{.5, 0, 0}},
		{mgl64.Vec3{0, 0, 0}, mgl64.Rotate3DY(mgl64.DegToRad(-90)).Mul3x1(mgl64.Vec3{1, 0, 0}), 2, mgl64.Vec3{0, 0, 2}},
	}

	for _, test := range tests {
		out := Ray(test.rayFrom, test.dir, test.t)
		if !out.ApproxEqualThreshold(test.out, 1e-3) { // No joke 1e-4 is too sensitive for test 2
			t.Errorf("Ray at %v pointing %v results in vector %v and not the expected %v", test.rayFrom, test.dir, out, test.out)
		}
	}
}
