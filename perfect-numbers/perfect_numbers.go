package perfect

import (
	"log"
)

// Classification type of number.
type Classification int

//PerfectErrors Errors to be thrown by Package
type PerfectErrors string

func (e PerfectErrors) Error() string {
	return string(e)
}

// ErrOnlyPositive Positive numbers only
const ErrOnlyPositive PerfectErrors = " positive numbers only"

//NotPerfect Invalid number
const NotPerfect PerfectErrors = " invalid number"

const (
	// ClassificationDeficient Deficient number
	ClassificationDeficient Classification = iota + 1
	// ClassificationPerfect Perfect number
	ClassificationPerfect
	// ClassificationAbundant Abundant number
	ClassificationAbundant
	// ClassificationUnknown Unknown
	ClassificationUnknown
)

// Classify classifies the number
func Classify(num int) (Classification, error) {
	log.Println(" Num: ", num)

	if num <= 0 {
		return ClassificationUnknown, ErrOnlyPositive
	}
	sum := 0
	for i := 1; i <= (num / 2); i++ {
		if num%i == 0 {
			sum += i
		}
	}
	switch {
	case sum == num:
		return ClassificationPerfect, nil
	case sum > num:
		return ClassificationAbundant, nil
	default:
		return ClassificationDeficient, nil
	}
}
