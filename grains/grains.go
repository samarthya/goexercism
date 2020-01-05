package grains

import (
	"errors"
	// "math"
	"fmt"
)


func Total() (sum uint64) {
	sum = 1
	fmt.Println(" Total: ", (sum << 64))
	return (sum << 64) - 1
}

func Square(squares int) (grains uint64, err error) {
	if squares <= 0 || squares > 64 {
		return 0, errors.New("invalid input")
	}

	grains = 1
	//grains = uint64(1 * (math.Pow(2, float64(squares - 1))))

	return grains << (squares - 1), nil
}