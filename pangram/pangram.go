package pangram

import "strings"

// Alphabets all the characters.
const Alphabets string = "abcdefghijklmnopqrstuvqxyz"

// IsPangram checks for pangram in the inputs string.
func IsPangram(in string) bool {
	in = strings.ToLower(strings.TrimSpace(in))
	if in == "" {
		return false
	}

	for _, v := range Alphabets {
		if strings.IndexRune(in, v) < 0 {
			return false
		}
	}
	return true
}
