package minmax

import (
	"math"
	"reflect"
	"testing"
)

func TestMinMaxIntExtremes(t *testing.T) {
	cases := []struct {
		values   []int
		min, max int
	}{
		{[]int{math.MinInt}, math.MinInt, math.MinInt},
		{[]int{math.MaxInt}, math.MaxInt, math.MaxInt},
		{[]int{math.MinInt, math.MaxInt}, math.MinInt, math.MaxInt},
		{[]int{math.MaxInt, math.MinInt}, math.MinInt, math.MaxInt},
		{[]int{0, math.MinInt, math.MaxInt, 0}, math.MinInt, math.MaxInt},
		{[]int{math.MinInt + 1, math.MinInt, math.MaxInt - 1, math.MaxInt}, math.MinInt, math.MaxInt},
	}
	for _, tc := range cases {
		input := append([]int(nil), tc.values...)
		before := append([]int(nil), input...)
		min, max, err := MinMax(input)
		if err != nil || min != tc.min || max != tc.max {
			t.Errorf("MinMax(%v) = %d, %d, %v; want %d, %d, nil", tc.values, min, max, err, tc.min, tc.max)
		}
		if !reflect.DeepEqual(input, before) {
			t.Errorf("input modified: before %v, after %v", before, input)
		}
	}
}

func TestMinMaxAllZeros(t *testing.T) {
	for _, input := range [][]int{{0}, {0, 0}, {0, 0, 0, 0, 0}} {
		in := append([]int(nil), input...)
		before := append([]int(nil), in...)
		min, max, err := MinMax(in)
		if err != nil || min != 0 || max != 0 {
			t.Errorf("MinMax(%v) = %d, %d, %v; want 0, 0, nil", input, min, max, err)
		}
		if !reflect.DeepEqual(in, before) {
			t.Errorf("input modified: before %v, after %v", before, in)
		}
	}
}

func TestMinMaxInputUnchanged(t *testing.T) {
	input := []int{5, -7, 3, 9, -1, 0, 9, -7}
	want := append([]int(nil), input...)
	min, max, err := MinMax(input)
	if err != nil || min != -7 || max != 9 {
		t.Errorf("MinMax = %d, %d, %v; want -7, 9, nil", min, max, err)
	}
	if !reflect.DeepEqual(input, want) {
		t.Errorf("input modified: before %v, after %v", want, input)
	}
}
