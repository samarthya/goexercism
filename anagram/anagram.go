package anagram

import (
	"strings"
)

func sumOfString(s string) (out int) {
	s = strings.ToLower(s)
	for _, r := range s {
		out += int(r)
	}
	return
}

// Detect the candidates anagrams
func Detect(s string, c []string) (result []string) {
	result = make([]string, 0)

	if s == "" {
		return
	}

	oV := sumOfString(s)
	for _, v := range c {
		if len(v) == len(s) {
			// val contains all the characters of s only then
			if oV == sumOfString(strings.ToLower(v)) {
				result = append(result, v)
			}
		}
	}
	return
}
