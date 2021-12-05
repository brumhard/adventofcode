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

func (l Line) CoveredCoordinates(withDiagonals bool) []Coordinate {
	if l.start.x == l.end.x && l.start.y == l.end.y {
		return []Coordinate{l.start}
	}

	var coordinates []Coordinate
	if l.start.x == l.end.x {
		// either step from higher to lower or from lower to higher
		min, max := l.start.y, l.end.y
		if l.end.y < l.start.y {
			min, max = l.end.y, l.start.y
		}

		for i := min; i <= max; i++ {
			coordinates = append(coordinates, Coordinate{x: l.start.x, y: i})
		}
	}

	if l.start.y == l.end.y {
		// either step from higher to lower or from lower to higher
		min, max := l.start.x, l.end.x
		if l.end.x < l.start.x {
			min, max = l.end.x, l.start.x
		}

		for i := min; i <= max; i++ {
			coordinates = append(coordinates, Coordinate{x: i, y: l.start.y})
		}
	}

	// diagonal
	if withDiagonals && math.Abs(float64(l.start.x)-float64(l.end.x)) == math.Abs(float64(l.start.y)-float64(l.end.y)) {
		x, y := l.start.x, l.start.y
		for {
			coordinates = append(coordinates, Coordinate{x: x, y: y})

			if x == l.end.x {
				break
			}

			if l.start.x < l.end.x {
				x++
			} else {
				x--
			}

			if l.start.y < l.end.y {
				y++
			} else {
				y--
			}
		}
	}

	return coordinates
}

func inputFromString(inputStr string) ([]Line, error) {
	lineStrs := aocconv.StrToStrSlice(inputStr)

	lines := make([]Line, 0, len(lineStrs))
	for _, lineStr := range lineStrs {
		startEnd := strings.Split(lineStr, " -> ")
		if len(startEnd) != 2 {
			return nil, errors.New("unexpected input")
		}

		startxy, err := aocconv.StrToIntSlice(startEnd[0], aocconv.WithDelimeter(","))
		if err != nil {
			return nil, err
		}

		endxy, err := aocconv.StrToIntSlice(startEnd[1], aocconv.WithDelimeter(","))
		if err != nil {
			return nil, err
		}

		lines = append(lines, Line{start: Coordinate{x: startxy[0], y: startxy[1]}, end: Coordinate{x: endxy[0], y: endxy[1]}})
	}

	return lines, nil
}

func SolvePart1(input []Line) int {
	coveredCountMap := map[Coordinate]int{}
	for _, line := range input {
		for _, coord := range line.CoveredCoordinates(false) {
			current, ok := coveredCountMap[coord]
			if !ok {
				current = 0
			}

			coveredCountMap[coord] = current + 1
		}
	}

	countCoveredByAtLeastTwo := 0
	for _, v := range coveredCountMap {
		if v >= 2 {
			countCoveredByAtLeastTwo++
		}
	}

	return countCoveredByAtLeastTwo
}

func SolvePart2(input []Line) int {
	coveredCountMap := map[Coordinate]int{}
	for _, line := range input {
		for _, coord := range line.CoveredCoordinates(true) {
			current, ok := coveredCountMap[coord]
			if !ok {
				current = 0
			}

			coveredCountMap[coord] = current + 1
		}
	}

	countCoveredByAtLeastTwo := 0
	for _, v := range coveredCountMap {
		if v >= 2 {
			countCoveredByAtLeastTwo++
		}
	}

	return countCoveredByAtLeastTwo
}
