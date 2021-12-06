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

// SolvePart1 is the naive solution to the problem without optimzations.
func SolvePart1(input []int) int {
	localInput := make([]int, len(input))
	copy(localInput, input)

	for y := 0; y < 80; y++ {
		var toAppend []int
		for i := range localInput {
			if localInput[i] == 0 {
				toAppend = append(toAppend, 8)
				localInput[i] = 6
				continue
			}

			localInput[i]--
		}
		localInput = append(localInput, toAppend...)
	}

	return len(localInput)
}

func SolvePart2(input []int) int {
	return solveAnyPartWithCache(input, 256)
}

func solveAnyPartWithCache(input []int, days int) int {
	sum := 0
	cache := map[int]int{}
	for _, fish := range input {
		sum += calcProduced(days, fish, cache)
	}

	return sum
}

// calcProduced returns itself and all fishes produced from itself and children.
// cache is used to save solutions that already have been calculated.
// If cache is nil it will not be used.
func calcProduced(daysLeft int, initialTimer int, cache map[int]int) int {
	sum := 1
	for i := daysLeft - initialTimer; i > 0; i -= 7 {
		if cache == nil {
			sum += calcProduced(i-1, 8, cache)
			continue
		}

		_, ok := cache[i]
		if !ok {
			cache[i] = calcProduced(i-1, 8, cache)
		}
		sum += cache[i]
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
