all:	vectors matrices

vectors: vectorf vectord

vectorf: vectorf.go
	gofmt -w=true vectorf.go

vectord: vectorf
	cp vectorf.go vectord.go
	gofmt -w=true -r="float32 -> float64" vectord.go
	gofmt -w=true -r="Vec2f -> Vec2d" vectord.go
	gofmt -w=true -r="Vec3f -> Vec3d" vectord.go
	gofmt -w=true -r="Vec4f -> Vec4d" vectord.go
	gofmt -w=true -r="Vecf -> Vecd" vectord.go
	gofmt -w=true -r="Matrix2f -> Matrix2d" vectord.go
	gofmt -w=true -r="Matrix3f -> Matrix3d" vectord.go
	gofmt -w=true -r="Matrix4f -> Matrix4d" vectord.go
	gofmt -w=true -r="Matrix2x3f -> Matrix2x3d" vectord.go
	gofmt -w=true -r="Matrix2x4f -> Matrix2x4d" vectord.go
	gofmt -w=true -r="Matrix3x2f -> Matrix3x2d" vectord.go
	gofmt -w=true -r="Matrix3x4f -> Matrix3x4d" vectord.go
	gofmt -w=true -r="Matrix4x2f -> Matrix4x2d" vectord.go
	gofmt -w=true -r="Matrix4x3f -> Matrix4x3d" vectord.go
	gofmt -w=true -r="Matrixf -> Matrixd" vectord.go
	gofmt -w=true -r="Ident2f -> Ident2d" vectord.go
	gofmt -w=true -r="Ident3f -> Ident3d" vectord.go
	gofmt -w=true -r="Ident4f -> Ident4d" vectord.go
	gofmt -w=true -r="Identf -> Identd" vectord.go
	
matrices: matrixf matrixd

matrixf: matrixf.go
	gofmt -w=true matrixf.go

matrixd: matrixf
	cp matrixf.go matrixd.go
	gofmt -w=true -r="float32 -> float64" matrixd.go
	gofmt -w=true -r="Vec2f -> Vec2d" matrixd.go
	gofmt -w=true -r="Vec3f -> Vec3d" matrixd.go
	gofmt -w=true -r="Vec4f -> Vec4d" matrixd.go
	gofmt -w=true -r="Vecf -> Vecd" matrixd.go
	gofmt -w=true -r="Matrix2f -> Matrix2d" matrixd.go
	gofmt -w=true -r="Matrix3f -> Matrix3d" matrixd.go
	gofmt -w=true -r="Matrix4f -> Matrix4d" matrixd.go
	gofmt -w=true -r="Matrix2x3f -> Matrix2x3d" matrixd.go
	gofmt -w=true -r="Matrix2x4f -> Matrix2x4d" matrixd.go
	gofmt -w=true -r="Matrix3x2f -> Matrix3x2d" matrixd.go
	gofmt -w=true -r="Matrix3x4f -> Matrix3x4d" matrixd.go
	gofmt -w=true -r="Matrix4x2f -> Matrix4x2d" matrixd.go
	gofmt -w=true -r="Matrix4x3f -> Matrix4x3d" matrixd.go
	gofmt -w=true -r="Matrixf -> Matrixd" matrixd.go
	gofmt -w=true -r="Ident2f -> Ident2d" matrixd.go
	gofmt -w=true -r="Ident3f -> Ident3d" matrixd.go
	gofmt -w=true -r="Ident4f -> Ident4d" matrixd.go
	gofmt -w=true -r="Identf -> Identd" matrixd.go