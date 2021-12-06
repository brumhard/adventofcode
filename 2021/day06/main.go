package main

import (
	_ "embed"
	"fmt"
	"math"
	"os"
	"runtime"

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

func inputFromString(inputStr string) ([]int, error) {
	return aocconv.StrToIntSlice(inputStr, aocconv.WithDelimeter(","))
}

// func SolvePart1(input []int) int {
// 	for y := 0; y < 80; y++ {
// 		var toAppend []int
// 		for i := range input {
// 			if input[i] == 0 {
// 				toAppend = append(toAppend, 8)
// 				input[i] = 6
// 				continue
// 			}

// 			input[i]--
// 		}
// 		input = append(input, toAppend...)
// 	}

// 	return len(input)
// }

func SolvePart1(input []int) int {
	return solveAnyPart(input, 80)
}

func SolvePart2(input []int) int {
	return solveAnyPart(input, 256)
}

func solveAnyPart(input []int, days int) int {
	sum := 0
	for _, fish := range input {
		sum += calcProduced(days, fish)
	}

	return sum
}

// daysleft to produced
var cache = map[int]int{}

// returns itself and all fishes produced from itself and children
func calcProduced(daysLeft int, state int) int {
	sum := 1
	for i := daysLeft - state; i > 0; i -= 7 {
		produced, ok := cache[i]
		if !ok {
			produced = calcProduced(i-1, 8)
			cache[i] = produced
		}
		sum += produced
	}
	return sum
}

func SolvePart2Concurrent(input []int) int {
	numCPU := runtime.NumCPU()
	sumChan := make(chan int)
	workerlen := int(math.Ceil(float64(len(input)) / float64(numCPU)))
	workerIndex := 0
	for i := 0; i < numCPU; i++ {
		if workerIndex > len(input)-1 {
			break
		}

		lastIndex := workerIndex + workerlen
		if lastIndex >= len(input)-1 {
			lastIndex = len(input) - 1
		}

		go func(input []int) {
			for y := 0; y < 256; y++ {
				var toAppend []int
				for i := range input {
					if input[i] == 0 {
						toAppend = append(toAppend, 8)
						input[i] = 6
						continue
					}

					input[i]--
				}
				input = append(input, toAppend...)
			}
			sumChan <- len(input)
		}(input[workerIndex:lastIndex])

		workerIndex += workerlen
	}

	sum := 0
	for i := 0; i < numCPU; i++ {
		sum += <-sumChan
	}

	return sum
}
