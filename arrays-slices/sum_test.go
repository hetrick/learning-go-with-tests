package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	result := Sum(numbers)
	expected := 15

	if result != expected {
		t.Errorf("got %d, but expected %d, given %v", result, expected, numbers)
	}
}

func TestSumAll(t *testing.T) {
	result := SumAll([]int{1, 2, 3}, []int{4, 5, 6})
	expected := []int{6, 15}

	if !slices.Equal(result, expected) {
		t.Errorf("got %v, but expected %v", result, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, result, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("got %v, but expected %v", result, expected)
		}
	}
	t.Run("sum tails", func(t *testing.T) {
		checkSums(t, SumAllTails([]int{1, 2}, []int{3, 4, 5}), []int{2, 9})
	})
	t.Run("sum tails with empty slice", func(t *testing.T) {
		checkSums(t, SumAllTails([]int{}, []int{1, 2, 3}), []int{0, 6})
	})
}
