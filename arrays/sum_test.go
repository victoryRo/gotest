package arrays

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

// INFO: go test -cover
// prueba la covertura del codigo en %

func TestSum(t *testing.T) {
	t.Run("collection of any sizes", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got: %d want: %d given: %v", got, want, numbers)
		}
	})
}

func ExampleSum() {
	sum := Sum([]int{7, 7, 7})
	fmt.Println(sum)

	// Output: 21
}

func BenchmarkSum(b *testing.B) {
	for b.Loop() {
		Sum([]int{7, 7, 7})
	}
}

// --------------------------------------------------------------

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !slices.Equal(got, want) {
		t.Errorf("got: %v want: %v", got, want)
	}
}

func ExampleSumAll() {
	sum := SumAll([]int{3, 3, 3}, []int{10, 10, 10})
	fmt.Println(sum)

	// Output: [9 30]
}

func BenchmarkSumAll(b *testing.B) {
	for b.Loop() {
		SumAll([]int{3, 3, 3}, []int{10, 10, 10})
	}
}

// --------------------------------------------------------------

func TestSumAllTails(t *testing.T) {
	// funcion auxiliar dentro del test
	// se usara unicamente dentro del test
	checkSum := func(t testing.TB, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v want: %v", got, want)
		}
	}

	t.Run("make the sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSum(t, got, want)
	})

	t.Run("safely sum empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSum(t, got, want)
	})
}

func ExampleSumAllTails() {
	sum := SumAllTails([]int{1, 3, 3}, []int{2, 4, 4}, []int{3, 5, 5})
	fmt.Println(sum)

	// Output: [6 8 10]
}

func BenchmarkSumAllTails(b *testing.B) {
	for b.Loop() {
		SumAllTails([]int{1, 3, 3}, []int{2, 4, 4}, []int{3, 5, 5})
	}
}

// --------------------------------------------------------------
