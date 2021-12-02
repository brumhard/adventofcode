package main

import (
	_ "embed"
	"errors"
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

type Move struct {
	Direction string
	Units     int
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

func inputFromString(inputStr string) ([]Move, error) {
	moveStrings := aocconv.StrToStrSlice(inputStr)
	input := make([]Move, 0, len(moveStrings))

	for _, str := range moveStrings {
		moveArr := strings.Split(str, " ")
		if len(moveArr) != 2 {
			return nil, errors.New("whoops")
		}

		direction := moveArr[0]
		units, err := strconv.Atoi(moveArr[1])
		if err != nil {
			return nil, err
		}

		input = append(input, Move{Direction: direction, Units: units})
	}

	return input, nil
}

func SolvePart1(input []Move) int {
	var hor, depth int

	for _, mv := range input {
		switch mv.Direction {
		case "forward":
			hor += mv.Units
		case "down":
			depth += mv.Units
		case "up":
			depth -= mv.Units
		}
	}

	return hor * depth
}

func SolvePart2(input []Move) int {
	var hor, depth, aim int

	for _, mv := range input {
		switch mv.Direction {
		case "forward":
			hor += mv.Units
			depth += aim * mv.Units
		case "down":
			aim += mv.Units
		case "up":
			aim -= mv.Units
		}
	}

	return hor * depth
}
