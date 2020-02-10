package cryptosquare

import (
	"fmt"
	"strings"
	"unicode"
)

// Normalize Normalizes the input
func Normalize(in string) string {
	var b strings.Builder

	for _, r := range in {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			fmt.Fprintf(&b, "%s", string(unicode.ToLower(r)))
		}
	}

	return b.String()
}

// getRect Calculates the rxc
func getRect(in string) (r, c int) {
	strLen := len(in)

	r = 1
	c = r + 1

	for {
		if (r * c) >= strLen {
			break
		}
		if r < c {
			r++
		} else {
			c++
		}
	}
	return
}

// Encode encodes the input string.
func Encode(in string) string {
	var r, c int
	var b strings.Builder

	in = Normalize(strings.TrimSpace(in))
	strLen := len(in)

	if strLen <= 1 {
		return in
	}

	r, c = getRect(in)

	// Append spaces to create a perfect Rectangle
	for strLen != r*c {
		in += " "
		strLen = len(in)
	}

	runes := []rune(in)

	step := 0
	for i := 0; i < c; i++ {
		step = i
		for step < strLen {
			fmt.Fprintf(&b, "%s", string(runes[step]))
			step += c
		}
		if i+1 < c {
			fmt.Fprintf(&b, "%s", " ")
		}
	}

	return b.String()
}
