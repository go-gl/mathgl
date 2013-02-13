package mathgl

import "errors"

type Vec2f [2]float32
type Vec2d [2]float64
type Vec2i [2]int32
type Vec2l [2]int64

type Vec3f [3]float32
type Vec3d [3]float64
type Vec3i [3]int32
type Vec3l [3]int64

type Vecf  []float32
type Vecd []float64
type Veci []int32
type Vecl []int64

func (v Vec2f) Dot2f(u Vec2f) float32 {
	return v[0] * u[0] + v[1] * u[1]
}

func (v Vec2d) Dot2d(u Vec2d) float64 {
	return v[0] * u[0] + v[1] * u[1]
}

func (v Vec2i) Dot2i(u Vec2i) int32 {
	return v[0] * u[0] + v[1] * u[1]
}

func (v Vec2l) Dot2l(u Vec2l) int64 {
	return v[0] * u[0] + v[1] * u[1]
}

func (v Vec3f) Dot3f(u Vec3f) float32 {
	return v[0] * u[0] + v[1] * u[1] + v[2] * u[2]
}

func (v Vec3d) Dot3d(u Vec3d) float64 {
	return v[0] * u[0] + v[1] * u[1] + v[2] * u[2]
}

func (v Vec3i) Dot3i(u Vec3i) int32 {
	return v[0] * u[0] + v[1] * u[1] + v[2] * u[2]
}

func (v Vec3l) Dot3l(u Vec3l) int64 {
	return v[0] * u[0] + v[1] * u[1] + v[2] * u[2]
}

func (v Vecf) Dotf(u Vecf) (f float32, err error) {
	if len(v) != len(u) {
		return 0.0 ,errors.New("Vectors aren't the same length")
	}
	
	f = 0.0
	for i := range v {
		f = f + u[i]*v[i]
	}
	
	return f,nil
}

func (v Vecd) Dotd(u Vecd) (f float64, err error) {
	if len(v) != len(u) {
		return 0.0 ,errors.New("Vectors aren't the same length")
	}
	
	f = 0.0
	for i := range v {
		f = f + u[i]*v[i]
	}
	
	return f,nil
}

func (v Veci) Doti(u Veci) (f int32, err error) {
	if len(v) != len(u) {
		return 0.0 ,errors.New("Vectors aren't the same length")
	}
	
	f = 0.0
	for i := range v {
		f = f + u[i]*v[i]
	}
	
	return f,nil
}

func (v Vecl) Dotl(u Vecl) (f int64, err error) {
	if len(v) != len(u) {
		return 0.0 ,errors.New("Vectors aren't the same length")
	}
	
	f = 0.0
	for i := range v {
		f = f + u[i]*v[i]
	}
	
	return f,nil
}

func (v Vec3f) Cross3f(u Vec3f) (r Vec3f) {
	r[0],r[1],r[2] = v[1]*u[2]-v[2]*u[1], v[2]*u[0]-v[0]*u[2], v[0]*u[1]-v[1]*u[0]
	
	return r
}

func (v Vec3d) Cross3d(u Vec3d) (r Vec3d) {
	r[0],r[1],r[2] = v[1]*u[2]-v[2]*u[1], v[2]*u[0]-v[0]*u[2], v[0]*u[1]-v[1]*u[0]
	
	return r
}

func (v Vec3i) Cross3i(u Vec3i) (r Vec3i) {
	r[0],r[1],r[2] = v[1]*u[2]-v[2]*u[1], v[2]*u[0]-v[0]*u[2], v[0]*u[1]-v[1]*u[0]
	
	return r
}

func (v Vec3l) Cross3l(u Vec3l) (r Vec3l) {
	r[0],r[1],r[2] = v[1]*u[2]-v[2]*u[1], v[2]*u[0]-v[0]*u[2], v[0]*u[1]-v[1]*u[0]
	
	return r
}

// Allow 7-dimensional cross product? (It's also binary).
func (v Vecf) Crossf(u Vecf) (r Vecf, err error) {
	if len(v) != 3 || len(u) != 3 {
		return nil, errors.New("Cross product requires both inputs to be 3-dimensional")
	}
	
	r = make([]float32,3,3)
	r[0],r[1],r[2] = v[1]*u[2]-v[2]*u[1], v[2]*u[0]-v[0]*u[2], v[0]*u[1]-v[1]*u[0]
	return r, nil
}

func (v Vecd) Crossd(u Vecd) (r Vecd, err error) {
	if len(v) != 3 || len(u) != 3 {
		return nil, errors.New("Cross product requires both inputs to be 3-dimensional")
	}
	
	r = make([]float64,3,3)
	r[0],r[1],r[2] = v[1]*u[2]-v[2]*u[1], v[2]*u[0]-v[0]*u[2], v[0]*u[1]-v[1]*u[0]
	return r, nil
}

func (v Veci) Crossf(u Veci) (r Veci, err error) {
	if len(v) != 3 || len(u) != 3 {
		return nil, errors.New("Cross product requires both inputs to be 3-dimensional")
	}
	
	r = make([]int32,3,3)
	r[0],r[1],r[2] = v[1]*u[2]-v[2]*u[1], v[2]*u[0]-v[0]*u[2], v[0]*u[1]-v[1]*u[0]
	return r, nil
}

func (v Vecl) Crossf(u Vecl) (r Vecl, err error) {
	if len(v) != 3 || len(u) != 3 {
		return nil, errors.New("Cross product requires both inputs to be 3-dimensional")
	}
	
	r = make([]int64,3,3)
	r[0],r[1],r[2] = v[1]*u[2]-v[2]*u[1], v[2]*u[0]-v[0]*u[2], v[0]*u[1]-v[1]*u[0]
	return r, nil
}