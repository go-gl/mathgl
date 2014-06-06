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
	"github.com/go-gl/mathgl/examples/opengl-tutorial/input"
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

	window, err := glfw.CreateWindow(1024, 768, "Tutorial 4", nil, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	window.MakeContextCurrent()

	gl.Init()
	gl.GetError() // Ignore error
	window.SetInputMode(glfw.StickyKeys, 1)
	window.SetCursorPosition(1024/2, 768/2)
	window.SetInputMode(glfw.Cursor, glfw.CursorHidden)

	gl.ClearColor(0., 0., 0.4, 0.)

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	gl.Enable(gl.CULL_FACE)

	camera := input.NewCamera(window)

	vertexArray := gl.GenVertexArray()
	defer vertexArray.Delete()
	vertexArray.Bind()

	prog := helper.MakeProgram("TransformVertexShader.vertexshader", "TextureFragmentShader.fragmentshader")
	defer prog.Delete()

	matrixID := prog.GetUniformLocation("MVP")

	texture, err := helper.TextureFromDDS("uvtemplate.DDS")
	if err != nil {
		fmt.Printf("Couldn't make texture: %v\n", err)
		return
	}
	defer texture.Delete()
	texSampler := prog.GetUniformLocation("myTextureSampler")

	vBufferData := [...]float32{
		-1.0, -1.0, -1.0,
		-1.0, -1.0, 1.0,
		-1.0, 1.0, 1.0,
		1.0, 1.0, -1.0,
		-1.0, -1.0, -1.0,
		-1.0, 1.0, -1.0,
		1.0, -1.0, 1.0,
		-1.0, -1.0, -1.0,
		1.0, -1.0, -1.0,
		1.0, 1.0, -1.0,
		1.0, -1.0, -1.0,
		-1.0, -1.0, -1.0,
		-1.0, -1.0, -1.0,
		-1.0, 1.0, 1.0,
		-1.0, 1.0, -1.0,
		1.0, -1.0, 1.0,
		-1.0, -1.0, 1.0,
		-1.0, -1.0, -1.0,
		-1.0, 1.0, 1.0,
		-1.0, -1.0, 1.0,
		1.0, -1.0, 1.0,
		1.0, 1.0, 1.0,
		1.0, -1.0, -1.0,
		1.0, 1.0, -1.0,
		1.0, -1.0, -1.0,
		1.0, 1.0, 1.0,
		1.0, -1.0, 1.0,
		1.0, 1.0, 1.0,
		1.0, 1.0, -1.0,
		-1.0, 1.0, -1.0,
		1.0, 1.0, 1.0,
		-1.0, 1.0, -1.0,
		-1.0, 1.0, 1.0,
		1.0, 1.0, 1.0,
		-1.0, 1.0, 1.0,
		1.0, -1.0, 1.0,
	}

	uvBufferData := [...]float32{
		0.000059, 1.0 - 0.000004,
		0.000103, 1.0 - 0.336048,
		0.335973, 1.0 - 0.335903,
		1.000023, 1.0 - 0.000013,
		0.667979, 1.0 - 0.335851,
		0.999958, 1.0 - 0.336064,
		0.667979, 1.0 - 0.335851,
		0.336024, 1.0 - 0.671877,
		0.667969, 1.0 - 0.671889,
		1.000023, 1.0 - 0.000013,
		0.668104, 1.0 - 0.000013,
		0.667979, 1.0 - 0.335851,
		0.000059, 1.0 - 0.000004,
		0.335973, 1.0 - 0.335903,
		0.336098, 1.0 - 0.000071,
		0.667979, 1.0 - 0.335851,
		0.335973, 1.0 - 0.335903,
		0.336024, 1.0 - 0.671877,
		1.000004, 1.0 - 0.671847,
		0.999958, 1.0 - 0.336064,
		0.667979, 1.0 - 0.335851,
		0.668104, 1.0 - 0.000013,
		0.335973, 1.0 - 0.335903,
		0.667979, 1.0 - 0.335851,
		0.335973, 1.0 - 0.335903,
		0.668104, 1.0 - 0.000013,
		0.336098, 1.0 - 0.000071,
		0.000103, 1.0 - 0.336048,
		0.000004, 1.0 - 0.671870,
		0.336024, 1.0 - 0.671877,
		0.000103, 1.0 - 0.336048,
		0.336024, 1.0 - 0.671877,
		0.335973, 1.0 - 0.335903,
		0.667969, 1.0 - 0.671889,
		1.000004, 1.0 - 0.671847,
		0.667979, 1.0 - 0.335851,
	}

	// Invert V because we're using a compressed texture
	for i := 1; i < len(uvBufferData); i += 2 {
		uvBufferData[i] = 1.0 - uvBufferData[i]
	}

	vertexBuffer := gl.GenBuffer()
	defer vertexBuffer.Delete()
	vertexBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(vBufferData)*4, &vBufferData, gl.STATIC_DRAW)

	uvBuffer := gl.GenBuffer()
	defer uvBuffer.Delete()
	uvBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(uvBufferData)*4, &uvBufferData, gl.STATIC_DRAW)

	// Equivalent to a do... while
	for ok := true; ok; ok = (window.GetKey(glfw.KeyEscape) != glfw.Press && !window.ShouldClose()) {
		func() {
			gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

			prog.Use()
			defer gl.ProgramUnuse()

			view, proj := camera.ComputeViewPerspective()
			model := mgl32.Ident4()

			MVP := proj.Mul4(view).Mul4(model)
			//mvpArray := mvp.AsCMOArray(mathgl.FLOAT32).([16]float32)

			matrixID.UniformMatrix4fv(false, MVP)

			gl.ActiveTexture(gl.TEXTURE0)
			texture.Bind(gl.TEXTURE_2D)
			defer texture.Unbind(gl.TEXTURE_2D)
			texSampler.Uniform1i(0)

			vertexAttrib := gl.AttribLocation(0)
			vertexAttrib.EnableArray()
			defer vertexAttrib.DisableArray()
			vertexBuffer.Bind(gl.ARRAY_BUFFER)
			defer vertexBuffer.Unbind(gl.ARRAY_BUFFER)
			vertexAttrib.AttribPointer(3, gl.FLOAT, false, 0, nil)

			uvAttrib := gl.AttribLocation(1)
			uvAttrib.EnableArray()
			defer uvAttrib.DisableArray()
			uvBuffer.Bind(gl.ARRAY_BUFFER)
			defer uvBuffer.Unbind(gl.ARRAY_BUFFER)
			uvAttrib.AttribPointer(2, gl.FLOAT, false, 0, nil)

			gl.DrawArrays(gl.TRIANGLES, 0, 12*3)

			window.SwapBuffers()
			glfw.PollEvents()
		}() // Defers unbinds and disables to here, end of the loop
	}

}
