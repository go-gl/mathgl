// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package input

import (
	glfw "github.com/go-gl/glfw3"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

type Camera struct {
	pos            mgl32.Vec3
	hAngle, vAngle float64
	time           float64
	window         *glfw.Window
}

const (
	speed      float64 = 3.0
	mouseSpeed         = .005
	width              = 1024.0
	height             = 768.0
	initialFOV         = 45.0
)

func NewCamera(window *glfw.Window) *Camera {
	return &Camera{pos: mgl32.Vec3{0, 0, 5}, hAngle: math.Pi, vAngle: 0.0, time: -1.0, window: window} // Make time -1 since it will never naturally be, this acts as a "first time?" flag
}

// Since go has multiple return values, I just went ahead and made it return the view and perspective matrices (in that order) rather than messing with getter methods
func (c *Camera) ComputeViewPerspective() (mgl32.Mat4, mgl32.Mat4) {
	if mgl64.FloatEqual(-1.0, c.time) {
		c.time = glfw.GetTime()
	}

	currTime := glfw.GetTime()
	deltaT := currTime - c.time

	xPos, yPos := c.window.GetCursorPosition()
	c.window.SetCursorPosition(width/2.0, height/2.0)

	c.hAngle += mouseSpeed * ((width / 2.0) - float64(xPos))
	c.vAngle += mouseSpeed * ((height / 2.0) - float64(yPos))

	dir := mgl32.Vec3{
		float32(math.Cos(c.vAngle) * math.Sin(c.hAngle)),
		float32(math.Sin(c.vAngle)),
		float32(math.Cos(c.vAngle) * math.Cos(c.hAngle))}

	right := mgl32.Vec3{
		float32(math.Sin(c.hAngle - math.Pi/2.0)),
		0.0,
		float32(math.Cos(c.hAngle - math.Pi/2.0))}

	up := right.Cross(dir)

	if c.window.GetKey(glfw.KeyUp) == glfw.Press || c.window.GetKey('W') == glfw.Press {
		c.pos = c.pos.Add(dir.Mul(float32(deltaT * speed)))
	}

	if c.window.GetKey(glfw.KeyDown) == glfw.Press || c.window.GetKey('S') == glfw.Press {
		c.pos = c.pos.Sub(dir.Mul(float32(deltaT * speed)))
	}

	if c.window.GetKey(glfw.KeyRight) == glfw.Press || c.window.GetKey('D') == glfw.Press {
		c.pos = c.pos.Add(right.Mul(float32(deltaT * speed)))
	}

	if c.window.GetKey(glfw.KeyLeft) == glfw.Press || c.window.GetKey('A') == glfw.Press {
		c.pos = c.pos.Sub(right.Mul(float32(deltaT * speed)))
	}

	// Adding to the original tutorial, Space goes up
	if c.window.GetKey(glfw.KeySpace) == glfw.Press {
		c.pos = c.pos.Add(up.Mul(float32(deltaT * speed)))
	}

	// Adding to the original tutorial, left control goes down
	if c.window.GetKey(glfw.KeyLeftControl) == glfw.Press {
		c.pos = c.pos.Sub(up.Mul(float32(deltaT * speed)))
	}

	fov := initialFOV //- 5.0*float64(glfw.MouseWheel())

	proj := mgl32.Perspective(float32(fov), 4.0/3.0, 0.1, 100.0)
	view := mgl32.LookAtV(c.pos, c.pos.Add(dir), up)

	c.time = currTime

	return view, proj
}
