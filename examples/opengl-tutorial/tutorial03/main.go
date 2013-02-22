package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"github.com/go-gl/glfw"
	"github.com/go-gl/gl"
	"github.com/Jragonmiris/mathgl"
)

func main() {
	if err := glfw.Init(); err != nil {
		fmt.Fprintf(os.Stderr,"%s\n",err.Error())
		return
	}
	
	defer glfw.Terminate()
	
	glfw.OpenWindowHint(glfw.FsaaSamples, 4)
	glfw.OpenWindowHint(glfw.OpenGLVersionMajor, 3)
	glfw.OpenWindowHint(glfw.OpenGLVersionMinor, 3)
	glfw.OpenWindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	
	if err := glfw.OpenWindow(1024,768, 0,0,0,0, 32,0, glfw.Windowed); err != nil {
		fmt.Fprintf(os.Stderr,"%s\n",err.Error())
		return
	}
	
	// glewExperimental=true ?
	gl.Init() // Can't find gl.GLEW_OK or any variation, not sure how to check if this worked
	
	
	glfw.SetWindowTitle("Tutorial 03")
	
	glfw.Enable(glfw.StickyKeys)
	gl.ClearColor(0.,0.,0.4,0.)
	
	vertexArray := gl.GenVertexArray()
	defer vertexArray.Delete()
	vertexArray.Bind()
	
	prog := MakeProgram("SimpleTransform.vertexshader", "SingleColor.fragmentshader")
	defer prog.Delete()
	
	matrixID := prog.GetUniformLocation("MVP")
	
	Projection := mathgl.Perspective(45.0, 4.0/3.0, 0.1, 100.0)
	View := mathgl.LookAt(4.0,3.0,3.0, 0.0,0.0,0.0, 0.0,1.0,0.0)
	Model := mathgl.Identity(4,mathgl.FLOAT64)
	MVP := Projection.Mul(View).Mul(Model)
	mvpArray := MVP.AsArray(mathgl.FLOAT32).([16]float32)
	
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
		
		matrixID.UniformMatrix4fv(false, mvpArray)
		
		
		attribLoc := gl.AttribLocation(0)
		attribLoc.EnableArray()
		buffer.Bind(gl.ARRAY_BUFFER)
		attribLoc.AttribPointer(3, gl.FLOAT, false, 0, nil)
		
		gl.DrawArrays(gl.TRIANGLES, 0, 3)
		
		attribLoc.DisableArray()
		
		glfw.SwapBuffers()
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
	
	
	vertShader,fragShader := gl.CreateShader(gl.VERTEX_SHADER), gl.CreateShader(gl.FRAGMENT_SHADER)
	vertShader.Source(string(vertSource))
	fragShader.Source(string(fragSource))
	
	vertShader.Compile()
	fragShader.Compile()
	
	prog := gl.CreateProgram()
	prog.AttachShader(vertShader)
	prog.AttachShader(fragShader)
	prog.Link()
	prog.Validate()
	fmt.Println(prog.GetInfoLog())
	
	return prog
}