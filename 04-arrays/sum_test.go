package arrays

import (
	"fmt"
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

	checkSums := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 9})

		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})

}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum([]int{1, 2, 3, 4, 5, 6})
	}
}

func BenchmarkSumAllTails(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAllTails([]int{2, 5, 7}, []int{8, 3, 9})
	}
}

func ExampleSum() {
	sum := Sum([]int{2, 3, 5})

	fmt.Println(sum)
	// Output: 10
}

func ExampleSumAllTails() {
	sums := SumAllTails([]int{}, []int{2, 4, 5}, []int{1, 0, 2, 5})

	fmt.Printf("%v", sums)
	// Output: [0 9 7]

}
