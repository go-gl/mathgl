package mathgl

import (
	"math"
)

func Ortho(left, right, bottom, top, near, far float64) Matrix {
	rml, tmb, fmn := (right - left), (top - bottom), (far - near)
	mp, _ := InferMatrixFromSlice([]interface{}{
		2. / rml, 0., 0., -(right + left) / rml,
		0., 2. / tmb, 0., -(top + bottom) / tmb,
		0., 0., -2. / fmn, -(far + near) / fmn,
		0., 0., 0., 1.}, 4, 4)

	return *mp
}

func Ortho2D(left, right, top, bottom float64) Matrix {
	return Ortho(left, right, top, bottom, -1, 1)
}

func Perspective(fovy, aspect, near, far float64) Matrix {
	fovy = (fovy * math.Pi) / 180.0 // convert from degrees to radians
	nmf, f := near-far, 1./math.Tan(fovy/2)

	mp, _ := InferMatrixFromSlice([]interface{}{
		f / aspect, 0., 0., 0.,
		0., f, 0., 0.,
		0., 0., (near + far) / nmf, (2 * far * near) / nmf,
		0., 0., -1., 0.}, 4, 4)

	return *mp
}

func Frustum(left, right, bottom, top, near, far float64) Matrix {
	rml, tmb, fmn := (right - left), (top - bottom), (far - near)
	A, B, C, D := (right+left)/rml, (top+bottom)/tmb, -(far+near)/fmn, (2*far*near)/fmn

	mp, _ := InferMatrixFromSlice([]interface{}{
		(2. * near) / rml, 0., A, 0.,
		0., (2. * near) / tmb, B, 0.,
		0., 0., C, D,
		0., 0., -1., 0.}, 4, 4)

	return *mp
}

func LookAt(eyeX, eyeY, eyeZ, centerX, centerY, centerZ, upX, upY, upZ float64) Matrix {
	F, _ := InferVectorOf(
		centerX-eyeX,
		centerY-eyeY,
		centerZ-eyeZ)

	f := F.Normalize()

	Up, _ := InferVectorOf(
		upX,
		upY,
		upZ)

	Upp := Up.Normalize()

	s := f.Cross(Upp)
	ss := s.AsSlice()
	us := s.Cross(f).AsSlice()
	fs := f.ScalarMul(MakeScalar(-1., FLOAT64)).AsSlice()
	z, _ := InferScalar(0.0)
	o, _ := InferScalar(1.0)

	M := *unsafeMatrixFromSlice([]Scalar{
		ss[0], ss[1], ss[2], z,
		us[0], us[1], us[2], z,
		fs[0], fs[1], fs[2], z,
		z, z, z, o}, FLOAT64, 4, 4)

	return M.Mul(Translate3D(-eyeX, -eyeY, -eyeZ))
}

func LookAtV(eye, center, up Vector) Matrix {
	F := center.Sub(eye)

	f := F.Normalize()

	Upp := up.Normalize()

	s := f.Cross(Upp)
	ss := s.AsSlice()
	us := s.Cross(f).AsSlice()
	fs := f.ScalarMul(MakeScalar(-1., FLOAT64)).AsSlice()
	z, _ := InferScalar(0.0)
	o, _ := InferScalar(1.0)

	M := *unsafeMatrixFromSlice([]Scalar{
		ss[0], ss[1], ss[2], z,
		us[0], us[1], us[2], z,
		fs[0], fs[1], fs[2], z,
		z, z, z, o}, FLOAT64, 4, 4)

	eyeX, eyeY, eyeZ := eye.GetElement(0), eye.GetElement(1), eye.GetElement(2)
	return M.Mul(Translate3D(-eyeX.Fl64(), -eyeY.Fl64(), -eyeZ.Fl64()))
}
