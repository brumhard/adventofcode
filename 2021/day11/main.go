package main

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/brumhard/adventofcode/aocconv"
	"github.com/brumhard/adventofcode/coords"
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

	input2 := make([][]int, len(input))
	for i := range input {
		input2[i] = make([]int, len(input[i]))
		copy(input2[i], input[i])
	}

	fmt.Printf("Part 1: %v\n", SolvePart1(input))
	fmt.Printf("Part 2: %v\n", SolvePart2(input2))

	return nil
}

func inputFromString(inputStr string) ([][]int, error) {
	lines := aocconv.StrToStrSlice(inputStr)
	input := make([][]int, 0, len(lines))

	for _, line := range lines {
		intsInLine, err := aocconv.StrToIntSlice(line, aocconv.WithDelimeter(""))
		if err != nil {
			return nil, err
		}

		input = append(input, intsInLine)
	}

	return input, nil
}

func SolvePart1(input [][]int) int {
	flashes := 0
	for s := 0; s < 100; s++ {
		flashed := map[coords.Point]struct{}{}
		for y := range input {
			for x := range input[y] {
				flashes += bumpEnergyLevel(input, coords.Point{Y: y, X: x}, flashed)
			}
		}
	}

	return flashes
}

func bumpEnergyLevel(input [][]int, pos coords.Point, flashed map[coords.Point]struct{}) int {
	y, x := pos.Y, pos.X

	if _, ok := flashed[pos]; ok {
		return 0
	}
	if y < 0 || y > len(input)-1 || x < 0 || x > len(input[y])-1 {
		return 0
	}

	input[y][x]++

	if input[y][x] <= 9 {
		return 0
	}

	// set to zero
	input[y][x] = 0
	flashed[coords.Point{Y: y, X: x}] = struct{}{}

	sorrounding := map[coords.Point]struct{}{
		{Y: y - 1, X: x - 1}: {},
		{Y: y - 1, X: x}:     {},
		{Y: y - 1, X: x + 1}: {},
		{Y: y, X: x - 1}:     {},
		{Y: y, X: x + 1}:     {},
		{Y: y + 1, X: x - 1}: {},
		{Y: y + 1, X: x}:     {},
		{Y: y + 1, X: x + 1}: {},
	}

	flashes := 1
	for p := range sorrounding {
		flashes += bumpEnergyLevel(input, p, flashed)
	}

	return flashes
}

func SolvePart2(input [][]int) int {
	step := 1
	for {
		flashes := 0
		flashed := map[coords.Point]struct{}{}
		for y := range input {
			for x := range input[y] {
				flashes += bumpEnergyLevel(input, coords.Point{Y: y, X: x}, flashed)
			}
		}
		if flashes == len(input)*len(input[0]) {
			break
		}
		step++
	}

	return step
}
