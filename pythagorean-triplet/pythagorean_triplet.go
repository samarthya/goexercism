package pythagorean

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

	min := 1
	max := int(p / 2)
	triples = make([]Triplet, 0)
	allElements := Range(min, max)

	for _, v := range allElements {
		if v[0]+v[1]+v[2] == p {
			triples = append(triples, v)
		}
	}
	return triples
}
