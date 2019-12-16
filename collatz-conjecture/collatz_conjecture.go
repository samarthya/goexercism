package collatzconjecture

import (
	"errors"
	"fmt"
)

// CollatzConjecture Start the seed probelm 3x + 1 problem.
func CollatzConjecture(seed int) (int, error) {
	var result int = seed
	var seq int = 0

	if seed <= 0 {
		return seed, errors.New(" negative numbers not allowed")
	}

	for result != 1 {
		if result%2 == 0 {
			result = (result / 2)

		} else {
			result = ((result * 3) + 1)
		}
		seq++
	}

	fmt.Println(seq, ". ", seed)
	return seq, nil
}
