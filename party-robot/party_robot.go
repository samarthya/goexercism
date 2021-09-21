package partyrobot

import (
	"fmt"
	"log"
	"strings"
)

const (
	Message = "Welcome to my party, %s!"
	Happy   = "Happy birthday %s! You are now %d years old!"
	Table   = "\nYou have been assigned to table %X. Your table is %s, exactly %.1F meters from here.\n"
	Next    = "You will be sitting next to %s."
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf(Message, name)
}

// HappyBirthday wishes happy birthday to the birthday person and stands out their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf(Happy, name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbour, direction string, distance float64) string {
	var b strings.Builder

	fmt.Fprintf(&b, Message, name)
	fmt.Fprintf(&b, Table, table, direction, distance)
	fmt.Fprintf(&b, Next, neighbour)
	log.Println(b.String())
	return b.String()
}
