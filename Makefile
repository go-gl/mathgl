all:	vectors matrices quaternions transforms
	go fmt

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
	gofmt -w=true -r="FloatEqual32 -> FloatEqual" vectord.go
	gofmt -w=true -r="FloatEqualThreshold32 -> FloatEqualThreshold" vectord.go
	
matrices: matrixf matrixd

matrixf: matrixf.go
	gofmt -w=true matrixf.go

matrixd: matrixf
	cp matrixf.go matrixd.go
	gofmt -w=true -r="float32 -> float64" matrixd.go
	gofmt -w=true -r="FloatEqual32 -> FloatEqual" matrixd.go
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
	gofmt -w=true -r="FloatEqual32 -> FloatEqual" matrixd.go
	gofmt -w=true -r="FloatEqualThreshold32 -> FloatEqualThreshold" matrixd.go

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
	gofmt -w=true -r="FloatEqual32 -> FloatEqual" quatd.go
	gofmt -w=true -r="FloatEqualThreshold32 -> FloatEqualThreshold" quatd.go
	gofmt -w=true -r="EulerToQuatf -> EulerToQuatd" quatd.go

transforms: projectf projectd transformf transformd

projectf:
	gofmt -w=true projectf.go

transformf:
	gofmt -w=true transformf.go

projectd:
	cp projectf.go projectd.go
	gofmt -w=true -r="float32 -> float64" projectd.go
	gofmt -w=true -r="FloatEqual32 -> FloatEqual" projectd.go
	gofmt -w=true -r="Vec2f -> Vec2d" projectd.go
	gofmt -w=true -r="Vec3f -> Vec3d" projectd.go
	gofmt -w=true -r="Vec4f -> Vec4d" projectd.go
	gofmt -w=true -r="Vecf -> Vecd" projectd.go
	gofmt -w=true -r="Mat2f -> Mat2d" projectd.go
	gofmt -w=true -r="Mat3f -> Mat3d" projectd.go
	gofmt -w=true -r="Mat4f -> Mat4d" projectd.go
	gofmt -w=true -r="Mat2x3f -> Mat2x3d" projectd.go
	gofmt -w=true -r="Mat2x4f -> Mat2x4d" projectd.go
	gofmt -w=true -r="Mat3x2f -> Mat3x2d" projectd.go
	gofmt -w=true -r="Mat3x4f -> Mat3x4d" projectd.go
	gofmt -w=true -r="Mat4x2f -> Mat4x2d" projectd.go
	gofmt -w=true -r="Mat4x3f -> Mat4x3d" projectd.go
	gofmt -w=true -r="Matf -> Matd" projectd.go
	gofmt -w=true -r="Ident2f -> Ident2d" projectd.go
	gofmt -w=true -r="Ident3f -> Ident3d" projectd.go
	gofmt -w=true -r="Ident4f -> Ident4d" projectd.go
	gofmt -w=true -r="Identf -> Identd" projectd.go
	gofmt -w=true -r="FloatEqual32 -> FloatEqual" projectd.go
	gofmt -w=true -r="FloatEqualThreshold32 -> FloatEqualThreshold" projectd.go
	gofmt -w=true -r="Quatf -> Quatd" projectd.go
	gofmt -w=true -r="QuatIdentf -> QuatIdentd" projectd.go
	gofmt -w=true -r="QuatRotatef -> QuatRotated" projectd.go
	gofmt -w=true -r="QuatSlerpf -> QuatSlerpd" projectd.go
	gofmt -w=true -r="QuatLerpf -> QuatLerpd" projectd.go
	gofmt -w=true -r="QuatNlerpf -> QuatNlerpd" projectd.go
	gofmt -w=true -r="Clampf -> Clampd" projectd.go
	gofmt -w=true -r="Rotate2D -> Rotate2Dd" projectd.go
	gofmt -w=true -r="Rotate3DX -> Rotate3DXd" projectd.go
	gofmt -w=true -r="Rotate3DY -> Rotate3DYd" projectd.go
	gofmt -w=true -r="Rotate3DZ -> Rotate3DZd" projectd.go
	gofmt -w=true -r="Translate2D -> Translate2Dd" projectd.go
	gofmt -w=true -r="Translate3D -> Translate3Dd" projectd.go
	gofmt -w=true -r="HomogRotate2D -> HomogRotate2Dd" projectd.go
	gofmt -w=true -r="HomogRotate3DX -> HomogRotate3DXd" projectd.go
	gofmt -w=true -r="HomogRotate3DY -> HomogRotate3DYd" projectd.go
	gofmt -w=true -r="HomogRotate3DZ -> HomogRotate3DZd" projectd.go
	gofmt -w=true -r="Scale2D -> Scale2Dd" projectd.go
	gofmt -w=true -r="Scale3D -> Scale3Dd" projectd.go
	gofmt -w=true -r="ShearX2D -> ShearX2Dd" projectd.go
	gofmt -w=true -r="ShearY2D -> ShearY2Dd" projectd.go
	gofmt -w=true -r="ShearX3D -> ShearX3Dd" projectd.go
	gofmt -w=true -r="ShearY3D -> ShearY3Dd" projectd.go
	gofmt -w=true -r="ShearZ3D -> ShearZ3Dd" projectd.go
	gofmt -w=true -r="HomogRotate3D -> HomogRotate3Dd" projectd.go
	gofmt -w=true -r="Ortho -> Orthod" projectd.go
	gofmt -w=true -r="Ortho2D -> Ortho2Dd" projectd.go
	gofmt -w=true -r="Perspective -> Perspectived" projectd.go
	gofmt -w=true -r="Frustum -> Frustumd" projectd.go
	gofmt -w=true -r="LookAt -> LookAtd" projectd.go
	gofmt -w=true -r="LookAtV -> LookAtVd" projectd.go
	gofmt -w=true -r="Projectf -> Projectd" projectd.go
	gofmt -w=true -r="UnProjectf -> UnProjectd" projectd.go

transformd:
	cp transformf.go transformd.go
	gofmt -w=true -r="float32 -> float64" transformd.go
	gofmt -w=true -r="FloatEqual32 -> FloatEqual" transformd.go
	gofmt -w=true -r="Vec2f -> Vec2d" transformd.go
	gofmt -w=true -r="Vec3f -> Vec3d" transformd.go
	gofmt -w=true -r="Vec4f -> Vec4d" transformd.go
	gofmt -w=true -r="Vecf -> Vecd" transformd.go
	gofmt -w=true -r="Mat2f -> Mat2d" transformd.go
	gofmt -w=true -r="Mat3f -> Mat3d" transformd.go
	gofmt -w=true -r="Mat4f -> Mat4d" transformd.go
	gofmt -w=true -r="Mat2x3f -> Mat2x3d" transformd.go
	gofmt -w=true -r="Mat2x4f -> Mat2x4d" transformd.go
	gofmt -w=true -r="Mat3x2f -> Mat3x2d" transformd.go
	gofmt -w=true -r="Mat3x4f -> Mat3x4d" transformd.go
	gofmt -w=true -r="Mat4x2f -> Mat4x2d" transformd.go
	gofmt -w=true -r="Mat4x3f -> Mat4x3d" transformd.go
	gofmt -w=true -r="Matf -> Matd" transformd.go
	gofmt -w=true -r="Ident2f -> Ident2d" transformd.go
	gofmt -w=true -r="Ident3f -> Ident3d" transformd.go
	gofmt -w=true -r="Ident4f -> Ident4d" transformd.go
	gofmt -w=true -r="Identf -> Identd" transformd.go
	gofmt -w=true -r="FloatEqual32 -> FloatEqual" transformd.go
	gofmt -w=true -r="FloatEqualThreshold32 -> FloatEqualThreshold" transformd.go
	gofmt -w=true -r="Quatf -> Quatd" transformd.go
	gofmt -w=true -r="QuatIdentf -> QuatIdentd" transformd.go
	gofmt -w=true -r="QuatRotatef -> QuatRotated" transformd.go
	gofmt -w=true -r="QuatSlerpf -> QuatSlerpd" transformd.go
	gofmt -w=true -r="QuatLerpf -> QuatLerpd" transformd.go
	gofmt -w=true -r="QuatNlerpf -> QuatNlerpd" transformd.go
	gofmt -w=true -r="Clampf -> Clampd" transformd.go
	gofmt -w=true -r="Rotate2D -> Rotate2Dd" transformd.go
	gofmt -w=true -r="Rotate3DX -> Rotate3DXd" transformd.go
	gofmt -w=true -r="Rotate3DY -> Rotate3DYd" transformd.go
	gofmt -w=true -r="Rotate3DZ -> Rotate3DZd" transformd.go
	gofmt -w=true -r="Translate2D -> Translate2Dd" transformd.go
	gofmt -w=true -r="Translate3D -> Translate3Dd" transformd.go
	gofmt -w=true -r="HomogRotate2D -> HomogRotate2Dd" transformd.go
	gofmt -w=true -r="HomogRotate3DX -> HomogRotate3DXd" transformd.go
	gofmt -w=true -r="HomogRotate3DY -> HomogRotate3DYd" transformd.go
	gofmt -w=true -r="HomogRotate3DZ -> HomogRotate3DZd" transformd.go
	gofmt -w=true -r="Scale2D -> Scale2Dd" transformd.go
	gofmt -w=true -r="Scale3D -> Scale3Dd" transformd.go
	gofmt -w=true -r="ShearX2D -> ShearX2Dd" transformd.go
	gofmt -w=true -r="ShearY2D -> ShearY2Dd" transformd.go
	gofmt -w=true -r="ShearX3D -> ShearX3Dd" transformd.go
	gofmt -w=true -r="ShearY3D -> ShearY3Dd" transformd.go
	gofmt -w=true -r="ShearZ3D -> ShearZ3Dd" transformd.go
	gofmt -w=true -r="HomogRotate3D -> HomogRotate3Dd" transformd.go
	gofmt -w=true -r="Ortho -> Orthod" transformd.go
	gofmt -w=true -r="Ortho2D -> Ortho2Dd" transformd.go
	gofmt -w=true -r="Perspective -> Perspectived" transformd.go
	gofmt -w=true -r="Frustum -> Frustumd" transformd.go
	gofmt -w=true -r="LookAt -> LookAtd" transformd.go
	gofmt -w=true -r="LookAtV -> LookAtVd" transformd.go
	gofmt -w=true -r="Projectf -> Projectd" transformd.go
	gofmt -w=true -r="UnProjectf -> UnProjectd" transformd.go