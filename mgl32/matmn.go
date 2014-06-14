package mgl32

// An arbitrary mxn matrix backed by a slice of floats.
//
// This is emphatically not recommended for hardcore n-dimensional
// linear algebra. For that purpose I recommend github.com/gonum/matrix or
// well-tested C libraries such as BLAS or LAPACK.
//
// This is meant to complement future algorithms that may require matrices larger than
// 4x4, but still relatively small (e.g. Jacobeans for inverse kinematics).
//
// It makes use of the same realloc callback as VecN, for use in memory pools if you
// want to avoid garbage collection.
//
// MatMN will always check if the receiver is nil on any method. Meaning MathMN(nil).Add(dst,m2)
// should always work. Except for the Reshape function, the semantics of this is to "propogate" nils
// forward, so if an invalid operation occurs in a long chain of matrix operations, the overall result will be nil.
type MatMxN struct {
	m, n int
	dat  []float32
}

// Creates a matrix backed by a new slice of size m*n
func NewMatrix(m, n int) (mat *MatMxN) {
	return &MatMxN{m: m, n: n, dat: make([]float32, m*n)}
}

// Returns a matrix backed by the slice dat,
// with dimensions m and n.
//
// For instance, to create a 3x3 MatMN from a Mat3
//
//    m1 := mgl32.Rotate3DX(3.14159)
//    mat := mgl32.NewBackedMatrix(m1[:],3,3)
//
// will create an MN matrix backed by the initial
// mat3 that still acts as a 3D rotation matrix.
//
// If m*n > cap(dat), this function will panic.
func NewBackedMatrix(dat []float32, m, n int) *MatMxN {
	mat := &MatMxN{m: m, n: n, dat: dat[:m*n]}
	return mat
}

// Copies src into dst. This Reshapes dst
// to the same size as src.
func CopyMatMN(dst, src *MatMxN) {
	dst.Reshape(src.m, src.n)
	copy(dst.dat, src.dat)
}

// Grows the underlying slice by the desired amount
func (mat *MatMxN) grow(size int) *MatMxN {
	if mat == nil {
		return &MatMxN{m: 0, n: 0, dat: make([]float32, size, size)}
	}

	// This matches Go's reallocation semantics when append is used.
	if len(mat.dat)+size > cap(mat.dat) {
		newCap := len(mat.dat) * 2
		if len(mat.dat)+size > 2*len(mat.dat) {
			newCap = len(mat.dat) + size
		}

		tmp := make([]float32, size, newCap)
		copy(tmp, mat.dat)
		if reallocCallback != nil {
			reallocCallback(mat.dat)
		}

		mat.dat = tmp

		return mat
	}

	mat.dat = mat.dat[:len(mat.dat)+size]

	return mat
}

// Returns the underlying matrix slice via the callback
// if it exists
func (mat *MatMxN) destroy() {
	if mat == nil {
		return
	}

	if reallocCallback != nil && mat.dat != nil {
		reallocCallback(mat.dat)
	}
	mat.m, mat.n = 0, 0
	mat.dat = nil
}

// Reshapes the matrix to the desired dimensions.
// If the overall size of the new matrix (m*n) is bigger
// than the current size, the underlying slice will
// be grown, reallocating if the needed memory exceeds its cap.
//
// If the caller is a nil pointer, the return value will be a new
// matrix, as if NewMatrix(m,n) had been called. Otherwise it's
// simply the caller.
func (mat *MatMxN) Reshape(m, n int) *MatMxN {
	if mat == nil {
		return NewMatrix(m, n)
	}

	if m*n <= len(mat.dat) {
		if mat.dat != nil {
			mat.dat = mat.dat[:m*n]
		} else {
			mat.dat = []float32{}
		}
		mat.m, mat.n = m, n
		return mat
	}

	mat.grow(m*n - len(mat.dat))
	mat.m, mat.n = m, n

	return mat
}

// Takes the transpose of mat and puts it in dst.
//
// If dst is not of the correct dimensions, it will be Reshaped,
// if dst and mat are the same, a temporary matrix of the correct size will
// be allocated; these resources will be released via the ReallocCallback if
// it is registered. This should be improved in the future.
func (mat *MatMxN) Transpose(dst *MatMxN) (t *MatMxN) {
	if mat == nil {
		return nil
	}

	if dst == mat {
		dst = NewMatrix(mat.n, mat.m)

		// Copy data to correct matrix,
		// delete temporary buffer,
		// and set the return value to the
		// correct one
		defer func() {
			copy(mat.dat, dst.dat)

			mat.m, mat.n = mat.n, mat.m

			dst.destroy()
			t = mat
		}()

		return mat
	} else {
		dst = dst.Reshape(mat.n, mat.m)
	}

	for r := 0; r < mat.m; r++ {
		for c := 0; c < mat.n; c++ {
			dst.dat[r*dst.m+c] = mat.dat[c*mat.m+r]
		}
	}

	return dst
}

// Returns the element at the given row and column.
// This is garbage in/garbage out and does no bounds
// checking. If the computation happens to lead to an invalid
// element, it will be returned; or it may panic.
func (mat *MatMxN) At(row, col int) float32 {
	return mat.dat[col*mat.m+row]
}

// Sets the element at the given row and column.
// This is garbage in/garbage out and does no bounds
// checking. If the computation happens to lead to an invalid
// element, it will be set; or it may panic.
func (mat *MatMxN) Set(row, col int, val float32) {
	mat.dat[col*mat.m+row] = val
}

func (mat *MatMxN) Add(dst *MatMxN, addend *MatMxN) *MatMxN {
	if mat == nil || addend == nil || mat.m != addend.m || mat.n != addend.n {
		return nil
	}

	dst = dst.Reshape(mat.m, mat.n)

	// No need to care about rows and columns
	// since it's element-wise anyway
	for i, el := range mat.dat {
		dst.dat[i] = el + addend.dat[i]
	}

	return dst
}

func (mat *MatMxN) Sub(dst *MatMxN, minuend *MatMxN) *MatMxN {
	if mat == nil || minuend == nil || mat.m != minuend.m || mat.n != minuend.n {
		return nil
	}

	dst = dst.Reshape(mat.m, mat.n)

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
func (mat *MatMxN) Mul(dst *MatMxN, mul *MatMxN) *MatMxN {
	if mat == nil || mul == nil || mat.n != mul.m {
		return nil
	}

	if dst == mul {
		mul = &MatMxN{m: mul.m, n: mul.n, dat: make([]float32, mul.m*mul.n)}
		copy(mul.dat, dst.dat)

		// If mul==dst==mul, we need to change
		// mat too or we have a bug
		if mul == dst {
			mat = mul
		}

		defer mul.destroy()
	} else if dst == mat {
		mat = &MatMxN{m: mat.m, n: mat.n, dat: make([]float32, mat.m*mat.n)}
		copy(mat.dat, dst.dat)

		defer mat.destroy()
	}

	dst = dst.Reshape(mat.m, mul.n)
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

func (mat *MatMxN) ApproxEqual(m2 *MatMxN) bool {
	if mat == m2 {
		return true
	}
	if mat.m != m2.m || mat.n != m2.n {
		return false
	}

	for i, el := range mat.dat {
		if !FloatEqual(el, m2.dat[i]) {
			return false
		}
	}

	return true
}

func (mat *MatMxN) ApproxEqualThreshold(m2 *MatMxN, epsilon float32) bool {
	if mat == m2 {
		return true
	}
	if mat.m != m2.m || mat.n != m2.n {
		return false
	}

	for i, el := range mat.dat {
		if !FloatEqualThreshold(el, m2.dat[i], epsilon) {
			return false
		}
	}

	return true
}

func (mat *MatMxN) ApproxEqualFunc(m2 *MatMxN, comp func(float32, float32) bool) bool {
	if mat == m2 {
		return true
	}
	if mat.m != m2.m || mat.n != m2.n {
		return false
	}

	for i, el := range mat.dat {
		if !comp(el, m2.dat[i]) {
			return false
		}
	}

	return true
}
