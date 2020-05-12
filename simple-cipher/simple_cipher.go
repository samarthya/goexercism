package cipher

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

// Vigenere Stores the key
type Vigenere struct {
	key string
}

// RegularExp Elements to find
var RegularExp = regexp.MustCompile(`[^A-Za-z]`)

// NewCaesar Object that implements the interface Cipher
func NewCaesar() Cipher {
	return Vigenere{key: "d"}
}

// NewShift Object that implements the interface Cipher
func NewShift(distance int) Cipher {
	// Corner case when of 0, 26, -26
	if distance <= -26 || distance >= 26 {
		return nil
	}
	return NewVigenere(string(mapChar('a', int32(distance))))
}

// NewVigenere Object that implements the interface Cipher
func NewVigenere(k string) Cipher {

	if regexp.MustCompile(`[^a-z]`).MatchString(k) || regexp.MustCompile(`^a*$`).MatchString(k) {
		return nil
	}

	return Vigenere{key: k}
}

func mapChar(s rune, r int32) (e rune) {
	switch {
	case s < 'a'-r:
		e = s + 26 + r
	case s > 'z'-r:
		e = (s - 26 + r)
	default:
		e = s + r
	}
	return e
}

// stripString replace all the other elements
func stripString(s string) (o string) {
	o = RegularExp.ReplaceAllString(s, "")
	log.Printf(" String: %s", o)
	o = strings.ToLower(o)
	return
}

// Encode Implements the interface Cipher
func (e Vigenere) Encode(s string) (o string) {

	var sb strings.Builder

	o = ""
	s = stripString(s)

	if s == "" || len(s) <= 0 {
		return
	}

	log.Printf(" Encode: %s", s)
	// Iterate throught the runes

	for i, m := range s {
		newChar := mapChar(m, int32(e.key[i%len(e.key)]-'a'))
		fmt.Fprintf(&sb, "%c", newChar)
	}

	o = sb.String()
	log.Printf(" Encoded: %s", o)
	return
}

// Decode decides the encrypted string
func (e Vigenere) Decode(s string) (o string) {
	var sb strings.Builder

	o = ""

	log.Printf(" Decode: %s", s)

	if s == "" || len(s) <= 0 {
		return
	}

	for i, m := range s {
		newChar := mapChar(m, -int32(e.key[i%len(e.key)]-'a'))
		fmt.Fprintf(&sb, "%c", newChar)
	}

	o = sb.String()
	log.Printf(" Decoded: %s", o)

	return
}
