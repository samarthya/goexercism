package cipher

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

// RegularExp Elements to find
var RegularExp = regexp.MustCompile(`[^A-Za-z]`)

// Encrypt struct
type Encrypt struct {
	distance int
	key      string
}

// NewCaesar Object that implements the interface Cipher
func NewCaesar() Cipher {
	return Encrypt{distance: 3, key: ""}
}

// NewShift Object that implements the interface Cipher
func NewShift(distance int) Cipher {
	// Corner case when of 0, 26, -26
	if distance < -26 || distance > 26 || distance == 0 {
		return nil
	}
	return Encrypt{distance, ""}
}

// NewVigenere Object that implements the interface Cipher
func NewVigenere(key string) Cipher {

	if key == "" || checkKeyInvalid(key) {
		return nil
	}
	return Encrypt{0, key}
}

// checkAllAs Checks for all A's in the key.
func checkKeyInvalid(s string) (f bool) {

	// If any invalid character matches
	if matched, _ := regexp.MatchString(`[^a-z]`, s); matched {
		return true
	}

	for _, m := range s {
		switch m {
		case 'a':
			f = true
		default:
			f = false
			return
		}
	}
	return
}

// removeAllOtherCharcterBeforeEncoding replace all the other elements
func removeAllOtherCharcterBeforeEncoding(s string) (o string) {
	o = RegularExp.ReplaceAllString(s, "")
	log.Printf(" String: %s", o)
	o = strings.ToLower(o)
	return
}

func transformWithString(s string, e string, flag bool) string {
	var sb strings.Builder

	em := []rune(e)

	for i, sm := range s {
		// The key has to be repeated
		var d rune = 0

		if i < len(em) {
			d = em[i] - 'a'
		} else {
			d = em[(i%len(em))] - 'a'
		}

		if flag {
			switch {
			case sm >= 'a' && sm <= 'z':
				fmt.Fprintf(&sb, "%s", string('a'+(sm-'a'+d)%26))
			}
		} else {
			switch {
			case sm >= 'a' && sm <= 'z':
				diff := ((sm - 'a' - d) % 26)
				newChar := ""

				if diff < 0 {
					newChar = string('z' + diff + 1)
				} else {
					newChar = string('a' + diff)
				}
				fmt.Fprintf(&sb, "%s", newChar)
			}
		}
	}
	return sb.String()
}

// transformWithDistance Allows transformation based on distance.
func transformWithDistance(s string, d int) string {

	rotateFunction := func(r rune) rune {
		switch {
		case r >= 'a' && r <= 'z':
			val := (r - 'a' + rune(d))
			if val < 0 && d < 0 {
				return 'z' + ((val + 1) % 26)
			}
			return 'a' + (val % 26)

		default:
			log.Printf(" Not a charcter: %s", string(r))
		}
		return ' '
	}

	return strings.Map(rotateFunction, s)

}

// Encode Implements the interface Cipher
func (e Encrypt) Encode(s string) (o string) {
	o = ""
	s = removeAllOtherCharcterBeforeEncoding(s)

	if s == "" || len(s) <= 0 {
		return
	}

	// Iterate throught the runes

	switch {
	case e.distance != 0:
		o = strings.ReplaceAll(transformWithDistance(s, e.distance), " ", "")
	default:
		o = transformWithString(s, e.key, true)
	}

	return
}

// Decode decides the encrypted string
func (e Encrypt) Decode(s string) (o string) {
	o = ""

	log.Printf(" Decode: %s", s)

	if s == "" || len(s) <= 0 {
		return
	}

	switch {
	case e.distance != 0:
		log.Printf(" Distance: %d", e.distance)
		o = transformWithDistance(s, -e.distance)
	default:
		o = transformWithString(s, e.key, false)
	}

	return
}
