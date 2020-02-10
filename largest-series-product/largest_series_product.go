package lsproduct

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// productCompute Computes the max product
func productCompute(series string, max int) (int, error) {
	in := strings.Split(series, "")

	result := 0
	out := 1
	end := len(in)

	for i := range in {
		skip := i
		if (skip + max) <= end {
			for {
				o, err := strconv.Atoi(in[skip])
				if err != nil {
					return 0, errors.New("digits input must only contain digits")
				}
				fmt.Println(" Mul: ", o)
				out *= o
				skip++
				if (skip >= end) || (skip > ((max - 1) + i)) {
					break
				}
			}
			fmt.Println(" Result: ", out)
			if out > result {
				result = out
			}
			out = 1
		}
	}

	return result, nil
}

// LargestSeriesProduct Returns the {max} digit product in the {series} of input number
func LargestSeriesProduct(series string, max int) (int, error) {
	if series == "" && len(series) < max || (len(series) < max) {
		return 0, errors.New("span must be smaller than string length")
	}

	if max < 0 {
		return 0, errors.New("span must be greater than zero")
	}

	if max == 0 {
		return 1, nil
	}

	return productCompute(series, max)
}
