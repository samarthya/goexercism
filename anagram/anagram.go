package anagram

import (
	"strings"
)

func mapOfRune(s string) (out map[rune]int) {
	out = make(map[rune]int, 0)
	s = strings.ToLower(s)
	for _, r := range s {
		out[r]++
	}
	return
}

// Detect the candidates anagrams
func Detect(s string, c []string) (result []string) {
	result = make([]string, 0)

	if s == "" {
		return
	}

	oMap := mapOfRune(s)
	for _, v := range c {
		if len(v) == len(s) && !strings.Contains(strings.ToLower(s), strings.ToLower(v)) {
			// val contains all the characters of s only then
			oV := mapOfRune(strings.ToLower(v))
			found := true
			for key, val := range oV {
				if oMap[key] == val {
					continue
				}
				found = false
				break
			}

			if found {
				result = append(result, v)
			}

		}
	}
	return
}
