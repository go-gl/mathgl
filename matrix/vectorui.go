package mathgl

import "errors"

type Vec2ui [2]uint32
type Vec3ui [3]uint32
type Vec4ui [4]uint32
type Vecui []uint32

func (v Vec2ui) Dot(u Vec2ui) uint32 {
	return v[0]*u[0] + v[1]*u[1]
}

func (v Vec3ui) Dot(u Vec3ui) uint32 {
	return v[0]*u[0] + v[1]*u[1] + v[2]*u[2]
}

func (v Vec4ui) Dot(u Vec4ui) uint32 {
	return v[0]*u[0] + v[1]*u[1] + v[2]*u[2] + v[3]*u[3]
}

func (v Vecui) Dot(u Vecui) (f uint32, err error) {
	if len(v) != len(u) {
		return 0.0, errors.New("Vectors aren't the same length")
	}

	f = 0.0
	for i := range v {
		f = f + u[i]*v[i]
	}

	return f, nil
}

func (v Vec3ui) Cross(u Vec3ui) (r Vec3ui) {
	r[0], r[1], r[2] = v[1]*u[2]-v[2]*u[1], v[2]*u[0]-v[0]*u[2], v[0]*u[1]-v[1]*u[0]

	return r
}

// Allow 7-dimensional cross product? (It's also binary).
func (v Vecui) Cross(u Vecui) (r Vecui, err error) {
	if len(v) != 3 || len(u) != 3 {
		return nil, errors.New("Cross product requires both inputs to be 3-dimensional")
	}

	r = make([]uint32, 3, 3)
	r[0], r[1], r[2] = v[1]*u[2]-v[2]*u[1], v[2]*u[0]-v[0]*u[2], v[0]*u[1]-v[1]*u[0]
	return r, nil
}

func (v Vec2ui) Add(u Vec2ui) (r Vec2ui) {
	r[0], r[1] = v[0]+u[0], v[1]+u[1]

	return r
}

func (v Vec3ui) Add(u Vec3ui) (r Vec3ui) {
	r[0], r[1], r[2] = v[0]+u[0], v[1]+u[1], v[2]+u[2]

	return r
}

func (v Vec4ui) Add(u Vec4ui) (r Vec4ui) {
	r[0], r[1], r[2], r[3] = v[0]+u[0], v[1]+u[1], v[2]+u[2], v[3]+u[3]

	return r
}

func (v Vecui) Add(u Vecui) (r Vecui, err error) {
	if len(v) != len(u) {
		return nil, errors.New("Vectors must be of the same length to add")
	}

	r = make([]uint32, len(v))
	for i := range v {
		r[i] = v[i] + u[i]
	}

	return r, nil
}

func (v Vecui) Sub(u Vecui) (r Vecui, err error) {
	if len(v) != len(u) {
		return nil, errors.New("Vectors must be of the same length to add")
	}

	r = make([]uint32, len(v))
	for i := range v {
		r[i] = v[i] - u[i]
	}

	return r, nil
}

func (v Vec2ui) Sub(u Vec2ui) (r Vec2ui) {
	r[0], r[1] = v[0]-u[0], v[1]-u[1]

	return r
}

func (v Vec3ui) Sub(u Vec3ui) (r Vec3ui) {
	r[0], r[1], r[2] = v[0]-u[0], v[1]-u[1], v[2]-u[2]

	return r
}

func (v Vec4ui) Sub(u Vec4ui) (r Vec4ui) {
	r[0], r[1], r[2], r[3] = v[0]-u[0], v[1]-u[1], v[2]-u[2], v[3]-u[3]

	return r
}

func (v Vecui) Mul(c uint32) (r Vecui) {

	r = make([]uint32, len(v))
	for i := range v {
		r[i] = v[i] * c
	}

	return r
}

func (v Vec2ui) Mul(c uint32) (r Vec2ui) {
	r[0], r[1] = v[0]*c, v[1]*c

	return r
}

func (v Vec3ui) Mul(c uint32) (r Vec3ui) {
	r[0], r[1], r[2] = v[0]*c, v[1]*c, v[2]*c

	return r
}

func (v Vec4ui) Mul(c uint32) (r Vec4ui) {
	r[0], r[1], r[2], r[3] = v[0]*c, v[1]*c, v[2]*c, v[3]*c

	return r
}
