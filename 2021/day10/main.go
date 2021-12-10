package main

import (
	_ "embed"
	"fmt"
	"os"
	"sort"

	"github.com/brumhard/adventofcode/aocconv"
)

//go:embed input.txt
var inputFile string

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "an error occured: %s\n", err.Error())
		os.Exit(1)
	}
}

func run() error {
	input, err := inputFromString(string(inputFile))
	if err != nil {
		return err
	}

	fmt.Printf("Part 1: %v\n", SolvePart1(input))
	fmt.Printf("Part 2: %v\n", SolvePart2(input))

	return nil
}

func inputFromString(inputStr string) ([]string, error) {
	return aocconv.StrToStrSlice(inputStr), nil
}

func isOpening(char rune) bool {
	switch char {
	case '(', '[', '{', '<':
		return true
	}

	return false
}

func closing(opening rune) rune {
	switch opening {
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	case '<':
		return '>'
	}

	panic("invalid sign")
}

type stack []rune

func (s *stack) Push(element rune) {
	*s = append(*s, element)
}

func (s *stack) Pop() rune {
	n := len(*s) - 1
	element := (*s)[n]
	*s = (*s)[:n]

	return element
}

var scoreMap1 = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func SolvePart1(input []string) int {
	sum := 0
	for _, line := range input {
		openingStack := stack{}
		for _, char := range line {
			if isOpening(char) {
				openingStack.Push(char)
				continue
			}

			if !(closing(openingStack.Pop()) == char) {
				sum += scoreMap1[char]
			}
		}
	}

	return sum
}

var scoreMap2 = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func SolvePart2(input []string) int {
	var scores []int
	for _, line := range input {
		openingStack := stack{}
		corrupted := false
		for _, char := range line {
			if isOpening(char) {
				openingStack.Push(char)
				continue
			}

			if !(closing(openingStack.Pop()) == char) {
				corrupted = true
				break
			}
		}

		if corrupted {
			continue
		}

		var fixRunes []rune
		for len(openingStack) > 0 {
			fixRunes = append(fixRunes, closing(openingStack.Pop()))
		}

		score := 0
		for _, char := range fixRunes {
			score *= 5
			score += scoreMap2[char]
		}
		scores = append(scores, score)
	}

	sort.Ints(scores)

	return scores[(len(scores)-1)/2]
}
