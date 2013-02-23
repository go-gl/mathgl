package main

import (
	"fmt"
	"os"
	"github.com/go-gl/glfw"
	"github.com/go-gl/gl"
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
	
	gl.Init()
	
	glfw.SetWindowTitle("Tutorial 01")
	glfw.Enable(glfw.StickyKeys)
	
	gl.ClearColor(0.,0.,0.4,0.)
	// Equivalent to a do... while
	for ok := true; ok; ok = (glfw.Key(glfw.KeyEsc) != glfw.KeyPress && glfw.WindowParam(glfw.Opened) == gl.TRUE) {
		glfw.SwapBuffers()
	}
	
}
