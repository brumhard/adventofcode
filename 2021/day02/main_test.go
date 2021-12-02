package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	inputStr := `forward 5
down 5
forward 8
up 3
down 8
forward 2`

	input, err := inputFromString(inputStr)
	if err != nil {
		t.Errorf("failed to load inputs %v", err)
	}

	fmt.Println(input)

	t.Run("Test Part1", func(t *testing.T) {
		got := SolvePart1(input)
		expected := 150

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test Part2", func(t *testing.T) {
		got := SolvePart2(input)
		expected := 900

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}
