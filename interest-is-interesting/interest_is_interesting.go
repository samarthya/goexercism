package interest

import (
	"log"
)

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {

	switch {
	case balance < 0:
		return 3.213
	case balance < 1000:
		return 0.5
	case balance < 5000:
		return 1.621
	case balance > 5000:
		fallthrough
	default:
		return 2.475
	}
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return (balance * float64(InterestRate(balance))) / 100
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	f := balance

	if targetBalance < balance {
		return 0
	}

	log.Println(balance, targetBalance)

	for i := 1; ; {
		if f = AnnualBalanceUpdate(f); f >= targetBalance {
			return i
		}
		log.Println(f, targetBalance)
		i++
	}
}
