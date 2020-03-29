package say

import (
	"fmt"
	"log"
	"strings"
	"strconv"
)

//Say reads out the number
func Say(i int64) (string, bool) {
	var s strings.Builder

	log.Printf(" Ele: %d", i)

	if i > 999999999999 || i < 0 {
		return "", false
	}

	for {
		switch {
		case >= 1000:
			//num := string(i)
			for {
				num := (i/1000)
			}
		// case i >= 1000000000 && i < 999999999999:
		// 	r := int(i / 1000000000)
		// 	if val, ok := billion(r); ok {
		// 		fmt.Fprintf(&s, "%s ", val)
		// 	}
		// 	i -= int64(r * 1000000000)
		// case i >= 1000000 && i < 1000000000:
		// 	r := int(i / 1000000)
		// 	if val, ok := million(r); ok {
		// 		fmt.Fprintf(&s, "%s ", val)
		// 	}
		// 	i -= int64(r * 1000000)
		// case i >= 1000:
		// 	r := int(i / 1000)
		// 	if val, ok := thousands(r); ok {
		// 		fmt.Fprintf(&s, "%s ", val)
		// 	}
		// 	i -= int64(r * 1000)
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
