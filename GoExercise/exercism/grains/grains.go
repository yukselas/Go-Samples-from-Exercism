package grains

import "errors"

const boardSize = 64

//n << x is "n times 2, x times". And y >> z is "y divided by 2, z times

func Square(number int) (uint64, error) {
	if number < 1 || number > boardSize {
		return 0, errors.New("square umbrt out of range")
	}

	return 1 << (number - 1), nil
}

func Total() uint64 {
	return (1 << boardSize) - 1
}
