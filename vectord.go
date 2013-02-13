package mathgl

import "errors"

type Vec2d [2]float64
type Vec3d [3]float64
type Vec4d [4]float64
type Vecd []float64

func (v Vec2d) Dot(u Vec2d) float64 {
	return v[0]*u[0] + v[1]*u[1]
}

func (v Vec3d) Dot(u Vec3d) float64 {
	return v[0]*u[0] + v[1]*u[1] + v[2]*u[2]
}

func (v Vec4d) Dot(u Vec4d) float64 {
	return v[0]*u[0] + v[1]*u[1] + v[2]*u[2] + v[3]*u[3]
}

func (v Vecd) Dot(u Vecd) (f float64, err error) {
	if len(v) != len(u) {
		return 0.0, errors.New("Vectors aren't the same length")
	}

	f = 0.0
	for i := range v {
		f = f + u[i]*v[i]
	}

	return f, nil
}

func (v Vec3d) Cross(u Vec3d) (r Vec3d) {
	r[0], r[1], r[2] = v[1]*u[2]-v[2]*u[1], v[2]*u[0]-v[0]*u[2], v[0]*u[1]-v[1]*u[0]

	return r
}

// Allow 7-dimensional cross product? (It's also binary).
func (v Vecd) Cross(u Vecd) (r Vecd, err error) {
	if len(v) != 3 || len(u) != 3 {
		return nil, errors.New("Cross product requires both inputs to be 3-dimensional")
	}

	r = make([]float64, 3, 3)
	r[0], r[1], r[2] = v[1]*u[2]-v[2]*u[1], v[2]*u[0]-v[0]*u[2], v[0]*u[1]-v[1]*u[0]
	return r, nil
}

func (v Vec2d) Add(u Vec2d) (r Vec2d) {
	r[0], r[1] = v[0]+u[0], v[1]+u[1]

	return r
}

func (v Vec3d) Add(u Vec3d) (r Vec3d) {
	r[0], r[1], r[2] = v[0]+u[0], v[1]+u[1], v[2]+u[2]

	return r
}

func (v Vec4d) Add(u Vec4d) (r Vec4d) {
	r[0], r[1], r[2], r[3] = v[0]+u[0], v[1]+u[1], v[2]+u[2], v[3]+u[3]

	return r
}

func (v Vecd) Add(u Vecd) (r Vecd, err error) {
	if len(v) != len(u) {
		return nil, errors.New("Vectors must be of the same length to add")
	}

	r = make([]float64, len(v))
	for i := range v {
		r[i] = v[i] + u[i]
	}

	return r, nil
}

func (v Vecd) Sub(u Vecd) (r Vecd, err error) {
	if len(v) != len(u) {
		return nil, errors.New("Vectors must be of the same length to add")
	}

	r = make([]float64, len(v))
	for i := range v {
		r[i] = v[i] - u[i]
	}

	return r, nil
}

func (v Vec2d) Sub(u Vec2d) (r Vec2d) {
	r[0], r[1] = v[0]-u[0], v[1]-u[1]

	return r
}

func (v Vec3d) Sub(u Vec3d) (r Vec3d) {
	r[0], r[1], r[2] = v[0]-u[0], v[1]-u[1], v[2]-u[2]

	return r
}

func (v Vec4d) Sub(u Vec4d) (r Vec4d) {
	r[0], r[1], r[2], r[3] = v[0]-u[0], v[1]-u[1], v[2]-u[2], v[3]-u[3]

	return r
}

func (v Vecd) Mul(c float64) (r Vecd) {

	r = make([]float64, len(v))
	for i := range v {
		r[i] = v[i] * c
	}

	return r
}

func (v Vec2d) Mul(c float64) (r Vec2d) {
	r[0], r[1] = v[0]*c, v[1]*c

	return r
}

func (v Vec3d) Mul(c float64) (r Vec3d) {
	r[0], r[1], r[2] = v[0]*c, v[1]*c, v[2]*c

	return r
}

func (v Vec4d) Mul(c float64) (r Vec4d) {
	r[0], r[1], r[2], r[3] = v[0]*c, v[1]*c, v[2]*c, v[3]*c

	return r
}
