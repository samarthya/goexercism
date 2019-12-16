package dna

import (
	"errors"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {

	var h Histogram = make(map[rune]int)

	h['A'] = 0
	h['C'] = 0
	h['G'] = 0
	h['T'] = 0

	if d == "" {
		return h, nil
	}

	// h['A'] = strings.Count(string(d), "A")
	// h['C'] = strings.Count(string(d), "C")
	// h['G'] = strings.Count(string(d), "G")
	// h['T'] = strings.Count(string(d), "T")

	for _, val := range d {
		switch val {
		case 'C', 'A', 'G', 'T':
			h[val] += 1

		default:
			return h, errors.New("invalid nucleoid")
		}
	}

	return h, nil
}
