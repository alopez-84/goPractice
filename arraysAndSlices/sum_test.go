package main

import (
	"reflect"
	"testing"
)

func TestSlice(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		received := Sum(numbers)
		expected := 6

		if received != expected {
			t.Errorf("Wanted %d, got %d, given: %v", expected, received, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("compute sume of 2 Slices", func(t *testing.T) {

		received := SumAll([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}

		if !reflect.DeepEqual(received, expected) {
			t.Errorf("Wanted %v, got %v", expected, received)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	check := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Wanted %v, got %v", want, got)
		}
	}
	t.Run("Compute the sum of all tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		check(t, got, want)
	})

	t.Run("Safely run empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		check(t, got, want)
	})
}
