package sieve

import "sort"

// Sieve calculates prime
func Sieve(limit int) (out []int) {
	out = make([]int, 0)
	myMap := make(map[int]bool, 0)

	if limit <= 1 {
		return
	}

	// Populate all the numbers
	for i := 2; i <= limit; i++ {
		myMap[i] = true
	}

	// Check the numbers and false all the multiples
	for i := 2; i <= limit; i++ {
		for j := 2; (j * i) <= limit; j++ {
			if myMap[(i * j)] {
				myMap[(i * j)] = false
			}
		}
	}

	// Append only the positive found
	for k := range myMap {
		if myMap[k] {
			out = append(out, k)
		}
	}

	sort.Ints(out)

	return
}
