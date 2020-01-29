package summultiples

func sumOfFactors(fac, num int, countMap map[int]int) {

	mul := fac

	if fac <= 0 {
		return
	}

	for {
		if fac > num || mul >= num {
			break
		}
		countMap[mul] = mul
		mul += fac
	}
}

// SumMultiples Sum of multiples
func SumMultiples(number int, multiples ...int) (sum int) {
	var countMap map[int]int = make(map[int]int)
	if number <= 0 {
		return 0
	}

	for _, val := range multiples {
		sumOfFactors(val, number, countMap)
	}

	for _, val := range countMap {
		sum += val
	}

	return sum
}
