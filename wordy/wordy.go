package wordy

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

const (
	// NUMBERS all numbers in the system
	NUMBERS = `-?\d+`

	//Operators Operators
	Operators = `[\+-/*]?`

	//Question Question
	Question = `\d+\?$`
)

// getNum Get the number
func getNum(n string) (int, error) {
	val, err := strconv.Atoi(n)
	return val, err
}

//isQuestion Check if it is a question
func isQuestion(q string) bool {
	question := regexp.MustCompile(Question)
	return question.Match([]byte(q))
}

// Answer Answers the statement.
func Answer(q string) (int, bool) {
	var stack []string = make([]string, 0)
	var sum int = 0
	numbers := regexp.MustCompile(NUMBERS)
	operators := regexp.MustCompile(Operators)

	// Matches the query
	q = strings.TrimSpace(q)

	if !isQuestion(q) {
		return 0, false
	}

	var elements []string = strings.Split(q, " ")

	for _, v := range elements {
		switch v {
		case "plus":
			stack = append(stack, "+")
		case "minus":
			stack = append(stack, "-")
		case "multiplied":
			stack = append(stack, "*")
		case "divided":
			stack = append(stack, "/")
		default:
			n := numbers.FindAllString(v, -1)
			if len(n) == 1 {
				stack = append(stack, n[0])
			}
		}
	}

	// No valid data found
	if len(stack) == 0 {
		return 0, false
	}

	for i, el := range stack {
		// If it is the first element
		if i == 0 {
			val, err := getNum(el)
			if err != nil {
				return 0, false
			}
			sum = val
		} else {

			if operators.Match([]byte(el)) {
				switch el {
				case "+":
					val, err := getNum(stack[i+1])
					if err != nil {
						return 0, false
					}
					sum += val
				case "-":
					val, err := getNum(stack[i+1])
					if err != nil {
						return 0, false
					}
					sum -= val
				case "/":
					val, err := getNum(stack[i+1])
					if err != nil {
						return 0, false
					}
					sum /= val
				case "*":
					val, err := getNum(stack[i+1])
					if err != nil {
						return 0, false
					}
					sum *= val
				default:
					if i%2 != 0 {
						return 0, false
					}
				}
			}
		}
	}

	log.Println(" Read: ", stack)
	return sum, true
}
