package secret

//Handshakes identifies the options
var Handshakes []string = []string{"wink", "double blink", "close your eyes", "jump"}

func reverse(a []string) {
	j := len(a)
	temp := ""
	for i, val := range a {
		if j%2 == 0 && i > (j/2-1) {
			break
		} else if j%2 != 0 && i == int(j/2) {
			break
		}

		temp = a[j-(i+1)]
		a[j-(i+1)] = val
		a[i] = temp
	}
}

//Handshake computes the secret handshake
func Handshake(input uint) (handshake []string) {
	handshake = make([]string, 0)
	if input > 0 {
		for i := 0; i < 4; i++ {
			if input&(1<<i) == (1 << i) {
				handshake = append(handshake, Handshakes[i])
			}
		}
	}

	if (input & (1 << 4)) == 16 {
		reverse(handshake[:])
	}

	return handshake
}
