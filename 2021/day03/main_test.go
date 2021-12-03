package main

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	inputStr := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

	input, err := inputFromString(inputStr)
	if err != nil {
		t.Errorf("failed to load inputs %v", err)
	}

	t.Run("Test Part1", func(t *testing.T) {
		got := SolvePart1(input)
		expected := 198

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test Part2", func(t *testing.T) {
		got := SolvePart2(input)
		expected := 230

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("BitSliceToNumber", func(t *testing.T) {
		got := BitSliceToNumber([]int{1, 1, 0})
		expected := 6

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("filterSliceByKeep true", func(t *testing.T) {
		got := filterBitSlices(input, true)
		expected := []int{1, 0, 1, 1, 1}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})

	t.Run("filterSliceByKeep false", func(t *testing.T) {
		got := filterBitSlices(input, false)
		expected := []int{0, 1, 0, 1, 0}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("expected '%v' but got '%v'", expected, got)
		}
	})
}
