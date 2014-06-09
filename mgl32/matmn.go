package mgl32

// An arbitrary mxn matrix backed by a slice of floats.
//
// This is emphatically not recommended to be used for hardcore n-dimensional
// linear algebra. For that purpose in I recommend github.com/gonum/matrix or
// well-tested C libraries such as BLAS or LAPACK.
//
// This is meant to complement future algorithms that may require matrices larger than
// 4x4, but still relatively small (e.g. Jacobeans for inverse kinematics).
//
// It makes use of the same realloc callback as VecN, for use in memory pools if you
// want to avoid garbage collection.
type MatMN struct {
	m, n int
	dat  []float32
}

// Grows the underlying slice by the desired amount
func (mat *MatMN) grow(size int) {
	if mat == nil {
		mat = &MatMN{m: 0, n: 0, dat: make([]float32, 0, size)}
	}

	if len(mat.dat)+size > cap(mat.dat) {
		tmp := make([]float32, len(mat.dat), len(mat.dat)*2)
		copy(tmp, mat.dat)
		if reallocCallback != nil {
			reallocCallback(mat.dat)
		}

		mat.dat = tmp

		return
	}

	mat.dat = mat.dat[:len(mat.dat)+size]
}

func (mat *MatMN) destroy() {
	if reallocCallback != nil {
		reallocCallback(mat.dat)
	}
	mat.m, mat.n = 0, 0
	mat.dat = nil
}

// Reshapes the matrix to the desired dimensions.
// If the overall size of the new matrix (m*n) is bigger
// than the current size, the underlying slice will
// be grown, reallocating if the needed memory exceeds its cap.
func (mat *MatMN) Reshape(m, n int) {
	if mat == nil {
		mat = &MatMN{m: m, n: n, dat: make([]float32, m*n)}
		return
	}

	if m*n <= len(mat.dat) {
		mat.dat = mat.dat[:m*n]
		mat.m, mat.n = m, n
		return
	}

	mat.grow(m*n - len(mat.dat))
	mat.m, mat.n = m, n
}

// Takes the transpose of mat and puts it in dst.
// Currently dst may NOT be the same as mat, due
// to the difficulty of the in-place transpose problem.
// (This functionality will be added in the future)
//
// If dst is not of the correct dimensions, it will be Reshaped
func (mat *MatMN) Transpose(dst *MatMN) *MatMN {
	if mat == nil || dst == mat {
		return nil
	}

	dst.Reshape(mat.n, mat.m)

	for r := 0; r < mat.m; r++ {
		for c := 0; c < mat.n; c++ {
			dst.dat[r*dst.m+c] = mat.dat[c*mat.m+r]
		}
	}

	return dst
}

func (mat *MatMN) Add(dst *MatMN, addend *MatMN) *MatMN {
	if mat == nil || addend == nil || mat.m != addend.m || mat.n != addend.n {
		return nil
	}

	dst.Reshape(mat.m, mat.n)

	// No need to care about rows and columns
	// since it's element-wise anyway
	for i, el := range mat.dat {
		dst.dat[i] = el + addend.dat[i]
	}

	return dst
}

func (mat *MatMN) Sub(dst *MatMN, minuend *MatMN) *MatMN {
	if mat == nil || minuend == nil || mat.m != minuend.m || mat.n != minuend.n {
		return nil
	}

	dst.Reshape(mat.m, mat.n)

	// No need to care about rows and columns
	// since it's element-wise anyway
	for i, el := range mat.dat {
		dst.dat[i] = el - minuend.dat[i]
	}

	return dst
}

// Performs matrix multiplication on MxN matrix mat and NxO matrix mul, storing the result in dst.
// This returns dst, or nil if the operation is not able to be performed.
//
// If mat == dst, or mul == dst a temporary matrix will be allocated, the temporary slice will be returned
// to the user after use if the reallocation callback is registered.
//
// This uses the naive algorithm (though on smaller matrices,
// this can actually be faster; about len(mat)+len(mul) < ~100)
func (mat *MatMN) Mul(dst *MatMN, mul *MatMN) *MatMN {
	if mat == nil || mul == nil || mat.n != mul.m {
		return nil
	}

	if dst == mul {
		mul = &MatMN{m: mul.m, n: mul.n, dat: make([]float32, mul.m*mul.n)}
		copy(mul.dat, dst.dat)

		// If mul==dst==mul, we need to change
		// mat too or we have a bug
		if mul == dst {
			mat = mul
		}

		defer mul.destroy()
	} else if dst == mat {
		mat = &MatMN{m: mat.m, n: mat.n, dat: make([]float32, mat.m*mat.n)}
		copy(mat.dat, dst.dat)

		defer mat.destroy()
	}

	dst.Reshape(mat.m, mul.n)
	for r1 := 0; r1 < mat.m; r1++ {
		for c2 := 0; c2 < mul.n; c2++ {

			dst.dat[c2*mat.m+r1] = 0
			for i := 0; i < mat.n; i++ {
				dst.dat[c2*mat.m+r1] += mat.dat[i*mat.m+r1] * mat.dat[c2*mul.m+i]
			}

		}
	}

	return dst
}
