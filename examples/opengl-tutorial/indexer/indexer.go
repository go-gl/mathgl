// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package indexer

import (
	"github.com/Jragonmiris/mathgl"
)

// Note: is_near is implemented in mathgl already as FloatEqual

type PackedVertex struct {
	Position mathgl.Vec3f
	UV       mathgl.Vec2f
	Norm     mathgl.Vec3f
}

// Only implementing the fast version
func IndexVBO(vertices []mathgl.Vec3f, uvs []mathgl.Vec2f, normals []mathgl.Vec3f) (outIndices []uint16, outVertices []mathgl.Vec3f, outUVs []mathgl.Vec2f, outNorms []mathgl.Vec3f) {
	vertToOutIndex := make(map[PackedVertex]uint16, 0)

	for i := range vertices {
		packed := PackedVertex{vertices[i], uvs[i], normals[i]}

		index, ok := vertToOutIndex[packed]
		if ok {
			outIndices = append(outIndices, index)
		} else {
			outVertices = append(outVertices, vertices[i])
			outUVs = append(outUVs, uvs[i])
			outNorms = append(outNorms, normals[i])
			index = uint16(len(outVertices) - 1)
			outIndices = append(outIndices, index)
			vertToOutIndex[packed] = index
		}
	}

	return
}

func IndexVBOSlow(vertices []mathgl.Vec3f, uvs []mathgl.Vec2f, normals []mathgl.Vec3f) (outIndices []uint16, outVertices []mathgl.Vec3f, outUVs []mathgl.Vec2f, outNorms []mathgl.Vec3f) {

	for i := range vertices {

		index, ok := SimilarVertexIndexSlow(vertices[i], uvs[i], normals[i], outVertices, outUVs, outNorms)
		if ok {
			outIndices = append(outIndices, index)
		} else {
			outVertices = append(outVertices, vertices[i])
			outUVs = append(outUVs, uvs[i])
			outNorms = append(outNorms, normals[i])
			index = uint16(len(outVertices) - 1)
			outIndices = append(outIndices, index)
		}
	}

	return
}

func SimilarVertexIndexSlow(vertex mathgl.Vec3f, uv mathgl.Vec2f, normal mathgl.Vec3f, vertices []mathgl.Vec3f, uvs []mathgl.Vec2f, normals []mathgl.Vec3f) (index uint16, found bool) {
	// Lame linear search
	for i := range vertices {
		if mathgl.FloatEqualThreshold32(vertex[0], vertices[i][0], .01) && mathgl.FloatEqualThreshold32(vertex[1], vertices[i][1], .01) && mathgl.FloatEqualThreshold32(vertex[2], vertices[i][2], .01) &&
			mathgl.FloatEqualThreshold32(uv[0], uvs[i][0], .01) && mathgl.FloatEqualThreshold32(uv[1], uvs[i][1], .01) &&
			mathgl.FloatEqualThreshold32(normal[0], normals[i][0], .01) && mathgl.FloatEqualThreshold32(normal[1], normals[i][1], .01) && mathgl.FloatEqualThreshold32(normal[2], normals[i][2], .01) {
			return uint16(i), true
		}
	}
	// No other vertex could be used instead.
	// Looks like we'll have to add it to the VBO.
	return uint16(0), false
}
