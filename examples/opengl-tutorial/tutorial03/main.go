package main

import (
	"fmt"
	"github.com/Jragonmiris/mathgl"
	"github.com/go-gl/gl"
	"github.com/go-gl/glfw"
	"github.com/Jragonmiris/mathgl/examples/opengl-tutorial/helper"
	"os"
)

func main() {
	if err := glfw.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}

	defer glfw.Terminate()

	glfw.OpenWindowHint(glfw.FsaaSamples, 4)
	glfw.OpenWindowHint(glfw.OpenGLVersionMajor, 3)
	glfw.OpenWindowHint(glfw.OpenGLVersionMinor, 3)
	glfw.OpenWindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	if err := glfw.OpenWindow(1024, 768, 0, 0, 0, 0, 32, 0, glfw.Windowed); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return
	}

	gl.Init() // Can't find gl.GLEW_OK or any variation, not sure how to check if this worked
	gl.GetError() // Ignore error

	glfw.SetWindowTitle("Tutorial 03")

	glfw.Enable(glfw.StickyKeys)
	gl.ClearColor(0., 0., 0.4, 0.)

	vertexArray := gl.GenVertexArray()
	defer vertexArray.Delete()
	vertexArray.Bind()

	prog := helper.MakeProgram("SimpleTransform.vertexshader", "SingleColor.fragmentshader")
	defer prog.Delete()

	matrixID := prog.GetUniformLocation("MVP")

	Projection := mathgl.Perspective(45.0, 4.0/3.0, 0.1, 100.0)
	//Projection := mathgl.Identity(4,mathgl.FLOAT64)
	//Projection := mathgl.Ortho2D(-5,5,-5,5)
	View := mathgl.LookAt(4.0, 3.0, 3.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0)
	//View := mathgl.Identity(4,mathgl.FLOAT64)

	Model := mathgl.Ident4f()
	//Model := mathgl.Scale3D(2.,2.,2.).Mul(mathgl.HomogRotate3DX(25.0)).Mul(mathgl.Translate3D(.5,.2,-.7))
	MVP := Projection.Mul4(View).Mul4(Model)                   // Remember, transform multiplication order is "backwards"

	vBufferData := [...]float32{
		-1., -1., 0.,
		1., -1., 0.,
		0., 1., 0.}
	//elBufferData := [...]uint8{0, 1, 2} // Not sure why this is here

	buffer := gl.GenBuffer()
	defer buffer.Delete()
	buffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(vBufferData)*4, &vBufferData, gl.STATIC_DRAW)

	// Equivalent to a do... while
	for ok := true; ok; ok = (glfw.Key(glfw.KeyEsc) != glfw.KeyPress && glfw.WindowParam(glfw.Opened) == gl.TRUE) {
		gl.Clear(gl.COLOR_BUFFER_BIT)

		prog.Use()

		matrixID.UniformMatrix4fv(false, MVP)

		attribLoc := gl.AttribLocation(0)
		attribLoc.EnableArray()
		buffer.Bind(gl.ARRAY_BUFFER)
		attribLoc.AttribPointer(3, gl.FLOAT, false, 0, nil)

		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		attribLoc.DisableArray()

		glfw.SwapBuffers()
	}

}
