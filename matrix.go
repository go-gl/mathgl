package mathgl

import (
	"errors"
	"math"
)

// MatrixMultiplyable let's us make multiplication between vectors and matrices easier (especially for the batch multiplyer)
type MatrixMultiplyable interface {
	Mul(m MatrixMultiplyable) Matrix
}

// A Matrix is another slice wrapper like Vector, but along with its VecType, it also
// encodes the number of rows and columns it would have were it written out in a non-flattened fashion
//
// All matrices are treated as if they are in Row Major Order
type Matrix struct {
	m, n int // an m x n matrix
	typ  VecType
	dat  []Scalar
}

// NewMatrix returns an mxn matrix with an underlying slice of m*n nil-valued Scalars.
// (You probably shouldn't use this)
func NewMatrix(m, n int, typ VecType) *Matrix {
	return &Matrix{m: m, n: n, typ: typ, dat: make([]Scalar, m*n)}
}

// Returns a size x size identity of type typ. The identity Matrix is defined as a square
// Matrix (same number of rows and columns, nxn) with 1's along the main diagonal, and zeros elsewhere
// Multiplying by the identity matrix is guaranteed to keep the matrix it's multiplied by the same (hence "identity")
//
// For instance, a size of 2 yields
// 	[[1 0]]
// 	[[0 1]]
func Identity(size int, typ VecType) Matrix {
	dat := make([]Scalar, size*size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == j {
				dat[i*size+j] = MakeScalar(1, typ)
			} else {
				dat[i*size+j] = vecNumZero(typ)
			}
		}
	}
	return *unsafeMatrixFromSlice(dat, typ, size, size)
}

func InferMatrixFromCols(el [][]interface{}) (mat *Matrix, err error) {
	in := make([][]Scalar, len(el))
	initial,typ := InferScalarSlice(el[0])
	in[0] = initial
	for i,sl := range el[1:] {
		tmp,typ2 := InferScalarSlice(sl)
		if typ != typ2 {
			return nil,errors.New("Types didn't match")
		}
		in[i+1] = tmp
	}
	return unsafeMatrixFromCols(in,typ), nil
}

func InferMatrixFromRows(el [][]interface{}) (mat *Matrix, err error) {
	in := make([][]Scalar, len(el))
	initial,typ := InferScalarSlice(el[0])
	in[0] = initial
	for i,sl := range el[1:] {
		tmp,typ2 := InferScalarSlice(sl)
		if typ != typ2 {
			return nil,errors.New("Types didn't match")
		}
		in[i+1] = tmp
	}
	return unsafeMatrixFromRows(in,typ), nil
}

func InferMatrixFromSlice(el []interface{}, m,n int) (mat *Matrix, err error) {
	in,typ := InferScalarSlice(el)
	return unsafeMatrixFromSlice(in, typ, m, n), nil
}

// MatrixFromCols takes in a slice of matrix columns (represented as slices of Scalars). m and n are inferred from the sizes of the slice
// For instance:
//
// 	[[1, 1]]
// 	[[0, 1]] Would be entered as a 2D array [[1,0],[1,1]]
//
// But will be flattened to Row Major Order, [1,1,0,1]
//
// It returns nil and an error if the types of the underlying slice elements don't match
// typ
func MatrixFromCols(el [][]Scalar, typ VecType) (mat *Matrix, err error) {
	mat = &Matrix{}
	mat.typ = typ

	mat.n = len(el)
	mat.m = len(el[0])
	mat.dat = make([]Scalar, mat.m*mat.n)

	for i := 0; i < mat.m; i++ {
		for j := 0; j < mat.n; j++ {
			if !checkType(mat.typ, el[j][i]) {
				return nil, errors.New("Element type does not match matrix")
			}
			mat.dat[i*mat.n+j] = el[j][i]
		}
	}

	return mat, nil
}

func unsafeMatrixFromCols(el [][]Scalar, typ VecType) (mat *Matrix) {
	mat = &Matrix{}
	mat.typ = typ

	mat.n = len(el)
	mat.m = len(el[0])
	mat.dat = make([]Scalar, mat.m*mat.n)

	for i := 0; i < mat.m; i++ {
		for j := 0; j < mat.n; j++ {
			mat.dat[i*mat.n+j] = el[j][i]
		}
	}

	return mat
}

// MatrixFromRows is MatrixFromCols, except is takes a slice of matrix columns (as Scalar slices), and makes them a matrix. m and n are inferred
// For instance:
//
// 	[[1, 1]]
// 	S[[0, 1]] is entered as [[1,1],[0,1]]
//
// And will be flattened to [1,1,0,1]
//
// It returns nil and an error if the types of the underlying slice elements don't
// match typ
func MatrixFromRows(el [][]Scalar, typ VecType) (mat *Matrix, err error) {
	mat = &Matrix{}
	mat.typ = typ

	mat.m = len(el)
	mat.n = len(el[0])
	mat.dat = make([]Scalar, mat.m*mat.n)

	for i := 0; i < mat.m; i++ {
		for j := 0; j < mat.n; j++ {
			if !checkType(mat.typ, el[i][j]) {
				return nil, errors.New("Element type does not match matrix")
			}
			mat.dat[i*mat.n+j] = el[i][j]
		}
	}

	return mat, nil
}

func unsafeMatrixFromRows(el [][]Scalar, typ VecType) (mat *Matrix) {
	mat = &Matrix{}
	mat.typ = typ

	mat.m = len(el)
	mat.n = len(el[0])
	mat.dat = make([]Scalar, mat.m*mat.n)

	for i := 0; i < mat.m; i++ {
		for j := 0; j < mat.n; j++ {
			mat.dat[i*mat.n+j] = el[i][j]
		}
	}

	return mat
}

// MatrixFromSlice takes in a slice of scalars, a type-checking argument, and its dimensions and gives a pointer to a Matrix
// Matrix treats its input as if it were in Row Major Order, see MatrixFromRows comments, specifically how it's flattened, for reference
//
//  It returns nil and an error if the length of the slice isn't m*n or any of el's elements fail to match typ
func MatrixFromSlice(el []Scalar, typ VecType, m, n int) (mat *Matrix, err error) {
	if m*n != len(el) {
		return nil, errors.New("Matrix dimensions do not match data passed in")
	}

	mat = &Matrix{}
	mat.typ = typ
	mat.m = m
	mat.n = n

	for _, e := range el {
		if !checkType(mat.typ, e) {
			return nil, errors.New("Type of at least one element does not match declared type")
		}
	}

	mat.dat = el

	return mat, nil
}

// Quick and dirty internal function to make a matrix without spending time checking types
func unsafeMatrixFromSlice(el []Scalar, typ VecType, m, n int) (mat *Matrix) {
	mat = &Matrix{}
	mat.typ = typ
	mat.m = m
	mat.n = n
	mat.dat = el

	return mat
}

// Returns the vector's underlying VecType
func (m Matrix) Type() VecType {
	return m.typ
}

// Does an element-wise matrix comparison and returns true if it passes.
// That is, for each ordinal in the underlying slice it compares the values of both matrices
//
// If the matrices' types are different, or their dimensions are different (meaning their m's, n's, or both are different)
// it will return false automatically.
//
// As with other Equal functions in this package, for integers it's a straight == comparison, but floats are compared as in
// util.go:FloatEquals
func (m1 Matrix) Equal(m2 Matrix) (eq bool) {
	if m1.typ != m2.typ || m1.n != m2.n || m1.m != m2.m {
		return false
	}

	for i := 0; i < len(m1.dat); i++ {
		eq = m1.dat[i].Equal(m2.dat[i])
		if !eq {
			break
		}
	}

	return eq
}

// AsSlice() simply returns the Marix's underlying slice
func (m Matrix) AsSlice() []Scalar {
	return m.dat
}

// SetElement set's the ith row and jth column of the Matrix to the element el.
//
// It returns an error if i or j are out of bounds, or the type of el doesn't match that of the Matrix
func (mat *Matrix) SetElement(i, j int, el Scalar) error {
	if i > mat.m || j > mat.n || i < 0 || j < 0 {
		return errors.New("Dimensions out of bounds")
	}

	if !checkType(mat.typ, el) {
		return errors.New("Type of element does not match matrix's type")
	}

	mat.dat[i*mat.n+j] = el

	return nil
}

// GetElement returns the Scalar element at the ith row and jth column of the matrix, or nil if i or j are out of bounds
func (m Matrix) GetElement(i, j int) Scalar {
	if i > m.m || j > m.n || i < 0 || j < 0 {
		return nil
	}

	return m.dat[i*m.n+j]
}

// AsVector recognizes that a 1xn or mx1 Matrix is equivalent to a Vector.
// This explicitly converts the matrix to a vector if those conditions are true, if not, it returns the zero-type for a Vector
func (mat Matrix) AsVector() (v Vector) {
	if mat.m != 1 && mat.n != 1 {
		return v
	}

	vPoint, err := VectorOf(mat.dat, mat.typ)
	if err != nil {
		return v
	}

	return *vPoint
}

// ToScalar operates exactly like Vector's ToScalar. If a matrix is a 1x1 matrix, it returns m.dat[0]
func (mat Matrix) ToScalar() Scalar {
	if mat.m != 1 || mat.n != 1 {
		return nil
	}

	return mat.dat[0]
}

// Add does a traditional element-wise addition of two Matrices.
//
// 	[[a, b]]   [[e, f]]   [[a+e, b+f]]
// 	[[c, d]] + [[g, h]] = [[c+g, d+h]]
//
// It returns the zero-type for a Matrix if the number of rows, columns, or underlying types don't match
func (m1 Matrix) Add(m2 Matrix) (m3 Matrix) {
	if m1.typ != m2.typ || m1.m != m2.m || m1.n != m2.n {
		return
	}

	m3.typ = m1.typ
	m3.dat = make([]Scalar, len(m1.dat))
	m3.m = m1.m
	m3.n = m2.n

	for i := range m1.dat {
		m3.dat[i] = m1.dat[i].Add(m2.dat[i])
	}

	return m3
}

// Sub does a traditional element-wise subtraction of two Matrices.
//
// 	[[a, b]]   [[e, f]]   [[a-e, b-f]]
// 	[[c, d]] - [[g, h]] = [[c-g, d-h]]
//
// It returns the zero-type for a Matrix if the number of rows, columns, or underlying types don't match
func (m1 Matrix) Sub(m2 Matrix) (m3 Matrix) {
	if m1.typ != m2.typ || len(m1.dat) != len(m2.dat) {
		return
	}

	m3.typ = m1.typ
	m3.dat = make([]Scalar, len(m1.dat))
	m3.m = m1.m
	m3.n = m2.n

	for i := range m1.dat {
		m3.dat[i] = m1.dat[i].Sub(m2.dat[i])
	}

	return m3
}

// ScalarMul does an element-wise multiplcation of a Matrix's elements and a scalar value
//
//       [[w, x]]    [[c*w, c*x]]
//     c*[[y, z]] =  [[c*y, c*z]]
//
// It returns the Matrix zero-type if c is not of the same underlying type as the matrix
func (m1 Matrix) ScalarMul(c Scalar) (mat Matrix) {
	if !checkType(m1.typ, c) {
		return
	}

	dat := make([]Scalar, len(m1.dat))
	for i := range m1.dat {
		dat[i] = m1.dat[i].Mul(c)
	}

	return *unsafeMatrixFromSlice(dat, m1.typ, m1.m, m1.n)

}

// TODO: Maybe look into the Strassen algorithm for large square matrices and hard code common cases

// Mul does a simple textbook Matrix multiplication between MatrixMultiplyables
// In the case m2 is a Vector, it is multiplied with as if it were a nx1 matrix (row vector)
// In the case m2 is a Matrix, textbook matrix multiplication happens.
//
// In the case m1 is a 1x1 matrix, it will multiply with m2 (treated as a row vector if it's a vector) as if using ScalarMul
// on m1.ToScalar()
//
// It returns the zero-type for Matrix if m.n and m2.m (or vec.Size()) don't match, or if the underlying types don't match.
func (m1 Matrix) Mul(m2 MatrixMultiplyable) (m3 Matrix) {
	var indat []Scalar
	m, n, o := m1.m, m1.n, 0

	if vec, ok := m2.(Vector); ok {
		if vec.typ != m1.typ || n != len(vec.dat) {
			return
		}
		if m1.m == 1 && m1.n == 1 {
			m3 = vec.ScalarMul(m1.ToScalar()).AsMatrix(false)
			return m3
		}
		indat = vec.dat
		o = 1
	} else {
		mat := m2.(Matrix)
		if m1.typ != mat.typ || n != mat.m {
			return
		}
		if m1.m == 1 && m1.n == 1 {
			return mat.ScalarMul(m1.ToScalar())
		}
		indat = mat.dat
		o = mat.n
	}

	dat := make([]Scalar, m*o)

	for j := 0; j < o; j++ { // Columns of m2 and m3
		for i := 0; i < m; i++ { // Rows of m1 and m3
			for k := 0; k < n; k++ { // Columns of m1, rows of m2
				if dat[i*o+j] == nil {
					dat[i*o+j] = MakeScalar(0, m1.typ)
				}
				dat[i*o+j] = dat[i*o+j].Add(m1.dat[i*n+k].Mul(indat[k*o+j])) // I think, needs testing
			}
		}
	}

	return *unsafeMatrixFromSlice(dat, m1.typ, m, o)
}

/*
Batch Multiply, as its name implies, is supposed to Multiply a huge amount of matrices at once
Since matrix Mmltiplication is associative, it can do pieces of the problem at the same time.
Since starting a goroutine has some overhead, I'd wager it's probably not worth it to use this function unless you have
6-8+ larger (i.e. 3x3 or 4x4) matrices, but I haven't benchmarked it so until then who knows?

Make sure the matrices are in order, and that they can indeed be Multiplied. If not you'll end up with an untyped 0x0 matrix (a.k.a the "zero type" for a Matrix struct)

If the only input is a single vector, it will return it as a vector as if you called vector.AsMatrix(true), meaning as a *ROW* vector

Yes, this means that for vectors it assumes inner product rather than outer product.

Be careful with vectors either way. I'd *recommend* only using Matrices and converting yourself.
*/
func BatchMultiply(args []MatrixMultiplyable) Matrix {
	// Fun fact: Since in (Go's) integer division 3/2=1, in the case where you have an odd number
	// of matrices to Multiply, the only way a vector can end up alone is when it's on the left, meaning it has to be a row vector.
	// this allows us to not need a special case for a slice length of 3!
	if len(args) == 1 {
		// We're expected to return a Matrix, so if it's suddenly a vector Go will panic from bad typing
		if vec, ok := args[0].(Vector); ok {
			ret := vec.AsMatrix(true)
			return ret
		}
		return args[0].(Matrix)
	}

	var m1, m2 MatrixMultiplyable
	if len(args) > 2 {
		ch1 := make(chan Matrix)
		ch2 := make(chan Matrix)

		// Split up the work, matrix Mult is associative
		go batchMultHelper(ch1, args[0:len(args)/2])
		go batchMultHelper(ch2, args[len(args)/2:len(args)])

		m1 = <-ch1
		m2 = <-ch2
	} else {
		m1 = args[0]
		m2 = args[1]
	}

	return m1.Mul(m2)
}

// Simple wrapper for channels
func batchMultHelper(ch chan<- Matrix, args []MatrixMultiplyable) {
	ch <- BatchMultiply(args)
	close(ch)
}

// TODO: Well duh, Laplace expansion is O(n!) -- look into alt methods like LU Decomposition

// Det computes the Determinant for a square matrix. For a 2x2, 3x3, and 4x4 matrix the return operation is hard-coded
// For above that, it uses a (very, very slow) Laplace Expansion method. I don't advise using this for >4x4 matrices
//
// It is returned as a float64 for similar reasons to Vector.Len(), an integer values determinant isn't a huge travesty, but isn't
// really worth it either.
//
// It returns 0 if your matrix isn't square, so be careful.
func (m1 Matrix) Det() (det float64) {
	if m1.m != m1.n { // Determinants are only for square matrices
		return
	}

	if m1.m == 2 {
		return m1.dat[0].Fl64()*m1.dat[3].Fl64() - m1.dat[1].Fl64()*m1.dat[2].Fl64()
	} else if m1.m == 3 {
		return m1.dat[0].Fl64()*(m1.dat[4].Fl64()*m1.dat[8].Fl64()-m1.dat[5].Fl64()*m1.dat[7].Fl64()) -
			m1.dat[1].Fl64()*(m1.dat[3].Fl64()*m1.dat[8].Fl64()-m1.dat[5].Fl64()*m1.dat[6].Fl64()) +
			m1.dat[2].Fl64()*(m1.dat[3].Fl64()*m1.dat[7].Fl64()-m1.dat[4].Fl64()*m1.dat[6].Fl64())
	} else if m1.m == 4 {
		return m1.dat[0].Fl64()*
			(m1.dat[5].Fl64()*(m1.dat[10].Fl64()*m1.dat[15].Fl64()-m1.dat[14].Fl64()*m1.dat[11].Fl64())-
				m1.dat[6].Fl64()*(m1.dat[9].Fl64()*m1.dat[15].Fl64()-m1.dat[13].Fl64()*m1.dat[11].Fl64())+
				m1.dat[7].Fl64()*(m1.dat[9].Fl64()*m1.dat[14].Fl64()-m1.dat[13].Fl64()*m1.dat[10].Fl64())) -
			m1.dat[1].Fl64()*
				(m1.dat[4].Fl64()*(m1.dat[10].Fl64()*m1.dat[15].Fl64()-m1.dat[14].Fl64()*m1.dat[11].Fl64())-
					m1.dat[6].Fl64()*(m1.dat[8].Fl64()*m1.dat[15].Fl64()-m1.dat[12].Fl64()*m1.dat[11].Fl64())+
					m1.dat[7].Fl64()*(m1.dat[8].Fl64()*m1.dat[14].Fl64()-m1.dat[12].Fl64()*m1.dat[10].Fl64())) +
			m1.dat[2].Fl64()*
				(m1.dat[4].Fl64()*(m1.dat[9].Fl64()*m1.dat[15].Fl64()-m1.dat[13].Fl64()*m1.dat[11].Fl64())-
					m1.dat[5].Fl64()*(m1.dat[8].Fl64()*m1.dat[15].Fl64()-m1.dat[12].Fl64()*m1.dat[11].Fl64())+
					m1.dat[7].Fl64()*(m1.dat[8].Fl64()*m1.dat[13].Fl64()-m1.dat[12].Fl64()*m1.dat[9].Fl64())) -
			m1.dat[3].Fl64()*
				(m1.dat[4].Fl64()*(m1.dat[9].Fl64()*m1.dat[14].Fl64()-m1.dat[13].Fl64()*m1.dat[10].Fl64())-
					m1.dat[5].Fl64()*(m1.dat[8].Fl64()*m1.dat[14].Fl64()-m1.dat[12].Fl64()*m1.dat[10].Fl64())+
					m1.dat[6].Fl64()*(m1.dat[8].Fl64()*m1.dat[13].Fl64()-m1.dat[12].Fl64()*m1.dat[9].Fl64()))

	}
	for i := 0; i < m1.n; i++ {
		det += m1.GetElement(i, 0).Fl64() * m1.cofactor(i, 0)
	}

	return det
}

// Transpose returns the matrix's transpose. A transpose is where a Matrix's rows and columns are swapped
//
//               T    [[a, d]]
//    [[a, b, c]]     [[b, e]]
//    [[d, e, f]]  =  [[c, f]]
//
// This is also useful for converting the underlying slice to Column Major Order if you ever need to
func (m Matrix) Transpose() Matrix {
	dat := make([]Scalar, len(m.dat))

	for i := 0; i < m.n; i++ {
		for j := 0; j < m.m; j++ {
			dat[j+i*m.m] = m.dat[j*m.n+i] // Basically convert to CMO
		}
	}

	return *unsafeMatrixFromSlice(dat, m.typ, m.n, m.m)
}

// Inverse creates the inverse Matrix
//
//     -1           T
//    M  = 1/Det * M
//
// That is, the inverse is one divided by the determinant times the matrix's transpose.
//
// In the case your determinant is Zero (also known as "being non-Singular"), or you have a non-square Matrix
// it will return the zero type.
func (m Matrix) Inverse() (m2 Matrix) {
	det := m.Det()
	if FloatEqual(det, 0.) {
		return
	}
	return m.Transpose().floatScale(float64(1.0) / det)
}

// INTERNAL: Same as floatScale on a vec, allows us to ease things over with integer-values matrices
func (m Matrix) floatScale(c float64) Matrix {
	dat := make([]Scalar, len(m.dat))
	for i, el := range m.dat {
		dat[i] = el.mulFl64(c)
	}

	return *unsafeMatrixFromSlice(dat, m.typ, m.m, m.n)
}

// AsArray returns the Matrix as a 1D array of the appropriate size and same Go type as the underlying Matrix
// As with Vector, it has to return an interface{} because of the face that in Go arrays have a fixes length and type,
// thus the user must cast the result to the correct type
//
// The array returned, as with the underlying slice, is in Row Major Order. If this doesn't work for you, call AsArray on the Matrix's transpose.
// This function is useful for passing into OpenGL functions that take arrays
//
// It returns nil should the matrix have more than four rows or columns.
func (m Matrix) AsArray(typ VecType) interface{} {
	if m.n < 1 || m.m < 1 || m.m > 4 || m.n > 4 {
		return nil
	}

	switch typ {
	case INT32:
		switch m.m * m.n {
		case 1:
			return [1]int32{m.dat[0].Int32()}
		case 2:
			return [2]int32{m.dat[0].Int32(), m.dat[1].Int32()}
		case 3:
			return [3]int32{m.dat[0].Int32(), m.dat[1].Int32(), m.dat[2].Int32()}
		case 4:
			return [4]int32{m.dat[0].Int32(), m.dat[1].Int32(), m.dat[2].Int32(), m.dat[3].Int32()}
		case 6:
			return [6]int32{m.dat[0].Int32(), m.dat[1].Int32(), m.dat[2].Int32(), m.dat[3].Int32(), m.dat[4].Int32(), m.dat[5].Int32()}
		case 8:
			return [8]int32{m.dat[0].Int32(), m.dat[1].Int32(), m.dat[2].Int32(), m.dat[3].Int32(), m.dat[4].Int32(), m.dat[5].Int32(), m.dat[6].Int32(), m.dat[7].Int32()}
		case 9:
			return [9]int32{m.dat[0].Int32(), m.dat[1].Int32(), m.dat[2].Int32(), m.dat[3].Int32(), m.dat[4].Int32(), m.dat[5].Int32(), m.dat[6].Int32(), m.dat[7].Int32(), m.dat[8].Int32()}
		case 12:
			return [12]int32{m.dat[0].Int32(), m.dat[1].Int32(), m.dat[2].Int32(), m.dat[3].Int32(), m.dat[4].Int32(), m.dat[5].Int32(), m.dat[6].Int32(), m.dat[7].Int32(), m.dat[8].Int32(),
				m.dat[9].Int32(), m.dat[10].Int32(), m.dat[11].Int32()}
		case 16:
			return [16]int32{m.dat[0].Int32(), m.dat[1].Int32(), m.dat[2].Int32(), m.dat[3].Int32(), m.dat[4].Int32(), m.dat[5].Int32(), m.dat[6].Int32(), m.dat[7].Int32(), m.dat[8].Int32(),
				m.dat[9].Int32(), m.dat[10].Int32(), m.dat[11].Int32(), m.dat[12].Int32(), m.dat[13].Int32(), m.dat[14].Int32(), m.dat[15].Int32()}
		}

	case UINT32:
		switch m.m * m.n {
		case 1:
			return [1]uint32{m.dat[0].Uint32()}
		case 2:
			return [2]uint32{m.dat[0].Uint32(), m.dat[1].Uint32()}
		case 3:
			return [3]uint32{m.dat[0].Uint32(), m.dat[1].Uint32(), m.dat[2].Uint32()}
		case 4:
			return [4]uint32{m.dat[0].Uint32(), m.dat[1].Uint32(), m.dat[2].Uint32(), m.dat[3].Uint32()}
		case 6:
			return [6]uint32{m.dat[0].Uint32(), m.dat[1].Uint32(), m.dat[2].Uint32(), m.dat[3].Uint32(), m.dat[4].Uint32(), m.dat[5].Uint32()}
		case 8:
			return [8]uint32{m.dat[0].Uint32(), m.dat[1].Uint32(), m.dat[2].Uint32(), m.dat[3].Uint32(), m.dat[4].Uint32(), m.dat[5].Uint32(), m.dat[6].Uint32(), m.dat[7].Uint32()}
		case 9:
			return [9]uint32{m.dat[0].Uint32(), m.dat[1].Uint32(), m.dat[2].Uint32(), m.dat[3].Uint32(), m.dat[4].Uint32(), m.dat[5].Uint32(), m.dat[6].Uint32(), m.dat[7].Uint32(), m.dat[8].Uint32()}
		case 12:
			return [12]uint32{m.dat[0].Uint32(), m.dat[1].Uint32(), m.dat[2].Uint32(), m.dat[3].Uint32(), m.dat[4].Uint32(), m.dat[5].Uint32(), m.dat[6].Uint32(), m.dat[7].Uint32(), m.dat[8].Uint32(),
				m.dat[9].Uint32(), m.dat[10].Uint32(), m.dat[11].Uint32()}
		case 16:
			return [16]uint32{m.dat[0].Uint32(), m.dat[1].Uint32(), m.dat[2].Uint32(), m.dat[3].Uint32(), m.dat[4].Uint32(), m.dat[5].Uint32(), m.dat[6].Uint32(), m.dat[7].Uint32(), m.dat[8].Uint32(),
				m.dat[9].Uint32(), m.dat[10].Uint32(), m.dat[11].Uint32(), m.dat[12].Uint32(), m.dat[13].Uint32(), m.dat[14].Uint32(), m.dat[15].Uint32()}
		}

	case FLOAT32:
		switch m.m * m.n {
		case 1:
			return [1]float32{m.dat[0].Fl32()}
		case 2:
			return [2]float32{m.dat[0].Fl32(), m.dat[1].Fl32()}
		case 3:
			return [3]float32{m.dat[0].Fl32(), m.dat[1].Fl32(), m.dat[2].Fl32()}
		case 4:
			return [4]float32{m.dat[0].Fl32(), m.dat[1].Fl32(), m.dat[2].Fl32(), m.dat[3].Fl32()}
		case 6:
			return [6]float32{m.dat[0].Fl32(), m.dat[1].Fl32(), m.dat[2].Fl32(), m.dat[3].Fl32(), m.dat[4].Fl32(), m.dat[5].Fl32()}
		case 8:
			return [8]float32{m.dat[0].Fl32(), m.dat[1].Fl32(), m.dat[2].Fl32(), m.dat[3].Fl32(), m.dat[4].Fl32(), m.dat[5].Fl32(), m.dat[6].Fl32(), m.dat[7].Fl32()}
		case 9:
			return [9]float32{m.dat[0].Fl32(), m.dat[1].Fl32(), m.dat[2].Fl32(), m.dat[3].Fl32(), m.dat[4].Fl32(), m.dat[5].Fl32(), m.dat[6].Fl32(), m.dat[7].Fl32(), m.dat[8].Fl32()}
		case 12:
			return [12]float32{m.dat[0].Fl32(), m.dat[1].Fl32(), m.dat[2].Fl32(), m.dat[3].Fl32(), m.dat[4].Fl32(), m.dat[5].Fl32(), m.dat[6].Fl32(), m.dat[7].Fl32(), m.dat[8].Fl32(),
				m.dat[9].Fl32(), m.dat[10].Fl32(), m.dat[11].Fl32()}
		case 16:
			return [16]float32{m.dat[0].Fl32(), m.dat[1].Fl32(), m.dat[2].Fl32(), m.dat[3].Fl32(), m.dat[4].Fl32(), m.dat[5].Fl32(), m.dat[6].Fl32(), m.dat[7].Fl32(), m.dat[8].Fl32(),
				m.dat[9].Fl32(), m.dat[10].Fl32(), m.dat[11].Fl32(), m.dat[12].Fl32(), m.dat[13].Fl32(), m.dat[14].Fl32(), m.dat[15].Fl32()}
		}

	case FLOAT64:
		switch m.m * m.n {
		case 1:
			return [1]float64{m.dat[0].Fl64()}
		case 2:
			return [2]float64{m.dat[0].Fl64(), m.dat[1].Fl64()}
		case 3:
			return [3]float64{m.dat[0].Fl64(), m.dat[1].Fl64(), m.dat[2].Fl64()}
		case 4:
			return [4]float64{m.dat[0].Fl64(), m.dat[1].Fl64(), m.dat[2].Fl64(), m.dat[3].Fl64()}
		case 6:
			return [6]float64{m.dat[0].Fl64(), m.dat[1].Fl64(), m.dat[2].Fl64(), m.dat[3].Fl64(), m.dat[4].Fl64(), m.dat[5].Fl64()}
		case 8:
			return [8]float64{m.dat[0].Fl64(), m.dat[1].Fl64(), m.dat[2].Fl64(), m.dat[3].Fl64(), m.dat[4].Fl64(), m.dat[5].Fl64(), m.dat[6].Fl64(), m.dat[7].Fl64()}
		case 9:
			return [9]float64{m.dat[0].Fl64(), m.dat[1].Fl64(), m.dat[2].Fl64(), m.dat[3].Fl64(), m.dat[4].Fl64(), m.dat[5].Fl64(), m.dat[6].Fl64(), m.dat[7].Fl64(), m.dat[8].Fl64()}
		case 12:
			return [12]float64{m.dat[0].Fl64(), m.dat[1].Fl64(), m.dat[2].Fl64(), m.dat[3].Fl64(), m.dat[4].Fl64(), m.dat[5].Fl64(), m.dat[6].Fl64(), m.dat[7].Fl64(), m.dat[8].Fl64(),
				m.dat[9].Fl64(), m.dat[10].Fl64(), m.dat[11].Fl64()}
		case 16:
			return [16]float64{m.dat[0].Fl64(), m.dat[1].Fl64(), m.dat[2].Fl64(), m.dat[3].Fl64(), m.dat[4].Fl64(), m.dat[5].Fl64(), m.dat[6].Fl64(), m.dat[7].Fl64(), m.dat[8].Fl64(),
				m.dat[9].Fl64(), m.dat[10].Fl64(), m.dat[11].Fl64(), m.dat[12].Fl64(), m.dat[13].Fl64(), m.dat[14].Fl64(), m.dat[15].Fl64()}
		}
	}

	return nil
}

func (m Matrix) AsCMOArray(typ VecType) interface{} {
	if m.n < 1 || m.m < 1 || m.m > 4 || m.n > 4 {
		return nil
	}
	m = m.Transpose() // TODO: Spell out the CMO literally if benchmarks prove this too slow

	switch typ {
	case INT32:
		switch m.m * m.n {
		case 1:
			return [1]int32{m.dat[0].Int32()}
		case 2:
			return [2]int32{m.dat[0].Int32(), m.dat[1].Int32()}
		case 3:
			return [3]int32{m.dat[0].Int32(), m.dat[1].Int32(), m.dat[2].Int32()}
		case 4:
			return [4]int32{m.dat[0].Int32(), m.dat[1].Int32(), m.dat[2].Int32(), m.dat[3].Int32()}
		case 6:
			return [6]int32{m.dat[0].Int32(), m.dat[1].Int32(), m.dat[2].Int32(), m.dat[3].Int32(), m.dat[4].Int32(), m.dat[5].Int32()}
		case 8:
			return [8]int32{m.dat[0].Int32(), m.dat[1].Int32(), m.dat[2].Int32(), m.dat[3].Int32(), m.dat[4].Int32(), m.dat[5].Int32(), m.dat[6].Int32(), m.dat[7].Int32()}
		case 9:
			return [9]int32{m.dat[0].Int32(), m.dat[1].Int32(), m.dat[2].Int32(), m.dat[3].Int32(), m.dat[4].Int32(), m.dat[5].Int32(), m.dat[6].Int32(), m.dat[7].Int32(), m.dat[8].Int32()}
		case 12:
			return [12]int32{m.dat[0].Int32(), m.dat[1].Int32(), m.dat[2].Int32(), m.dat[3].Int32(), m.dat[4].Int32(), m.dat[5].Int32(), m.dat[6].Int32(), m.dat[7].Int32(), m.dat[8].Int32(),
				m.dat[9].Int32(), m.dat[10].Int32(), m.dat[11].Int32()}
		case 16:
			return [16]int32{m.dat[0].Int32(), m.dat[1].Int32(), m.dat[2].Int32(), m.dat[3].Int32(), m.dat[4].Int32(), m.dat[5].Int32(), m.dat[6].Int32(), m.dat[7].Int32(), m.dat[8].Int32(),
				m.dat[9].Int32(), m.dat[10].Int32(), m.dat[11].Int32(), m.dat[12].Int32(), m.dat[13].Int32(), m.dat[14].Int32(), m.dat[15].Int32()}
		}

	case UINT32:
		switch m.m * m.n {
		case 1:
			return [1]uint32{m.dat[0].Uint32()}
		case 2:
			return [2]uint32{m.dat[0].Uint32(), m.dat[1].Uint32()}
		case 3:
			return [3]uint32{m.dat[0].Uint32(), m.dat[1].Uint32(), m.dat[2].Uint32()}
		case 4:
			return [4]uint32{m.dat[0].Uint32(), m.dat[1].Uint32(), m.dat[2].Uint32(), m.dat[3].Uint32()}
		case 6:
			return [6]uint32{m.dat[0].Uint32(), m.dat[1].Uint32(), m.dat[2].Uint32(), m.dat[3].Uint32(), m.dat[4].Uint32(), m.dat[5].Uint32()}
		case 8:
			return [8]uint32{m.dat[0].Uint32(), m.dat[1].Uint32(), m.dat[2].Uint32(), m.dat[3].Uint32(), m.dat[4].Uint32(), m.dat[5].Uint32(), m.dat[6].Uint32(), m.dat[7].Uint32()}
		case 9:
			return [9]uint32{m.dat[0].Uint32(), m.dat[1].Uint32(), m.dat[2].Uint32(), m.dat[3].Uint32(), m.dat[4].Uint32(), m.dat[5].Uint32(), m.dat[6].Uint32(), m.dat[7].Uint32(), m.dat[8].Uint32()}
		case 12:
			return [12]uint32{m.dat[0].Uint32(), m.dat[1].Uint32(), m.dat[2].Uint32(), m.dat[3].Uint32(), m.dat[4].Uint32(), m.dat[5].Uint32(), m.dat[6].Uint32(), m.dat[7].Uint32(), m.dat[8].Uint32(),
				m.dat[9].Uint32(), m.dat[10].Uint32(), m.dat[11].Uint32()}
		case 16:
			return [16]uint32{m.dat[0].Uint32(), m.dat[1].Uint32(), m.dat[2].Uint32(), m.dat[3].Uint32(), m.dat[4].Uint32(), m.dat[5].Uint32(), m.dat[6].Uint32(), m.dat[7].Uint32(), m.dat[8].Uint32(),
				m.dat[9].Uint32(), m.dat[10].Uint32(), m.dat[11].Uint32(), m.dat[12].Uint32(), m.dat[13].Uint32(), m.dat[14].Uint32(), m.dat[15].Uint32()}
		}

	case FLOAT32:
		switch m.m * m.n {
		case 1:
			return [1]float32{m.dat[0].Fl32()}
		case 2:
			return [2]float32{m.dat[0].Fl32(), m.dat[1].Fl32()}
		case 3:
			return [3]float32{m.dat[0].Fl32(), m.dat[1].Fl32(), m.dat[2].Fl32()}
		case 4:
			return [4]float32{m.dat[0].Fl32(), m.dat[1].Fl32(), m.dat[2].Fl32(), m.dat[3].Fl32()}
		case 6:
			return [6]float32{m.dat[0].Fl32(), m.dat[1].Fl32(), m.dat[2].Fl32(), m.dat[3].Fl32(), m.dat[4].Fl32(), m.dat[5].Fl32()}
		case 8:
			return [8]float32{m.dat[0].Fl32(), m.dat[1].Fl32(), m.dat[2].Fl32(), m.dat[3].Fl32(), m.dat[4].Fl32(), m.dat[5].Fl32(), m.dat[6].Fl32(), m.dat[7].Fl32()}
		case 9:
			return [9]float32{m.dat[0].Fl32(), m.dat[1].Fl32(), m.dat[2].Fl32(), m.dat[3].Fl32(), m.dat[4].Fl32(), m.dat[5].Fl32(), m.dat[6].Fl32(), m.dat[7].Fl32(), m.dat[8].Fl32()}
		case 12:
			return [12]float32{m.dat[0].Fl32(), m.dat[1].Fl32(), m.dat[2].Fl32(), m.dat[3].Fl32(), m.dat[4].Fl32(), m.dat[5].Fl32(), m.dat[6].Fl32(), m.dat[7].Fl32(), m.dat[8].Fl32(),
				m.dat[9].Fl32(), m.dat[10].Fl32(), m.dat[11].Fl32()}
		case 16:
			return [16]float32{m.dat[0].Fl32(), m.dat[1].Fl32(), m.dat[2].Fl32(), m.dat[3].Fl32(), m.dat[4].Fl32(), m.dat[5].Fl32(), m.dat[6].Fl32(), m.dat[7].Fl32(), m.dat[8].Fl32(),
				m.dat[9].Fl32(), m.dat[10].Fl32(), m.dat[11].Fl32(), m.dat[12].Fl32(), m.dat[13].Fl32(), m.dat[14].Fl32(), m.dat[15].Fl32()}
		}

	case FLOAT64:
		switch m.m * m.n {
		case 1:
			return [1]float64{m.dat[0].Fl64()}
		case 2:
			return [2]float64{m.dat[0].Fl64(), m.dat[1].Fl64()}
		case 3:
			return [3]float64{m.dat[0].Fl64(), m.dat[1].Fl64(), m.dat[2].Fl64()}
		case 4:
			return [4]float64{m.dat[0].Fl64(), m.dat[1].Fl64(), m.dat[2].Fl64(), m.dat[3].Fl64()}
		case 6:
			return [6]float64{m.dat[0].Fl64(), m.dat[1].Fl64(), m.dat[2].Fl64(), m.dat[3].Fl64(), m.dat[4].Fl64(), m.dat[5].Fl64()}
		case 8:
			return [8]float64{m.dat[0].Fl64(), m.dat[1].Fl64(), m.dat[2].Fl64(), m.dat[3].Fl64(), m.dat[4].Fl64(), m.dat[5].Fl64(), m.dat[6].Fl64(), m.dat[7].Fl64()}
		case 9:
			return [9]float64{m.dat[0].Fl64(), m.dat[1].Fl64(), m.dat[2].Fl64(), m.dat[3].Fl64(), m.dat[4].Fl64(), m.dat[5].Fl64(), m.dat[6].Fl64(), m.dat[7].Fl64(), m.dat[8].Fl64()}
		case 12:
			return [12]float64{m.dat[0].Fl64(), m.dat[1].Fl64(), m.dat[2].Fl64(), m.dat[3].Fl64(), m.dat[4].Fl64(), m.dat[5].Fl64(), m.dat[6].Fl64(), m.dat[7].Fl64(), m.dat[8].Fl64(),
				m.dat[9].Fl64(), m.dat[10].Fl64(), m.dat[11].Fl64()}
		case 16:
			return [16]float64{m.dat[0].Fl64(), m.dat[1].Fl64(), m.dat[2].Fl64(), m.dat[3].Fl64(), m.dat[4].Fl64(), m.dat[5].Fl64(), m.dat[6].Fl64(), m.dat[7].Fl64(), m.dat[8].Fl64(),
				m.dat[9].Fl64(), m.dat[10].Fl64(), m.dat[11].Fl64(), m.dat[12].Fl64(), m.dat[13].Fl64(), m.dat[14].Fl64(), m.dat[15].Fl64()}
		}
	}

	return nil
}

// Unexported because it may be buggy
func (m1 Matrix) cofactor(i, j int) float64 {
	return math.Pow(float64(-1), float64(i+j)) * m1.minorMatrix(i, j).Det() // C^(i,j) * Minor(i,j)
}

// Unexported because it may be buggy, or possibly just slow
func (m1 Matrix) minorMatrix(i, j int) Matrix {
	dat := make([]Scalar, (m1.m-1)*(m1.n-1))

	for k := 0; k < m1.m; k++ {
		var q int
		if k < j {
			q = k
		} else if k == j {
			continue
		} else {
			q = k - q
		}
		for l := 0; l < m1.n; k++ {
			var r int
			if l < i {
				r = l
			} else if l == i {
				continue
			} else {
				r = l - 1
			}

			dat[r*(m1.n-1)+q] = m1.dat[l*m1.n+k]
		}
	}

	return *unsafeMatrixFromSlice(dat, m1.typ, m1.m-1, m1.n-1)
}
