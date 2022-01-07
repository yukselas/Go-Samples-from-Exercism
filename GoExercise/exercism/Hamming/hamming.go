package hamming

import "errors"

func Distance(a, b string) (int, error) {

	distancecount := 0

	//panic("Implement the Distance function")

	if len(a) != len(b) {

		return -1, errors.New("inputs must have the same length")

	}

	for i := range a {

		if b[i] != a[i] {

			distancecount++

		}

	}

	return distancecount, nil

}
