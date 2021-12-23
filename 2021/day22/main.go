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

type instruction struct {
	itype      string // on or off
	xMin, xMax int
	yMin, yMax int
	zMin, zMax int
}

func inputFromString(inputStr string) ([]instruction, error) {
	lines := aocconv.StrToStrSlice(inputStr)
	instructions := make([]instruction, 0, len(lines))
	for _, line := range lines {
		split := aocconv.StrToStrSlice(line, aocconv.WithDelimeter(" "))
		xyz := aocconv.StrToStrSlice(split[1], aocconv.WithDelimeter(","))
		xMin, xMax, err := aocconv.IntTuple(strings.TrimPrefix(xyz[0], "x="), aocconv.WithDelimeter(".."))
		if err != nil {
			return nil, err
		}
		yMin, yMax, err := aocconv.IntTuple(strings.TrimPrefix(xyz[1], "y="), aocconv.WithDelimeter(".."))
		if err != nil {
			return nil, err
		}
		zMin, zMax, err := aocconv.IntTuple(strings.TrimPrefix(xyz[2], "z="), aocconv.WithDelimeter(".."))
		if err != nil {
			return nil, err
		}
		instructions = append(instructions, instruction{split[0], xMin, xMax, yMin, yMax, zMin, zMax})
	}

	return instructions, nil
}

func getOverlap(i1 instruction, i2 instruction) *cube {
	overlap := instruction{
		itype: i2.itype,
		xMax:  min(i1.xMax, i2.xMax),
		xMin:  max(i1.xMin, i2.xMin),
		yMax:  min(i1.yMax, i2.yMax),
		yMin:  max(i1.yMin, i2.yMin),
		zMax:  min(i1.zMax, i2.zMax),
		zMin:  max(i1.zMin, i2.zMin),
	}

	if overlap.xMax < overlap.xMin || overlap.yMax < overlap.yMin || overlap.zMax < overlap.zMin {
		return nil
	}

	return &cube{instruction: overlap}
}

func min(i1, i2 int) int {
	return int(math.Min(float64(i1), float64(i2)))
}

func max(i1, i2 int) int {
	return int(math.Max(float64(i1), float64(i2)))
}

type cube struct {
	instruction
	substractions []cube
}

func NewCube(in instruction, existingCubes []cube) *cube {
	for i := range existingCubes {
		existingCubes[i].Substract(cube{instruction: in})
	}
	if in.itype == "off" {
		return nil
	}

	newCube := cube{
		instruction: in,
	}

	return &newCube
}

func (c *cube) Substract(newSubstraction cube) {
	overlap := getOverlap(c.instruction, newSubstraction.instruction)
	if overlap == nil {
		return
	}
	for _, existingSubstraction := range c.substractions {
		overlap.Substract(existingSubstraction)
	}
	c.substractions = append(c.substractions, *overlap)
}

func (c *cube) Volume() int {
	if c == nil {
		return 0
	}
	volume := (c.zMax + 1 - c.zMin) * (c.yMax + 1 - c.yMin) * (c.xMax + 1 - c.xMin)
	for _, sub := range c.substractions {
		volume -= sub.Volume()
	}
	return volume
}

func SolvePart2(input []instruction) int {
	var cubes []cube
	for _, instr := range input {
		cube := NewCube(instr, cubes)
		if cube != nil {
			cubes = append(cubes, *cube)
		}
	}

	completeVolume := 0
	for i := range cubes {
		completeVolume += cubes[i].Volume()
	}

	return completeVolume
}

func SolvePart1(input []instruction) int {
	grid := make([][][]bool, 101)
	for z := range grid {
		grid[z] = make([][]bool, 101)
		for y := range grid[z] {
			grid[z][y] = make([]bool, 101)
		}
	}

	for _, instr := range input {
		for z := instr.zMin; z <= instr.zMax; z++ {
			if z < -50 || z > 50 {
				continue
			}
			for y := instr.yMin; y <= instr.yMax; y++ {
				if y < -50 || y > 50 {
					continue
				}
				for x := instr.xMin; x <= instr.xMax; x++ {
					if x < -50 || x > 50 {
						continue
					}
					if instr.itype == "on" {
						grid[z+50][y+50][x+50] = true
					} else {
						grid[z+50][y+50][x+50] = false
					}
				}
			}
		}
	}

	onCount := 0
	for z := range grid {
		for y := range grid[z] {
			for x := range grid[z][y] {
				if grid[z][y][x] {
					onCount++
				}
			}
		}
	}

	return onCount
}
