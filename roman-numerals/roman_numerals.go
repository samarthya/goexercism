package romannumerals

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// IVXCLDM

// arabicMap for representation of roman numerals.
var arabicMap = map[int]string{
	1:    "I",
	5:    "V",
	10:   "X",
	40:   "XL",
	50:   "L",
	60:   "LX",
	70:   "LXX",
	80:   "XXC",
	90:   "XC",
	100:  "C",
	200:  "CC",
	300:  "CCD",
	400:  "CD",
	500:  "D",
	600:  "DC",
	700:  "DCC",
	800:  "CCM",
	900:  "CM",
	1000: "M",
}

func onesConversion(a int) string {
	var s strings.Builder
	if a <= 10 {
		if arabicMap[a] != "" {
			fmt.Fprintf(&s, "%s", arabicMap[a])
		} else {
			switch {
			case a > 5 && a <= 8:
				fmt.Fprintf(&s, "%s", arabicMap[5])
				fmt.Fprintf(&s, "%s", strings.Repeat(arabicMap[1], (a-5)))
			case a > 8 && a < 10:
				fmt.Fprintf(&s, "%s", strings.Repeat(arabicMap[1], (10-a)))
				fmt.Fprintf(&s, "%s", arabicMap[10])
			case a > 1 && a <= 3:
				fmt.Fprintf(&s, "%s", strings.Repeat(arabicMap[1], a))
			case a == 4:
				fmt.Fprintf(&s, "%s", strings.Repeat(arabicMap[1], 1))
				fmt.Fprintf(&s, "%s", arabicMap[5])
			}
		}
	}
	return s.String()
}

// ToRomanNumeral Converts the int to Roman numeral.
func ToRomanNumeral(a int) (string, error) {
	var s strings.Builder

	if a > 3000 || a <= 0 {
		return "", errors.New(" Invalid input")
	}
	// Convert the int to String.
	// E.g. 123 => "[1,2,3]"
	fvs := strings.Split(strconv.Itoa(a), "")

	lenFVS := len(fvs)
	index := 0
	for pos := lenFVS; pos >= 2; pos-- {

		placeValue := int(math.Pow10(pos - 1))

		// How many times to repeat e.g. 100 * 2
		repeatTimes, _ := strconv.Atoi(fvs[index])
		if arabicMap[(placeValue*repeatTimes)] != "" {
			fmt.Fprintf(&s, "%s", arabicMap[(placeValue*repeatTimes)])
		} else {
			fmt.Fprintf(&s, "%s", strings.Repeat(arabicMap[placeValue], repeatTimes))
		}

		index++
	}

	val, _ := strconv.Atoi(fvs[index])
	fmt.Fprintf(&s, "%s", onesConversion(val))
	return s.String(), nil
}
