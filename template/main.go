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
	input, err := inputFromString(string(inputFile))
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
	var solution int

	return solution
}

func SolvePart2(input []int) int {
	var solution int

	return solution
}
