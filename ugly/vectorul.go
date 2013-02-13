package mathgl

import "errors"

type Vec2ul [2]uint64
type Vec3ul [3]uint64
type Vec4ul [4]uint64
type Vecul []uint64

func (v Vec2ul) Dot(u Vec2ul) uint64 {
	return v[0]*u[0] + v[1]*u[1]
}

func (v Vec3ul) Dot(u Vec3ul) uint64 {
	return v[0]*u[0] + v[1]*u[1] + v[2]*u[2]
}

func (v Vec4ul) Dot(u Vec4ul) uint64 {
	return v[0]*u[0] + v[1]*u[1] + v[2]*u[2] + v[3]*u[3]
}

func (v Vecul) Dot(u Vecul) (f uint64, err error) {
	if len(v) != len(u) {
		return 0.0, errors.New("Vectors aren't the same length")
	}

	f = 0.0
	for i := range v {
		f = f + u[i]*v[i]
	}

	return f, nil
}

func (v Vec3ul) Cross(u Vec3ul) (r Vec3ul) {
	r[0], r[1], r[2] = v[1]*u[2]-v[2]*u[1], v[2]*u[0]-v[0]*u[2], v[0]*u[1]-v[1]*u[0]

	return r
}

// Allow 7-dimensional cross product? (It's also binary).
func (v Vecul) Cross(u Vecul) (r Vecul, err error) {
	if len(v) != 3 || len(u) != 3 {
		return nil, errors.New("Cross product requires both inputs to be 3-dimensional")
	}

	r = make([]uint64, 3, 3)
	r[0], r[1], r[2] = v[1]*u[2]-v[2]*u[1], v[2]*u[0]-v[0]*u[2], v[0]*u[1]-v[1]*u[0]
	return r, nil
}

func (v Vec2ul) Add(u Vec2ul) (r Vec2ul) {
	r[0], r[1] = v[0]+u[0], v[1]+u[1]

	return r
}

func (v Vec3ul) Add(u Vec3ul) (r Vec3ul) {
	r[0], r[1], r[2] = v[0]+u[0], v[1]+u[1], v[2]+u[2]

	return r
}

func (v Vec4ul) Add(u Vec4ul) (r Vec4ul) {
	r[0], r[1], r[2], r[3] = v[0]+u[0], v[1]+u[1], v[2]+u[2], v[3]+u[3]

	return r
}

func (v Vecul) Add(u Vecul) (r Vecul, err error) {
	if len(v) != len(u) {
		return nil, errors.New("Vectors must be of the same length to add")
	}

	r = make([]uint64, len(v))
	for i := range v {
		r[i] = v[i] + u[i]
	}

	return r, nil
}

func (v Vecul) Sub(u Vecul) (r Vecul, err error) {
	if len(v) != len(u) {
		return nil, errors.New("Vectors must be of the same length to add")
	}

	r = make([]uint64, len(v))
	for i := range v {
		r[i] = v[i] - u[i]
	}

	return r, nil
}

func (v Vec2ul) Sub(u Vec2ul) (r Vec2ul) {
	r[0], r[1] = v[0]-u[0], v[1]-u[1]

	return r
}

func (v Vec3ul) Sub(u Vec3ul) (r Vec3ul) {
	r[0], r[1], r[2] = v[0]-u[0], v[1]-u[1], v[2]-u[2]

	return r
}

func (v Vec4ul) Sub(u Vec4ul) (r Vec4ul) {
	r[0], r[1], r[2], r[3] = v[0]-u[0], v[1]-u[1], v[2]-u[2], v[3]-u[3]

	return r
}

func (v Vecul) Mul(c uint64) (r Vecul) {

	r = make([]uint64, len(v))
	for i := range v {
		r[i] = v[i] * c
	}

	return r
}

func (v Vec2ul) Mul(c uint64) (r Vec2ul) {
	r[0], r[1] = v[0]*c, v[1]*c

	return r
}

func (v Vec3ul) Mul(c uint64) (r Vec3ul) {
	r[0], r[1], r[2] = v[0]*c, v[1]*c, v[2]*c

	return r
}

func (v Vec4ul) Mul(c uint64) (r Vec4ul) {
	r[0], r[1], r[2], r[3] = v[0]*c, v[1]*c, v[2]*c, v[3]*c

	return r
}
