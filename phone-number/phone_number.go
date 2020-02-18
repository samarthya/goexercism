package phonenumber

import (
	"errors"
	"fmt"
	"strings"
)

//NotAllowed characters that are not allowed
var NotAllowed = []rune("!@#$%^&*")

// Number - Gets the complete digits of the phone number.
func Number(s string) (o string, err error) {

	// 1. If empty, nothing to do; return the string itself
	if s == "" {
		return s, nil
	}

	// 2. Check for not allowed characters
	for _, v := range NotAllowed {
		if strings.ContainsRune(s, v) {
			return s, errors.New(" invalid characters")
		}
	}

	// 3. Making an output array.
	out := make([]rune, 0)

	// 4. Get all the numbers
	for _, v := range s {
		switch {
		case '9' >= v && '0' <= v:
			// Append only numbers
			out = append(out, v)
		default:
			// Skip other characters
			// Remove Punctuation and Country code.
			continue
		}
	}

	// Total number of elements
	length := len(out)

	// 5. If lenght is less than 10 or greater than 11
	if length < 10 || length > 11 {
		return s, errors.New(" phone number should be 10")
	}

	// 6. If country code is included.
	if length == 11 {
		if out[0] == '1' {
			// Remove the first element
			out = out[1:]
		} else {
			return s, errors.New(" Cannot be anything other than 1")
		}
	}

	// Get the area Code
	areadCode := out[0:3]

	// 7. Validate area code.
	if areadCode[0] <= '1' {
		// Invalid area code
		return s, errors.New(" Invalid area code")
	}

	// 8. Validate exchange code.
	exchangeCode := out[3:6]
	if exchangeCode[0] <= '1' {
		return s, errors.New(" Invalid exchange code")
	}

	// 9. All valid return the number
	return string(out), nil
}

//AreaCode gets the areacode
func AreaCode(s string) (string, error) {
	out, err := Number(s)

	if err != nil {
		return s, err
	}

	return string(out[0:3]), nil
}

// Format Formats the input string
func Format(s string) (string, error) {
	out, err := Number(s)

	if err != nil {
		return out, err
	}

	areaCode := out[0:3]
	exchangeCode := out[3:6]
	subscriberCode := out[6:]

	return fmt.Sprintf("(%v) %v-%v", areaCode, exchangeCode, subscriberCode), nil
}
