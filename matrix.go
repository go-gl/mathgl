package mathgl

import (
	"errors"
)

type Matrix struct {
	m, n int // an m x n matrix
	typ  VecType
	dat  []VecNum
}

func NewMatrix(m, n int, typ VecType) *Matrix {
	return &Matrix{m: m, n: n, typ: typ, dat: make([]VecNum, 0, 2)}
}

// [1, 1]
// [0, 1] Would be entered as a 2D array [[1,1],[0,1]] -- but converted to RMO
//
// This may seem confusing, but it's because it's easier to type out and visualize things in CMO
// So it's easier to type write your matrix as a slice in CMO, and pass it into this method
func MatrixFromCols(typ VecType, el [][]VecNum) (mat *Matrix, err error) {
	mat.typ = typ

	mat.m = len(el)
	mat.n = len(el[0])
	mat.dat = make([]VecNum, 0, mat.m*mat.n)

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
func MatrixFromRows(typ VecType, el [][]VecNum) (mat *Matrix, err error) {
	mat.typ = typ

	mat.m = len(el)
	mat.n = len(el[0])
	mat.dat = make([]VecNum, 0, mat.m*mat.n)

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
func MatrixFromSlice(typ VecType, el []VecNum, m, n int) (mat *Matrix, err error) {
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
func unsafeMatrixFromSlice(typ VecType, el []VecNum, m, n int) (mat *Matrix, err error) {
	mat.typ = typ
	mat.m = m
	mat.n = n

	/*if mat.m * mat.n != len(el) {
		return nil, errors.New("Matrix dimensions do not match data passed in")
	}

	for _,e := range el {
		if !checkType(mat.typ, e) {
			return nil, errors.New("Type of at least one element does not match declared type")
		}
	}*/

	mat.dat = el

	return mat, nil
}

// TODO: "Add" or "Append" data method (expand the matrix)

func (mat *Matrix) SetElement(i, j int, el VecNum) error {
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

func (mat Matrix) ToScalar() VecNum {
	if mat.m != 1 || mat.n != 1 {
		return nil
	}

	/*switch mat.typ {
	case INT32:
		return mat.dat[0].(int32)
	case UINT32:
		return mat.dat[0].(uint32)
	case FLOAT32:
		return mat.dat[0].(float32)
	case FLOAT64:
		return mat.dat[0].(float64)
	}*/

	return mat.dat[0]
}

func (m1 Matrix) Add(m2 Matrix) (m3 Matrix) {
	if m1.typ != m2.typ || len(m1.dat) != len(m2.dat) {
		return
	}

	m3.typ = m1.typ
	m3.dat = make([]VecNum, len(m1.dat))

	for i := range m1.dat {
		m3.dat[i] = m1.dat[i].add(m2.dat[i])
		/*switch m1.typ {
		case INT32:
			m3.dat[i] = m1.dat[i].(int32) + m2.dat[i].(int32)
		case UINT32:
			m3.dat[i] = m1.dat[i].(uint32) + m2.dat[i].(uint32)
		case FLOAT32:
			m3.dat[i] = m1.dat[i].(float32) + m2.dat[i].(float32)
		case FLOAT64:
			m3.dat[i] = m1.dat[i].(float64) + m2.dat[i].(float64)
		}*/
	}

	return m3
}

func (m1 Matrix) Sub(m2 Matrix) (m3 Matrix) {
	if m1.typ != m2.typ || len(m1.dat) != len(m2.dat) {
		return
	}

	m3.typ = m1.typ
	m3.dat = make([]VecNum, len(m1.dat))

	for i := range m1.dat {
		m3.dat[i] = m1.dat[i].sub(m2.dat[i])
		/*switch m1.typ {
		case INT32:
			m3.dat[i] = m1.dat[i].(int32) - m2.dat[i].(int32)
		case UINT32:
			m3.dat[i] = m1.dat[i].(uint32) - m2.dat[i].(uint32)
		case FLOAT32:
			m3.dat[i] = m1.dat[i].(float32) - m2.dat[i].(float32)
		case FLOAT64:
			m3.dat[i] = m1.dat[i].(float64) - m2.dat[i].(float64)
		}*/
	}

	return m3
}

func (m1 Matrix) Mul(m2 Matrix) (m3 Matrix) {
	if m1.n != m2.m || m1.typ != m2.typ {
		return
	}
	dat := make([]VecNum, m1.m*m2.n)

	for j := 0; j < m2.n; j++ { // Columns of m2 and m3
		for i := 0; i < m1.m; i++ { // Rows of m1 and m3
			for k := 0; k < m1.n; k++ { // Columns of m1, rows of m2
				dat[j*m2.n+i] = dat[j*m2.n+i].add(m1.dat[k*m1.n+i].mul(m2.dat[j*m2.n+k])) // I think, needs testing
			}
		}
	}
	/*switch m1.typ {
	case INT32:
		for j := 0; j < m2.n; j++ { // Columns of m2 and m3
			for i := 0; i < m1.m; i++ { // Rows of m1 and m3
				for k := 0; k < m1.n; k++ { // Columns of m1, rows of m2
					dat[j * m2.n + i] = dat[j * m2.n + i].(int32) + m1.dat[k * m1.n + i].(int32) * m2.dat[j * m2.n + k].(int32) // I think, needs testing
				}
			}
		}
	case UINT32:
		for j := 0; j < m2.n; j++ { // Columns of m2 and m3
			for i := 0; i < m1.m; i++ { // Rows of m1 and m3
				for k := 0; k < m1.n; k++ { // Columns of m1, rows of m2
					dat[j * m2.n + i] = dat[j * m2.n + i].(uint32) + m1.dat[k * m1.n + i].(uint32) * m2.dat[j * m2.n + k].(uint32) // I think, needs testing
				}
			}
		}
	case FLOAT32:
		for j := 0; j < m2.n; j++ { // Columns of m2 and m3
			for i := 0; i < m1.m; i++ { // Rows of m1 and m3
				for k := 0; k < m1.n; k++ { // Columns of m1, rows of m2
					dat[j * m2.n + i] = dat[j * m2.n + i].(float32) + m1.dat[k * m1.n + i].(float32) * m2.dat[j * m2.n + k].(float32) // I think, needs testing
				}
			}
		}
	case FLOAT64:
		for j := 0; j < m2.n; j++ { // Columns of m2 and m3
			for i := 0; i < m1.m; i++ { // Rows of m1 and m3
				for k := 0; k < m1.n; k++ { // Columns of m1, rows of m2
					dat[j * m2.n + i] = dat[j * m2.n + i].(float64) + m1.dat[k * m1.n + i].(float64) * m2.dat[j * m2.n + k].(float64) // I think, needs testing
				}
			}
		}
	}*/

	mat, err := unsafeMatrixFromSlice(m1.typ, dat, m1.m, m2.n)
	if err != nil {
		return
	}

	return *mat
}

func BatchMultiply(args []Matrix) Matrix {
	if len(args) == 1 {
		return args[0]
	}
	
	var m1,m2 Matrix
	if len(args) > 2 {
			ch1 := make(chan Matrix)
			ch2 := make(chan Matrix)
			
			// Split up the work, matrix mult is associative so this will work
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

// Wrapper so we can use multiply concurrently. Code duplication might be faster (if concurrency is faster here at all, that is). We'll need benchmarks to be sure
func batchMultHelper(ch chan<- Matrix, args[]Matrix) { 
	ch <- BatchMultiply(args)
	close(ch)
}

// INCOMPLETE DO NOT USE
// Need better function to make matrices for recursion
func (m1 Matrix) Det() interface{} {
	if m1.m != m1.n { // Determinants are only for square matrices
		return nil
	}

	return nil
}
