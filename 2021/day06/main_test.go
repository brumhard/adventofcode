package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	inputStr := `3,4,3,1,2`

	input, err := inputFromString(inputStr)
	if err != nil {
		t.Errorf("failed to load inputs %v", err)
	}

	t.Run("Test Part1", func(t *testing.T) {
		got := SolvePart1(input)
		expected := 5934

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test Part2", func(t *testing.T) {
		got := SolvePart2(input)
		expected := 26984457539

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("calcProduced", func(t *testing.T) {
		tests := []struct {
			name                             string
			daysleft, initialTimer, expected int
		}{
			{
				name:         "no time to reproduce",
				daysleft:     0,
				initialTimer: 0,
				expected:     1,
			},
			{
				name:         "time to reproduce",
				daysleft:     1,
				initialTimer: 0,
				expected:     2,
			},
			{
				name:         "time to reproduce 2",
				daysleft:     2,
				initialTimer: 1,
				expected:     2,
			},
			{
				name:         "no time to reproduce twice",
				daysleft:     8,
				initialTimer: 1,
				expected:     2,
			},
			{
				name:         "time to reproduce twice",
				daysleft:     9,
				initialTimer: 1,
				expected:     3,
			},
			{
				name:         "no time to reproduce three times",
				daysleft:     17,
				initialTimer: 8,
				expected:     3,
			},
			{
				name:         "time to reproduce three times",
				daysleft:     18,
				initialTimer: 8,
				expected:     4,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				got := calcProduced(tt.daysleft, tt.initialTimer, nil)

				if got != tt.expected {
					t.Errorf("expected '%d' but got '%d'", tt.expected, got)
				}
			})
		}
	})
}
