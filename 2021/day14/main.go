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

func SolvePart2(input input) int {
	const steps = 40

	templateRunes := []rune(input.template)
	countMap := map[rune]int{}
	for y := range templateRunes {
		count := countMap[rune(templateRunes[y])]
		countMap[rune(templateRunes[y])] = count + 1

		if y+1 == len(templateRunes) {
			break
		}

		toInsert := solvePair(string(templateRunes[y])+string(templateRunes[y+1]), input.insertionMap, steps)
		mergeIntoMap(countMap, toInsert)
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

var cache = map[string]map[int]ResultMap{}

// solvePair calculates what runes need to be inserted in between the given pair after n iterations.
func solvePair(pair string, insertionMap map[string]string, iterations int) (res map[rune]int) {
	// caching
	if cached, ok := cache[pair][iterations]; ok {
		return cached
	}

	defer func() {
		if _, ok := cache[pair]; ok {
			cache[pair][iterations] = res
		} else {
			cache[pair] = map[int]ResultMap{iterations: res}
		}
	}()

	// logic
	resultMap := map[rune]int{
		rune(insertionMap[pair][0]): 1,
	}

	if iterations <= 1 {
		return resultMap
	}

	withInsertion := string(pair[0]) + insertionMap[pair] + string(pair[1])
	for y := range withInsertion {
		if y+1 >= len(withInsertion) {
			break
		}

		toInsert := solvePair(string(withInsertion[y])+string(withInsertion[y+1]), insertionMap, iterations-1)
		mergeIntoMap(resultMap, toInsert)
	}

	return resultMap
}

func mergeIntoMap(resultMap map[rune]int, toMerge map[rune]int) {
	for k, v := range toMerge {
		count := resultMap[k]
		resultMap[k] = count + v
	}
}
