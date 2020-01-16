package matrix

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//
// 1. You must implement a constructor and methods Rows() and Cols() as
//    described in the README, but Rows() and Cols must return results that
//    are independent from the original matrix.  That is, you should be able
//    to do as you please with the results without affecting the matrix.
//
// 2. You must implement a method Set(row, col, val) for setting a matrix
//    element.
//
// 3. As usual in Go, you must detect and return error conditions.

// Interfaces are implemented implicitly.
type IMATRIX interface {
	Rows() []int
	Cols() []int
}

type ArrayOfInts []int

type Matrix struct {
	rows map[int]ArrayOfInts
	cols map[int]ArrayOfInts
}

func (m *Matrix) Set(row, col, val int) bool {
	result := m

	cols := result.rows[row]
	rows := result.cols[col]
	if cols != nil && rows != nil {
		cols[col] = val
		rows[row] = val
	} else {
		return false
	}

	return true
}

func (m *Matrix) Rows() (result [][]int) {

	result = make([][]int, len(m.rows))

	for j := range m.rows {
		result[j] = make([]int, len(m.cols))
		for i := range m.rows[j] {
			result[j][i] = m.rows[j][i]
		}
		// result[key] = m.rows[key]
	}

	return result
}

func (m *Matrix) Cols() (result [][]int) {

	result = make([][]int, len(m.cols))

	for i := range m.cols {
		result[i] = make([]int, len(m.rows))
		fmt.Println(" COLS:", m.cols[i])
		for j, val := range m.cols[i] {
			fmt.Println(" COLS[]:", val)
			result[i][j] = val
		}
		//result[key] = m.cols[key]
	}

	return result
}

// Constructor
func New(input string) (*Matrix, error) {
	fmt.Println("\n", input)
	matrix := &Matrix{rows: make(map[int]ArrayOfInts, 0), cols: make(map[int]ArrayOfInts, 0)}

	if input == "" || len(input) == 0 {
		return nil, errors.New(" invalid input")
	}

	for j, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		for i, row := range strings.Split(line, " ") {
			// Time to break the space separated column values.
			val, err := strconv.Atoi(strings.TrimSpace(row))
			if err != nil {
				return nil, errors.New(" cannot convert.")
			}

			matrix.cols[i] = append(matrix.cols[i], val)
			matrix.rows[j] = append(matrix.rows[j], val)
		}
		if (j > 0) && len(matrix.rows[j-1]) != len(matrix.rows[j]) {
			return nil, errors.New(" inequal rows length.")
		}
	}

	return matrix, nil
}
