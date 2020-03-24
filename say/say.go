package say

import (
	"fmt"
	"log"
	"strings"
)

var myMap = map[int]string{
	0:          "zero",
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

func tens(i int) (string, bool) {

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

		if val, ok := myMap[int(i)]; ok {
			return val, true
		}
	}

	return "", false
}

func hundreds(r int) (string, bool) {

	switch r {
	case 0:
		return fmt.Sprintf("one %s", myMap[100]), true
	default:
		return fmt.Sprintf("%s %s", myMap[r], myMap[100]), true

	}
}

func thousands(r int) (string, bool) {
	switch {
	case r == 0:
		return fmt.Sprintf("one %s", myMap[1000]), true
	case r >= 1 && r <= 10:
		if val, ok := tens(r); ok {
			return fmt.Sprintf("%s %s", val, myMap[1000]), true
		}
		return "", false
	case r > 10 && r <= 1000:
		if val, ok := hundreds(r); ok {
			return fmt.Sprintf("%s %s", val, myMap[1000]), true
		}
		return "", false
	default:
		return fmt.Sprintf("%s %s", myMap[r], myMap[1000]), true
	}
}

func million(r int) (string, bool) {
	switch {
	case r == 0:
		return fmt.Sprintf("one %s", myMap[1000000]), true
	case r >= 1 && r <= 10:
		if val, ok := tens(r); ok {
			return fmt.Sprintf("%s %s", val, myMap[1000000]), true
		}
		return "", false
	case r > 10 && r <= 100:
		if val, ok := hundreds(r); ok {
			return fmt.Sprintf("%s %s", val, myMap[1000000]), true
		}
		return "", false
	case r > 100 && r <= 1000:
		if val, ok := thousands(r); ok {
			return fmt.Sprintf("%s %s", val, myMap[1000000]), true
		}
		return "", false
	default:
		return fmt.Sprintf("%s %s", myMap[r], myMap[1000]), true
	}
}

func billion(r int) (string, bool) {
	switch {
	case r == 0:
		return fmt.Sprintf("one %s", myMap[1000000000]), true
	case r >= 1 && r <= 10:
		if val, ok := tens(r); ok {
			return fmt.Sprintf("%s %s", val, myMap[1000000000]), true
		}
		return "", false
	case r > 10 && r <= 1000:
		if val, ok := hundreds(r); ok {
			return fmt.Sprintf("%s %s", val, myMap[1000000000]), true
		}
		return "", false
	case r >= 1000 && r <= 1000000:
		if val, ok := thousands(r); ok {
			return fmt.Sprintf("%s %s", val, myMap[1000000000]), true
		}
		return "", false
	case r > 1000000 && r <= 1000000000:
		if val, ok := million(r); ok {
			return fmt.Sprintf("%s %s", val, myMap[1000000000]), true
		}
		return "", false
	default:
		return fmt.Sprintf("%s %s", myMap[r], myMap[1000000000]), true
	}
}

//Say reads out the number
func Say(i int64) (string, bool) {
	var s strings.Builder

	log.Printf(" Ele: %d", i)

	if i > 999999999999 || i < 0 {
		return "", false
	}

	for {
		switch {
		case i >= 1000000000 && i < 999999999999:
			r := int(i / 1000000000)
			if val, ok := billion(r); ok {
				fmt.Fprintf(&s, "%s ", val)
			}
			i -= int64(r * 1000000000)
		case i >= 1000000 && i < 1000000000:
			r := int(i / 1000000)
			if val, ok := million(r); ok {
				fmt.Fprintf(&s, "%s ", val)
			}
			i -= int64(r * 1000000)
		case i >= 1000 && i < 100000:
			r := int(i / 1000)
			if val, ok := thousands(r); ok {
				fmt.Fprintf(&s, "%s ", val)
			}
			i -= int64(r * 1000)
		case i >= 100 && i < 1000:
			r := int(i / 100)
			if val, ok := hundreds(r); ok {
				fmt.Fprintf(&s, "%s ", val)
			}
			i -= int64(r * 100)
		case i > 0 && i < 100:
			if val, ok := tens(int(i)); ok {
				fmt.Fprintf(&s, "%s ", val)
			}
			return strings.TrimSpace(s.String()), true
		case i == 0:
			if len(s.String()) == 0 {
				if val, ok := tens(int(i)); ok {
					fmt.Fprintf(&s, "%s ", val)
				}
			}
			return strings.TrimSpace(s.String()), true
		default:
			if val, ok := myMap[int(i)]; ok {
				return val, true
			}
		}
	}

	return "", false
}
