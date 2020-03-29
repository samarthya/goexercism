package say

import (
	"log"
)

//Say reads out the number
func Say(i int64) (string, bool) {
	log.Println(" DEBUG: ", i)
	if i > 999999999999 || i < 0 {
		return "", false
	}

	return NumberConversion(i)
}
