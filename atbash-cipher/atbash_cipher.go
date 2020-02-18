package atbash

import (
	"regexp"
	"strings"
	"unicode"
)

//Atbash Crypt function
func Atbash(s string) (out string) {
	var replaceWhiteSpace = regexp.MustCompile(`\W`)
	// var addChunk = regexp.MustCompile(`\w{5}`)

	rot13 := func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			switch {
			case r >= 'a' && r <= 'z':
				return ('z' - r) + 'a'
			default:
				return r
			}
		}
		return ' '
	}

	// s = sanitizeString(s)
	out = strings.Map(rot13, strings.ToLower(s))
	out = replaceWhiteSpace.ReplaceAllString(out, "")
	encodedString := ""
	sc := 0
	for _, val := range out {
		encodedString += string(val)
		if sc > 0 && sc%4 == 0 {
			encodedString += " "
			sc = 0
			continue
		}
		sc++
	}
	out = strings.TrimSpace(encodedString)
	return
}
