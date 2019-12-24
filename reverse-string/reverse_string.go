package reverse

func Reverse(s string) string {
	chars := []rune(s)
	size := len(chars)

	for i := 0; i < size/2; i++ {
		chars[i], chars[size-i-1] = chars[size-i-1], chars[i]
	}

	return string(chars)
}
