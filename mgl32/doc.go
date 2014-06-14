// Copyright 2014 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package mgl[32|64] (an abbreviation of mathgl since the packages were split between 32 and 64-bit versions)
is a pure Go math package specialized for 3D math, with inspiration from GLM. It provides statically-sized vectors and matrices with
compile-time generated calculations for most basic math operations. It also provides several basic graphics utilities such as bezier curves and surfaces,
generation of basic primitives like circles, easy creation of common matrices such as perspective or rotation, and common operations like converting
to/from screen/OpenGL coordinates or Projecting/Unprojecting from an MVP matrix. Quaternions are also supported.

The basic vectors and matrices are written with code generation, so looking directly at the source will probably be a bit confusing. I recommend looking at the Godoc
instead, as all basic functions are documented.

This package is written in Column Major Order to make it easier with OpenGL. This means for uniform blocks you can use the default ordering, and when you call
pass-in functions you can leave the "transpose" argument as false.
*/
package mgl32
