package say

import (
	"fmt"
	"strings"
)

//myMap defines the constant string values
var myMap = map[int]string{
	0:          "zero", // 0 - Starts with zero.
	1:          "one",
	2:          "two",
	3:          "three",
	4:          "four",
	5:          "five",
	6:          "six",
	7:          "seven",
	8:          "eight",
	9:          "nine",
	10:         "ten",
	11:         "eleven",
	12:         "twelve",
	13:         "thirteen",
	20:         "twenty",
	30:         "thirty",
	40:         "fourty",
	50:         "fifty",
	60:         "sixty",
	70:         "seventy",
	80:         "eighty",
	90:         "ninety",
	100:        "hundred",
	1000:       "thousand",
	1000000:    "million",
	1000000000: "billion",
}

//tens It is a function that will convert the numeric into string representation
func tens(i int) (string, bool) {

	// Valid for numbers less than hundred only.
	switch {
	case i > 13 && i < 20:
		return fmt.Sprintf("%s%s", myMap[i-10], "teen"), true
	case i > 20 && i < 29:
		return fmt.Sprintf("%s%s", "twenty-", myMap[i-20]), true
	case i > 30 && i < 40:
		return fmt.Sprintf("%s%s", "thirty-", myMap[i-30]), true
	case i > 40 && i < 50:
		return fmt.Sprintf("%s%s", "forty-", myMap[i-40]), true
	case i > 50 && i < 60:
		return fmt.Sprintf("%s%s", "fifty-", myMap[i-50]), true
	case i > 60 && i < 70:
		return fmt.Sprintf("%s%s", "sixty-", myMap[i-60]), true
	case i > 70 && i < 80:
		return fmt.Sprintf("%s%s", "seventy-", myMap[i-70]), true
	case i > 80 && i < 90:
		return fmt.Sprintf("%s%s", "eighty-", myMap[i-80]), true
	case i > 90 && i < 100:
		return fmt.Sprintf("%s%s", "ninety-", myMap[i-90]), true
	default:
		// If there is a numeric equivalent available in map.
		if val, ok := myMap[int(i)]; ok {
			return val, true
		}
	}

	//No match found in map or can be calculated
	return "", false
}

// hundreds Allows to compute the numeric equivalent of the 100-999
func hundreds(r int) (string, bool) {

	switch {
	case r == 0:
		return fmt.Sprintf("one %s", myMap[100]), true
	default:
		return fmt.Sprintf("%s %s", myMap[r], myMap[100]), true
	}
}

//NumberConversion Converts the number
func NumberConversion(i int64) (string, bool) {
	var s strings.Builder

	for i >= 0 {
		switch {
		case i == 0:
			if len(s.String()) == 0 {
				if val, ok := tens(int(i)); ok {
					fmt.Fprintf(&s, " %s", val)
				}
			}
			return strings.TrimSpace(s.String()), true
		case i < 100:
			if val, ok := tens(int(i)); ok {
				fmt.Fprintf(&s, " %s", val)
			}
			return strings.TrimSpace(s.String()), true
		case i >= 100 && i < 1000:
			r := int(i / 100)
			if val, ok := hundreds(r); ok {
				fmt.Fprintf(&s, " %s", val)
			}
			i %= 100
		case i >= 1000 && i < 1000000:
			r := i / 1000
			if val, ok := NumberConversion(r); ok {
				fmt.Fprintf(&s, " %s %s", val, myMap[1000])
			}
			i %= 1000
		case i >= 1000000 && i < 1000000000:
			r := i / 1000000
			if val, ok := NumberConversion(r); ok {
				fmt.Fprintf(&s, " %s %s", val, myMap[1000000])
			}
			i %= 1000000
		case i >= 1000000000 && i < 1000000000000:
			r := i / 1000000000
			if val, ok := NumberConversion(r); ok {
				fmt.Fprintf(&s, "%s %s", val, myMap[1000000000])
			}
			i %= 1000000000
		}
	}
	return "", false
}
