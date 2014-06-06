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

Feel free to submit pull requests for features and bug fixes. Do note that, aside from documentation bugs, meta (travis.yml etc) fixes, and example code, pull requests will not be accepted without tests corresponding to the new code. If it's a bug fix, the test should test the bug.
