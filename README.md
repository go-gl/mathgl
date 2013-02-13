MathGL
======

This is a Go matrix and vector math library specialized for Open GL graphics capabilities. Hopefully to one day be good enough for inclusion with go-gl.

# Questions

## Matters of syntax

Which is more pleasing: a := v.Dot(u) or a := Dot(u,v)? If we use the latter, should something like v.Dot(u) take a pointer receiver and modify v in place, or not exist at all?

Obviously, for operations that return vectors you can chain things like v.Add(u).Cross(c).Sub(n).Dot(j), which may or may not be more clear than Dot(Sub(Cross(Add(v,u),c),n),j)