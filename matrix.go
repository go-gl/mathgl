package mathgl

import (
	"errors"
)

type MatrixMultiplyable interface {
	Mul(m MatrixMultiplyable) Matrix
}

type Matrix struct {
	m, n int // an m x n matrix
	typ  VecType
	dat  []Scalar
}

func NewMatrix(m, n int, typ VecType) *Matrix {
	return &Matrix{m: m, n: n, typ: typ, dat: make([]Scalar, 0, 2)}
}

func Identity(size int, typ VecType) Matrix {
	dat := make([]Scalar, size*size)
	
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if i == j {
				dat[i * size + j] = MakeScalar(1, typ)
			} else {
				dat[i * size + j] = vecNumZero(typ)
			}
		}
	}
	return *unsafeMatrixFromSlice(typ, dat, size, size)
}

// [1, 1]
// [0, 1] Would be entered as a 2D array [[1,1],[0,1]] -- but converted to RMO
//
// This may seem confusing, but it's because it's easier to type out and visualize things in CMO
// So it's easier to type write your matrix as a slice in CMO, and pass it into this method
func MatrixFromCols(typ VecType, el [][]Scalar) (mat *Matrix, err error) {
	mat.typ = typ

	mat.m = len(el)
	mat.n = len(el[0])
	mat.dat = make([]Scalar, 0, mat.m*mat.n)

	// Row Major Order, like in OpenGL
	for i := 0; i < mat.n; i++ {
		for j := 0; j < mat.m; j++ {
			if !checkType(mat.typ, el[j][i]) {
				return nil, errors.New("Element type does not match matrix")
			}
			mat.dat = append(mat.dat, el[j][i])
		}
	}

	return mat, nil
}


// This function is MatrixOf, except it takes a list of row "vectors" instead of row "vectors" (really slices)
func MatrixFromRows(typ VecType, el [][]Scalar) (mat *Matrix, err error) {
	mat.typ = typ

	mat.m = len(el)
	mat.n = len(el[0])
	mat.dat = make([]Scalar, 0, mat.m*mat.n)

	// Row Major Order, like in OpenGL
	for i := 0; i < mat.m; i++ {
		for j := 0; j < mat.n; j++ {
			if !checkType(mat.typ, el[j][i]) {
				return nil, errors.New("Element type does not match matrix")
			}
			mat.dat = append(mat.dat, el[j][i])
		}
	}

	return mat, nil
}

// Slice-format data should be in Row Major Order
func MatrixFromSlice(typ VecType, el []Scalar, m, n int) (mat *Matrix, err error) {
	mat = &Matrix{}
	mat.typ = typ
	mat.m = m
	mat.n = n

	if mat.m*mat.n != len(el) {
		return nil, errors.New("Matrix dimensions do not match data passed in")
	}

	for _, e := range el {
		if !checkType(mat.typ, e) {
			return nil, errors.New("Type of at least one element does not match declared type")
		}
	}

	mat.dat = el

	return mat, nil
}

// Quick and dirty internal function to make a matrix without spending time checking types
func unsafeMatrixFromSlice(typ VecType, el []Scalar, m, n int) (mat *Matrix) {
	mat = &Matrix{}
	mat.typ = typ
	mat.m = m
	mat.n = n

	mat.dat = el

	return mat
}

func (m Matrix) Type() VecType {
	return m.typ
}

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

func (m Matrix) AsSlice() []Scalar {
	return m.dat
}

// TODO: "Append" data method (expand the matrix)

func (mat *Matrix) SetElement(i, j int, el Scalar) error {
	if i < mat.m || j < mat.n {
		return errors.New("Dimensions out of bounds")
	}

	if !checkType(mat.typ, el) {
		return errors.New("Type of element does not match matrix's type")
	}

	mat.dat[mat.m*j+i] = el

	return nil
}

func (mat Matrix) AsVector() (v Vector, err error) {
	if mat.m != 1 && mat.n != 1 {
		return v, errors.New("Matrix is not 1-dimensional in either direction.")
	}

	vPoint, err := VectorOf(mat.typ, mat.dat)
	if err != nil {
		return v, err
	}

	return *vPoint, nil
}

func (mat Matrix) ToScalar() Scalar {
	if mat.m != 1 || mat.n != 1 {
		return nil
	}

	return mat.dat[0]
}

func (m1 Matrix) Add(m2 Matrix) (m3 Matrix) {
	if m1.typ != m2.typ || len(m1.dat) != len(m2.dat) {
		return
	}

	m3.typ = m1.typ
	m3.dat = make([]Scalar, len(m1.dat))

	for i := range m1.dat {
		m3.dat[i] = m1.dat[i].Add(m2.dat[i])
	}

	return m3
}

func (m1 Matrix) Sub(m2 Matrix) (m3 Matrix) {
	if m1.typ != m2.typ || len(m1.dat) != len(m2.dat) {
		return
	}

	m3.typ = m1.typ
	m3.dat = make([]Scalar, len(m1.dat))

	for i := range m1.dat {
		m3.dat[i] = m1.dat[i].Sub(m2.dat[i])
	}

	return m3
}

func (m1 Matrix) ScalarMul(c Scalar) (mat Matrix) {
	if !checkType(m1.typ, c) {
		return
	}
	
	dat := make([]Scalar, len(m1.dat))
	for i := range m1.dat {
		dat[i] = m1.dat[i].Mul(c)
	}
	
	return *unsafeMatrixFromSlice(m1.typ, dat, m1.m, m1.n)
	
}

func (m1 Matrix) Mul(m2 MatrixMultiplyable) (m3 Matrix) {
	var indat []Scalar
	m,n,o := m1.m, m1.n, 0
	
	if vec,ok := m2.(Vector); ok {
		if vec.typ != m1.typ || n != len(vec.dat) {
			return
		}
		if m1.m == 1 && m1.n == 1 {
			m3,_ = vec.ScalarMul(m1.ToScalar()).AsMatrix(false)
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
				dat[j*o+i] = dat[j*o+i].Add(m1.dat[k*n+i].Mul(indat[j*o+k])) // I think, needs testing
			}
		}
	}

	return *unsafeMatrixFromSlice(m1.typ, dat, m, o)
}


/* 
Batch Multiply, as its name implies, is supposed to Multiply a huge amount of matrices at once
Since matrix Multiplication is associative, it can do pieces of the problem at the same time.
Since starting a goroutine has some overhead, I'd wager it's probably not worth it to use this function unless you have
8+ matrices, but I haven't benchmarked it so until then who knows?

Make sure the matrices are in order, and that they can indeed be Multiplied. If not you'll end up with an untyped 0x0 matrix (a.k.a the "zero type" for a Matrix struct)

If the only input is a single vector, it will return it as a vector as if you called vector.AsMatrix(true), meaning as a *ROW* vector
*/
func BatchMultiply(args []MatrixMultiplyable) Matrix {
	// Fun fact: Since in (Go's) integer division 3/2=1, in the case where you have an odd number 
	// of matrices to Multiply, the only way a vector can end up alone is when it's on the left, meaning it has to be a row vector.
	// this allows us to not need a special case for a slice length of 3!
	if len(args) == 1 {
		// We're expected to return a Matrix, so if it's suddenly a vector Go will panic from bad typing
		if vec,ok := args[0].(Vector); ok {
			ret,_ := vec.AsMatrix(true)
			return ret
		}
		return args[0].(Matrix)
	}
	
	var m1,m2 MatrixMultiplyable
	if len(args) > 2 {
			ch1 := make(chan Matrix)
			ch2 := make(chan Matrix)
			
			// Split up the work, matrix Mult is associative
			go batchMultHelper(ch1, args[0:len(args)/2])
			go batchMultHelper(ch2, args[len(args)/2:len(args)])
			
			m1 = <- ch1
			m2 = <- ch2
	} else {
		m1 = args[0]
		m2 = args[1]
	}
	
	
	return m1.Mul(m2)
}

func batchMultHelper(ch chan<- Matrix, args []MatrixMultiplyable) {
	ch <- BatchMultiply(args)
	close(ch)
}

// INCOMPLETE DO NOT USE
// Need better function to make matrices for recursion
func (m1 Matrix) Det() Scalar {
	if m1.m != m1.n { // Determinants are only for square matrices
		return nil
	}

	return nil
}

func (m Matrix) Transpose() Matrix {
	dat := make([]Scalar, len(m.dat))
	
	for i := 0; i < m.n; i++ {
		for j := 0; j < m.m; j++ {
			dat[j + i * m.m] = m.dat[j * m.n + i] // Basically convert to CMO
		}
	}
	
	return *unsafeMatrixFromSlice(m.typ, dat, m.n, m.m)
}
