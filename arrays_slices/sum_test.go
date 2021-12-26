package arrays_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}
		got := Sum(numbers)
		want := 10

		if got != want {
			t.Errorf("got: %d, want %d, given %v", got, want, numbers)
		}

	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 3}, []int{0, 2})
	want := []int{4, 2}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSumTrail(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want %v", got, want)
		}
	}

	t.Run("get sum of some slices", func(t *testing.T) {
		got := SumAllTrails([]int{1, 3}, []int{0, 9})
		want := []int{3, 9}
		checkSums(t, got, want)
	})

	t.Run("get sum of empty slice", func(t *testing.T) {
		got := SumAllTrails([]int{}, []int{2, 3, 4})
		want := []int{0, 7}
		checkSums(t, got, want)
	})

}
