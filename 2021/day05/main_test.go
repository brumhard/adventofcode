package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	inputStr := `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

	input, err := inputFromString(inputStr)
	if err != nil {
		t.Errorf("failed to load inputs %v", err)
	}

	t.Run("Test Part1", func(t *testing.T) {
		got := SolvePart1(input)
		expected := 5

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test Part2", func(t *testing.T) {
		got := SolvePart2(input)
		expected := 12

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("CoveredCoordinates", func(t *testing.T) {
		t.Run("horizontal", func(t *testing.T) {
			line := Line{start: Coordinate{x: 1, y: 2}, end: Coordinate{x: 3, y: 2}}
			coords := line.CoveredCoordinates(false)

			if len(coords) != 3 {
				t.Errorf("should cover exactly 3 coordinates, but got %d", len(coords))
			}
		})

		t.Run("vertical", func(t *testing.T) {
			line := Line{start: Coordinate{x: 1, y: 4}, end: Coordinate{x: 1, y: 2}}
			coords := line.CoveredCoordinates(false)

			if len(coords) != 3 {
				t.Errorf("should cover exactly 3 coordinates, but got %d", len(coords))
			}
		})

		t.Run("diagonal", func(t *testing.T) {
			line := Line{start: Coordinate{x: 1, y: 1}, end: Coordinate{x: 3, y: 3}}
			coords := line.CoveredCoordinates(true)

			if len(coords) != 3 {
				t.Errorf("should cover exactly 3 coordinates, but got %d", len(coords))
			}
		})

		t.Run("diagonal 2", func(t *testing.T) {
			line := Line{start: Coordinate{x: 9, y: 7}, end: Coordinate{x: 7, y: 9}}
			coords := line.CoveredCoordinates(true)

			if len(coords) != 3 {
				t.Errorf("should cover exactly 3 coordinates, but got %d", len(coords))
			}
		})
	})
}
