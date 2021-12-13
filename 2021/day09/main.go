package main

import (
	_ "embed"
	"fmt"
	"os"
	"sort"

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

func SolvePart2(input [][]int) int {
	var lowPoints []coords.Point

	for y := range input {
		for x := range input[y] {
			if y-1 >= 0 && input[y-1][x] <= input[y][x] {
				continue
			}
			if y+1 <= len(input)-1 && input[y+1][x] <= input[y][x] {
				continue
			}
			if x-1 >= 0 && input[y][x-1] <= input[y][x] {
				continue
			}
			if x+1 <= len(input[y])-1 && input[y][x+1] <= input[y][x] {
				continue
			}
			lowPoints = append(lowPoints, coords.Point{X: x, Y: y})
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

func sumSorroundingLowPoints(input [][]int, p coords.Point) map[coords.Point]struct{} {
	y, x := p.Y, p.X

	if input[y][x] == 9 {
		return nil
	}

	sum := map[coords.Point]struct{}{p: {}}
	if y-1 >= 0 && input[y-1][x] > input[y][x] {
		for k := range sumSorroundingLowPoints(input, coords.Point{Y: y - 1, X: x}) {
			sum[k] = struct{}{}
		}
	}
	if y+1 <= len(input)-1 && input[y+1][x] > input[y][x] {
		for k := range sumSorroundingLowPoints(input, coords.Point{Y: y + 1, X: x}) {
			sum[k] = struct{}{}
		}
	}
	if x-1 >= 0 && input[y][x-1] > input[y][x] {
		for k := range sumSorroundingLowPoints(input, coords.Point{Y: y, X: x - 1}) {
			sum[k] = struct{}{}
		}
	}
	if x+1 <= len(input[y])-1 && input[y][x+1] > input[y][x] {
		for k := range sumSorroundingLowPoints(input, coords.Point{Y: y, X: x + 1}) {
			sum[k] = struct{}{}
		}
	}
	return sum
}
