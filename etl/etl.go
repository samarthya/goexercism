package etl

import "strings"

// Transform ALlows you transform the input
func Transform(input map[int][]string) (output map[string]int) {
	output = make(map[string]int)
	for i, val := range input {
		// i Denotes the Score (key)
		for _, ele := range val {
			output[strings.ToLower(ele)] = i
		}
	}
	return output
}
