package main

import (
	_ "embed"
	"errors"
	"fmt"
	"math"
	"os"
	"strings"

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

type Line struct {
	start coords.Point
	end   coords.Point
}

func (l Line) CoveredCoordinates(enableDiagonalCheck bool) []coords.Point {
	var coordinates []coords.Point

	isHorizontal := l.start.X == l.end.X
	isVertical := l.start.Y == l.end.Y
	is45Diagonal := math.Abs(float64(l.start.X)-float64(l.end.X)) == math.Abs(float64(l.start.Y)-float64(l.end.Y))
	if isHorizontal || isVertical || (is45Diagonal && enableDiagonalCheck) {
		x, y := l.start.X, l.start.Y
		for {
			coordinates = append(coordinates, coords.Point{X: x, Y: y})

			if x == l.end.X && y == l.end.Y {
				break
			}

			// if end is greater than current increase
			// if it's the same do nothing
			// if end is smaller than current decrease
			// -> go towards end
			x += sign(l.end.X - x)
			y += sign(l.end.Y - y)
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

		lines = append(lines, Line{start: coords.Point{X: startx, Y: starty}, end: coords.Point{X: endx, Y: endy}})
	}

	return lines, nil
}

func coveredCount(lines []Line, withDiagonals bool) int {
	countCoveredByAtLeastTwo := 0
	coveredCountMap := map[coords.Point]int{}

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
