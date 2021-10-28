package foodchain

import "bytes"

var Creatures = []string{"fly", "spider", "bird", "cat", "dog", "goat", "cow", "horse"}
var Line1 = "I know an old lady who swallowed a"
var LastLine = "I don't know why she swallowed the fly. Perhaps she'll die."
var Line2 = map[string]string{
	"fly":    LastLine,
	"spider": "It wriggled and jiggled and tickled inside her.",
	"bird":   "How absurd to swallow a bird!",
	"cat":    "Imagine that, to swallow a cat!",
	"dog":    "What a hog, to swallow a dog!",
	"goat":   "Just opened her throat and swallowed a goat!",
	"cow":    "I don't know how she swallowed a cow!",
	"horse":  "She's dead, of course!",
}
var Line3_1 = "She swallowed the"
var Line3_2 = "to catch the"
var SpiderSpecial = "that wriggled and jiggled and tickled inside her"

func Verse(i int) string {
	var b bytes.Buffer
	b.WriteString(Line1)
	b.WriteString(" ")
	b.WriteString(Creatures[i-1])
	b.WriteString(".\n")
	b.WriteString(Line2[Creatures[i-1]])
	if i > 1 && i < 8 {
		for j := i; j >= 2; j-- {
			b.WriteString("\n")
			b.WriteString(Line3_1)
			b.WriteString(" ")
			b.WriteString(Creatures[j-1])
			b.WriteString(" ")
			b.WriteString(Line3_2)
			b.WriteString(" ")
			b.WriteString(Creatures[j-2])
			if j-1 == 2 {
				b.WriteString(" ")
				b.WriteString(SpiderSpecial)
			}
			b.WriteString(".")
		}
		b.WriteString("\n")
		b.WriteString(Line2["fly"])
	}
	return b.String()
}
func Verses(i, j int) string {
	var b bytes.Buffer
	for k := i; k <= j; k++ {
		b.WriteString(Verse(k))
		if k != j {
			b.WriteString("\n")
			b.WriteString("\n")
		}
	}
	return b.String()
}
func Song() string {
	return Verses(1, 8)
}
