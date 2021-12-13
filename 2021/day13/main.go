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
	fmt.Printf("Part 2: \n%v", SolvePart2(input))

	return nil
}

type point struct {
	x, y int
}

type fold struct {
	foldType string // x or y
	position int
}

type input struct {
	points map[point]struct{}
	folds  []fold
}

func inputFromString(inputStr string) (input, error) {
	split := aocconv.StrToStrSlice(inputStr, aocconv.WithDelimeter("\n\n"))

	points := map[point]struct{}{}
	for _, line := range aocconv.StrToStrSlice(split[0]) {
		x, y, err := aocconv.IntTuple(line, aocconv.WithDelimeter(","))
		if err != nil {
			return input{}, err
		}

		points[point{x, y}] = struct{}{}
	}

	var folds []fold
	for _, line := range aocconv.StrToStrSlice(split[1]) {
		splitLine := aocconv.StrToStrSlice(line, aocconv.WithDelimeter(" "))

		directions := aocconv.StrToStrSlice(splitLine[2], aocconv.WithDelimeter("="))

		position, err := strconv.Atoi(directions[1])
		if err != nil {
			return input{}, err
		}

		folds = append(folds, fold{
			foldType: directions[0],
			position: position,
		})
	}

	return input{
		points: points,
		folds:  folds,
	}, nil
}

func calcNewPoints(input input) map[point]struct{} {
	currentPoints := input.points

	for _, fold := range input.folds {
		newPoints := make(map[point]struct{}, len(currentPoints))

		for point := range currentPoints {
			if fold.foldType == "x" && point.x > fold.position {
				dist := point.x - fold.position
				point.x = fold.position - dist
			}
			if fold.foldType == "y" && point.y > fold.position {
				dist := point.y - fold.position
				point.y = fold.position - dist
			}
			newPoints[point] = struct{}{}
		}

		currentPoints = newPoints
	}

	return currentPoints
}

func SolvePart1(input input) int {
	input.folds = input.folds[:1]
	points := calcNewPoints(input)
	return len(points)
}

func SolvePart2(input input) string {
	points := calcNewPoints(input)

	var maxX, maxY int
	for p := range points {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	matrix := NewMatrix(maxX+1, maxY+1).FillWith(" ")

	for p := range points {
		matrix[p.y][p.x] = "#"
	}

	return matrix.String()
}

type Matrix [][]string

func NewMatrix(x, y int) Matrix {
	visualize := make([][]string, y)
	for i := range visualize {
		visualize[i] = make([]string, x)
	}

	return visualize
}

func (m Matrix) FillWith(char string) Matrix {
	for y := range m {
		for x := range m[y] {
			m[y][x] = char
		}
	}

	return m
}

func (m Matrix) String() string {
	builder := strings.Builder{}
	for y := range m {
		for x := range m[y] {
			builder.WriteString(m[y][x])
		}
		builder.WriteString("\n")
	}

	return builder.String()
}
