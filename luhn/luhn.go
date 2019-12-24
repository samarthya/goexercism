package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

const (
	ONLY_NUMERIC = "0123456789"
)

/**
 * First Benchmark
 * BenchmarkValid-16        2339478               506 ns/op             248 B/op          5 allocs/op
 * PASS
 * ok      luhn    2.914s
 */
func Valid(number string) bool {

	// Remove any trailing spaces.
	number = strings.TrimSpace(number)

	if number == "" || len(number) <= 1 {
		return false
	}

	var mk []int

	for _, val := range number {
		if strings.IndexRune(ONLY_NUMERIC, val) != -1 {
			number, err := strconv.Atoi(string(val))
			if err == nil {
				mk = append(mk, number)
			}
		} else if unicode.IsSpace(val) {
			continue
		} else {
			return false
		}
	}

	var sum int

	for i := (len(mk) - 1); i > 0; i -= 2 {
		if (mk[i-1] * 2) > 9 {
			mk[i-1] = ((mk[i-1] * 2) - 9)
		} else {
			mk[i-1] *= 2
		}

		sum += (mk[i-1] + mk[i])
	}

	if sum%10 == 0 {
		return true
	}

	return false
}
