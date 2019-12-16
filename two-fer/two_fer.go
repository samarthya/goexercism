// Package twofer that shall contain only ShareWith function.
package twofer

import (
	"fmt"
)

// ShareWith utilises the 'name' parameter to return the greetings.
func ShareWith(name string) string {
	// If no 'name' is passed return the default.
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
