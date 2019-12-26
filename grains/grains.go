package grains

import (
	"errors"
	"math"
)


func Total() (sum uint64) {
	for i := 1; i< 65; i++ {
		sum += uint64(math.Pow(2, float64(i - 1)))
	}
	return sum
}

func Square(squares int) (grains uint64, err error) {
	if squares <= 0 || squares > 64{
		return 0, errors.New("invalid input")
	}

	
	grains = uint64(1 * (math.Pow(2, float64(squares - 1))))

	return grains, nil
}