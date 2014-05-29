package mathgl

import (
	"testing"
)

func TestProject(t *testing.T) {
	obj := Vec3d{1002, 960, 0}
	modelview := Mat4d{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 203, 1, 0, 1}
	projection := Mat4d{0.0013020833721384406, 0, 0, 0, -0, -0.0020833334419876337, -0, -0, -0, -0, -1, -0, -1, 1, 0, 1}
	initialX, initialY, width, height := 0, 0, 1536, 960
	win := Projectd(obj, modelview, projection, initialX, initialY, width, height)
	answer := Vec3d{1205.0000359117985, -1.0000501200556755, 0.5} // From glu.Project()

	if !win.ApproxEqual(answer) {
		t.Errorf("Project does something weird, differs from expected by of %v", win.Sub(answer).Len())
	}
}
