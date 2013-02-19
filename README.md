MathGL
======

This is a Go matrix and vector math library intended for Open GL graphics capabilities. Hopefully to one day be good enough for inclusion with go-gl.

# Mathglfast

This (sub-)package is also in development, to be a closer-to-OpenGL library. It is unfortunately a complete pain to write. I'm working on a code generation script at the moment for it.


# TODO:

Benchmarks (important!)

Make more unit tests. Right now basic functionality is essentially tested, but a lot of corner cases potentially aren't covered.

Look into an alternate determinant method (standard method is O(n!)). Low priority, since most people probably won't be taking >4x4 matrix determinants anyway.

Documentation.

Tranformation library functions

Perhaps improve API for making Scalars/Scalar slices. It can look a bit ugly right now. Maybe simplify it so you don't have to pass in the variable type to the matrix/vector creation methods?

Matrix->Array method