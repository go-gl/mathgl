package mgl32

type AABB struct {
	Min, Max Vec3
}

func PointInAABB(aabb AABB, point Vec3) bool {
	for i := 0; i < 3; i++ {
		if aabb.Min[i] > point[i] || aabb.Max[i] < point[i] {
			return false
		}
	}

	return true
}

func AABBsCollide(aabb1 AABB, aabb2 AABB) bool {
	for i := 0; i < 3; i++ {
		if aabb1.Min[i] > aabb2.Max[i] || aabb1.Max[i] < aabb2.Min[i] {
			return false
		}
	}

	return true
}

func TriangleCollidesAABB(aabb AABB, p1, p2, p3 Vec3) {
	for i := 0; i < 3; i++ {
		if Min(Min(p1[i], p2[i]), p3[i]) > aabb.Max[i] || Max(Max(p1[i], p2[i]), p3[i]) < aabb.Min[i] {
			return false
		}
	}

	return true
}

func InverseRayIntersectsAABB(rayOrigin, inverseRayDirection Vec3, raySign [3]uint, bounds [2]Vec3, lambdaMin, lambdaMax float32) (bool, float32) {
	tmin := (bounds[raySign[0]].X() - rayOrigin.X()) * inverseRayDirection.X()
}

func TransformAABB(aabb AABB, transformation Mat3) AABB {
	center, extent := aabb.Min.Add(aabb.Max).Mul(.5), aabb.Max.Sub(aabb.Min).Mul(.5)

}
