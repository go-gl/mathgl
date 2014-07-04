package mgl32

type AABB struct {
	Min, Max Vec3
}

func (aabb *AABB) IsPointInside(point Vec3) bool {
	for i := 0; i < 3; i++ {
		if aabb.Min[i] > point[i] || aabb.Max[i] < point[i] {
			return false
		}
	}

	return true
}

func (aabb1 *AABB) Collides(aabb2 AABB) bool {
	for i := 0; i < 3; i++ {
		if aabb1.Min[i] > aabb2.Max[i] || aabb1.Max[i] < aabb2.Min[i] {
			return false
		}
	}

	return true
}

func (aabb *AABB) CollidesTriangle(p1, p2, p3 Vec3) bool {
	for i := 0; i < 3; i++ {
		if Min(Min(p1[i], p2[i]), p3[i]) > aabb.Max[i] || Max(Max(p1[i], p2[i]), p3[i]) < aabb.Min[i] {
			return false
		}
	}

	return true
}

// Computes the inverse direction of a ray. This is equivalent
// to the vector where each element is the reciprocal of the elements
// of the input vector.
func InverseDirection(dir Vec3) Vec3 {
	return Vec3{1 / dir[0], 1 / dir[1], 1 / dir[2]}
}

// Computes whether a ray with the given origin and direction intersects an AABB. This also returns the
// parametric description of the ray at the point it intersects the box.
//
// This is a simple passthrough to the Precomputed version -- it computes the inverse direction, the sign
// array, and unpacks the AABB into the bounds array.
func (aabb *AABB) IntersectsRay(rayOrigin, rayDirection Vec3) (intersects bool, tmin, tmax float32) {
	invDir := InverseDirection(rayDirection)
	return rayIntersectsAABBInverseDir(rayOrigin, invDir, aabb)
}

// Here in case we want to export it. I thought 3 types of ray/direction tests were too much.
func rayIntersectsAABBInverseDir(rayOrigin, inverseRayDirection Vec3, aabb *AABB) (intersects bool, tmin, tmax float32) {
	bounds := [2]Vec3{aabb.Min, aabb.Max}
	sign := [3]int{}
	for i := range sign {
		if inverseRayDirection[i] < 0 {
			sign[i] = 1
		} else {
			sign[i] = 0
		}
	}

	return RayIntersectsAABBFromPrecomputedInfo(rayOrigin, inverseRayDirection, sign, bounds)
}

// This is equivalent to aabb.IntersectsRay, but has a precomputed sign vector and uses a bounds array.
//
// The bounds array is equivalent to [2]{aabb.Min, aabb.Max} and the sign is equivalent to whether each component inverseRayDirection is < 0
// that is, if inverseRayDirection[i] < 0, then sign[i] = 1, otherwise sign[i] = 0.
//
// This returns whether there was an intersection. If so, the values of tmin and tmax will be parametric description of the ray where it intersects the box.
func RayIntersectsAABBFromPrecomputedInfo(rayOrigin, inverseRayDirection Vec3, sign [3]int, bounds [2]Vec3) (intersects bool, tmin, tmax float32) {
	// This is effectively equivalent to
	/* if inverseRayDirection.X() < 0 {
		tmin = (aabb.Min.X() - rayOrigin.X()) * inverseRayDirection.X()
	} else {
		tmin = (aabb.Max.X - rayOrigin.X()) * inverseRayDirection.X()
	}
	*/
	// Effectively, these checks control for the direction of the ray, swapping max and min
	// depending on the direction
	tmin = (bounds[sign[0]].X() - rayOrigin.X()) * inverseRayDirection.X()
	tmax = (bounds[1-sign[0]].X() - rayOrigin.X()) * inverseRayDirection.X()

	tymin := (bounds[sign[1]].Y() - rayOrigin.X()) * inverseRayDirection.X()
	tymax := (bounds[1-sign[1]].Y() - rayOrigin.X()) * inverseRayDirection.X()

	if tmin > tymax || tymin > tmax {
		return false, tmin, tmax
	}

	if tymin > tmin {
		tmin = tymin
	}
	if tymax < tmax {
		tmax = tymax
	}

	tzmin := (bounds[sign[2]].Z() - rayOrigin.Z()) * inverseRayDirection.Z()
	tzmax := (bounds[1-sign[2]].Z() - rayOrigin.Z()) * inverseRayDirection.Z()
	if tmin > tzmax || tzmin > tmax {
		return false, tmin, tmax
	}
	if tzmin > tmin {
		tmin = tzmin
	}
	if tzmax < tmax {
		tmax = tzmax
	}

	return true, tmin, tmax
}

// Transforms the AABB with the given transformation matrix.
func (aabb *AABB) Transform(transformation Mat4) *AABB {
	center, extent := aabb.Min.Add(aabb.Max).Mul(.5), aabb.Max.Sub(aabb.Min).Mul(.5)

	newCenter := TransformCoordinate(center, transformation)
	newExtent := TransformNormal(extent, transformation.Abs())

	return &AABB{Min: newCenter.Sub(newExtent), Max: newCenter.Add(newExtent)}
}
