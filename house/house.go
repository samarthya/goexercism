package house

import (
	"fmt"
	"strings"
)

const (
	This string = "This is the"
)

type Character struct {
	name      string
	adjective string
	verb      string
}

var Lines map[int]string = map[int]string{
	0:  "house",
	1:  "malt",
	2:  "rat",
	3:  "cat",
	4:  "dog",
	5:  "cow with the crumpled horn",
	6:  "maiden all forlorn",
	7:  "man all tattered and torn",
	8:  "priest all shaven and shorn",
	9:  "rooster that crowed in the morn",
	10: "farmer sowing his corn",
	11: "horse and the hound and the horn",
}
var Actions map[string]string = map[string]string{
	"house":                            "Jack built.",
	"malt":                             "lay in the house",
	"rat":                              "ate the malt",
	"cat":                              "killed the rat",
	"dog":                              "worried the cat",
	"cow with the crumpled horn":       "tossed the dog",
	"maiden all forlorn":               "milked the cow with the crumpled horn",
	"man all tattered and torn":        "kissed the maiden all forlorn",
	"priest all shaven and shorn":      "married the man all tattered and torn",
	"rooster that crowed in the morn":  "woke the priest all shaven and shorn",
	"farmer sowing his corn":           "kept the rooster that crowed in the morn",
	"horse and the hound and the horn": "belonged to the farmer sowing his corn",
}

func Verse(stanza int) (verse string) {
	var v strings.Builder

	stanza -= 1
	if stanza == 0 {
		fmt.Fprintf(&v, "%s %s that %s", This, Lines[stanza], Actions[Lines[stanza]])
	} else {
		/*
			1. This is the malt
			   that lay in the house that Jack built.
			4. This is the cat
			   that killed the rat
			   that ate the malt\nthat lay in the house that Jack built.
		*/
		for j := stanza; j >= 0; j-- {
			if j == stanza {
				fmt.Fprintf(&v, "%s %s\nthat %s", This, Lines[j], Actions[Lines[j]])
			} else if j == 0 {
				fmt.Fprintf(&v, " that %s", Actions[Lines[0]])
			} else {
				fmt.Fprintf(&v, "\nthat %s", Actions[Lines[j]])
			}
		}
	}
	verse = v.String()
	return verse
}

func Song() (song string) {
	var v strings.Builder
	for j := 1; j < 13; j++ {
		fmt.Fprintf(&v, "%s", Verse(j))
		if j < 12 {
			fmt.Fprintf(&v, "\n\n")
		}
	}

	song = v.String()
	return song
}
