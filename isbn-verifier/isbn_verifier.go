package isbn

import "strings"

// IsValidISBN Validates the ISBN
func IsValidISBN(s string) bool {

	// Remove spaces
	s = strings.TrimSpace(s)
	// Remove all `-`
	s = strings.ReplaceAll(s, "-", "")

	if len(s) < 10 || len(s) > 10 {
		return false
	}

	place := 10
	sum := 0
	for i, v := range s {

		if place == 0 {
			return false
		}
		switch {
		case v == 'X':
			if i != (len(s) - 1) {
				return false
			}
			sum += (place * 10)
			place--
		case v >= '0' && v <= '9':
			sum += (place * int(v-'0'))
			place--
		default:
			return false
		}
		continue
	}

	if sum%11 == 0 {
		return true
	}
	return false
}
