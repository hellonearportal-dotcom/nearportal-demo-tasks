package minmax

import (
	"math"
	"reflect"
	"testing"
)

func TestMinMaxIntExtremes(t *testing.T) {
	// math.MinInt and math.MaxInt are the extreme representable int values
	values := []int{math.MinInt, 0, math.MaxInt}
	gotMin, gotMax, err := MinMax(values)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMin != math.MinInt {
		t.Errorf("min: got %d, want %d", gotMin, math.MinInt)
	}
	if gotMax != math.MaxInt {
		t.Errorf("max: got %d, want %d", gotMax, math.MaxInt)
	}
}

func TestMinMaxAllZeros(t *testing.T) {
	// A slice containing only zeros: min and max must both be 0
	values := []int{0, 0, 0, 0}
	gotMin, gotMax, err := MinMax(values)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotMin != 0 || gotMax != 0 {
		t.Errorf("all-zeros: got min=%d, max=%d; want 0, 0", gotMin, gotMax)
	}
}

func TestMinMaxInputUnchanged(t *testing.T) {
	// Verify that MinMax leaves the input slice unchanged
	original := []int{5, -3, 8, 1, -7, 0}
	before := append([]int(nil), original...)
	_, _, err := MinMax(original)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !reflect.DeepEqual(original, before) {
		t.Errorf("input was modified: before %v, after %v", before, original)
	}
}
