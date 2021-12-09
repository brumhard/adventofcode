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

func inputFromString(inputStr string) ([][]int, error) {
	rows := aocconv.StrToStrSlice(inputStr)

	input := make([][]int, 0, len(rows))

	for _, row := range rows {
		inputRow, err := aocconv.StrToIntSlice(row, aocconv.WithDelimeter(""))
		if err != nil {
			return nil, err
		}

		input = append(input, inputRow)
	}

	return input, nil
}

func SolvePart1(input [][]int) int {
	var lowPoints []int

	for i := range input {
		for y := range input[i] {
			if i-1 >= 0 && input[i-1][y] <= input[i][y] {
				continue
			}
			if i+1 <= len(input)-1 && input[i+1][y] <= input[i][y] {
				continue
			}
			if y-1 >= 0 && input[i][y-1] <= input[i][y] {
				continue
			}
			if y+1 <= len(input[i])-1 && input[i][y+1] <= input[i][y] {
				continue
			}
			lowPoints = append(lowPoints, input[i][y])
		}
	}

	riskLevel := 0
	for _, p := range lowPoints {
		riskLevel += p + 1
	}

	return riskLevel
}

type point struct {
	i, y int
}

func SolvePart2(input [][]int) int {
	var lowPoints []point

	for i := range input {
		for y := range input[i] {
			if i-1 >= 0 && input[i-1][y] <= input[i][y] {
				continue
			}
			if i+1 <= len(input)-1 && input[i+1][y] <= input[i][y] {
				continue
			}
			if y-1 >= 0 && input[i][y-1] <= input[i][y] {
				continue
			}
			if y+1 <= len(input[i])-1 && input[i][y+1] <= input[i][y] {
				continue
			}
			lowPoints = append(lowPoints, point{i: i, y: y})
		}
	}

	var basins []int
	for _, p := range lowPoints {
		basins = append(basins, len(sumSorroundingLowPoints(input, p)))
	}

	sort.Ints(basins)

	product := 1
	for i := 1; i <= 3; i++ {
		product *= basins[len(basins)-i]
	}

	return product
}

func sumSorroundingLowPoints(input [][]int, p point) map[point]struct{} {
	i, y := p.i, p.y

	if input[i][y] == 9 {
		return nil
	}

	sum := map[point]struct{}{p: {}}
	if i-1 >= 0 && input[i-1][y] > input[i][y] {
		for k := range sumSorroundingLowPoints(input, point{i - 1, y}) {
			sum[k] = struct{}{}
		}
	}
	if i+1 <= len(input)-1 && input[i+1][y] > input[i][y] {
		for k := range sumSorroundingLowPoints(input, point{i + 1, y}) {
			sum[k] = struct{}{}
		}
	}
	if y-1 >= 0 && input[i][y-1] > input[i][y] {
		for k := range sumSorroundingLowPoints(input, point{i, y - 1}) {
			sum[k] = struct{}{}
		}
	}
	if y+1 <= len(input[i])-1 && input[i][y+1] > input[i][y] {
		for k := range sumSorroundingLowPoints(input, point{i, y + 1}) {
			sum[k] = struct{}{}
		}
	}
	return sum
}
