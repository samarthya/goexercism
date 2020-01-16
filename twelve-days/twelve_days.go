package twelve

import (
	"fmt"
	"strings"
)

const (
	Opening       string = "On the"
	OpeningSuffix string = "day of Christmas my true love gave to me: "
)

var Days map[int]string = map[int]string{
	0:  "first",
	1:  "second",
	2:  "third",
	3:  "fourth",
	4:  "fifth",
	5:  "sixth",
	6:  "seventh",
	7:  "eighth",
	8:  "ninth",
	9:  "tenth",
	10: "eleventh",
	11: "twelfth",
}

var Lines []string = []string{
	"a Partridge in a Pear Tree",
	"two Turtle Doves",
	"three French Hens",
	"four Calling Birds",
	"five Gold Rings",
	"six Geese-a-Laying",
	"seven Swans-a-Swimming",
	"eight Maids-a-Milking",
	"nine Ladies Dancing",
	"ten Lords-a-Leaping",
	"eleven Pipers Piping",
	"twelve Drummers Drumming",
}

// Stitch the song.
func SongPlay(stringSlice []string) string {
	stringSuffix := ""
	length := len(stringSlice) - 1

	for i, _ := range stringSlice {
		if i == 0 {
			stringSuffix += stringSlice[length]
		} else if i == length {
			stringSuffix += (", and " + stringSlice[0])
		} else {
			stringSuffix += ", " + stringSlice[length-i]
		}
	}
	stringSuffix += "."
	return stringSuffix
}

//Verse gets the song verse for the specified day, starting at day 1.
func Verse(num int) string {
	var lineSong strings.Builder

	fmt.Fprintf(&lineSong, "%s %s %s", Opening, Days[num-1], OpeningSuffix)
	fmt.Fprintf(&lineSong, "%s", SongPlay(Lines[:num]))

	return lineSong.String()
}

func Song() string {
	var completeSong strings.Builder

	for i := 0; i < 12; i++ {
		fmt.Fprintf(&completeSong, "%s\n", Verse(i+1))
	}

	return strings.Trim(completeSong.String(), "\n")
}
