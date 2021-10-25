package logs

import (
	"strings"
)

var AppID = "❗🔍☀"

var Applications = map[rune]string{
	'❗': "recommendation",
	'🔍': "search",
	'☀': "weather",
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, v := range AppID {
		if strings.IndexRune(log, v) != -1 {
			return Applications[v]
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var n []rune
	for _, v := range log {
		if v == oldRune {
			n = append(n, newRune)
		} else {
			n = append(n, v)
		}
	}
	return string(n)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	for _, v := range AppID {
		index := strings.IndexRune(log, v)
		if index != -1 && index < limit {
			return true
		}
	}

	return false
}
