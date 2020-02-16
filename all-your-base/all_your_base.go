package allyourbase

import (
	"errors"
)

func numberToBase(num int, b int) (out []int) {
	out = make([]int, 0)
	for num > 0 {
		v := (num % b)
		out = append([]int{v}, out...)
		num /= b
	}
	return
}

// ConvertToBase Conversion function
func ConvertToBase(b int, d []int, ob int) ([]int, error) {
	num := 0
	power := 1

	if b < 2 {
		return nil, errors.New("input base must be >= 2")
	}

	if ob < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	ln := (len(d) - 1)

	for i := range d {
		digit := d[ln-i]
		if digit < b && digit >= 0 {
			num += (digit * power)
			power *= b
		} else {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
	}

	if len(d) == 0 || num == 0 {
		return []int{0}, nil
	}

	return numberToBase(num, ob), nil
}
