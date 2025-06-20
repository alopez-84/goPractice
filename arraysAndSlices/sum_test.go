package main

import (
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
