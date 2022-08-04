package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}

func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := Sum(numbers)
	fmt.Println(sum)
	// Output: 15
}

func BenchmarkSum(b *testing.B) {
	numbers := []int{100, 200, 300, 400, 500}
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func ExampleSumAll() {
	sums := SumAll([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9})
	fmt.Println(sums)
	// Output: [6 15 24]
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9})
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9})
	want := []int{5, 11, 17}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
