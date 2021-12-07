package main

import (
	_ "embed"
	"fmt"
	"math"
	"os"
	"runtime"
	"sync"

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
	return solveAnyPartWithCache(input, 256, newSafeCache())
}

type safeCache struct {
	m  map[int]int
	mu sync.RWMutex
}

func newSafeCache() *safeCache {
	return &safeCache{
		m:  map[int]int{},
		mu: sync.RWMutex{},
	}
}

func (c *safeCache) Get(index int) (int, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	val, ok := c.m[index]
	return val, ok
}

func (c *safeCache) Set(index, val int) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.m[index] = val
}

func solveAnyPartWithCache(input []int, days int, cache *safeCache) int {
	sum := 0
	for _, fish := range input {
		sum += calcProduced(days, fish, cache)
	}

	return sum
}

// calcProduced returns itself and all fishes produced from itself and children.
// cache is used to save solutions that already have been calculated.
// If cache is nil it will not be used.
func calcProduced(daysLeft int, initialTimer int, cache *safeCache) int {
	sum := 1
	for i := daysLeft - initialTimer; i > 0; i -= 7 {
		if cache == nil {
			sum += calcProduced(i-1, 8, nil)
			continue
		}

		produced, ok := cache.Get(i)
		if !ok {
			produced = calcProduced(i-1, 8, cache)
			cache.Set(i, produced)
		}
		sum += produced
	}
	return sum
}

func SolvePart2Concurrently(input []int) int {
	numCPU := runtime.NumCPU()
	sumChan := make(chan int)
	cache := newSafeCache()

	workerlen := int(math.Ceil(float64(len(input)) / float64(numCPU)))
	workerIndex := 0
	workersStarted := 0
	for i := 0; i < numCPU; i++ {
		workersStarted++

		lastIndex := workerIndex + workerlen
		if lastIndex >= len(input)-1 {
			lastIndex = len(input) - 1
		}

		go func(input []int) {
			sumChan <- solveAnyPartWithCache(input, 256, cache)
		}(input[workerIndex:lastIndex])

		if lastIndex == len(input)-1 {
			break
		}

		workerIndex += workerlen
	}
	if workerlen*workersStarted < len(input) {
		panic("wtf")
	}

	sum := 0
	for i := 0; i < workersStarted; i++ {
		sum += <-sumChan
	}

	return sum
}
