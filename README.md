MathGL
======

This is a Go matrix and vector math library specialized for Open GL graphics capabilities. Hopefully to one day be good enough for inclusion with go-gl.

# Questions

## Matters of syntax

Which is more pleasing: a := v.Dot(u) or a := Dot(u,v)? If we use the latter, should something like v.Dot(u) take a pointer receiver and modify v in place, or not exist at all?

Obviously, for operations that return vectors you can chain things like v.Add(u).Cross(c).Sub(n).Dot(j), which may or may not be more clear than Dot(Sub(Cross(Add(v,u),c),n),j)

## This is the ugliest code ever what's wrong with you?

Somebody can do a benchmark to make sure, but I'm pretty sure using for loops to do things like multiplication etc on the explicit matrices (i.e. Matrix2x3) is a little slower than just passing back a constructor. If I'm proven wrong then by all means change them all to nice for loops.

As for why there's so much terrible code duplication, blame that on Go's lack of generics and its rules for arrays (as opposed to slices). You want a Matrix2x3 and a Matrix 2x4? I hope you feel like defining their functions longhand even if they're basically the same! If one has a float32 instead of a float64? Better copy it and change the 64 to a 32!

This is why I have the makefile (which, if you can clean that up into something with fewer lines with the same functionality) that would be grand. It's basically just so I only have to write each package for float32, then the Makefile automagically generates the matching, near-identical code needed for the corresponding matrix and vector files.

I also have a simple program that generates multiplcation functions (that I may upload if desired) for matrices of various sized, because it is REALLY tedious to write them out longhand. especially since my brain seems to not agree with Row Major Order intuitively.

### I still don't like it...

Yeah, I'm not particularly fond of it either. There are other ideas that may be slower (and hopefully not by much). They still involve some tricks, though. For instance, get rid off all of the Matrix<n> types, and just have Vec/Matrix[f|d|i|l|ul|ui]. Since go-gl's Uniform <blahblah> methods take in arrays, we can have a method called AsArray, which would basically be a giant switch statement that returns the correct sized vector/array (up to 4x4) as an interface{}. It's still messy, but it's way LESS messy.