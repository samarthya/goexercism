package isogram

import (
	"strings"
	"unicode"
)

/**
 * Checks whether an input strig is an Isogram or not.
 */
func IsIsogram(input string) bool {
	if input == "" {
		return true
	}

	input = strings.ToLower(input)
	for i, el := range input {
		if strings.LastIndexByte(strings.ToLower(input), input[i]) != i {
			if unicode.IsLetter(el) != false {
				return false
			}
		}
	}
	return true
}
