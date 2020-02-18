package rotationalcipher

// RotationalCipher Rotate the string with the cipher seed transition.
func RotationalCipher(s string, in int) string {
	out := make([]rune, 0)

	if s == "" || in == 0 {
		return s
	}

	rin := rune(in)

	for _, r := range s {
		switch {
		case r >= 'A' && r <= 'Z':
			out = append(out, 'A'+(r-'A'+rin)%26)
		case r >= 'a' && r <= 'z':
			out = append(out, 'a'+(r-'a'+rin)%26)
		default:
			out = append(out, r)
		}
	}

	return string(out)
}
