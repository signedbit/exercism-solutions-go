package hamming

import "errors"

func Distance(a, b string) (int, error) {
	r1 := []rune(a)
	r2 := []rune(b)

	if len(r1) != len(r2) {
		return 0, errors.New("different lengths")
	}

	dist := 0
	for i, c := range r1 {
		if r2[i] != c {
			dist++
		}
	}
	return dist, nil
}
