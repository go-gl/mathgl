package mathgl

import "errors"

type Vec2f [2]float32
type Vec3f [3]float32
type Vec4f [4]float32
type Vecf []float32

func (v Vec2f) Dot(u Vec2f) float32 {
	return v[0]*u[0] + v[1]*u[1]
}

func (v Vec3f) Dot(u Vec3f) float32 {
	return v[0]*u[0] + v[1]*u[1] + v[2]*u[2]
}

func (v Vec4f) Dot(u Vec4f) float32 {
	return v[0]*u[0] + v[1]*u[1] + v[2]*u[2] + v[3]*u[3]
}

func (v Vecf) Dot(u Vecf) (f float32, err error) {
	if len(v) != len(u) {
		return 0.0, errors.New("Vectors aren't the same length")
	}

	f = 0.0
	for i := range v {
		f = f + u[i]*v[i]
	}

	return f, nil
}

func (v Vec3f) Cross(u Vec3f) (r Vec3f) {
	r[0], r[1], r[2] = v[1]*u[2]-v[2]*u[1], v[2]*u[0]-v[0]*u[2], v[0]*u[1]-v[1]*u[0]

	return r
}

// Allow 7-dimensional cross product? (It's also binary).
func (v Vecf) Cross(u Vecf) (r Vecf, err error) {
	if len(v) != 3 || len(u) != 3 {
		return nil, errors.New("Cross product requires both inputs to be 3-dimensional")
	}

	r = make([]float32, 3, 3)
	r[0], r[1], r[2] = v[1]*u[2]-v[2]*u[1], v[2]*u[0]-v[0]*u[2], v[0]*u[1]-v[1]*u[0]
	return r, nil
}

func (v Vec2f) Add(u Vec2f) (r Vec2f) {
	r[0], r[1] = v[0]+u[0], v[1]+u[1]

	return r
}

func (v Vec3f) Add(u Vec3f) (r Vec3f) {
	r[0], r[1], r[2] = v[0]+u[0], v[1]+u[1], v[2]+u[2]

	return r
}

func (v Vec4f) Add(u Vec4f) (r Vec4f) {
	r[0], r[1], r[2], r[3] = v[0]+u[0], v[1]+u[1], v[2]+u[2], v[3]+u[3]

	return r
}

func (v Vecf) Add(u Vecf) (r Vecf, err error) {
	if len(v) != len(u) {
		return nil, errors.New("Vectors must be of the same length to add")
	}

	r = make([]float32, len(v))
	for i := range v {
		r[i] = v[i] + u[i]
	}

	return r, nil
}

func (v Vecf) Sub(u Vecf) (r Vecf, err error) {
	if len(v) != len(u) {
		return nil, errors.New("Vectors must be of the same length to add")
	}

	r = make([]float32, len(v))
	for i := range v {
		r[i] = v[i] - u[i]
	}

	return r, nil
}

func (v Vec2f) Sub(u Vec2f) (r Vec2f) {
	r[0], r[1] = v[0]-u[0], v[1]-u[1]

	return r
}

func (v Vec3f) Sub(u Vec3f) (r Vec3f) {
	r[0], r[1], r[2] = v[0]-u[0], v[1]-u[1], v[2]-u[2]

	return r
}

func (v Vec4f) Sub(u Vec4f) (r Vec4f) {
	r[0], r[1], r[2], r[3] = v[0]-u[0], v[1]-u[1], v[2]-u[2], v[3]-u[3]

	return r
}

func (v Vecf) Mul(c float32) (r Vecf) {

	r = make([]float32, len(v))
	for i := range v {
		r[i] = v[i] * c
	}

	return r
}

func (v Vec2f) Mul(c float32) (r Vec2f) {
	r[0], r[1] = v[0]*c, v[1]*c

	return r
}

func (v Vec3f) Mul(c float32) (r Vec3f) {
	r[0], r[1], r[2] = v[0]*c, v[1]*c, v[2]*c

	return r
}

func (v Vec4f) Mul(c float32) (r Vec4f) {
	r[0], r[1], r[2], r[3] = v[0]*c, v[1]*c, v[2]*c, v[3]*c

	return r
}
