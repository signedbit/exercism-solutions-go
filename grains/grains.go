package grains

import (
	"errors"
)

func Square(n int) (uint64, error) {
	if n <= 0 || n > 64 {
		return 0, errors.New("invalid")
	}
	return 1 << ((n) - 1), nil
}

func Total() uint64 {
	total, _ := Square(64)
	return total*2 - 1
}
