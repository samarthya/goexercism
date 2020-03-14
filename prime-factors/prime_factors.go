package prime

// Factors factors calculation
func Factors(i int64) []int64 {
	m := make([]int64, 0)

	var r int64 = 2
	for {
		if i%r == 0 {
			i /= r
			m = append(m, r)
		} else {
			r++
		}

		if i < r {
			break
		}
	}
	return m
}
