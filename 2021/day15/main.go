package main

import (
	_ "embed"
	"fmt"
	"math"
	"os"

	"github.com/brumhard/adventofcode/aocconv"
	"github.com/brumhard/adventofcode/coords"
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

	// graph, err := nodesFromMatrix(input)
	// if err != nil {
	// 	return err
	// }

	input2 := input2FromMatrix(input)
	graph2, err := nodesFromMatrix(input2)
	if err != nil {
		return err
	}

	// fmt.Printf("Part 1: %v\n", SolvePart12(graph))
	fmt.Printf("Part 2: %v\n", SolvePart12(graph2))

	return nil
}

type Node struct {
	neighbors map[coords.Point]int
	minDist   int
	visited   bool
}

func inputFromString(inputStr string) ([][]int, error) {
	rows := aocconv.StrToStrSlice(inputStr)

	input := make([][]int, 0, len(rows))

	for _, row := range rows {
		inputRow, err := aocconv.StrToIntSlice(row, aocconv.WithDelimeter(""))
		if err != nil {
			return nil, err
		}

		input = append(input, inputRow)
	}

	return input, nil
}

var MAXY, MAXX int

func nodesFromMatrix(matrix [][]int) (map[coords.Point]Node, error) {
	MAXY = len(matrix) - 1
	MAXX = len(matrix[0]) - 1

	input := map[coords.Point]Node{}
	for y := range matrix {
		for x := range matrix[y] {
			newNode := Node{
				neighbors: map[coords.Point]int{},
				minDist:   math.MaxInt,
				visited:   false,
			}

			sorrounding := []coords.Point{
				{X: x + 1, Y: y},
				{X: x - 1, Y: y},
				{X: x, Y: y + 1},
				{X: x, Y: y - 1},
			}

			for _, p := range sorrounding {
				if p.Y < 0 || p.Y >= len(matrix) || p.X < 0 || p.X >= len(matrix[0]) {
					continue
				}

				newNode.neighbors[p] = matrix[p.Y][p.X]
			}

			input[coords.Point{X: x, Y: y}] = newNode
		}
	}

	return input, nil
}

func SolvePart12(input map[coords.Point]Node) int {
	source := input[coords.Point{X: 0, Y: 0}]
	source.minDist = 0
	source.visited = true
	input[coords.Point{X: 0, Y: 0}] = source

	for !checkSolved(input) {
		for _, n := range input {
			if !n.visited {
				continue
			}

			for adjacentPoint, dist := range n.neighbors {
				nodeAtCoords := input[adjacentPoint]
				if distToNode := n.minDist + dist; distToNode < nodeAtCoords.minDist {
					nodeAtCoords.minDist = distToNode
				}

				input[adjacentPoint] = nodeAtCoords
			}
		}

		minDistUnvisited := math.MaxInt
		var nodesToVisit []coords.Point
		for c, n := range input {
			if n.visited {
				continue
			}

			if n.minDist == minDistUnvisited {
				nodesToVisit = append(nodesToVisit, c)
			}

			if n.minDist < minDistUnvisited {
				minDistUnvisited = n.minDist
				nodesToVisit = []coords.Point{c}
			}
		}

		for _, c := range nodesToVisit {
			toVisit := input[c]
			toVisit.visited = true
			input[c] = toVisit
		}

	}

	return input[coords.Point{X: MAXX, Y: MAXY}].minDist
}

func input2FromMatrix(matrix [][]int) [][]int {
	oldY := len(matrix)
	oldX := len(matrix[0])

	newMatrix := make([][]int, 5*oldY)
	for i := range newMatrix {
		newMatrix[i] = make([]int, 5*oldX)
	}

	for i := 0; i < 5; i++ {
		for f := 0; f < 5; f++ {
			toInsert := bumpMatrixTimes(matrix, i+f)
			for y := range toInsert {
				for x := range toInsert[y] {
					newMatrix[y+i*oldY][x+f*oldX] = toInsert[y][x]
				}
			}
		}
	}

	return newMatrix
}

func checkSolved(input map[coords.Point]Node) bool {
	for _, n := range input {
		if !n.visited {
			return false
		}
	}

	return true
}

func bumpMatrixTimes(matrix [][]int, times int) [][]int {
	newMatrix := matrix
	for i := 0; i < times; i++ {
		newMatrix = bumpMatrix(newMatrix)
	}
	return newMatrix
}

func bumpMatrix(matrix [][]int) [][]int {
	newMatrix := make([][]int, 0, len(matrix))
	for y := range matrix {
		newRow := make([]int, 0, len(matrix[y]))
		for x := range matrix[y] {
			if matrix[y][x] == 9 {
				newRow = append(newRow, 1)
				continue
			}
			newRow = append(newRow, matrix[y][x]+1)
		}
		newMatrix = append(newMatrix, newRow)
	}

	return newMatrix
}

func SolvePart2(input [][]int) int {
	var solution int

	return solution
}
