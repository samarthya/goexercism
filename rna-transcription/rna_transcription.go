package strand

import "strings"

type mappingFunction func(rune) rune

//ToRNA allows calculation of complement of the DNA.
func ToRNA(dna string) (rna string) {
	if dna == "" {
		rna = ""
	}

	var myFunction mappingFunction = func(r rune) rune {
		switch r {
		case 'G':
			return 'C'
		case 'C':
			return 'G'
		case 'T':
			return 'A'
		case 'A':
			return 'U'
		}
		return r
	}

	// Will use Map function to allow the mapping.
	rna = strings.Map(myFunction, dna)
	return rna
}
