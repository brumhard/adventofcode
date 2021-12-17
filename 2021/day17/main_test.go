package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	inputStr := `target area: x=20..30, y=-10..-5
`

	input, err := inputFromString(inputStr)
	if err != nil {
		t.Errorf("failed to load inputs %v", err)
	}

	t.Run("iterate", func(t *testing.T) {
		tests := []struct {
			xVel, yVel   int
			expectHit    bool
			expectHeight int
		}{
			{xVel: 6, yVel: 9, expectHit: true, expectHeight: 45},
			{xVel: 6, yVel: 3, expectHit: true, expectHeight: 6},
			{xVel: 9, yVel: 0, expectHit: true, expectHeight: 0},
			{xVel: 17, yVel: -4, expectHit: false},
		}

		target := target{
			xMin: 20, xMax: 30, yMin: -10, yMax: -5,
		}
		for _, tt := range tests {
			t.Run(fmt.Sprintf("xVel: %d,yVel: %d", tt.xVel, tt.yVel), func(t *testing.T) {
				maxHeight := iterate(target, tt.xVel, tt.yVel)

				if tt.expectHit != (maxHeight != nil) {
					t.Errorf("expected to hit: '%v', but didn't work", tt.expectHit)
				}

				if maxHeight == nil {
					return
				}

				if *maxHeight != tt.expectHeight {
					t.Errorf("expected '%d' but got '%d'", tt.expectHeight, *maxHeight)
				}
			})
		}
	})

	t.Run("Test Part1", func(t *testing.T) {
		got := SolvePart1(input)
		expected := 45

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test Part2", func(t *testing.T) {
		got := SolvePart2(input)
		expected := 112

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}
