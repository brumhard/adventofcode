package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	inputStr := `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`

	input, err := inputFromString(inputStr)
	if err != nil {
		t.Errorf("failed to load inputs %v", err)
	}

	t.Run("Test Part12", func(t *testing.T) {
		got := SolvePart12(nodesFromMatrix(input))
		expected := 40

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test Part2", func(t *testing.T) {
		got := SolvePart12(nodesFromMatrix(input2FromMatrix(input)))
		expected := 315

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}
