package go_recommend_me

import (
	"log"
	"math/rand"
	"time"
)

type matrix struct {
	n int
	m int
}

func (mat *matrix) numElements() int {
	return mat.n * mat.m
}

func (mat *matrix) Rows() int{
	return mat.n
}

func (mat *matrix) Cols() int{
	return mat.m
}

func (mat *matrix) Nil() bool {
	return mat == nil
}

func (mat *matrix) checkBounds(row int, col int)  {
	if !(row < mat.n && col < mat.m) {
		log.Fatalln("Out of Bounds Error")
	}
}

// A dense matrix ie a regular matrix format
type DenseMatrix struct{
	matrix
	elements	[][]float64
}

// Initialize a dense array with random values between 0 and 1
func randomMatInit(row int, col int) *DenseMatrix{

	d := &DenseMatrix{
		matrix{row, col},
		nil,
	}

	rand.Seed(time.Now().UnixNano())
	d.elements = make([][]float64, row)
	for i := range d.elements {
		d.elements[i] = make([]float64, col)
		for j := range d.elements[i]{
			d.elements[i][j] = rand.Float64()
		}
	}

	return d

}

// The sparse matrix format will follow the CSR or compressed sparse row format

type SparseMatrix struct {
	matrix
	vals 	[]float64
	rows	[]int
	cols	[]int
}


// Create a new Sparse Matrix with i rows and j cols
func (s *SparseMatrix) createNew(i int, j int)  {
	if (i < 1 || j < 1) {
		log.Fatalln("Can not have dimensions 0 or negative")
	}

	s.n = i
	s.m = j

	s.vals = nil
	s.cols = nil
	s.rows = make([]int, s.n + 1)

}

// get certain value from the Sparse Matrix

func (s *SparseMatrix) get(row int, col int) (float64, bool)  {

	s.checkBounds(row, col)

	rowElements := s.vals[s.rows[row]:s.rows[row + 1]]
	for i := 0; i < len(rowElements); i++ {
		if col == s.cols[i] {
			return s.vals[i], true
		}
	}
	return 0, false
}

// set a certain value in the matrix

func (s *SparseMatrix) set(val float64, row int, col int) {
	s.checkBounds(row, col)

	numElements := s.rows[row + 1]
	s.vals = append(s.vals, 0)
	s.cols = append(s.cols, 0)

	copy(s.vals[numElements + 1:], s.vals[numElements:])
	copy(s.cols[numElements + 1:], s.cols[numElements:])

	s.vals[numElements] = val
	s.cols[numElements] = col

	for i := row + 1; i < len(s.rows); i++ {
		s.rows[i] += 1
	}
}




// Check that the matrix contains an element
func (s *SparseMatrix)elementExists(row int, col int) bool {
	s.checkBounds(row, col)
	_, ok := s.get(row, col)

	return ok
}



// Get the average of row i (sum of filled values / number of filled values)
func (s *SparseMatrix) avgRowVals(row int) float64 {
	s.checkBounds(row, 0)

	sum := 0.0
	cnt := 0.0

	for i := s.rows[row]; i < s.rows[row + 1]; i++ {
		cnt += 1
		sum += s.vals[i]
	}

	return sum/cnt
}

// Get the average of column j (sum of filled values / number of filled values)
func (s *SparseMatrix) avgColsVals(col int) float64 {
	s.checkBounds(0, col)

	sum := 0.0
	cnt := 0.0

	for i := 0; i < len(s.cols); i++ {
		if s.cols[i] == col {
			cnt += 1
			sum += s.vals[i]
		}
	}

	if cnt == 0 {
		return 0
	}

	return sum/cnt
}

// add a row to the matrix
func (s *SparseMatrix) addRow() {
	s.n += 1
	s.rows = append(s.rows, s.rows[len(s.rows) - 1])
}

// add a col to the matrix
func (s *SparseMatrix) addCol() {
	s.m += 1
}

// add multiple rows to the matrix
func (s *SparseMatrix) addRows(num int) {
	for i := 0; i < num; i++ {
		s.addRow()
	}

}


// Get a row from the matrix
func (s *SparseMatrix) getRow(row int) []float64 {
	return s.vals[s.rows[row]:s.rows[row + 1]]
}



