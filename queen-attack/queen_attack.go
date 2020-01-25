package queenattack

import (
	"errors"
	"strings"
)

// CONSTANTS constants to find the index.
const CONSTANTS string = "abcdefgh"

// Columns Index of columns
const Columns string = "12345678"

func indexOfSquare(a, q string) int {
	return strings.Index(a, q)
}

func squareWidth(qb, qa string) int {
	return indexOfSquare(CONSTANTS, qb) - indexOfSquare(CONSTANTS, qa)
}

// CanQueenAttack Validates if qa can attack qb
func CanQueenAttack(qa string, qb string) (bool, error) {

	qa = strings.ToLower(qa)
	qb = strings.ToLower(qb)

	if len(qa) != len(qb) || strings.Compare(qa, qb) == 0 || len(qa) > 2 || len(qb) > 2 {
		return false, errors.New(" invalid position")
	}

	rowa := qa[0:1]
	rowb := qb[0:1]

	i := indexOfSquare(Columns, qa[1:2])
	j := indexOfSquare(Columns, qb[1:2])

	if i < 0 || j < 0 || indexOfSquare(CONSTANTS, rowa) < 0 || indexOfSquare(CONSTANTS, rowb) < 0 {
		return false, errors.New(" board is just 8 squares")
	}

	// Same column or row
	if i == j || strings.Compare(rowa, rowb) == 0 {
		return true, nil
	}

	if (squareWidth(rowb, rowa) == (j - i)) || (squareWidth(rowb, rowa) == (i - j)) {
		return true, nil
	}

	return false, nil
}
