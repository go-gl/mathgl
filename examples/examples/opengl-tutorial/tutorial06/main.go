package main

import (
	"fmt"
	"github.com/Jragonmiris/mathgl"
	"github.com/Jragonmiris/mathgl/examples/opengl-tutorial/input"
	"github.com/go-gl/gl"
	"github.com/go-gl/glfw"
	"github.com/go-gl/glh"
	"io/ioutil"
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

	glfw.SetSwapInterval(0)

	gl.GlewExperimental(true)
	gl.Init()     // Can't find gl.GLEW_OK or any variation, not sure how to check if this worked
	gl.GetError() // ignore error, since we're telling it to use CoreProfile above, we get "invalid enumerant" (GLError 1280) which freaks the OpenGLSentinel out
	// With go-gl we also apparently can't set glewExpe6rimental

	glfw.SetWindowTitle("Tutorial 05")

	glfw.Enable(glfw.StickyKeys)
	glfw.Disable(glfw.MouseCursor) // Not in the original tutorial, but IMO it SHOULD be there
	glfw.SetMousePos(1024.0/2.0, 768.0/2.0)

	gl.ClearColor(0., 0., 0.4, 0.)

	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	gl.Enable(gl.CULL_FACE)

	camera := input.NewCamera()

	vertexArray := gl.GenVertexArray()
	defer vertexArray.Delete()
	vertexArray.Bind()

	prog := MakeProgram("TransformVertexShader.vertexshader", "TextureFragmentShader.fragmentshader")
	defer prog.Delete()

	matrixID := prog.GetUniformLocation("MVP")

	texture := MakeTextureFromTGA("uvtemplate.tga")
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
	for ok := true; ok; ok = (glfw.Key(glfw.KeyEsc) != glfw.KeyPress && glfw.WindowParam(glfw.Opened) == gl.TRUE && glfw.Key('Q') != glfw.KeyPress) {
		func() {
			gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

			prog.Use()
			defer gl.ProgramUnuse()

			view, proj := camera.ComputeViewPerspective()
			model := mathgl.Identity(4, mathgl.FLOAT64)

			mvp := proj.Mul(view).Mul(model)
			mvpArray := mvp.AsCMOArray(mathgl.FLOAT32).([16]float32)

			matrixID.UniformMatrix4fv(false, mvpArray)

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

func MakeProgram(vertFname, fragFname string) gl.Program {
	vertSource, err := ioutil.ReadFile(vertFname)
	if err != nil {
		panic(err)
	}

	fragSource, err := ioutil.ReadFile(fragFname)
	if err != nil {
		panic(err)
	}
	return glh.NewProgram(glh.Shader{gl.VERTEX_SHADER, string(vertSource)}, glh.Shader{gl.FRAGMENT_SHADER, string(fragSource)})
}

func MakeTextureFromTGA(fname string) gl.Texture {
	tex := gl.GenTexture()

	tex.Bind(gl.TEXTURE_2D)
	glfw.LoadTexture2D(fname, 0)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
	gl.GenerateMipmap(gl.TEXTURE_2D)

	glh.OpenGLSentinel() // check for errors

	return tex
}
