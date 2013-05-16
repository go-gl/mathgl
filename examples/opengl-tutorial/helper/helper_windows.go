package helper

import (
	"github.com/go-gl/gl"
	"github.com/go-gl/glfw"
	"io/ioutil"
	"fmt"
)

func MakeProgram(vertFname, fragFname string) gl.Program {
	vertSource, err := ioutil.ReadFile(vertFname)
	if err != nil {
		panic(err)
	}

	fragSource, err := ioutil.ReadFile(fragFname)
	if err != nil {
		panic(err)
	}

	vertShader, fragShader := gl.CreateShader(gl.VERTEX_SHADER), gl.CreateShader(gl.FRAGMENT_SHADER)
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

func MakeTextureFromTGA(fname string) gl.Texture {
	tex := gl.GenTexture()

	tex.Bind(gl.TEXTURE_2D)
	glfw.LoadTexture2D(fname, 0)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR_MIPMAP_LINEAR)
	gl.GenerateMipmap(gl.TEXTURE_2D)

	//	glh.OpenGLSentinel() // check for errors

	return tex
}
