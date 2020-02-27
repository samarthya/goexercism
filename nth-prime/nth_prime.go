package prime

// isPrime checks if the number is prime
func isPrime(num int) bool {
	for d := 2; d <= num/2; d++ {
		if (num % d) == 0 {
			return false
		}
	}
	return true
}

// Nth - Returns the Nth prime number
func Nth(n int) (int, bool) {
	num := 2

	if n < 1 {
		return 0, false
	}

	for {
		if isPrime(num) {
			n--
		}
		if n == 0 {
			break
		}
		num++

	}

	return num, true
}
