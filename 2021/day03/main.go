package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"

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

func inputFromString(inputStr string) ([][]int, error) {
	inputstrings := aocconv.StrToStrSlice(inputStr)
	input := make([][]int, 0, len(inputstrings))
	for _, str := range inputstrings {
		binaryInts := strings.Split(str, "")
		ints := make([]int, 0, len(binaryInts))

		for _, intAsStr := range binaryInts {
			bin, err := strconv.Atoi(intAsStr)
			if err != nil {
				return nil, err
			}

			ints = append(ints, bin)
		}

		input = append(input, ints)
	}

	return input, nil
}

func SolvePart1(input [][]int) int {
	gamma := make([]int, len(input[0]))
	epsilon := make([]int, len(input[0]))
	for y := range input[0] {
		sum := 0
		for i := range input {
			sum += input[i][y]
		}

		if sum > (len(input) / 2) {
			gamma[y] = 1
		} else {
			epsilon[y] = 1
		}
	}

	return BitSliceToNumber(gamma) * BitSliceToNumber(epsilon)
}

func BitSliceToNumber(slice []int) int {
	decimal := 0
	for i := 0; i < len(slice); i++ {
		decimal += slice[len(slice)-1-i] << i
	}

	return decimal
}

func SolvePart2(input [][]int) int {
	oxygen := filterBitSlices(input, true)
	co2 := filterBitSlices(input, false)

	return BitSliceToNumber(oxygen) * BitSliceToNumber(co2)
}

func filterBitSlices(input [][]int, mostCommon bool) []int {
	toKeep := input

	y := 0
	for len(toKeep) > 1 {
		sorted := map[int][][]int{}
		for i := range toKeep {
			sorted[toKeep[i][y]] = append(sorted[toKeep[i][y]], toKeep[i])
		}

		if mostCommon {
			toKeep = sorted[0]
			if len(sorted[1]) >= len(sorted[0]) {
				toKeep = sorted[1]
			}
		} else {
			toKeep = sorted[1]
			if len(sorted[1]) >= len(sorted[0]) {
				toKeep = sorted[0]
			}
		}

		y++
	}

	return toKeep[0]
}
