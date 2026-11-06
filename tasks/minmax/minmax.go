package minmax

import "errors"

// MinMax returns the smallest and largest values without modifying values.
// For empty or nil input, return 0, 0 and a non-nil error.
func MinMax(values []int) (int, int, error) {
	if len(values) == 0 {
		return 0, 0, errors.New("minmax: empty input")
	}
	min, max := values[0], values[0]
	for _, v := range values[1:] {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max, nil
}
