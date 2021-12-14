package main

import (
	_ "embed"
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

type input struct {
	insertionMap map[string]string
	template     string
}

func inputFromString(inputStr string) (input, error) {
	lines := aocconv.StrToStrSlice(inputStr)

	insertionMap := make(map[string]string, len(lines[2:]))
	for _, line := range lines[2:] {
		keyval := strings.Split(line, " -> ")
		insertionMap[keyval[0]] = keyval[1]
	}

	return input{
		template:     lines[0],
		insertionMap: insertionMap,
	}, nil
}

func SolvePart1(input input) int {
	currentTemplate := input.template
	for i := 0; i < 10; i++ {
		newTemplateBuilder := strings.Builder{}

		templateRunes := []rune(currentTemplate)
		for y := range templateRunes {
			newTemplateBuilder.WriteRune(templateRunes[y])

			if y+1 == len(currentTemplate) {
				break
			}

			pair := string(templateRunes[y : y+2])
			insert := input.insertionMap[pair]
			newTemplateBuilder.WriteString(insert)
		}

		currentTemplate = newTemplateBuilder.String()
	}

	countMap := map[rune]int{}
	for _, r := range currentTemplate {
		currentCount := countMap[r]
		countMap[r] = currentCount + 1
	}

	min, max := math.MaxInt, 0

	for _, count := range countMap {
		if count < min {
			min = count
		}
		if count > max {
			max = count
		}
	}

	return max - min
}

var cache = map[string]map[int]ResultMap{}

func SolvePart2(input input) int {
	const steps = 10

	templateRunes := []rune(input.template)
	countMap := map[rune]int{}
	for y := range templateRunes {
		if y+1 == len(templateRunes) {
			break
		}

		solve2(string(templateRunes[y:y+2]), input.insertionMap, steps, y == 0, countMap)
	}

	min, max := math.MaxInt, 0

	for _, count := range countMap {
		if count < min {
			min = count
		}
		if count > max {
			max = count
		}
	}

	return max - min
}

type ResultMap map[rune]int

func solve2(double string, insertionMap map[string]string, iterations int, first bool, resultMap map[rune]int) {
	if cached, ok := cache[double][iterations]; ok {
		mergeIntoMap(resultMap, cached)
		return
	}

	doubleRunes := []rune(double)
	currentTemplate := string(doubleRunes[0]) + insertionMap[double] + string(doubleRunes[1])
	newResultMap := map[rune]int{}

	if iterations == 1 {
		for i, r := range currentTemplate {
			if i == 0 && !first {
				continue
			}
			count := newResultMap[r]
			newResultMap[r] = count + 1
		}
	}

	for i := 1; i < iterations; i++ {
		templateRunes := []rune(currentTemplate)
		for y := range templateRunes {
			if y+1 == len(templateRunes) {
				break
			}

			curDouble := string(templateRunes[y : y+2])
			solve2(curDouble, insertionMap, iterations-i, y == 0, newResultMap)
		}
	}

	if _, ok := cache[double]; ok {
		cache[double][iterations] = newResultMap
	} else {
		cache[double] = map[int]ResultMap{iterations: newResultMap}
	}

	mergeIntoMap(resultMap, newResultMap)
}

func mergeIntoMap(resultMap map[rune]int, toMerge map[rune]int) {
	for k, v := range toMerge {
		count := resultMap[k]
		resultMap[k] = count + v
	}
}
