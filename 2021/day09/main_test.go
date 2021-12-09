package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	inputStr := `2199943210
3987894921
9856789892
8767896789
9899965678`

	input, err := inputFromString(inputStr)
	if err != nil {
		t.Errorf("failed to load inputs %v", err)
	}

	t.Run("Test Part1", func(t *testing.T) {
		got := SolvePart1(input)
		expected := 15

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test Part2", func(t *testing.T) {
		got := SolvePart2(input)
		expected := 1134

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}
