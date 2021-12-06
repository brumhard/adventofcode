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

	t.Run("calcProduced", func(t *testing.T) {
		t.Run("test", func(t *testing.T) {
			got := calcProduced(0, 0)
			expected := 1

			if got != expected {
				t.Errorf("expected '%d' but got '%d'", expected, got)
			}
		})
		t.Run("test2", func(t *testing.T) {
			got := calcProduced(1, 0)
			expected := 2

			if got != expected {
				t.Errorf("expected '%d' but got '%d'", expected, got)
			}
		})
		t.Run("test22", func(t *testing.T) {
			got := calcProduced(2, 1)
			expected := 2

			if got != expected {
				t.Errorf("expected '%d' but got '%d'", expected, got)
			}
		})
		t.Run("test3", func(t *testing.T) {
			got := calcProduced(4, 3)
			expected := 2

			if got != expected {
				t.Errorf("expected '%d' but got '%d'", expected, got)
			}
		})
		t.Run("test4", func(t *testing.T) {
			got := calcProduced(8, 1)
			expected := 2

			if got != expected {
				t.Errorf("expected '%d' but got '%d'", expected, got)
			}
		})
		t.Run("test5", func(t *testing.T) {
			got := calcProduced(9, 1)
			expected := 3

			if got != expected {
				t.Errorf("expected '%d' but got '%d'", expected, got)
			}
		})
		t.Run("test6", func(t *testing.T) {
			got := calcProduced(17, 8)
			expected := 3

			if got != expected {
				t.Errorf("expected '%d' but got '%d'", expected, got)
			}
		})
		t.Run("test7", func(t *testing.T) {
			got := calcProduced(18, 8)
			expected := 4

			if got != expected {
				t.Errorf("expected '%d' but got '%d'", expected, got)
			}
		})
	})

	t.Run("Test Part2", func(t *testing.T) {
		got := SolvePart2(input)
		expected := 26984457539

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}
