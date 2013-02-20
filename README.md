MathGL
======

This is a Go matrix and vector math library intended to ease Open GL graphics computations (especially since GLRotate and so on have become deprecated with OpenGL 3+). Hopefully to one day be good enough for inclusion with go-gl. Feedback is very welcome, especially API improvements, and pull requests are welcome. Benchmarks, more test cases, and example code (probably something like following an OpenGL tutorial) are definitely needed.

# Documentation

The package (the base package, not mathglfast) is more or less fully documented at the moment. However, suggestions for improvement or correction of obvious errors are obviously welcome.

The package-level documentation is in vector.go

# Mathglfast

This (sub-)package is also in development, to be a closer-to-OpenGL library. It is unfortunately a complete pain to write. I'm working on a code generation script at the moment for it.


# TODO (Contribution welcome)

Benchmarks (important!)

Make more unit tests. Right now basic functionality is essentially tested, but a lot of corner cases potentially aren't covered. Additionally, no new content (Quaternions and anything in transform.go) is untested

Look into an alternate determinant method (standard method is O(n!)). Low priority, since most people probably won't be taking >4x4 matrix determinants anyway.

# Wishlist

I would love to make this have some more general functions, finding eigenvectors, putting matrices in echelon form. Perhaps even some utility methods to make finding a ray/plane intersection easier. This doesn't have to be Matlab for Go, but the more tested, optimized utility methods we get, the better this package will be.

Perhaps push parallel development on fastmathgl a little farther. At the moment I just don't have the heart to work on it. To really make it shine, it seems like you need to use code generation and text replacement tricks to minimize the amount of work you need. On the plus side, it's almost guaranteed to be faster (no for loops etc), requires less conversion from basic types to Scalar and such, and is much closer to GLM. (Also, suggestions to improve the name of fastmathgl are welcome, I have no idea what to call the thing)