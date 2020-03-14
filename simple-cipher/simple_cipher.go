package cipher

import (
	"log"
	"strings"
)

const (
	ceasarCipher          = 3
	englishAlphabetsRange = 'z' - 'a' + 1
)

// NewCaesar Default 3 bit cipher
func NewCaesar() Cipher {
	return NewShift(ceasarCipher)
}

//NewShift newshift
func NewShift(d int) Cipher {
	return
}

// NewVigenere new vigenere
func NewVigenere() {
	return
}

// Encode encodes the string
func (n *Cipher) Encode(e string) (s string) {
	e = strings.ToLower(strings.TrimSpace(e))

	fn := func(r rune) rune {
		switch {
		case r >= 'a' && r <= 'z':
			log.Println("Original: ", r, ": ", 'a'+(r+3)%26)
			return 'a' + (r+3)%26
		default:
			return r
		}
	}

	if len(s) > 0 {
		s = strings.Map(fn, e)
	}
	return
}

// Decode decodes the string
func (n *Cipher) Decode(d string) string {
	return " "
}
