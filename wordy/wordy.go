package wordy

import (
	"log"
	"regexp"
	"strconv"
)

const (
	plus = "plus"
	sub  = "minus"
	div  = "divided by"
	mul  = "multiplied by"
)

const (
	// PLUS constant
	PLUS string = `-?\d+\splus\s-?\d*`
	//SUBTRACT subtracts
	SUBTRACT = `-?\d+\sminus\s-?\d*`
	//MULTIPLY subtracts
	MULTIPLY = `-?\d+\smultiplied\sby\s-?\d*`
	//DIVIDE subtracts
	DIVIDE = `-?\d+\sdivided\sby\s-?\d*`
	//NUMERICQ numeric question
	NUMERICQ = `-?\d+.$`
	// NUMBERS all numbers in the system
	NUMBERS = `-?\d+`
)

// getNum Get the number
func getNum(n string) (int, error) {
	val, err := strconv.Atoi(n)
	return val, err
}

// Answer Answers the statement.
func Answer(q string) (int, bool) {
	onlyQuestion := regexp.MustCompile(NUMERICQ)
	plusQuery := regexp.MustCompile(PLUS)
	mulQuery := regexp.MustCompile(MULTIPLY)
	divQuery := regexp.MustCompile(DIVIDE)
	subQuery := regexp.MustCompile(SUBTRACT)
	numbersFound := regexp.MustCompile(NUMBERS)

	numbers := numbersFound.FindAllString(q, -1)

	switch {
	case len(plusQuery.FindAllString(q, -1)) >= 1:
		el := plusQuery.FindAllString(q, -1)
		sum := 0

		for _, val := range el {
			num := numbersFound.FindAllString(val, -1)
			// Find the 2 numbers to add.
			if len(num) < 2 {
				return 0, false
			}

			for _, n := range num {
				number, err := getNum(n)

				if err != nil {
					return 0, false
				}

				sum += number

			}
		}
		return sum, true

	case len(subQuery.FindAllString(q, -1)) >= 1:
		el := subQuery.FindAllString(q, -1)
		sum := 0

		for _, val := range el {
			num := numbersFound.FindAllString(val, -1)
			// Find the 2 numbers to add.
			if len(num) < 2 {
				return 0, false
			}

			for i, n := range num {
				number, err := getNum(n)

				if err != nil {
					return 0, false
				}
				if i == 0 {
					sum = number
				} else {
					sum -= number
				}

			}
		}
		return sum, true
	case len(mulQuery.FindAllString(q, -1)) >= 1:
		el := mulQuery.FindAllString(q, -1)
		sum := 0

		for _, val := range el {
			num := numbersFound.FindAllString(val, -1)
			// Find the 2 numbers to add.
			if len(num) < 2 {
				return 0, false
			}

			for i, n := range num {
				number, err := getNum(n)

				if err != nil {
					return 0, false
				}
				if i == 0 {
					sum = number
				} else {
					sum *= number
				}

			}
		}
		return sum, true
	case len(divQuery.FindAllString(q, -1)) >= 1:
		el := divQuery.FindAllString(q, -1)
		sum := 0

		for _, val := range el {
			num := numbersFound.FindAllString(val, -1)
			// Find the 2 numbers to add.
			if len(num) < 2 {
				return 0, false
			}

			for i, n := range num {
				number, err := getNum(n)

				if err != nil {
					return 0, false
				}
				if i == 0 {
					sum = number
				} else {
					sum /= number
				}

			}
		}
		return sum, true
	case len(onlyQuestion.FindAllString(q, -1)) >= 1:
		el := onlyQuestion.FindAllString(q, -1)
		if len(el) == 1 {

			// Found elements
			log.Println(" Elements: ", numbers)

			val, err := getNum(numbers[0])

			if err != nil {
				return 0, false
			}

			return val, true
		}
	}
	// Matches the query

	return 0, false
}
