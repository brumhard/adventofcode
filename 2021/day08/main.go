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

func run() error {
	input, err := inputFromString(string(inputFile))
	if err != nil {
		return err
	}

	fmt.Printf("Part 1: %v\n", SolvePart1(input))
	fmt.Printf("Part 2: %v\n", SolvePart2(input))

	return nil
}

type IO struct {
	input  []string
	output []string
}

func inputFromString(inputStr string) ([]IO, error) {
	ioStrs := aocconv.StrToStrSlice(inputStr)

	ios := make([]IO, 0, len(ioStrs))
	for _, str := range ioStrs {
		ioArr := strings.Split(str, " | ")
		if len(ioArr) != 2 {
			return nil, errors.New("malformed input")
		}
		ios = append(ios, IO{
			input:  aocconv.StrToStrSlice(ioArr[0], aocconv.WithWhitespace()),
			output: aocconv.StrToStrSlice(ioArr[1], aocconv.WithWhitespace()),
		})
	}

	return ios, nil
}

func SolvePart1(input []IO) int {
	uniqueDigitCount := 0
	for _, io := range input {
		for _, output := range io.output {
			switch len([]rune(output)) {
			case 2, 3, 4, 7:
				uniqueDigitCount++
			}
		}
	}

	return uniqueDigitCount
}

func SolvePart2(input []IO) int {
	sum := 0
	for _, io := range input {
		numToPattern := map[int]string{}
		// find unique ones in input
		for _, output := range io.input {
			switch len([]rune(output)) {
			case 2:
				numToPattern[1] = output
			case 3:
				numToPattern[7] = output
			case 4:
				numToPattern[4] = output
			case 7:
				numToPattern[8] = output
			}
		}

		sb := strings.Builder{}
		for _, output := range io.output {
			sb.WriteString(findNumForString(output, numToPattern))
		}
		outputAsInt, err := strconv.Atoi(sb.String())
		if err != nil {
			panic("malformed output string")
		}

		sum += outputAsInt
	}

	return sum
}

func findNumForString(str string, numToPattern map[int]string) string {
	switch len([]rune(str)) {
	case 2:
		return "1"
	case 3:
		return "7"
	case 4:
		return "4"
	case 7:
		return "8"
	case 5:
		if strContainsAll(str, numToPattern[1]) {
			return "3"
		}

		diff := substractString(str, numToPattern[4])
		if len([]rune(diff)) == 3 {
			return "2"
		}
		return "5"
	case 6:
		if !strContainsAll(str, numToPattern[1]) {
			return "6"
		}

		diff := substractString(numToPattern[8], str)
		if strContainsAll(numToPattern[4], diff) {
			return "0"
		}
		return "9"
	}
	panic("damn")
}

func substractString(a, b string) string {
	sb := strings.Builder{}
	for _, char := range a {
		if !strings.ContainsRune(b, char) {
			sb.WriteRune(char)
		}
	}

	return sb.String()
}

func strContainsAll(a, b string) bool {
	for _, char := range b {
		if !strings.ContainsRune(a, char) {
			return false
		}
	}
	return true
}
