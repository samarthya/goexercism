package lsproduct

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

// productCompute Computes the max product
func productCompute(series string, max int) (int, error) {
	in := strings.Split(series, "")

	sop := make([]int, 1)
	out := 1
	end := len(in)

	sop[0] = 0

	for i := range in {

		if (i + max) > end {
			break
		}

		newSeries := in[i : max+i]

		for i := range newSeries {
			o, err := strconv.Atoi(newSeries[i])
			if err != nil {
				return 0, errors.New("digits input must only contain digits")
			}
			out *= o
		}

		sop = append(sop, out)
		out = 1
	}

	sort.Slice(sop, func(i, j int) bool {
		return sop[i] > sop[j]
	})

	return sop[0], nil
}

// LargestSeriesProduct Returns the {max} digit product in the {series} of input number
func LargestSeriesProduct(series string, max int) (int, error) {
	if len(series) < max {
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
