package mgl64

import (
	"testing"
)

func TestMatrixLookAtV(t *testing.T) {
	doLookAtTests(func(eye, center, up Vec3) translateFunc {
		m := LookAtV(eye, center, up)
		return func(v Vec3) Vec3 {
			return m.Mul4x1(v.Vec4(0)).Vec3()
		}
	}, t)
}

func TestQuatLookAtVb(t *testing.T) {
	doLookAtTests(func(eye, center, up Vec3) translateFunc {
		q := QuatLookAtV(eye, center, up).Normalize()
		return func(v Vec3) Vec3 {
			return q.Rotate(v)
		}
	}, t)
}

func TestQuatLookAtOld(t *testing.T) {
	doLookAtTests(func(eye, center, up Vec3) translateFunc {
		q := QuatLookAtOld(eye, center, up).Normalize()
		return func(v Vec3) Vec3 {
			return q.Rotate(v)
		}
	}, t)
}

func TestQuatLookAtNew(t *testing.T) {
	doLookAtTests(func(eye, center, up Vec3) translateFunc {
		q := QuatLookAtNew(eye, center, up).Normalize()
		return func(v Vec3) Vec3 {
			return q.Rotate(v)
		}
	}, t)
}

func TestQuatLookAtOgre(t *testing.T) {
	doLookAtTests(func(eye, center, up Vec3) translateFunc {
		q := QuatLookAtOgre(eye, center, up).Normalize()
		return func(v Vec3) Vec3 {
			return q.Rotate(v)
		}
	}, t)
}

type translateFunc func(Vec3) Vec3
type lookAtFunc func(eye, center, up Vec3) translateFunc

func doLookAtTests(f lookAtFunc, t *testing.T) {
	/*
		right-hand rule for orientation and rotation

		y+
		^
		|
		+-->x+
		 \
		  v
		   z+

		4   +----+   0
		    |\   |\
		  5 | +----+   1
		7   +-|--+ | 3
		     \|   \|
		  6   +----+   2

	*/
	cube := []Vec3{
		Vec3{1, 1, -1},
		Vec3{1, 1, 1},
		Vec3{1, -1, 1},
		Vec3{1, -1, -1},

		Vec3{-1, 1, -1},
		Vec3{-1, 1, 1},
		Vec3{-1, -1, 1},
		Vec3{-1, -1, -1},
	}

	tests := map[string]struct {
		Eye, Center, Up Vec3
		Expected        []Vec3
	}{

		"identity rotation": {
			// looking from viewer into screen z-, up y+
			Vec3{0, 0, 0}, Vec3{0, 0, -1}, Vec3{0, 1, 0},
			// identiy, thus equals cube
			cube,
		},

		"look right": {
			// x+
			// rotate around y -45 deg
			Vec3{0, 0, 0}, Vec3{1, 0, 0}, Vec3{0, 1, 0},
			/*
				5   +----+   4
				    |\   |\
				  1 | +----+   0
				6   +-|--+ | 7
				     \|   \|
				  2   +----+   3
			*/
			[]Vec3{
				cube[4], cube[0], cube[3], cube[7],
				cube[5], cube[1], cube[2], cube[6],
			},
		},

		"look down": {
			// y-
			// rotate around x -45 deg
			// up toward z-
			Vec3{0, 0, 0}, Vec3{0, -1, 0}, Vec3{0, 0, -1},
			/*
				5   +----+   1
				    |\   |\
				  6 | +----+   2
				4   +-|--+ | 0
				     \|   \|
				  7   +----+   3
			*/
			[]Vec3{
				cube[1], cube[2], cube[3], cube[0],
				cube[5], cube[6], cube[7], cube[4],
			},
		},

		"half roll": { // immelmann turn without the half roll
			// looking from screen to viewer z+
			// upside down, y-
			Vec3{0, 0, 0}, Vec3{0, 0, 1}, Vec3{0, -1, 0},
			/*
				6   +----+   2
				    |\   |\
				  7 | +----+   3
				5   +-|--+ | 1
				     \|   \|
				  4   +----+   0
			*/
			[]Vec3{
				cube[2], cube[3], cube[0], cube[1],
				cube[6], cube[7], cube[4], cube[5],
			},
		},
	}

	for d, c := range tests {
		t.Log(d)
		tf := f(c.Eye, c.Center, c.Up)

		for i, v := range cube {
			t.Log("orig:", i, v)

			if r := tf(v); !c.Expected[i].ApproxEqual(r) {
				t.Log("rotated :", r)
				t.Log("expected:", c.Expected[i])

				t.Fatalf("%s failed: %v != %v", d, c.Expected[i], r)
			}
		}
	}

}
