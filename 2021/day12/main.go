package main

import (
	_ "embed"
	"errors"
	"fmt"
	"os"
	"strings"
	"unicode"

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

func inputFromString(inputStr string) (map[string][]string, error) {
	edgeMap := map[string][]string{}

	edges := aocconv.StrToStrSlice(inputStr)
	for _, edge := range edges {
		startEnd := strings.Split(edge, "-")
		if len(startEnd) != 2 {
			return nil, errors.New("malformed input")
		}

		edgeMap[startEnd[0]] = append(edgeMap[startEnd[0]], startEnd[1])
		edgeMap[startEnd[1]] = append(edgeMap[startEnd[1]], startEnd[0])
	}

	return edgeMap, nil
}

type path []string

func (p path) end() string {
	return p[len(p)-1]
}

func (p path) isInvalid(part2 bool) bool {
	lowercaseMap := map[string]struct{}{}
	oneLowercaseVisitedTwice := false
	for _, str := range p {
		if !isLower(str) {
			continue
		}
		if _, ok := lowercaseMap[str]; !ok {
			lowercaseMap[str] = struct{}{}
			continue
		}

		if !part2 {
			return true
		}

		// can't visit start twice
		if str == "start" {
			return true
		}

		if oneLowercaseVisitedTwice {
			return true
		}
		oneLowercaseVisitedTwice = true
	}

	return false
}

func (p path) isFinished() bool {
	return p.end() == "end"
}

func isLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// NOTE: sth is wrong here the last path element in some paths is replaced with "kj"
// in some cases for whatever reason. Still returns the right amount of paths tho.
func possiblePathsFrom(current path, edgeMap map[string][]string, part2 bool) []path {
	var paths []path

	if current.isInvalid(part2) {
		return nil
	}

	if current.isFinished() {
		return []path{current}
	}

	for _, n := range edgeMap[current.end()] {
		additionalPaths := possiblePathsFrom(append(current, n), edgeMap, part2)
		if additionalPaths == nil {
			continue
		}

		paths = append(paths, additionalPaths...)
	}

	return paths
}

func SolvePart1(input map[string][]string) int {
	paths := possiblePathsFrom(path{"start"}, input, false)

	return len(paths)
}

func SolvePart2(input map[string][]string) int {
	paths := possiblePathsFrom(path{"start"}, input, true)

	return len(paths)
}
