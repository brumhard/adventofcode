package main

import (
	_ "embed"
	"fmt"
	"os"

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
	input, err := inputFromString(inputFile)
	if err != nil {
		return err
	}

	fmt.Printf("Part 1: %v\n", SolvePart1(input))
	fmt.Printf("Part 2: %v\n", SolvePart2(input))

	return nil
}

func inputFromString(inputStr string) ([]int, error) {
	return aocconv.StrToIntSlice(inputStr)
}

func SolvePart1(input []int) int {
	increasedCount := 0
	for i := range input {
		if i == 0 {
			continue
		}

		if input[i] > input[i-1] {
			increasedCount++
		}
	}

	return increasedCount
}

func SolvePart2(input []int) int {
	increasedCount := 0
	for i := range input {
		if i+2 >= len(input)-1 {
			break
		}

		if sumFromIndex(input, i) < sumFromIndex(input, i+1) {
			increasedCount++
		}
	}

	return increasedCount
}

func sumFromIndex(inputs []int, i int) int {
	return inputs[i] + inputs[i+1] + inputs[i+2]
}
