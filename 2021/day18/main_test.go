package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("pairFromStr", func(t *testing.T) {
		tests := []struct {
			in       string
			expected pair
		}{
			{"[1,2]", pair{1, 2}},
			{"[[1,2],3]", pair{pair{1, 2}, 3}},
			{"[9,[8,7]]", pair{9, pair{8, 7}}},
			{"[[1,9],[8,5]]", pair{pair{1, 9}, pair{8, 5}}},
			{"[[[[1,2],[3,4]],[[5,6],[7,8]]],9]", pair{pair{pair{pair{1, 2}, pair{3, 4}}, pair{pair{5, 6}, pair{7, 8}}}, 9}},
		}

		for _, tt := range tests {
			t.Run(tt.in, func(t *testing.T) {
				got, err := pairFromStr(tt.in)
				if err != nil {
					t.Error(err)
				}

				if got != tt.expected {
					t.Errorf("expected %v, but got %v", tt.expected, got)
				}
			})
		}
	})

	t.Run("split", func(t *testing.T) {
		got := split(7)
		expected := pair{3, 4}

		if got != expected {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})

	t.Run("reduce", func(t *testing.T) {
		tests := []struct {
			in       string
			expected string
		}{
			{"[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]"},
			{"[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]"},
			{"[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]"},
			{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
			{"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"},
			{"[[[[0,7],4],[7,[[8,4],9]]],[1,1]]", "[[[[0,7],4],[15,[0,13]]],[1,1]]"},
			{"[[[[4,0],[5,4]],[[7,7],[6,0]]],[14,[[[3,7],[4,3]],[[6,3],[8,8]]]]]", "[[[[4,0],[5,4]],[[7,7],[6,0]]],[17,[[0,[11,3]],[[6,3],[8,8]]]]]"},
			{"[[[[4,0],[5,4]],[[7,7],[6,0]]],[17,[[11,0],[[9,3],[8,8]]]]]", "[[[[4,0],[5,4]],[[7,7],[6,0]]],[17,[[11,9],[0,[11,8]]]]]"},
			{"[[[[4,0],[5,4]],[[7,7],[6,0]]],[17,[[11,9],[11,0]]]]", "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,9],[[11,9],[11,0]]]]"},
		}

		for _, tt := range tests {
			t.Run(tt.in, func(t *testing.T) {
				p, _ := pairFromStr(tt.in)
				expected, _ := pairFromStr(tt.expected)

				got := reducePair(p)

				if got != expected {
					t.Errorf("expected %v, but got %v", expected, got)
				}
			})
		}
	})
	t.Run("repeatReduce", func(t *testing.T) {
		tests := []struct {
			in       string
			expected string
		}{
			{"[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"},
			{"[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]", "[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]"},
			{"[[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]],[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]]", "[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]]"},
			{"[[[[[4,0],[5,4]],[[7,7],[6,0]]],[[8,[7,7]],[[7,9],[5,0]]]],[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]]", "[[[[6,7],[6,7]],[[7,7],[0,7]]],[[[8,7],[7,7]],[[8,8],[8,0]]]]"},
		}

		for _, tt := range tests {
			t.Run(tt.in, func(t *testing.T) {
				p, _ := pairFromStr(tt.in)
				expected, _ := pairFromStr(tt.expected)

				got := repeatReduce(p)

				if got != expected {
					t.Errorf("expected %v, but got %v", expected, got)
				}
			})
		}
	})

	t.Run("sumList", func(t *testing.T) {
		inputStr := `[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
[7,[5,[[3,8],[1,4]]]]
[[2,[2,2]],[8,[8,1]]]
[2,9]
[1,[[[9,3],9],[[9,0],[0,7]]]]
[[[5,[7,4]],7],1]
[[[[4,2],2],6],[8,7]]`

		input, err := inputFromString(inputStr)
		if err != nil {
			t.Errorf("failed to load inputs %v", err)
		}

		got := sumList(input)
		expected, _ := pairFromStr("[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]")

		if got != expected {
			t.Errorf("expected %v, got %v", expected, got)
		}
	})

	t.Run("magnitude", func(t *testing.T) {
		tests := []struct {
			p        pair
			expected int
		}{
			{pair{pair{1, 2}, pair{pair{3, 4}, 5}}, 143},
			{pair{pair{pair{pair{8, 7}, pair{7, 7}}, pair{pair{8, 6}, pair{7, 7}}}, pair{pair{pair{0, 7}, pair{6, 6}}, pair{8, 7}}}, 3488},
		}

		for _, tt := range tests {
			t.Run(tt.p.String(), func(t *testing.T) {
				got := magnitude(tt.p)

				if got != tt.expected {
					t.Errorf("expected %v, got %v", tt.expected, got)
				}
			})
		}
	})

	inputStr := `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`

	input, err := inputFromString(inputStr)
	if err != nil {
		t.Errorf("failed to load inputs %v", err)
	}

	t.Run("Test Part1", func(t *testing.T) {
		got := SolvePart1(input)
		expected := 4140

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})

	t.Run("Test Part2", func(t *testing.T) {
		got := SolvePart2(input)
		expected := 3993

		if got != expected {
			t.Errorf("expected '%d' but got '%d'", expected, got)
		}
	})
}
