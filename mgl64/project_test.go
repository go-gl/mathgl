// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mgl64

import (
	"testing"
)

func TestProject(t *testing.T) {
	t.Parallel()

	obj := Vec3{1002, 960, 0}
	modelview := Mat4{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 203, 1, 0, 1}
	projection := Mat4{0.0013020833721384406, 0, 0, 0, -0, -0.0020833334419876337, -0, -0, -0, -0, -1, -0, -1, 1, 0, 1}
	initialX, initialY, width, height := 0, 0, 1536, 960
	win := Project(obj, modelview, projection, initialX, initialY, width, height)
	answer := Vec3{1205.0000359117985, -1.0000501200556755, 0.5} // From glu.Project()

	if !win.ApproxEqualThreshold(answer, 1e-4) {
		t.Errorf("Project does something weird, differs from expected by of %v", win.Sub(answer).Len())
	}
}
