package proverb

import "fmt"

const (
	// TemplateString Used for proverbs generation
	TemplateString string = "For want of a %s the %s was lost."
	// ConcludingString Used for concluding statement.
	ConcludingString string = "And all for the want of a %s."
)

// Proverb function.
func Proverb(rhyme []string) (ret []string) {

	if rhyme == nil || len(rhyme) == 0 {
		return []string{}
	}

	ret = make([]string, len(rhyme))

	switch {
	case len(rhyme) == 0:
		return []string{}
	case len(rhyme) <= 1:
		ret[0] = fmt.Sprintf(ConcludingString, rhyme[0])
	default:
		for i, val := range rhyme {
			if i == 0 {
				ret[len(rhyme)-1] = fmt.Sprintf(ConcludingString, val)
			} else {
				ret[i-1] = fmt.Sprintf(TemplateString, rhyme[i-1], val)
			}
		}
	}
	
	return ret
}
