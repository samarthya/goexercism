package house

import (
	"fmt"
	"strings"
)

const (
	This string = "This is the"
)

type Predicate struct {
	character string
	predicate string
}

type Character struct {
	adjective string
	pred      Predicate
}

var characters map[string]Character = map[string]Character{
	"house": {adjective: "", pred: Predicate{
		predicate: "built.", character: "Jack",
	},
	},
	"malt": {adjective: "", pred: Predicate{
		predicate: "lay in the", character: "house",
	},
	},
	"rat": {adjective: "", pred: Predicate{
		predicate: "ate the", character: "malt",
	},
	},
	"cat": {adjective: "", pred: Predicate{
		predicate: "killed the", character: "rat",
	},
	},
	"dog": {adjective: "", pred: Predicate{
		predicate: "worried the", character: "cat",
	},
	},
	"cow": {adjective: "with the crumpled horn", pred: Predicate{
		predicate: "tossed the", character: "dog",
	},
	},
	"maiden": {adjective: "all forlorn", pred: Predicate{
		predicate: "milked the", character: "cow",
	},
	},
	"man": {adjective: "all tattered and torn", pred: Predicate{
		predicate: "kissed the", character: "maiden",
	},
	},
	"priest": {adjective: "all shaven and shorn", pred: Predicate{
		predicate: "married the", character: "man",
	},
	},
	"rooster": {adjective: "that crowed in the morn", pred: Predicate{
		predicate: "woke the", character: "priest",
	},
	},
	"farmer": {adjective: "sowing his corn", pred: Predicate{
		predicate: "kept the", character: "rooster",
	},
	},
	"horse": {adjective: "and the hound and the horn", pred: Predicate{
		predicate: "belonged to the", character: "farmer",
	},
	},
}

var verses map[int]string = map[int]string{
	1: "house",
	2: "malt",
	3: "rat",
	4: "cat",
	5: "dog",
	6: "cow", 7: "maiden",
	8: "man", 9: "priest",
	10: "rooster", 11: "farmer",
	12: "horse",
}

func Verse(stanza int) (verse string) {
	var v strings.Builder
	var character string
	character = verses[stanza]
	if stanza == 1 {
		fmt.Fprintf(&v, "%s %s that %s %s", This, character, characters[character].pred.character, characters[character].pred.predicate)
	} else {
		for j := stanza; j > 0; j-- {
			character = verses[j]
			predChar := characters[character].pred.character
			if j == stanza {
				if characters[character].adjective == "" {
					fmt.Fprintf(&v, "%s %s\nthat %s %s", This, character, characters[character].pred.predicate, characters[character].pred.character)
				} else {
					fmt.Fprintf(&v, "%s %s %s\nthat %s %s", This, character, characters[character].adjective, characters[character].pred.predicate, characters[character].pred.character)
				}
			} else if j == 1 {
				fmt.Fprintf(&v, " that %s %s", characters[character].pred.character, characters[character].pred.predicate)
			} else {
				fmt.Fprintf(&v, "\nthat %s %s", characters[character].pred.predicate, characters[character].pred.character)
			}
			if predChar != "" && characters[predChar].adjective != "" {
				fmt.Fprintf(&v, " %s", characters[predChar].adjective)
			}
		}
	}

	verse = strings.TrimSpace(v.String())
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
