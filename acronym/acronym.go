// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
	"unicode"
	"fmt"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	s = strings.ReplaceAll(s,"-", " ")
	s = strings.ReplaceAll(s,"_", " ")

	elements := strings.Split(strings.TrimSpace(s)," ")
	var acronym [] rune
	fmt.Println(s)
	for _, value := range elements {
		value = strings.TrimSpace(value)

		if value != "" {
			acronym = append(acronym, unicode.ToUpper(rune(value[0])))
		}
	}
	return string(acronym)
}
