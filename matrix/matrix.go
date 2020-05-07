package matrix

import (
	"errors"
	"strconv"
	"strings"
)

/**
Not incorrect, but it's worth remembering that a [][]int already is a map[int][]int!
In other words, you can think of a [][]int as a data structure that takes an integer
index representing a row number, and gives you the slice of int representing the row.
So the struct type is unnecessary and simply wraps an existing built-in type that Go already provides.
*/

// Matrix Type matrix
type Matrix [][]int

// New function that creates
func New(input string) (*Matrix, error) {
	lines := strings.Split(input, "\n")
	lineLength := len(strings.Fields(lines[0]))

	matrix := make(Matrix, len(lines))

	for rowIndex, line := range lines {
		values := strings.Fields(line)
		matrix[rowIndex] = make([]int, len(values))

		if len(values) != lineLength {
			return nil, errors.New("must be of same length")
		}

		for columnIndex, value := range values {
			intValue, err := strconv.Atoi(value)

			if err != nil {
				return nil, errors.New("creation failed")
			}

			matrix.Set(rowIndex, columnIndex, intValue)
		}
	}

	return &matrix, nil
}

// Rows Returns a copy of the matrix rows.
func (matrix Matrix) Rows() [][]int {
	c := make(Matrix, len(matrix))

	for rowIndex := range c {
		c[rowIndex] = make([]int, len(matrix[0]))
	}

	for rowIndex, row := range matrix {
		for columnIndex, element := range row {
			c[rowIndex][columnIndex] = element
		}
	}

	return c
}

// Cols Returns a copy of the matrix columns.
func (matrix Matrix) Cols() [][]int {
	c := make(Matrix, len(matrix[0]))

	for rowIndex := range c {
		c[rowIndex] = make([]int, len(matrix))
	}

	for rowIndex, row := range matrix {
		for columnIndex, element := range row {
			c[columnIndex][rowIndex] = element
		}
	}

	return c
}

// Set Sets a (row, column)
func (matrix *Matrix) Set(row, column, value int) bool {
	newMatrix := *matrix

	if row >= 0 && row < len(newMatrix) && column >= 0 && column < len(newMatrix[0]) {
		newMatrix[row][column] = value
	} else {
		return false
	}
	return true
}
