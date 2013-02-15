MathGL
======

This is a Go matrix and vector math library specialized for Open GL graphics capabilities. Hopefully to one day be good enough for inclusion with go-gl.

# Questions

## Matters of syntax

Which is more pleasing: a := v.Dot(u) or a := Dot(u,v)? If we use the latter, should something like v.Dot(u) take a pointer receiver and modify v in place, or not exist at all?

Obviously, for operations that return vectors you can chain things like v.Add(u).Cross(c).Sub(n).Dot(j), which may or may not be more clear than Dot(Sub(Cross(Add(v,u),c),n),j)

## What's "ugly"?

Ugly was my first attempt at this. It explicitly defines things like Vec3 and Vec2. It requires code-generation scripts and makefiles. It's probably faster, but a lot uglier.

It will either go away or replace the current code once a decision is made on what's best. Or maybe I'll just leave it there if people want to import it


# TODO:

Make unit tests, verify row major order stuff is correct on the matrices (especially multiplication). RMO makes my brain overflow for some reason.

Get determinant function working

Documentation

## Performance Ideas:

Test concurrency. Theoretically things like vector addition and matrix multiplication can be done concurrently. In practice I'm fairly sure the overhead for creating channels and spawning goroutines will outweight any benefit. However, some operations may be able to be optimized with concurrency in mind. For instance, matrix multiplication is associative, so if you have 4 matrices that all need to be multiplied at once, it may be possible and worthwhile to break them into a pair of multiplications with four each, and multiply the result.

Make internal (unexported) methods that forego type checking. For instance, when multiplying we can be sure that the new matrix we're constructing has elements all of the correct type, so we don't need to go through and check.

# New additions

Hey, I just found out about that variable.(Type) returns "ok" as its second value, so we can check types without reflect!