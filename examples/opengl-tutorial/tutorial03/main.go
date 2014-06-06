// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"github.com/go-gl/mathgl/examples/opengl-tutorial/helper"
	"github.com/go-gl/mathgl/mgl32"
)

func main() {
	runtime.LockOSThread()

	if !glfw.Init() {
		fmt.Fprintf(os.Stderr, "Can't open GLFW")
		return
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Samples, 4)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenglProfile, glfw.OpenglCoreProfile)
	glfw.WindowHint(glfw.OpenglForwardCompatible, glfw.True) // needed for macs

	window, err := glfw.CreateWindow(1024, 768, "Tutorial 1", nil, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	window.MakeContextCurrent()

	gl.Init()
	gl.GetError() // Ignore error
	window.SetInputMode(glfw.StickyKeys, 1)

	gl.ClearColor(0., 0., 0.4, 0.)

	vertexArray := gl.GenVertexArray()
	defer vertexArray.Delete()
	vertexArray.Bind()

	prog := helper.MakeProgram("SimpleTransform.vertexshader", "SingleColor.fragmentshader")
	defer prog.Delete()

	matrixID := prog.GetUniformLocation("MVP")

	Projection := mgl32.Perspective(45.0, 4.0/3.0, 0.1, 100.0)
	//Projection := mathgl.Identity(4,mathgl.FLOAT64)
	//Projection := mathgl.Ortho2D(-5,5,-5,5)
	View := mgl32.LookAt(4.0, 3.0, 3.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0)
	//View := mathgl.Identity(4,mathgl.FLOAT64)

	Model := mgl32.Ident4()
	//Model := mathgl.Scale3D(2.,2.,2.).Mul(mathgl.HomogRotate3DX(25.0)).Mul(mathgl.Translate3D(.5,.2,-.7))
	MVP := Projection.Mul4(View).Mul4(Model) // Remember, transform multiplication order is "backwards"

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
	for ok := true; ok; ok = (window.GetKey(glfw.KeyEscape) != glfw.Press && !window.ShouldClose()) {
		gl.Clear(gl.COLOR_BUFFER_BIT)

		prog.Use()

		matrixID.UniformMatrix4fv(false, MVP)

		attribLoc := gl.AttribLocation(0)
		attribLoc.EnableArray()
		buffer.Bind(gl.ARRAY_BUFFER)
		attribLoc.AttribPointer(3, gl.FLOAT, false, 0, nil)

		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		attribLoc.DisableArray()

		window.SwapBuffers()
		glfw.PollEvents()
	}

}
