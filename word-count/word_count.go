package wordcount

import (
	"strings"
	"unicode"
)

// Frequency of the words
type Frequency map[string]int

// WordCount Count the words occurence.
func WordCount(in string) Frequency {
	var myMap Frequency = make(map[string]int, 0)

	in = strings.ToLower(in)
	inStrings := strings.FieldsFunc(in, func(c rune) bool {
		return c == ' ' || c == '\n' || c == '\t' || c == ','
	})

	for _, val := range inStrings {
		val = strings.TrimFunc(val, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsDigit(r)
		})
		myMap[val]++
	}
	return myMap
}
