package main

import (
	_ "embed"
	"fmt"
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

type target struct {
	xMin, xMax, yMin, yMax int
}

func inputFromString(inputStr string) (target, error) {
	inputStr = strings.TrimPrefix(inputStr, "target area: ")
	inputStrs := aocconv.StrToStrSlice(inputStr, aocconv.WithDelimeter(", "))
	xMin, xMax, err := aocconv.IntTuple(strings.TrimPrefix(inputStrs[0], "x="), aocconv.WithDelimeter(".."))
	if err != nil {
		return target{}, err
	}

	yMin, yMax, err := aocconv.IntTuple(strings.TrimPrefix(inputStrs[1], "y="), aocconv.WithDelimeter(".."))
	if err != nil {
		return target{}, err
	}

	return target{xMin, xMax, yMin, yMax}, nil
}

// Broodforce everything <3
func SolvePart1(input target) int {
	maxHeight := 0
	for x := 0; x < input.xMax; x++ {
		for y := 0; y < 1000; y++ {
			height := iterate(input, x, y)
			if height != nil && *height > maxHeight {
				maxHeight = *height
			}
		}
	}

	return maxHeight
}

func SolvePart2(input target) int {
	points := map[coords.Point]struct{}{}
	for x := -1000; x < 1000; x++ {
		for y := -1000; y < 1000; y++ {
			height := iterate(input, x, y)
			if height != nil {
				points[coords.Point{X: x, Y: y}] = struct{}{}
			}
		}
	}

	return len(points)
}

func iterate(target target, xVel, yVel int) *int {
	pos := coords.Point{X: 0, Y: 0}
	positiveXVel := xVel > 0
	peakReached := false
	highestYVal := 0
	for {
		if (positiveXVel && pos.X > target.xMax) || (!positiveXVel && pos.X < target.xMin) || (peakReached && pos.Y < target.yMin) {
			// went to far
			return nil
		}

		if pos.X >= target.xMin && pos.X <= target.xMax && pos.Y >= target.yMin && pos.Y <= target.yMax {
			return &highestYVal
		}

		pos.X += xVel
		pos.Y += yVel
		if pos.Y > highestYVal {
			highestYVal = pos.Y
		} else {
			peakReached = true
		}

		if xVel != 0 {
			if xVel > 0 {
				xVel -= 1
			} else {
				xVel += 1
			}
		}
		yVel -= 1
	}
}
