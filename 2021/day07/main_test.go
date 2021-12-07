package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	inputStr := `16,1,2,0,4,2,7,1,2,14`

	input, err := inputFromString(inputStr)
	if err != nil {
		t.Errorf("failed to load inputs %v", err)
	}

	t.Run("Test Part1", func(t *testing.T) {
		got := SolvePart1(input)
		expected := 37

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test Part2", func(t *testing.T) {
		got := SolvePart2(input)
		expected := 168

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}
