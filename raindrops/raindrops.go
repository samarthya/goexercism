package raindrops

import "strconv"

/**
 * Converts the raindrops.
 */
func Convert(drop int) (output string) {

	if drop%3 == 0 {
		output = "Pling"
	}

	if drop%5 == 0 {
		output += "Plang"
	}

	if drop%7 == 0 {
		output += "Plong"
	}

	if len(output) == 0 {
		output = strconv.Itoa(drop)
	}

	return output
}
