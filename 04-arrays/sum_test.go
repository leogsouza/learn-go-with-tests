package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got '%d' want '%d' given, %v", got, want, numbers)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2}, []int{3, 9})

	want := []int{2, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
