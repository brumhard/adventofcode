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

func SolvePart1(input []int) int {
	var max int
	for _, pos := range input {
		if pos > max {
			max = pos
		}
	}

	minFuelUsed := 0

outer:
	for i := 0; i <= max; i++ {
		fuelUsed := 0
		for _, pos := range input {
			fuelUsed += int(math.Abs(float64(pos - i)))
			if i != 0 && fuelUsed > minFuelUsed {
				continue outer
			}
		}

		minFuelUsed = fuelUsed
	}

	return minFuelUsed
}

func SolvePart2(input []int) int {
	var max int
	for _, pos := range input {
		if pos > max {
			max = pos
		}
	}

	minFuelUsed := 0

outer:
	for i := 0; i <= max; i++ {
		fuelUsed := 0
		for _, pos := range input {
			dist := int(math.Abs(float64(pos - i)))
			fuelUsed += (dist * (dist + 1)) / 2 // gauss formula for sum of 1+...+n-1+n
			if i != 0 && fuelUsed > minFuelUsed {
				continue outer
			}
		}

		minFuelUsed = fuelUsed
	}

	return minFuelUsed
}
