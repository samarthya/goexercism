package series

import (
	"strings"
)

// All returns the string array.
func All(n int, s string) (output []string) {
	output = make([] string, 0)

	var sb strings.Builder
	var lenOfS int = len(s)
	chars := []rune(s)
	i := 0
	
	if (lenOfS >= 0) && (n <= lenOfS) {
		for  {
			if (i+n) <= lenOfS {
				sb.WriteString(string(chars[i: i+n]))
				output = append(output, sb.String())
				sb.Reset()
				i++
			} else {
				break
			}
		}
	}
	return output
}

// UnsafeFirst First of the all.
func UnsafeFirst(n int, s string) string {
	return All(n, s)[0]
}
