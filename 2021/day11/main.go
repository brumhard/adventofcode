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

type point struct {
	i, y int
}

func SolvePart1(input [][]int) int {
	flashes := 0
	for s := 0; s < 100; s++ {
		flashed := map[point]struct{}{}
		for i := range input {
			for y := range input[i] {
				flashes += bumpEnergyLevel(input, point{i, y}, flashed)
			}
		}
	}

	return flashes
}

func bumpEnergyLevel(input [][]int, pos point, flashed map[point]struct{}) int {
	i, y := pos.i, pos.y

	if _, ok := flashed[pos]; ok {
		return 0
	}
	if i < 0 || i > len(input)-1 || y < 0 || y > len(input[i])-1 {
		return 0
	}

	input[i][y]++

	if input[i][y] <= 9 {
		return 0
	}

	// set to zero
	input[i][y] = 0
	flashed[point{i, y}] = struct{}{}

	sorrounding := map[point]struct{}{
		{i - 1, y - 1}: {},
		{i - 1, y}:     {},
		{i - 1, y + 1}: {},
		{i, y - 1}:     {},
		{i, y + 1}:     {},
		{i + 1, y - 1}: {},
		{i + 1, y}:     {},
		{i + 1, y + 1}: {},
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
		flashed := map[point]struct{}{}
		for i := range input {
			for y := range input[i] {
				flashes += bumpEnergyLevel(input, point{i, y}, flashed)
			}
		}
		if flashes == len(input)*len(input[0]) {
			break
		}
		step++
	}

	return step
}
