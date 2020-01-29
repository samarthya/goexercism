package pythagorean

import "fmt"

// Triplet Pythagorean triplet
type Triplet [3]int

func sq(x int) int {
	return x * x
}

// Range range.
func Range(min, max int) (triples []Triplet) {
	triples = make([]Triplet, 0)
	for i := min; i <= max; i++ {
		for j := i + 1; j <= max; j++ {
			k := j + 1
			for k <= max {
				if (sq(i) + sq(j)) == sq(k) {
					triples = append(triples, Triplet{i, j, k})
				}
				k++
			}
		}
	}
	return triples
}

// Sum sum.
func Sum(p int) (triples []Triplet) {
	triples = make([]Triplet, 0)

	min := 1
	max := int(p / 2)

	fmt.Println(min, max)
	for i := min; i <= max; i++ {
		for j := i + 1; j <= max; j++ {
			k := j + 1
			for k <= max {
				if ((sq(i) + sq(j)) == sq(k)) && ((i + j + k) == p) {
					triples = append(triples, Triplet{i, j, k})
				}
				k++
			}
		}
	}
	return triples
}
