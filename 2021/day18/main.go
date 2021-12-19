package main

import (
	_ "embed"
	"errors"
	"fmt"
	"math"
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

type pair struct {
	x, y interface{}
}

func (p pair) String() string {
	return fmt.Sprintf("[%s,%s]", elementString(p.x), elementString(p.y))
}

func elementString(element interface{}) string {
	if elementPair, ok := element.(pair); ok {
		return elementPair.String()
	}
	return fmt.Sprintf("%v", element)
}

func inputFromString(inputStr string) ([]pair, error) {
	pairStrs := aocconv.StrToStrSlice(inputStr)
	pairs := make([]pair, 0, len(pairStrs))

	for _, pairStr := range pairStrs {
		p, err := pairFromStr(pairStr)
		if err != nil {
			return nil, err
		}

		pairs = append(pairs, p)
	}

	return pairs, nil
}

func pairFromStr(in string) (pair, error) {
	contents := in[1 : len(in)-1]
	openedBracketsCount := 0
	splitPos := 0
	for i, char := range contents {
		switch char {
		case '[':
			openedBracketsCount++
		case ']':
			openedBracketsCount--
		case ',':
			if openedBracketsCount != 0 {
				continue
			}
			splitPos = i
		}
	}

	x, err := elementFromStr(contents[:splitPos])
	if err != nil {
		return pair{}, err
	}

	y, err := elementFromStr(contents[splitPos+1:])
	if err != nil {
		return pair{}, err
	}

	return pair{x, y}, nil
}

func elementFromStr(in string) (interface{}, error) {
	if len(in) < 1 {
		return nil, errors.New("can't handle empty string")
	}

	if !strings.ContainsAny(in, "[],") {
		return strconv.Atoi(in)
	}

	return pairFromStr(in)
}

func applyExplodes(p pair, nestLevel int) (interface{}, *pair) {
	newPair := pair{p.x, p.y}

	if nestLevel == 5 {
		return 0, &newPair
	}

	var exploded *pair
	var element interface{}

	if xPair, ok := p.x.(pair); ok {
		element, exploded = applyExplodes(xPair, nestLevel+1)
		newPair.x = element
		if exploded != nil {
			// can only directly apply the explodedPair to the other pair and bubble up the rest
			newPair.y = applyLeftover(p.y, exploded.y.(int), "y")
			// set to zero after apply to not bubble up anymore
			exploded.y = 0
		}
	}
	if newPair != p { // || exploded != nil
		return newPair, exploded
	}

	if yPair, ok := p.y.(pair); ok {
		element, exploded = applyExplodes(yPair, nestLevel+1)
		newPair.y = element
		if exploded != nil {
			// can only directly apply the explodedPair to the other pair and bubble up the rest
			newPair.x = applyLeftover(p.x, exploded.x.(int), "x")
			// set to zero after apply to not bubble up anymore
			exploded.x = 0
		}
	}
	return newPair, exploded
}

func applySplits(p pair) pair {
	newPair := pair{p.x, p.y}

	if xInt, ok := p.x.(int); ok && xInt > 9 {
		newPair.x = split(xInt)
		return newPair
	}

	if xPair, ok := p.x.(pair); ok {
		newPair.x = applySplits(xPair)
	}
	if newPair != p {
		return newPair
	}

	if yInt, ok := p.y.(int); ok && yInt > 9 {
		newPair.y = split(yInt)
		return newPair
	}

	if yPair, ok := p.y.(pair); ok {
		newPair.y = applySplits(yPair)
	}
	if newPair != p {
		return newPair
	}

	return newPair
}

// reducePair returns the new element to replace the current pair with, as well as the exploded pair if any exploded
func reducePair(p pair) pair {
	npInterface, _ := applyExplodes(p, 1)
	if npInterface.(pair) != p {
		return npInterface.(pair)
	}

	newPair := applySplits(p)
	if newPair != p {
		return newPair
	}

	return p
}

func applyLeftover(element interface{}, value int, ltype string) interface{} {
	if elementInt, ok := element.(int); ok {
		return elementInt + value
	}

	p := element.(pair)
	if ltype == "y" {
		return pair{applyLeftover(p.x, value, "y"), p.y}
	}
	if ltype == "x" {
		return pair{p.x, applyLeftover(p.y, value, "x")}
	}
	panic("unsupported leftover type")
}

func split(number int) pair {
	res := float64(number) / 2
	return pair{int(math.Floor(res)), int(math.Ceil(res))}
}

func repeatReduce(p pair) pair {
	var currentPair pair
	newPair := p

	for newPair != currentPair {
		currentPair = newPair
		newPair = reducePair(currentPair)
	}

	return newPair
}

func sumList(input []pair) pair {
	var currentSum pair
	for i := range input {
		if i == 0 {
			currentSum = input[i]
			continue
		}

		currentSum = repeatReduce(pair{currentSum, input[i]})
	}

	return currentSum
}

func magnitude(in interface{}) int {
	if inInt, ok := in.(int); ok {
		return inInt
	}

	inPair := in.(pair)
	return 3*magnitude(inPair.x) + 2*magnitude(inPair.y)
}

func SolvePart1(input []pair) int {
	res := sumList(input)
	return magnitude(res)
}

func SolvePart2(input []pair) int {
	highestMagnitude := 0
	for i := range input {
		for y := range input {
			if y == i {
				continue
			}

			magn1 := magnitude(repeatReduce(pair{input[i], input[y]}))

			if magn1 > highestMagnitude {
				highestMagnitude = magn1
			}
		}
	}

	return highestMagnitude
}
