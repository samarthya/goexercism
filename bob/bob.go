// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	. "unicode"
)

// isYellingQuestion: 'Calm down, I know what I'm doing!' 
func isYellingQuestion(remark string) bool {
	if isQuestion(remark) && isYelling(remark) {
		return true
	}
	return false
}
// isYelling : He answers 'Whoa, chill out!' if you YELL AT HIM (in all capitals).
func isYelling(remark string) bool {
	// Get the rune slice
	chars := [] rune(remark)
	noneFound := true
	for _, val := range chars {

		if IsLetter(val) {
			if IsUpper(val) {
				noneFound = false
				continue
			} else {
				noneFound = true
				break
			}
		} 
	}

	return !noneFound
}

//isQuestion Bob answers 'Sure.' if you ask him a question, such as "How are you?".
func isQuestion(remark string) bool {
	return strings.HasSuffix(remark,"?")
}
// Hey should have a comment documenting it.
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)

	if isYellingQuestion(remark) {
		return "Calm down, I know what I'm doing!"
	} else if (isQuestion(remark)) {
		return "Sure."
	}

	if isYelling(remark) {
		return "Whoa, chill out!"
	}

	if remark == "" || len(strings.TrimSpace(remark)) == 0 {
		return "Fine. Be that way!"
	}

	return "Whatever."
}
