package mathgl

import "errors"

type Vec2i [2]int32
type Vec3i [3]int32
type Vec4i [4]int32
type Veci []int32

func (v Vec2i) Dot(u Vec2i) int32 {
	return v[0]*u[0] + v[1]*u[1]
}

func (v Vec3i) Dot(u Vec3i) int32 {
	return v[0]*u[0] + v[1]*u[1] + v[2]*u[2]
}

func (v Vec4i) Dot(u Vec4i) int32 {
	return v[0]*u[0] + v[1]*u[1] + v[2]*u[2] + v[3]*u[3]
}

func (v Veci) Dot(u Veci) (f int32, err error) {
	if len(v) != len(u) {
		return 0.0, errors.New("Vectors aren't the same length")
	}

	f = 0.0
	for i := range v {
		f = f + u[i]*v[i]
	}

	return f, nil
}

func (v Vec3i) Cross(u Vec3i) (r Vec3i) {
	r[0], r[1], r[2] = v[1]*u[2]-v[2]*u[1], v[2]*u[0]-v[0]*u[2], v[0]*u[1]-v[1]*u[0]

	return r
}

// Allow 7-dimensional cross product? (It's also binary).
func (v Veci) Cross(u Veci) (r Veci, err error) {
	if len(v) != 3 || len(u) != 3 {
		return nil, errors.New("Cross product requires both inputs to be 3-dimensional")
	}

	r = make([]int32, 3, 3)
	r[0], r[1], r[2] = v[1]*u[2]-v[2]*u[1], v[2]*u[0]-v[0]*u[2], v[0]*u[1]-v[1]*u[0]
	return r, nil
}

func (v Vec2i) Add(u Vec2i) (r Vec2i) {
	r[0], r[1] = v[0]+u[0], v[1]+u[1]

	return r
}

func (v Vec3i) Add(u Vec3i) (r Vec3i) {
	r[0], r[1], r[2] = v[0]+u[0], v[1]+u[1], v[2]+u[2]

	return r
}

func (v Vec4i) Add(u Vec4i) (r Vec4i) {
	r[0], r[1], r[2], r[3] = v[0]+u[0], v[1]+u[1], v[2]+u[2], v[3]+u[3]

	return r
}

func (v Veci) Add(u Veci) (r Veci, err error) {
	if len(v) != len(u) {
		return nil, errors.New("Vectors must be of the same length to add")
	}

	r = make([]int32, len(v))
	for i := range v {
		r[i] = v[i] + u[i]
	}

	return r, nil
}

func (v Veci) Sub(u Veci) (r Veci, err error) {
	if len(v) != len(u) {
		return nil, errors.New("Vectors must be of the same length to add")
	}

	r = make([]int32, len(v))
	for i := range v {
		r[i] = v[i] - u[i]
	}

	return r, nil
}

func (v Vec2i) Sub(u Vec2i) (r Vec2i) {
	r[0], r[1] = v[0]-u[0], v[1]-u[1]

	return r
}

func (v Vec3i) Sub(u Vec3i) (r Vec3i) {
	r[0], r[1], r[2] = v[0]-u[0], v[1]-u[1], v[2]-u[2]

	return r
}

func (v Vec4i) Sub(u Vec4i) (r Vec4i) {
	r[0], r[1], r[2], r[3] = v[0]-u[0], v[1]-u[1], v[2]-u[2], v[3]-u[3]

	return r
}

func (v Veci) Mul(c int32) (r Veci) {

	r = make([]int32, len(v))
	for i := range v {
		r[i] = v[i] * c
	}

	return r
}

func (v Vec2i) Mul(c int32) (r Vec2i) {
	r[0], r[1] = v[0]*c, v[1]*c

	return r
}

func (v Vec3i) Mul(c int32) (r Vec3i) {
	r[0], r[1], r[2] = v[0]*c, v[1]*c, v[2]*c

	return r
}

func (v Vec4i) Mul(c int32) (r Vec4i) {
	r[0], r[1], r[2], r[3] = v[0]*c, v[1]*c, v[2]*c, v[3]*c

	return r
}
