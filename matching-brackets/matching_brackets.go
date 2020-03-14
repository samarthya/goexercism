package brackets

import (
	"strings"
)

const bracksOpen string = "{(["
const bracksClosed string = "})]"

var myMap = map[rune]rune{
	']': '[',
	'}': '{',
	')': '(',
}

// Replace all other characters
func replaceOthers(s string) string {
	var sb strings.Builder
	fn := func(r rune) rune {
		switch {
		case strings.ContainsRune(bracksOpen, r):
			return r
		case strings.ContainsRune(bracksClosed, r):
			return r
		default:
			return ' '
		}
	}
	sb.WriteString(strings.Map(fn, s))
	return sb.String()
}

//Bracket Match the content
func Bracket(s string) bool {
	openCount := []rune{}

	// Skip other elements
	s = replaceOthers(s)

	// Replace all the spaces
	s = strings.ReplaceAll(s, " ", "")

	// Break the string.
	runes := []rune(s)
	l := len(runes)

	if (l % 2) != 0 {
		return false
	}

	matchCount := 0
	for _, val := range s {
		switch {
		case strings.ContainsRune(bracksOpen, val):
			if len(openCount) == matchCount {
				openCount = append(openCount, val)
			} else {
				openCount[matchCount] = val
			}
			matchCount++

		case strings.ContainsRune(bracksClosed, val):
			if matchCount == 0 || myMap[val] != openCount[matchCount-1] {
				return false
			}
			matchCount--
		}
	}

	if matchCount == 0 {
		return true
	}

	return false
}
