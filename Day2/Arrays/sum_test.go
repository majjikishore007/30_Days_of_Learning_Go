package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assertCorrect := func(t testing.TB, got, want int, nums []int) {
		t.Helper()
		if got != want {
			t.Errorf("got %d  and want %d nums %v", got, want, nums)
		}
	}
	t.Run("Collection of size 5", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15
		assertCorrect(t, got, want, numbers)

	})
	t.Run("Dynamic sizes", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6
		assertCorrect(t, got, want, numbers)

	})
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	// want := "kishore"

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
