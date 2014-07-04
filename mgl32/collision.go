package mgl32

import (
	"fmt"
)

// AABB stands for Axis-Aligned Bounding Box -- a box whose planes
// are parallel with the 3 basis axes of your coordinate system.
//
// These are generally used for simple broad-phase collision tests.
// We often use them for quick and dirty culling in graphics, but this could
// be used as the foundation for the very basics of a physics engine.
type AABB struct {
	Min, Max Vec3
}

// Tests whether a point lies on or inside the AABB.
func (aabb *AABB) IsPointInside(point Vec3) bool {
	for i := 0; i < 3; i++ {
		if aabb.Min[i] > point[i] || aabb.Max[i] < point[i] {
			return false
		}
	}

	return true
}

// Tests whether or not two AABBs collide -- "collision" is defined
// here as they are touching OR one is completely engulfed by the other.
func (aabb1 *AABB) Collides(aabb2 AABB) bool {
	for i := 0; i < 3; i++ {
		if aabb1.Min[i] > aabb2.Max[i] || aabb1.Max[i] < aabb2.Min[i] {
			return false
		}
	}

	return true
}

// Tests whether or not a triangle given by the points p1, p2, and p3 collides with
// the AABB -- where "collision" is defined as they intersect at some point of one engulfs the
// other. The points p1, p2, and p3 are not required to be in any specific order.
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
// parametric description of the ray at the points it intersects the box (that is, the distance from the ray's origin),
// where tmin is the entry point and tmax the exit point.
//
// This is a simple passthrough to the Precomputed version -- it computes the inverse direction, the sign
// array, and unpacks the AABB into the bounds array.
//
// If the returned t values are negative, then the "intersection" is behind the ray. If the returned t values
// have different signs, the ray originated from inside the box.
func (aabb *AABB) IntersectsRay(rayOrigin, rayDirection Vec3) (intersects bool, tmin, tmax float32) {
	invDir := InverseDirection(rayDirection)
	return aabb.IntersectsRayInverseDir(rayOrigin, invDir)
}

// This is equivalent to aabb.IntersectsRay, except with a precomputed inverse direction. It will extract the AABB
// into the bounds array, and infer the sign vector from the inverse direction. This is much better than IntersectsRay,
// and you should probably use this if you don't want maintain references to ugly types like the sign or bounds arrays yourself.
func (aabb *AABB) IntersectsRayInverseDir(rayOrigin, inverseRayDirection Vec3) (intersects bool, tmin, tmax float32) {
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

// This is equivalent to aabb.IntersectsRay, but has a precomputed sign vector, inverse direction, and uses a bounds array for the aabb's min and max.
//
// The bounds array is equivalent to [2]float32{aabb.Min, aabb.Max} and the sign is equivalent to whether each component inverseRayDirection is < 0
// that is, if inverseRayDirection[i] < 0, then sign[i] = 1, otherwise sign[i] = 0.
//
// This returns whether there was an intersection. If so, the values of tmin and tmax will be parametric description of the ray where it intersects the box (roughly the distance at that point).
//
// You should probably use this rather than aabb.IntersectsRay, but the former is provided because it "looks" a lot cleaner. IntersectsRayInverseDir is
// a pretty happy medium, but still slower.
//
// If the returned t values are negative, then the "intersection" is behind the ray. If the returned t values
// have different signs, the ray originated from inside the box.
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

	tymin := (bounds[sign[1]].Y() - rayOrigin.Y()) * inverseRayDirection.Y()
	tymax := (bounds[1-sign[1]].Y() - rayOrigin.Y()) * inverseRayDirection.Y()

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

// Transforms the AABB with the given transformation matrix, this should correctly handle scaling.
func (aabb *AABB) Transform(transformation Mat4) *AABB {
	center, extent := aabb.Min.Add(aabb.Max).Mul(.5), aabb.Max.Sub(aabb.Min).Mul(.5)

	newCenter := TransformCoordinate(center, transformation)
	newExtent := TransformNormal(extent, transformation.Abs())

	fmt.Println(newCenter, newExtent)

	return &AABB{Min: newCenter.Sub(newExtent), Max: newCenter.Add(newExtent)}
}
