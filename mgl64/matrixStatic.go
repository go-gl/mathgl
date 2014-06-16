package mgl64

// Sets a Column within the Matrix, so it mutates the calling matrix.
func (m *Mat2) SetCol(col int, v Vec2) {
	m[col*2+0], m[col*2+1] = v[0], v[1]
}

// Sets a Column within the Matrix, so it mutates the calling matrix.
func (m *Mat2x3) SetCol(col int, v Vec2) {
	m[col*2+0], m[col*2+1] = v[0], v[1]
}

// Sets a Column within the Matrix, so it mutates the calling matrix.
func (m *Mat2x4) SetCol(col int, v Vec2) {
	m[col*2+0], m[col*2+1] = v[0], v[1]
}

// Sets a Column within the Matrix, so it mutates the calling matrix.
func (m *Mat3x2) SetCol(col int, v Vec3) {
	m[col*3+0], m[col*3+1], m[col*3+2] = v[0], v[1], v[2]
}

// Sets a Column within the Matrix, so it mutates the calling matrix.
func (m *Mat3) SetCol(col int, v Vec3) {
	m[col*3+0], m[col*3+1], m[col*3+2] = v[0], v[1], v[2]
}

// Sets a Column within the Matrix, so it mutates the calling matrix.
func (m *Mat3x4) SetCol(col int, v Vec3) {
	m[col*3+0], m[col*3+1], m[col*3+2] = v[0], v[1], v[2]
}

// Sets a Column within the Matrix, so it mutates the calling matrix.
func (m *Mat4x2) SetCol(col int, v Vec4) {
	m[col*4+0], m[col*4+1], m[col*4+2], m[col*4+3] = v[0], v[1], v[2], v[3]
}

// Sets a Column within the Matrix, so it mutates the calling matrix.
func (m *Mat4x3) SetCol(col int, v Vec4) {
	m[col*4+0], m[col*4+1], m[col*4+2], m[col*4+3] = v[0], v[1], v[2], v[3]
}

// Sets a Column within the Matrix, so it mutates the calling matrix.
func (m *Mat4) SetCol(col int, v Vec4) {
	m[col*4+0], m[col*4+1], m[col*4+2], m[col*4+3] = v[0], v[1], v[2], v[3]
}

// Sets a Row within the Matrix, so it mutates the calling matrix.
func (m *Mat2) SetRow(row int, v Vec2) {
	m[row+0], m[row+2] = v[0], v[1]
}

// Sets a Row within the Matrix, so it mutates the calling matrix.
func (m *Mat2x3) SetRow(row int, v Vec3) {
	m[row+0], m[row+2], m[row+4] = v[0], v[1], v[2]
}

// Sets a Row within the Matrix, so it mutates the calling matrix.
func (m *Mat2x4) SetRow(row int, v Vec4) {
	m[row+0], m[row+2], m[row+4], m[row+6] = v[0], v[1], v[2], v[3]
}

// Sets a Row within the Matrix, so it mutates the calling matrix.
func (m *Mat3x2) SetRow(row int, v Vec2) {
	m[row+0], m[row+3] = v[0], v[1]
}

// Sets a Row within the Matrix, so it mutates the calling matrix.
func (m *Mat3) SetRow(row int, v Vec3) {
	m[row+0], m[row+3], m[row+6] = v[0], v[1], v[2]
}

// Sets a Row within the Matrix, so it mutates the calling matrix.
func (m *Mat3x4) SetRow(row int, v Vec4) {
	m[row+0], m[row+3], m[row+6], m[row+9] = v[0], v[1], v[2], v[3]
}

// Sets a Row within the Matrix, so it mutates the calling matrix.
func (m *Mat4x2) SetRow(row int, v Vec2) {
	m[row+0], m[row+4] = v[0], v[1]
}

// Sets a Row within the Matrix, so it mutates the calling matrix.
func (m *Mat4x3) SetRow(row int, v Vec3) {
	m[row+0], m[row+4], m[row+8] = v[0], v[1], v[2]
}

// Sets a Row within the Matrix, so it mutates the calling matrix.
func (m *Mat4) SetRow(row int, v Vec4) {
	m[row+0], m[row+4], m[row+8], m[row+12] = v[0], v[1], v[2], v[3]
}

// Trace is a basic operation on a square matrix that simply
// sums up all elements on the main diagonal (meaning all elements such that row==col).
func (m Mat2) Diag() Vec2 {
	return Vec2{m[0], m[3]}
}

// Trace is a basic operation on a square matrix that simply
// sums up all elements on the main diagonal (meaning all elements such that row==col).
func (m Mat3) Diag() Vec3 {
	return Vec3{m[0], m[4], m[8]}
}

// Trace is a basic operation on a square matrix that simply
// sums up all elements on the main diagonal (meaning all elements such that row==col).
func (m Mat4) Diag() Vec4 {
	return Vec4{m[0], m[5], m[10], m[15]}
}

/*
func (m Mat2) Mat2x3() Mat2x3 {
	col0, col1 := m.Cols()
	return Mat2x3FromCols(
		col0,
		col1,
		Vec2{0, 0},
	)
}

func (m Mat2) Mat2x4() Mat2x4 {
	col0, col1 := m.Cols()
	return Mat2x4FromCols(
		col0,
		col1,
		Vec2{0, 0},
		Vec2{0, 0},
	)
}

func (m Mat2) Mat3x2() Mat3x2 {
	col0, col1 := m.Cols()
	return Mat3x2FromCols(
		col0.Vec3(0),
		col1.Vec3(0),
	)
}

func (m Mat2) Mat3x4() Mat3x4 {
	col0, col1 := m.Cols()
	return Mat3x4FromCols(
		col0.Vec3(0),
		col1.Vec3(0),
		Vec3{0, 0, 1},
		Vec3{0, 0, 0},
	)
}

func (m Mat2) Mat4x2() Mat4x2 {
	col0, col1 := m.Cols()
	return Mat4x2FromCols(
		col0.Vec4(0, 0),
		col1.Vec4(0, 0),
	)
}

func (m Mat2) Mat4x3() Mat4x3 {
	col0, col1 := m.Cols()
	return Mat4x3FromCols(
		col0.Vec4(0, 0),
		col1.Vec4(0, 0),
		Vec4{0, 0, 1, 0},
	)
}

// this becomes way too much code to do manually
*/

// only most important conversions

func (m Mat2) Mat3() Mat3 {
	col0, col1 := m.Cols()
	return Mat3FromCols(
		col0.Vec3(0),
		col1.Vec3(0),
		Vec3{0, 0, 1},
	)
}

func (m Mat2) Mat4() Mat4 {
	col0, col1 := m.Cols()
	return Mat4FromCols(
		col0.Vec4(0, 0),
		col1.Vec4(0, 0),
		Vec4{0, 0, 1, 0},
		Vec4{0, 0, 0, 1},
	)
}

func (m Mat3) Mat2() Mat2 {
	col0, col1, _ := m.Cols()
	return Mat2FromCols(
		col0.Vec2(),
		col1.Vec2(),
	)
}

func (m Mat3) Mat4() Mat4 {
	col0, col1, col2 := m.Cols()
	return Mat4FromCols(
		col0.Vec4(0),
		col1.Vec4(0),
		col2.Vec4(0),
		Vec4{0, 0, 0, 1},
	)
}

func (m Mat4) Mat2() Mat2 {
	col0, col1, _, _ := m.Cols()
	return Mat2FromCols(
		col0.Vec2(),
		col1.Vec2(),
	)
}

func (m Mat4) Mat3() Mat3 {
	col0, col1, col2, _ := m.Cols()
	return Mat3FromCols(
		col0.Vec3(),
		col1.Vec3(),
		col2.Vec3(),
	)
}
