package hamming

import (
	"errors"
)

// Distance Computes the hamming distance between two DNA's
func Distance(a, b string) (diff int, err error) {
	// If both are empty no difference
	if a == "" && b == "" {
		return 0, nil
	}

	// Inequal length implies error.
	if len(a) != len(b) {
		return 0, errors.New(" empty length not allowed")
	}

	diff = 0

	aR := []rune(a)
	bR := []rune(b)

	for i, val := range aR {
		if val != bR[i] {
			diff++
		}
	}

	return diff, nil

}
