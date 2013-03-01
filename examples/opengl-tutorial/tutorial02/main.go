package main

import (
	"fmt"
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

	glfw.SetWindowTitle("Tutorial 02")

	glfw.Enable(glfw.StickyKeys)
	gl.ClearColor(0., 0., 0.4, 0.)

	prog := helper.MakeProgram("SimpleVertexShader.vertexshader", "SimpleFragmentShader.fragmentshader")

	vBufferData := [...]float32{
		-1., -1., 0.,
		1., -1., 0.,
		0., 1., 0.}

	vertexArray := gl.GenVertexArray()
	vertexArray.Bind()
	buffer := gl.GenBuffer()
	buffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(vBufferData)*4, &vBufferData, gl.STATIC_DRAW)

	// Equivalent to a do... while
	for ok := true; ok; ok = (glfw.Key(glfw.KeyEsc) != glfw.KeyPress && glfw.WindowParam(glfw.Opened) == gl.TRUE) {
		gl.Clear(gl.COLOR_BUFFER_BIT)

		prog.Use()

		attribLoc := gl.AttribLocation(0)
		attribLoc.EnableArray()
		buffer.Bind(gl.ARRAY_BUFFER)
		attribLoc.AttribPointer(3, gl.FLOAT, false, 0, nil)

		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		attribLoc.DisableArray()

		glfw.SwapBuffers()
	}

}
