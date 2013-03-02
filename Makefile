all:	vectors matrices quaternions

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
	gofmt -w=true -r="Mat2f -> Mat2d" vectord.go
	gofmt -w=true -r="Mat3f -> Mat3d" vectord.go
	gofmt -w=true -r="Mat4f -> Mat4d" vectord.go
	gofmt -w=true -r="Mat2x3f -> Mat2x3d" vectord.go
	gofmt -w=true -r="Mat2x4f -> Mat2x4d" vectord.go
	gofmt -w=true -r="Mat3x2f -> Mat3x2d" vectord.go
	gofmt -w=true -r="Mat3x4f -> Mat3x4d" vectord.go
	gofmt -w=true -r="Mat4x2f -> Mat4x2d" vectord.go
	gofmt -w=true -r="Mat4x3f -> Mat4x3d" vectord.go
	gofmt -w=true -r="Matf -> Matd" vectord.go
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
	gofmt -w=true -r="Mat2f -> Mat2d" matrixd.go
	gofmt -w=true -r="Mat3f -> Mat3d" matrixd.go
	gofmt -w=true -r="Mat4f -> Mat4d" matrixd.go
	gofmt -w=true -r="Mat2x3f -> Mat2x3d" matrixd.go
	gofmt -w=true -r="Mat2x4f -> Mat2x4d" matrixd.go
	gofmt -w=true -r="Mat3x2f -> Mat3x2d" matrixd.go
	gofmt -w=true -r="Mat3x4f -> Mat3x4d" matrixd.go
	gofmt -w=true -r="Mat4x2f -> Mat4x2d" matrixd.go
	gofmt -w=true -r="Mat4x3f -> Mat4x3d" matrixd.go
	gofmt -w=true -r="Matf -> Matd" matrixd.go
	gofmt -w=true -r="Ident2f -> Ident2d" matrixd.go
	gofmt -w=true -r="Ident3f -> Ident3d" matrixd.go
	gofmt -w=true -r="Ident4f -> Ident4d" matrixd.go
	gofmt -w=true -r="Identf -> Identd" matrixd.go

quaternions: quatf quatd

quatf: quatf.go
	gofmt -w=true quatf.go

quatd: quatf
	cp quatf.go quatd.go
	gofmt -w=true -r="float32 -> float64" quatd.go
	gofmt -w=true -r="Vec3f -> Vec3d" quatd.go
	gofmt -w=true -r="Quatf -> Quatd" quatd.go
	gofmt -w=true -r="Mat4f -> Mat4d" quatd.go
	gofmt -w=true -r="QuatIdentf -> QuatIdentd" quatd.go
	gofmt -w=true -r="QuatRotatef -> QuatRotated" quatd.go
	gofmt -w=true -r="QuatSlerpf -> QuatSlerpd" quatd.go
	gofmt -w=true -r="QuatLerpf -> QuatLerpd" quatd.go
	gofmt -w=true -r="QuatNlerpf -> QuatNlerpd" quatd.go
	gofmt -w=true -r="Clampf -> Clampd" quatd.go
