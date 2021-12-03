package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"

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

	// fmt.Println(input)

	fmt.Printf("Part 1: %v\n", SolvePart1(input))
	fmt.Printf("Part 2: %v\n", SolvePart2(input))

	return nil
}

func inputFromString(inputStr string) ([]int, error) {
	inputstrings := aocconv.StrToStrSlice(inputStr)
	input := make([]int, 0, len(inputstrings))
	for _, str := range inputstrings {
		binary, err := strconv.ParseInt(str, 2, 0)
		if err != nil {
			return nil, err
		}

		input = append(input, int(binary))
	}

	return input, nil
}

// getNthBit uses a bitmask to find the nth bit
func getNthBit(binary int, n int) int {
	masked := binary & (1 << n)
	if masked > 0 {
		return 1
	}
	return 0
}

func SolvePart1(input []int) int {
	var gamma, epsilon int
	// hardcode 20 since it's hard to find the length of the input binary number
	for y := 20; y >= 0; y-- {
		sum := 0
		for _, binary := range input {
			sum += getNthBit(binary, y)
		}

		if sum == 0 {
			continue
		}

		if sum > (len(input) / 2) {
			gamma += 1 << y
		} else {
			epsilon += 1 << y
		}
	}

	return gamma * epsilon
}

func SolvePart2(input []int) int {
	oxygen := filterBitSlices(input, 0)
	co2 := filterBitSlices(input, 1)

	return oxygen * co2
}

func filterBitSlices(input []int, prefer int) int {
	toKeep := input

	y := 20
	for len(toKeep) > 1 {
		sum := 0
		sorted := map[int][]int{}
		for _, binary := range toKeep {
			bit := getNthBit(binary, y)
			sum += bit
			sorted[bit] = append(sorted[bit], binary)
		}

		if sum == 0 {
			y--
			continue
		}

		// this logic hardcodes that prefer==0 means that least common bit is preffered and the other way around
		toKeep = sorted[prefer^1]
		// if one is the most common bit or both counts are the same take the preferred bits
		if len(sorted[1]) >= len(sorted[0]) {
			toKeep = sorted[prefer]
		}

		y--
	}

	return toKeep[0]
}
