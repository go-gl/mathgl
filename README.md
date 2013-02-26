MathGL
======

This is a Go matrix and vector math library specialized for Open GL graphics capabilities. Hopefully to one day be good enough for inclusion with go-gl.

This package is the "fast" package. it's written with a code generation script (and thus, is liable to lack good documentation at least for a while) to dynamically generate fast functions for vectors up to size 4 and matrices up to 4x4. It is currently not as well developed as ./mathgl (the subpackage is also named mathgl so as to cause minimal refactoring of existing code). This one is much dirtier, much uglier, but probably much more efficient and it's much easier to see the true cost by looking directly at the code.
