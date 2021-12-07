package main

import (
	_ "embed"
	"fmt"
	"math"
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
	return aocconv.StrToIntSlice(inputStr, aocconv.WithDelimeter(","))
}

func solveAnyPart(input []int, calcFuelUsedFromDist func(dist int) int) int {
	var max int
	for _, pos := range input {
		if pos > max {
			max = pos
		}
	}

	minFuelUsed := math.MaxInt
	for i := 0; i <= max; i++ {
		fuelUsed := 0
		for _, pos := range input {
			dist := int(math.Abs(float64(pos - i)))
			fuelUsed += calcFuelUsedFromDist(dist)
		}

		if fuelUsed < minFuelUsed {
			minFuelUsed = fuelUsed
		}
	}

	return minFuelUsed
}

func SolvePart1(input []int) int {
	return solveAnyPart(input, func(dist int) int { return dist })
}

func SolvePart2(input []int) int {
	return solveAnyPart(input, func(dist int) int { return (dist * (dist + 1)) / 2 })
}
