package protein

import (
	"errors"
	"regexp"
	"strings"
)

var (
	// ErrStop stop error
	ErrStop = errors.New("STOP word has been found")

	// ErrInvalidBase invalid base
	ErrInvalidBase = errors.New("the base is invalid")

	// Proteins that are to mapped
	Proteins = map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "STOP",
		"UAG": "STOP",
		"UGA": "STOP",
	}
)

// FromRNA Returns proteins
func FromRNA(rna string) ([]string, error) {
	// 3 character long
	var rex = regexp.MustCompile(`...`)
	codons := rex.FindAllString(rna, -1)

	var proteins []string

	for _, codon := range codons {
		protein, err := FromCodon(codon)

		switch err {
		case ErrStop:
			return proteins, nil
		case ErrInvalidBase:
			return proteins, err
		}

		proteins = append(proteins, protein)
	}

	return proteins, nil
}

// FromCodon Returns proteins
func FromCodon(codon string) (string, error) {
	out := ""
	if codon != "" {

		codon = strings.ToUpper(codon)
		out = Proteins[codon]

		if strings.Compare(out, "STOP") == 0 {
			return "", ErrStop
		} else if out == "" {
			return "", ErrInvalidBase
		}
	}
	return out, nil
}
