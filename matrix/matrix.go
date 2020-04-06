package matrix

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

/**
Not incorrect, but it's worth remembering that a [][]int already is a map[int][]int! 
In other words, you can think of a [][]int as a data structure that takes an integer 
index representing a row number, and gives you the slice of int representing the row. 
So the struct type is unnecessary and simply wraps an existing built-in type that Go already provides.
*/

// Matrix rep of Matrix
type Matrix struct {
	rows map[int][]int
	cols map[int][]int
}

//Set rep set
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

// Rows function Rows
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

// Cols the columns
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

// New Allows you to contruct
func New(input string) (*Matrix, error) {
	fmt.Println("\n", input)
	matrix := &Matrix{rows: make(map[int][]int, 0), cols: make(map[int][]int, 0)}

	if input == "" || len(input) == 0 {
		return nil, errors.New(" invalid input")
	}

	for j, line := range strings.Split(input, "\n") {
		line = strings.TrimSpace(line)
		for i, row := range strings.Split(line, " ") {
			// Time to break the space separated column values.
			val, err := strconv.Atoi(strings.TrimSpace(row))
			if err != nil {
				return nil, errors.New(" cannot convert")
			}

			matrix.cols[i] = append(matrix.cols[i], val)
			matrix.rows[j] = append(matrix.rows[j], val)
		}
		if (j > 0) && len(matrix.rows[j-1]) != len(matrix.rows[j]) {
			return nil, errors.New(" inequal rows length")
		}
	}

	return matrix, nil
}
