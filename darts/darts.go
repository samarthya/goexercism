package darts

import "math"

// Score Allows to calculate the score of the dart coordinates in a circle.
func Score(x, y float64) int {
	switch {
	case math.Sqrt(math.Pow(x, 2)+math.Pow(y, 2)) <= 1:
		return 10
	case math.Sqrt(math.Pow(x, 2)+math.Pow(y, 2)) <= 5:
		return 5
	case math.Sqrt(math.Pow(x, 2)+math.Pow(y, 2)) <= 10:
		return 1
	default:
		return 0
	}
}
