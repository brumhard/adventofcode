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

	fmt.Printf("Part 1: %v\n", SolvePart1(input))
	fmt.Printf("Part 2: %v\n", SolvePart2(input))

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
		xy, err := aocconv.StrToIntSlice(line, aocconv.WithDelimeter(","))
		if err != nil {
			return input{}, err
		}

		points[point{xy[0], xy[1]}] = struct{}{}
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

func SolvePart1(input input) int {
	newMap := make(map[point]struct{}, len(input.points))

	foldType, pos := input.folds[0].foldType, input.folds[0].position
	for point := range input.points {
		if foldType == "x" {
			if point.x < pos {
				newMap[point] = struct{}{}
				continue
			}

			if point.x == pos {
				panic("what")
			}

			dist := point.x - pos
			point.x = pos - dist

			newMap[point] = struct{}{}
		}
		if foldType == "y" {
			if point.y < pos {
				newMap[point] = struct{}{}
				continue
			}

			if point.y == pos {
				panic("what")
			}

			dist := point.y - pos
			point.y = pos - dist

			newMap[point] = struct{}{}
		}
	}

	return len(newMap)
}

func SolvePart2(input input) int {
	curMap := input.points
	for _, fold := range input.folds {
		newMap := make(map[point]struct{}, len(input.points))

		foldType, pos := fold.foldType, fold.position
		for point := range curMap {
			if foldType == "x" {
				if point.x < pos {
					newMap[point] = struct{}{}
					continue
				}

				if point.x == pos {
					panic("what")
				}

				dist := point.x - pos
				point.x = pos - dist

				newMap[point] = struct{}{}
			}
			if foldType == "y" {
				if point.y < pos {
					newMap[point] = struct{}{}
					continue
				}

				if point.y == pos {
					panic("what")
				}

				dist := point.y - pos
				point.y = pos - dist

				newMap[point] = struct{}{}
			}
		}

		curMap = newMap
	}

	visualize := make([][]string, 100)
	for i := range visualize {
		visualize[i] = make([]string, 100)
	}

	for y := range visualize {
		for x := range visualize[y] {
			if _, ok := curMap[point{x: x, y: y}]; ok {
				visualize[y][x] = "#"
				fmt.Print("#")
				continue
			}

			visualize[y][x] = "-"
			fmt.Print("-")
		}
		fmt.Println()
	}

	fmt.Println(visualize)

	return 0
}
