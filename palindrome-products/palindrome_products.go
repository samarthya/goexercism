package palindrome

import (
	"errors"
	"log"
	"strconv"
)

// Product defines the structure
type Product struct {
	Product        int      // palindromic, of course
	Factorizations [][2]int //list of all possible two-factor factorizations of Product, within given limits, in order
}

// Products finds the product
func Products(fmin, fmax int) (pmin, pmax Product, err error) {
	pmin.Factorizations = make([][2]int, 0)
	pmin.Product = 0
	pmax.Product = 0

	if fmin > fmax {
		return pmin, pmax, errors.New("fmin > fmax...")
	}

	for fmin <= fmax {
		for x := fmin; x <= fmax; x++ {
			// Check the product if is palndrome
			prod := (fmin * x)

			// First check if the product is palindrome only 3 or more digits can be palindrome
			if isPalindrome(prod) {

				if pmin.Product > prod || pmin.Product == 0 {
					pmin.Product = prod
					pmin.Factorizations = append(pmin.Factorizations, [2]int{fmin, x})
				}

				switch {
				case pmax.Product < prod:
					pmax.Product = prod

					// It is replacement
					if len(pmax.Factorizations) > 0 {
						// Remove the old
						pmax.Factorizations = make([][2]int, 0)
					}
					pmax.Factorizations = append(pmax.Factorizations, [2]int{fmin, x})
				case pmax.Product == prod:
					pmax.Factorizations = append(pmax.Factorizations, [2]int{fmin, x})
				}
			}

		}
		fmin++
	}

	if pmax.Product == 0 {
		log.Println(" PMIN", pmin, " PMAX", pmax)
		return pmin, pmax, errors.New("no palindromes")
	}

	return pmin, pmax, nil
}

// Checks if it is a plindrome
func isPalindrome(num int) (palindrome bool) {
	//log.Println(" Palindrome check ", num)
	elements := []rune(strconv.Itoa(num))
	//log.Println(" Number: ", elements)
	palindrome = true
	for l, ele := range elements {
		if l >= (len(elements) / 2) {
			break
		}
		if ele != (elements[len(elements)-(l+1)]) {
			palindrome = false
		}
	}
	return
}
