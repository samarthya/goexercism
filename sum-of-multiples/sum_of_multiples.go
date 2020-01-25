package summultiples

import "fmt"

func sumOfFactors(fac, num int) (sum int) {

	sum = 0
	mul := fac
	for {
		if mul < num {
			sum += mul
			mul += fac
			continue
		} else {
			break
		}
	}
	return sum
}

// SumMultiples Sum of multiples
func SumMultiples(number int, multiples ...int) (sum int) {

	for _, val := range multiples {
		sum += sumOfFactors(val, number)
		fmt.Println(" sum: ", sum)
	}
	return sum
}
