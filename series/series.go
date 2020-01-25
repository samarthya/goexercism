package series

// All returns the string array.
func All(n int, s string) (output []string) {
	output = make([] string, 0)
	var lenOfS int = len(s)
	i := 0
	
	if (lenOfS >= 0) && (n <= lenOfS) {
		for  {
			if (i+n) <= lenOfS {
				output = append(output, s[i:i+n])
				i++
			} else {
				break
			}
		}
	}
	return output
}

// UnsafeFirst First of the all.
func UnsafeFirst(n int, s string) string {
	return s[:n]
}
