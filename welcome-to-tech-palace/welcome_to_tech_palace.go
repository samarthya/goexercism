package techpalace

import (
	"fmt"
	"log"
	"strings"
)

const (
	Welcome       = "Welcome to the Tech Palace, %s"
	StarryWelcome = "Welcome!"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf(Welcome, strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var b strings.Builder
	b.WriteString(strings.Repeat("*", numStarsPerLine) + "\n")
	b.WriteString(welcomeMsg + "\n")
	b.WriteString(strings.Repeat("*", numStarsPerLine))
	return b.String()
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	var b strings.Builder

	for _, v := range oldMsg {
		switch {
		case v != '*':
			b.WriteRune(v)
		}
	}

	str := strings.Trim(b.String(), "\n")
	log.Println(strings.TrimSpace(str))
	return strings.TrimSpace(str)
}
