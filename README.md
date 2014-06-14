MathGL [![Build Status](https://travis-ci.org/go-gl/mathgl.png?branch=master)](https://travis-ci.org/go-gl/mathgl)
======

This is a Go matrix and vector math library specialized for Open GL graphics capabilities.

This package is made primarily with code generation for the basic vector and matrix operations, though and functionality beyond that is handwritten.

Vectors and matrices are stored in Column Major Order, just like OpenGL, which means the "transpose" argument should be **false** when passing in vectors and matrices using this package.

This package is split into two sub-packages. The package `mgl32` deals with 32-bit floats, and `mgl64` deals with 64-bit ones. Generally you'll use the 32-bit ones with OpenGL, but the 64-bit one is available in case you use the double extension or simply want to do higher precision 3D math without OpenGL.

The old repository, before the split between the 32-bit and 64-bit subpackages, is kept at github.com/Jragonmiris/mathgl (the old repository path), but is no longer maintained.

The examples are now working! Go look at the examples folder for working examples of how to use the code!

Contributing
============

Feel free to submit pull requests for features and bug fixes. Do note that, aside from documentation bugs, meta (travis.yml etc) fixes, example code, and *extremely* trivial changes (basic accessors) pull requests will not be accepted without tests corresponding to the new code. If it's a bug fix, the test should test the bug.

Also note that since code generation is used in the files `matrix.go` and `vector.go`, no changes should be made to those files directly. Either changes should be made to `genprog/main.go` if you're brave enough to add to that mess, or (preferably), in a different file altogether. No such file currently exists, but something like `matrixStatic.go` would suffice.

API Changes
===========

From now on, no major API breaking changes will be made between major Go version releases. That means any time the "x" in Go1.x increases. Exceptions are made, of course, for bug fixes. If a bug fix necessitates changing semantics of calling software, it will be changed. (An example is the recent update of Transpose which was mistakenly using row major rules). Deprecated functions may also be nuked at major version released. Before any API breaking changes near major releases, the most recent non-breaking commit will be tagged with the major Go version number (e.g. Go1.2). If no such tag exists, one can assume nothing has been broken.

Currently Deprecated Functions
-------

EulerToQuat(ax,ay,az) ==USE INSTEAD==> AnglesToQuat(ax,ay,az,ZYX)