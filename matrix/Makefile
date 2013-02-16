all:	vectors matrices

vectors: vectorf vectord vectori vectorl vectorui vectorul

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

vectori: vectorf
	cp vectorf.go vectori.go
	gofmt -w=true -r="float32 -> int32" vectori.go
	gofmt -w=true -r="Vec2f -> Vec2i" vectori.go
	gofmt -w=true -r="Vec3f -> Vec3i" vectori.go
	gofmt -w=true -r="Vec4f -> Vec4i" vectori.go
	gofmt -w=true -r="Vecf -> Veci" vectori.go
	gofmt -w=true -r="Matrix2f -> Matrix2i" vectori.go
	gofmt -w=true -r="Matrix3f -> Matrix3i" vectori.go
	gofmt -w=true -r="Matrix4f -> Matrix4i" vectori.go
	gofmt -w=true -r="Matrix2x3f -> Matrix2x3i" vectori.go
	gofmt -w=true -r="Matrix2x4f -> Matrix2x4i" vectori.go
	gofmt -w=true -r="Matrix3x2f -> Matrix3x2i" vectori.go
	gofmt -w=true -r="Matrix3x4f -> Matrix3x4i" vectori.go
	gofmt -w=true -r="Matrix4x2f -> Matrix4x2i" vectori.go
	gofmt -w=true -r="Matrix4x3f -> Matrix4x3i" vectori.go
	gofmt -w=true -r="Matrixf -> Matrixi" vectori.go
	gofmt -w=true -r="Ident2f -> Ident2i" vectori.go
	gofmt -w=true -r="Ident3f -> Ident3i" vectori.go
	gofmt -w=true -r="Ident4f -> Ident4i" vectori.go
	gofmt -w=true -r="Identf -> Identi" vectori.go

vectorl: vectorf
	cp vectorf.go vectorl.go
	gofmt -w=true -r="float32 -> int64" vectorl.go
	gofmt -w=true -r="Vec2f -> Vec2l" vectorl.go
	gofmt -w=true -r="Vec3f -> Vec3l" vectorl.go
	gofmt -w=true -r="Vec4f -> Vec4l" vectorl.go
	gofmt -w=true -r="Vecf -> Vecl" vectorl.go
	gofmt -w=true -r="Matrix2f -> Matrix2l" vectorl.go
	gofmt -w=true -r="Matrix3f -> Matrix3l" vectorl.go
	gofmt -w=true -r="Matrix4f -> Matrix4l" vectorl.go
	gofmt -w=true -r="Matrix2x3f -> Matrix2x3l" vectorl.go
	gofmt -w=true -r="Matrix2x4f -> Matrix2x4l" vectorl.go
	gofmt -w=true -r="Matrix3x2f -> Matrix3x2l" vectorl.go
	gofmt -w=true -r="Matrix3x4f -> Matrix3x4l" vectorl.go
	gofmt -w=true -r="Matrix4x2f -> Matrix4x2l" vectorl.go
	gofmt -w=true -r="Matrix4x3f -> Matrix4x3l" vectorl.go
	gofmt -w=true -r="Matrixf -> Matrixl" vectorl.go
	gofmt -w=true -r="Ident2f -> Ident2l" vectorl.go
	gofmt -w=true -r="Ident3f -> Ident3l" vectorl.go
	gofmt -w=true -r="Ident4f -> Ident4l" vectorl.go
	gofmt -w=true -r="Identf -> Identl" vectorl.go

vectorui: vectorf
	cp vectorf.go vectorui.go
	gofmt -w=true -r="float32 -> uint32" vectorui.go
	gofmt -w=true -r="Vec2f -> Vec2ui" vectorui.go
	gofmt -w=true -r="Vec3f -> Vec3ui" vectorui.go
	gofmt -w=true -r="Vec4f -> Vec4ui" vectorui.go
	gofmt -w=true -r="Vecf -> Vecui" vectorui.go
	gofmt -w=true -r="Matrix2f -> Matrix2ui" vectorui.go
	gofmt -w=true -r="Matrix3f -> Matrix3ui" vectorui.go
	gofmt -w=true -r="Matrix4f -> Matrix4ui" vectorui.go
	gofmt -w=true -r="Matrix2x3f -> Matrix2x3ui" vectorui.go
	gofmt -w=true -r="Matrix2x4f -> Matrix2x4ui" vectorui.go
	gofmt -w=true -r="Matrix3x2f -> Matrix3x2ui" vectorui.go
	gofmt -w=true -r="Matrix3x4f -> Matrix3x4ui" vectorui.go
	gofmt -w=true -r="Matrix4x2f -> Matrix4x2ui" vectorui.go
	gofmt -w=true -r="Matrix4x3f -> Matrix4x3ui" vectorui.go
	gofmt -w=true -r="Matrixf -> Matrixui" vectorui.go
	gofmt -w=true -r="Ident2f -> Ident2ui" vectorui.go
	gofmt -w=true -r="Ident3f -> Ident3ui" vectorui.go
	gofmt -w=true -r="Ident4f -> Ident4ui" vectorui.go
	gofmt -w=true -r="Identf -> Identui" vectorui.go
	
vectorul: vectorf
	cp vectorf.go vectorul.go
	gofmt -w=true -r="float32 -> uint64" vectorul.go
	gofmt -w=true -r="Vec2f -> Vec2ul" vectorul.go
	gofmt -w=true -r="Vec3f -> Vec3ul" vectorul.go
	gofmt -w=true -r="Vec4f -> Vec4ul" vectorul.go
	gofmt -w=true -r="Vecf -> Vecul" vectorul.go
	gofmt -w=true -r="Matrix2f -> Matrix2ul" vectorul.go
	gofmt -w=true -r="Matrix3f -> Matrix3ul" vectorul.go
	gofmt -w=true -r="Matrix4f -> Matrix4ul" vectorul.go
	gofmt -w=true -r="Matrix2x3f -> Matrix2x3ul" vectorul.go
	gofmt -w=true -r="Matrix2x4f -> Matrix2x4ul" vectorul.go
	gofmt -w=true -r="Matrix3x2f -> Matrix3x2ul" vectorul.go
	gofmt -w=true -r="Matrix3x4f -> Matrix3x4ul" vectorul.go
	gofmt -w=true -r="Matrix4x2f -> Matrix4x2ul" vectorul.go
	gofmt -w=true -r="Matrix4x3f -> Matrix4x3ul" vectorul.go
	gofmt -w=true -r="Matrixf -> Matrixul" vectorul.go
	gofmt -w=true -r="Ident2f -> Ident2ul" vectorul.go
	gofmt -w=true -r="Ident3f -> Ident3ul" vectorul.go
	gofmt -w=true -r="Ident4f -> Ident4ul" vectorul.go
	gofmt -w=true -r="Identf -> Identul" vectorul.go
	
matrices: matrixf matrixd matrixi matrixl matrixui matrixul

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

matrixi: matrixf
	cp matrixf.go matrixi.go
	gofmt -w=true -r="float32 -> int32" matrixi.go
	gofmt -w=true -r="Vec2f -> Vec2i" matrixi.go
	gofmt -w=true -r="Vec3f -> Vec3i" matrixi.go
	gofmt -w=true -r="Vec4f -> Vec4i" matrixi.go
	gofmt -w=true -r="Vecf -> Veci" matrixi.go
	gofmt -w=true -r="Matrix2f -> Matrix2i" matrixi.go
	gofmt -w=true -r="Matrix3f -> Matrix3i" matrixi.go
	gofmt -w=true -r="Matrix4f -> Matrix4i" matrixi.go
	gofmt -w=true -r="Matrix2x3f -> Matrix2x3i" matrixi.go
	gofmt -w=true -r="Matrix2x4f -> Matrix2x4i" matrixi.go
	gofmt -w=true -r="Matrix3x2f -> Matrix3x2i" matrixi.go
	gofmt -w=true -r="Matrix3x4f -> Matrix3x4i" matrixi.go
	gofmt -w=true -r="Matrix4x2f -> Matrix4x2i" matrixi.go
	gofmt -w=true -r="Matrix4x3f -> Matrix4x3i" matrixi.go
	gofmt -w=true -r="Matrixf -> Matrixi" matrixi.go
	gofmt -w=true -r="Ident2f -> Ident2i" matrixi.go
	gofmt -w=true -r="Ident3f -> Ident3i" matrixi.go
	gofmt -w=true -r="Ident4f -> Ident4i" matrixi.go
	gofmt -w=true -r="Identf -> Identi" matrixi.go

matrixl: matrixf
	cp matrixf.go matrixl.go
	gofmt -w=true -r="float32 -> int64" matrixl.go
	gofmt -w=true -r="Vec2f -> Vec2l" matrixl.go
	gofmt -w=true -r="Vec3f -> Vec3l" matrixl.go
	gofmt -w=true -r="Vec4f -> Vec4l" matrixl.go
	gofmt -w=true -r="Vecf -> Vecl" matrixl.go
	gofmt -w=true -r="Matrix2f -> Matrix2l" matrixl.go
	gofmt -w=true -r="Matrix3f -> Matrix3l" matrixl.go
	gofmt -w=true -r="Matrix4f -> Matrix4l" matrixl.go
	gofmt -w=true -r="Matrix2x3f -> Matrix2x3l" matrixl.go
	gofmt -w=true -r="Matrix2x4f -> Matrix2x4l" matrixl.go
	gofmt -w=true -r="Matrix3x2f -> Matrix3x2l" matrixl.go
	gofmt -w=true -r="Matrix3x4f -> Matrix3x4l" matrixl.go
	gofmt -w=true -r="Matrix4x2f -> Matrix4x2l" matrixl.go
	gofmt -w=true -r="Matrix4x3f -> Matrix4x3l" matrixl.go
	gofmt -w=true -r="Matrixf -> Matrixl" matrixl.go
	gofmt -w=true -r="Ident2f -> Ident2l" matrixl.go
	gofmt -w=true -r="Ident3f -> Ident3l" matrixl.go
	gofmt -w=true -r="Ident4f -> Ident4l" matrixl.go
	gofmt -w=true -r="Identf -> Identl" matrixl.go

matrixui: matrixf
	cp matrixf.go matrixui.go
	gofmt -w=true -r="float32 -> uint32" matrixui.go
	gofmt -w=true -r="Vec2f -> Vec2ui" matrixui.go
	gofmt -w=true -r="Vec3f -> Vec3ui" matrixui.go
	gofmt -w=true -r="Vec4f -> Vec4ui" matrixui.go
	gofmt -w=true -r="Vecf -> Vecui" matrixui.go
	gofmt -w=true -r="Matrix2f -> Matrix2ui" matrixui.go
	gofmt -w=true -r="Matrix3f -> Matrix3ui" matrixui.go
	gofmt -w=true -r="Matrix4f -> Matrix4ui" matrixui.go
	gofmt -w=true -r="Matrix2x3f -> Matrix2x3ui" matrixui.go
	gofmt -w=true -r="Matrix2x4f -> Matrix2x4ui" matrixui.go
	gofmt -w=true -r="Matrix3x2f -> Matrix3x2ui" matrixui.go
	gofmt -w=true -r="Matrix3x4f -> Matrix3x4ui" matrixui.go
	gofmt -w=true -r="Matrix4x2f -> Matrix4x2ui" matrixui.go
	gofmt -w=true -r="Matrix4x3f -> Matrix4x3ui" matrixui.go
	gofmt -w=true -r="Matrixf -> Matrixui" matrixui.go
	gofmt -w=true -r="Ident2f -> Ident2ui" matrixui.go
	gofmt -w=true -r="Ident3f -> Ident3ui" matrixui.go
	gofmt -w=true -r="Ident4f -> Ident4ui" matrixui.go
	gofmt -w=true -r="Identf -> Identui" matrixui.go
	
matrixul: matrixf
	cp matrixf.go matrixul.go
	gofmt -w=true -r="float32 -> uint64" matrixul.go
	gofmt -w=true -r="Vec2f -> Vec2ul" matrixul.go
	gofmt -w=true -r="Vec3f -> Vec3ul" matrixul.go
	gofmt -w=true -r="Vec4f -> Vec4ul" matrixul.go
	gofmt -w=true -r="Vecf -> Vecul" matrixul.go
	gofmt -w=true -r="Matrix2f -> Matrix2ul" matrixul.go
	gofmt -w=true -r="Matrix3f -> Matrix3ul" matrixul.go
	gofmt -w=true -r="Matrix4f -> Matrix4ul" matrixul.go
	gofmt -w=true -r="Matrix2x3f -> Matrix2x3ul" matrixul.go
	gofmt -w=true -r="Matrix2x4f -> Matrix2x4ul" matrixul.go
	gofmt -w=true -r="Matrix3x2f -> Matrix3x2ul" matrixul.go
	gofmt -w=true -r="Matrix3x4f -> Matrix3x4ul" matrixul.go
	gofmt -w=true -r="Matrix4x2f -> Matrix4x2ul" matrixul.go
	gofmt -w=true -r="Matrix4x3f -> Matrix4x3ul" matrixul.go
	gofmt -w=true -r="Matrixf -> Matrixul" matrixul.go
	gofmt -w=true -r="Ident2f -> Ident2ul" matrixul.go
	gofmt -w=true -r="Ident3f -> Ident3ul" matrixul.go
	gofmt -w=true -r="Ident4f -> Ident4ul" matrixul.go
	gofmt -w=true -r="Identf -> Identul" matrixul.go