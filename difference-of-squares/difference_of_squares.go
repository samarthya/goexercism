package diffsquares

func SquareOfSum(endT int) (sum int) {
	sum = 0
	element := 1

	for element <= endT {
		sum += element
		element++
	}

	return (sum * sum)
}

func SumOfSquares(endT int) (sum int) {
	sum = 0
	element := 1

	for element <= endT {
		sum += (element * element)
		element++
	}

	return sum
}

func Difference(endT int) int {
	return SquareOfSum(endT) - SumOfSquares(endT)
}
