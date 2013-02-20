MathGL
======

This is a Go matrix and vector math library intended for Open GL graphics capabilities. Hopefully to one day be good enough for inclusion with go-gl. Feedback is very welcome, especially API improvements, and pull requests are welcome. Benchmarks, more test cases, and example code (probably something like following an OpenGL tutorial) are definitely needed.

# Documentation

The package (the base package, not mathglfast) is more or less fully documented at the moment. However, suggestions for improvement or correction of obvious errors are obviously welcome.

The package-level documentation is in vector.go

# Mathglfast

This (sub-)package is also in development, to be a closer-to-OpenGL library. It is unfortunately a complete pain to write. I'm working on a code generation script at the moment for it.


# TODO (Contribution welcome)

Benchmarks (important!)

Make more unit tests. Right now basic functionality is essentially tested, but a lot of corner cases potentially aren't covered. Additionally, no new content (Quaternions and anything in transform.go) is untested

Look into an alternate determinant method (standard method is O(n!)). Low priority, since most people probably won't be taking >4x4 matrix determinants anyway.

Perhaps improve API for making Scalars/Scalar slices. It can look a bit ugly right now.

Perhaps Euler Angles should be supported -- maybe at least minimally add functions to construct Quaternions and Rotation Matrices from some similar representation