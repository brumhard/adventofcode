package main

import (
	_ "embed"
	"errors"
	"fmt"
	"math"
	"os"
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

type Coordinate struct {
	x int
	y int
}

type Line struct {
	start Coordinate
	end   Coordinate
}

func (l Line) CoveredCoordinates(enableDiagonalCheck bool) []Coordinate {
	var coordinates []Coordinate

	isHorizontal := l.start.x == l.end.x
	isVertical := l.start.y == l.end.y
	is45Diagonal := math.Abs(float64(l.start.x)-float64(l.end.x)) == math.Abs(float64(l.start.y)-float64(l.end.y))
	if isHorizontal || isVertical || (is45Diagonal && enableDiagonalCheck) {
		x, y := l.start.x, l.start.y
		for {
			coordinates = append(coordinates, Coordinate{x: x, y: y})

			if x == l.end.x && y == l.end.y {
				break
			}

			// if end is greater than current increase
			// if it's the same do nothing
			// if end is smaller than current decrease
			// -> go towards end
			x += sign(l.end.x - x)
			y += sign(l.end.y - y)
		}
	}

	return coordinates
}

// sign returns -1 for negative, 1 for positive and 0 for 0.
func sign(x int) int {
	if x == 0 {
		return 0
	}
	if x < 0 {
		return -1
	}
	return 1
}

func inputFromString(inputStr string) ([]Line, error) {
	lineStrs := aocconv.StrToStrSlice(inputStr)

	lines := make([]Line, 0, len(lineStrs))
	for _, lineStr := range lineStrs {
		startEnd := strings.Split(lineStr, " -> ")
		if len(startEnd) != 2 {
			return nil, errors.New("unexpected input")
		}

		startx, starty, err := aocconv.IntTuple(startEnd[0], aocconv.WithDelimeter(","))
		if err != nil {
			return nil, err
		}

		endx, endy, err := aocconv.IntTuple(startEnd[1], aocconv.WithDelimeter(","))
		if err != nil {
			return nil, err
		}

		lines = append(lines, Line{start: Coordinate{x: startx, y: starty}, end: Coordinate{x: endx, y: endy}})
	}

	return lines, nil
}

func coveredCount(lines []Line, withDiagonals bool) int {
	countCoveredByAtLeastTwo := 0
	coveredCountMap := map[Coordinate]int{}

	for _, line := range lines {
		for _, coord := range line.CoveredCoordinates(withDiagonals) {
			current := coveredCountMap[coord]
			coveredCountMap[coord] = current + 1

			if coveredCountMap[coord] == 2 {
				countCoveredByAtLeastTwo++
			}
		}
	}

	return countCoveredByAtLeastTwo
}

func SolvePart1(input []Line) int {
	return coveredCount(input, false)
}

func SolvePart2(input []Line) int {
	return coveredCount(input, true)
}
