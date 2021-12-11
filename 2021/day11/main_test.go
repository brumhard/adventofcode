package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	inputStr := `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

	t.Run("Test Part1", func(t *testing.T) {
		input, err := inputFromString(inputStr)
		if err != nil {
			t.Errorf("failed to load inputs %v", err)
		}

		got := SolvePart1(input)
		expected := 1656

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test Part2", func(t *testing.T) {
		input, err := inputFromString(inputStr)
		if err != nil {
			t.Errorf("failed to load inputs %v", err)
		}

		got := SolvePart2(input)
		expected := 195

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}
