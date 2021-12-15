package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	inputStr := `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`

	input, err := inputFromString(inputStr)
	if err != nil {
		t.Errorf("failed to load inputs %v", err)
	}

	t.Run("Test Part1", func(t *testing.T) {
		got := SolvePart1(input)
		expected := 1588

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test Part2", func(t *testing.T) {
		got := SolvePart2(input)
		expected := 2188189693529

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("NB", func(t *testing.T) {
		tests := []struct {
			double     string
			iterations int
			expected   map[rune]int
		}{
			{"NB", 1, map[rune]int{'B': 1}},
			{"NB", 2, map[rune]int{'B': 2, 'N': 1}},
			{"NB", 3, map[rune]int{'B': 5, 'N': 2}},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("%s with %d iterations", tt.double, tt.iterations), func(t *testing.T) {
				got := solvePair(tt.double, input.insertionMap, tt.iterations)

				if !reflect.DeepEqual(got, tt.expected) {
					t.Errorf("expected %v but got %v", tt.expected, got)
				}
			})
		}
	})
}
