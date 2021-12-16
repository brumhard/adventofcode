package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test Part1", func(t *testing.T) {
		tests := []struct {
			in          string
			expectedSum int
		}{
			{"8A004A801A8002F478", 16},
			{"620080001611562C8802118E34", 12},
			{"C0015000016115A2E0802F182340", 23},
			{"A0016C880162017C3686B18A3D4780", 31},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("Test in %s", tt.in), func(t *testing.T) {
				input, err := inputFromString(tt.in)
				if err != nil {
					t.Errorf("failed to load inputs %v", err)
				}

				got := SolvePart1(input)
				if got != tt.expectedSum {
					t.Errorf("expected '%d' but got '%d'", tt.expectedSum, got)
				}
			})
		}
	})

	t.Run("Test Part2", func(t *testing.T) {
		tests := []struct {
			in          string
			expectedSum int
		}{
			{"C200B40A82", 3},
			{"04005AC33890", 54},
			{"880086C3E88112", 7},
			{"CE00C43D881120", 9},
			{"D8005AC2A8F0", 1},
			{"F600BC2D8F", 0},
			{"9C005AC2F8F0", 0},
			{"9C0141080250320F1802104A08", 1},
		}

		for _, tt := range tests {
			t.Run(fmt.Sprintf("Test in %s", tt.in), func(t *testing.T) {
				input, err := inputFromString(tt.in)
				if err != nil {
					t.Errorf("failed to load inputs %v", err)
				}

				got := SolvePart2(input)
				if got != tt.expectedSum {
					t.Errorf("expected '%d' but got '%d'", tt.expectedSum, got)
				}
			})
		}
	})
}
