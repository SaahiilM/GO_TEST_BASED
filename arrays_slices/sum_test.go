package arraysslices

import (
	"testing"
)

func TestSum(t *testing.T) {
	numbers := [5]int{1, 3, 5, 7, 9}

	got := Sum(numbers)
	want := 25

	if got != want {
		t.Errorf("got %d, want %d, given %v", got, want, numbers)
	}
}
