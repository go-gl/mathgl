package mathgl

import "errors"

type Vec2l [2]int64
type Vec3l [3]int64
type Vec4l [4]int64
type Vecl []int64

func (v Vec2l) Dot(u Vec2l) int64 {
	return v[0]*u[0] + v[1]*u[1]
}

func (v Vec3l) Dot(u Vec3l) int64 {
	return v[0]*u[0] + v[1]*u[1] + v[2]*u[2]
}

func (v Vec4l) Dot(u Vec4l) int64 {
	return v[0]*u[0] + v[1]*u[1] + v[2]*u[2] + v[3]*u[3]
}

func (v Vecl) Dot(u Vecl) (f int64, err error) {
	if len(v) != len(u) {
		return 0.0, errors.New("Vectors aren't the same length")
	}

	f = 0.0
	for i := range v {
		f = f + u[i]*v[i]
	}

	return f, nil
}

func (v Vec3l) Cross(u Vec3l) (r Vec3l) {
	r[0], r[1], r[2] = v[1]*u[2]-v[2]*u[1], v[2]*u[0]-v[0]*u[2], v[0]*u[1]-v[1]*u[0]

	return r
}

// Allow 7-dimensional cross product? (It's also binary).
func (v Vecl) Cross(u Vecl) (r Vecl, err error) {
	if len(v) != 3 || len(u) != 3 {
		return nil, errors.New("Cross product requires both inputs to be 3-dimensional")
	}

	r = make([]int64, 3, 3)
	r[0], r[1], r[2] = v[1]*u[2]-v[2]*u[1], v[2]*u[0]-v[0]*u[2], v[0]*u[1]-v[1]*u[0]
	return r, nil
}

func (v Vec2l) Add(u Vec2l) (r Vec2l) {
	r[0], r[1] = v[0]+u[0], v[1]+u[1]

	return r
}

func (v Vec3l) Add(u Vec3l) (r Vec3l) {
	r[0], r[1], r[2] = v[0]+u[0], v[1]+u[1], v[2]+u[2]

	return r
}

func (v Vec4l) Add(u Vec4l) (r Vec4l) {
	r[0], r[1], r[2], r[3] = v[0]+u[0], v[1]+u[1], v[2]+u[2], v[3]+u[3]

	return r
}

func (v Vecl) Add(u Vecl) (r Vecl, err error) {
	if len(v) != len(u) {
		return nil, errors.New("Vectors must be of the same length to add")
	}

	r = make([]int64, len(v))
	for i := range v {
		r[i] = v[i] + u[i]
	}

	return r, nil
}

func (v Vecl) Sub(u Vecl) (r Vecl, err error) {
	if len(v) != len(u) {
		return nil, errors.New("Vectors must be of the same length to add")
	}

	r = make([]int64, len(v))
	for i := range v {
		r[i] = v[i] - u[i]
	}

	return r, nil
}

func (v Vec2l) Sub(u Vec2l) (r Vec2l) {
	r[0], r[1] = v[0]-u[0], v[1]-u[1]

	return r
}

func (v Vec3l) Sub(u Vec3l) (r Vec3l) {
	r[0], r[1], r[2] = v[0]-u[0], v[1]-u[1], v[2]-u[2]

	return r
}

func (v Vec4l) Sub(u Vec4l) (r Vec4l) {
	r[0], r[1], r[2], r[3] = v[0]-u[0], v[1]-u[1], v[2]-u[2], v[3]-u[3]

	return r
}

func (v Vecl) Mul(c int64) (r Vecl) {

	r = make([]int64, len(v))
	for i := range v {
		r[i] = v[i] * c
	}

	return r
}

func (v Vec2l) Mul(c int64) (r Vec2l) {
	r[0], r[1] = v[0]*c, v[1]*c

	return r
}

func (v Vec3l) Mul(c int64) (r Vec3l) {
	r[0], r[1], r[2] = v[0]*c, v[1]*c, v[2]*c

	return r
}

func (v Vec4l) Mul(c int64) (r Vec4l) {
	r[0], r[1], r[2], r[3] = v[0]*c, v[1]*c, v[2]*c, v[3]*c

	return r
}
