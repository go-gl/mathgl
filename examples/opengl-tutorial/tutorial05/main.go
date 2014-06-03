// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/Jragonmiris/mathgl"
	"github.com/Jragonmiris/mathgl/examples/opengl-tutorial/helper"
	"github.com/go-gl/gl"
	"github.com/go-gl/glfw"
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

	//gl.GlewExperimental(true)
	gl.Init()     // Can't find gl.GLEW_OK or any variation, not sure how to check if this worked
	gl.GetError() // ignore error, since we're telling it to use CoreProfile above, we get "invalid enumerant" (GLError 1280) which freaks the OpenGLSentinel out
	// With go-gl we also apparently can't set glewExperimental

	glfw.SetWindowTitle("Tutorial 05")

	glfw.Enable(glfw.StickyKeys)
	gl.ClearColor(0., 0., 0.4, 0.)

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	vertexArray := gl.GenVertexArray()
	defer vertexArray.Delete()
	vertexArray.Bind()

	prog := helper.MakeProgram("TransformVertexShader.vertexshader", "TextureFragmentShader.fragmentshader")
	defer prog.Delete()

	matrixID := prog.GetUniformLocation("MVP")

	Projection := mathgl.Perspective(45.0, 4.0/3.0, 0.1, 100.0)

	View := mathgl.LookAt(4.0, 3.0, 3.0, 0.0, 0.0, 0.0, 0.0, 1.0, 0.0)

	Model := mathgl.Ident4f()

	MVP := Projection.Mul4(View).Mul4(Model) // Remember, transform multiplication order is "backwards"

	texture := helper.MakeTextureFromTGA("uvtemplate.tga")
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
		1.0, -1.0, 1.0}

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
		0.667979, 1.0 - 0.335851}

	vertexBuffer := gl.GenBuffer()
	defer vertexBuffer.Delete()
	vertexBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(vBufferData)*4, &vBufferData, gl.STATIC_DRAW)

	uvBuffer := gl.GenBuffer()
	defer uvBuffer.Delete()
	uvBuffer.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(uvBufferData)*4, &uvBufferData, gl.STATIC_DRAW)

	// Equivalent to a do... while
	for ok := true; ok; ok = (glfw.Key(glfw.KeyEsc) != glfw.KeyPress && glfw.WindowParam(glfw.Opened) == gl.TRUE) {
		func() {
			gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

			prog.Use()
			defer gl.ProgramUnuse()

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

			glfw.SwapBuffers()
		}() // Defers unbinds and disables to here, end of the loop
	}

}
