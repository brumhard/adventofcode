package main

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"

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
	fmt.Printf("Part 2: \n%v", SolvePart2(input))

	return nil
}

type fold struct {
	foldType string // x or y
	position int
}

type input struct {
	points map[coords.Point]struct{}
	folds  []fold
}

func inputFromString(inputStr string) (input, error) {
	split := aocconv.StrToStrSlice(inputStr, aocconv.WithDelimeter("\n\n"))

	points := map[coords.Point]struct{}{}
	for _, line := range aocconv.StrToStrSlice(split[0]) {
		x, y, err := aocconv.IntTuple(line, aocconv.WithDelimeter(","))
		if err != nil {
			return input{}, err
		}

		points[coords.Point{x, y}] = struct{}{}
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

func calcNewPoints(input input) map[coords.Point]struct{} {
	currentPoints := input.points

	for _, fold := range input.folds {
		newPoints := make(map[coords.Point]struct{}, len(currentPoints))

		for point := range currentPoints {
			if fold.foldType == "x" && point.X > fold.position {
				dist := point.X - fold.position
				point.X = fold.position - dist
			}
			if fold.foldType == "y" && point.Y > fold.position {
				dist := point.Y - fold.position
				point.Y = fold.position - dist
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

	return coords.VisualizePointSet(points)
}
