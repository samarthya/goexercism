package pascal

func fac(n int) int {
	if n <= 0 {
		return 1
	} else {
		return n * fac(n-1)
	}
}

func nCr(n, r int) int {
	return (fac(n) / (fac(n-r) * fac(r)))
}

func pascalsRow(i int) (output []int) {
	for j := 0; j <= i; j++ {
		output = append(output, nCr(i, j))
	}
	return output
}
func Triangle(n int) [][]int {
	// Array of arrays
	var output [][]int
	output = make([][]int, n)
	for i := 0; i < n; i++ {
		output[i] = make([]int, i)
		output[i] = pascalsRow(i)
	}
	return output
}
