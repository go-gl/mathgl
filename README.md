MathGL
======

This is a Go matrix and vector math library specialized for Open GL graphics capabilities. Hopefully to one day be good enough for inclusion with go-gl.

# Questions

## Matters of syntax

Which is more pleasing: a := v.Dot(u) or a := Dot(u,v)? If we use the latter, should something like v.Dot(u) take a pointer receiver and modify v in place, or not exist at all?

Obviously, for operations that return vectors you can chain things like v.Add(u).Cross(c).Sub(n).Dot(j), which may or may not be more clear than Dot(Sub(Cross(Add(v,u),c),n),j)

## What's "ugly"

Ugly was my first attempt at this. It explicitly defines things like Vec3 and Vec2. It requires code-generation scripts and makefiles. It's probably faster, but a lot uglier.

It will either go away or replace the current code once a decision is made on what's best. Or maybe I'll just leave it there if people want to import it

## Reflection? Really? In a MATH package?

This attempt is basically the polar opposite of the last attempt, forget speed in an attempt to generalize code. This is still very early on in the project's lifecycle, I figure eventually we'll decide on a healthy mix of the two.

# TODO:

Make unit tests, verify row major order stuff is correct on the matrices (especially multiplication). RMO makes my brain overflow for some reason.

Get determinant function working

Allow multiplication between matrices and vectors without using AsVector() or AsMatrix()

Documentation