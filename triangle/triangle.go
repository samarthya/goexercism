// Package triangle All you need to know aboout a triangle.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import (
	"fmt"
	"math"
)

// Kind : KindFromSides() returns this type.
type Kind string

// Pick values for the following identifiers used by the test program.
const (
	// NaT not a triangle
	NaT Kind = "Nat" // not a triangle
	// Equ
	Equ Kind = "Equ" // equilateral
	// Iso
	Iso Kind = "Iso" // isosceles
	// Sca
	Sca Kind = "Sca" // scalene
)

func thirdSideRule(a, b, c float64) bool {
	var twoConditons int32 = 0

	if (a + b) >= c {
		twoConditons++
	}

	if (a + c) >= b {
		twoConditons++
	}

	if (b + c) >= a {
		twoConditons++
	}

	if twoConditons > 2 {
		return true
	}

	return false
}

/**
 * For a shape to be a triangle at all, all sides have to be of length > 0, and the sum
 * of the lengths of any two sides must be greater than or equal to the length of the
 * third side. See Triangle Inequality.
 */
func isTriangle(a, b, c float64) bool {

	if math.IsInf(math.Abs(a), 1 ) || math.IsInf(math.Abs(b), 1) || math.IsInf(math.Abs(c), 1) {
		return false
	}
	
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}

	return thirdSideRule(a, b, c)

}

func isISC(a, b, c float64) bool {
	if a == b || b == c || c == a {
		return true
	}
	return false
}

func isEQU(a, b, c float64) bool {
	if a == b && b == c {
		return true
	}
	return false
}

func isSCA(a, b, c float64) bool {
	if a != b && b != c && c != a {
		return true
	}
	return false
}

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind = NaT

	fmt.Println(a, b, c)
	if isTriangle(a, b, c) {
		if isSCA(a, b, c) {
			return Sca
		}

		if isEQU(a, b, c) {
			return Equ
		}

		if isISC(a, b, c) {
			return Iso
		}

	}
	return k
}
