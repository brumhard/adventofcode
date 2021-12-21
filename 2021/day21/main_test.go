package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	input := input{
		p1: 4,
		p2: 8,
	}

	t.Run("Test Part1", func(t *testing.T) {
		got := SolvePart1(input)
		expected := 739785

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test Part2", func(t *testing.T) {
		got := SolvePart2(input)
		expected := 444356092776315

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}
