package armstrong

import (
	"strconv"
)

// IsNumber - Checks for the armstrong number
func IsNumber(i int) bool {
	in := strconv.Itoa(i)
	l := len(in)

	if l <= 1 {
		return true
	}

	sum := 0

	for _, v := range in {
		pow := 1

		for p := l; p > 0; p-- {
			pow *= int(v - '0')
		}
		sum += pow
	}

	if (sum - i) == 0 {
		return true
	}
	return false
}
