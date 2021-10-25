package lasagna

import (
	"log"
	"strings"
)

const (
	Fifty = 50
	Sauce = .2
)

// 'PreparationTime()' function
func PreparationTime(l []string, t int) int {
	if t == 0 {
		t = 2
	}
	return len(l) * t
}

// 'Quantities()' function
func Quantities(l []string) (q int, r float64) {
	for _, v := range l {

		if strings.Compare(v, "noodles") == 0 {
			q += Fifty
		}

		if strings.Compare(v, "sauce") == 0 {
			r += Sauce
		}
	}

	return
}

// AddSecretIngredient(): The function should generate a new slice
// and add the last item from your friends list to the end of your list.
// Neither argument should be modified.
func AddSecretIngredient(a, b []string) (r []string) {
	r = b
	r = append(r, a[len(a)-1])
	log.Println(" Elements >", r)
	return
}

// ScaleRecipe : should return a slice of float64 of the amounts needed for the desired number of portions.
func ScaleRecipe(s []float64, n int) (r []float64) {
	r = make([]float64, len(s))

	log.Println(s, n)

	m := 2.0000

	if n != 2 {
		m = float64(n) / m
	}

	for i, v := range s {
		r[i] = v * m
	}

	return
}
