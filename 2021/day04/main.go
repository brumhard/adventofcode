package main

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/brumhard/adventofcode/aocconv"
)

type Bingo struct {
	inputs []int
	boards []Board
}

type Board struct {
	fields [][]Field
	done   bool
}

func (b *Board) SelectNumber(number int) {
	for i := range b.fields {
		for y := range b.fields[i] {
			if b.fields[i][y].number == number {
				b.fields[i][y].selected = true
			}
		}
	}
}

func (b Board) CheckBingo() bool {
	// this assumes that the board is i x i fields and not i x y
	for i := range b.fields {
		if checkVertical(b.fields, i) || checkHorizontal(b.fields, i) {
			return true
		}
	}

	return false
}

func (b Board) SumUnmarked() int {
	var sum int
	for _, fieldRow := range b.fields {
		for _, field := range fieldRow {
			if field.selected {
				continue
			}

			sum += field.number
		}
	}

	return sum
}

func checkVertical(fields [][]Field, index int) bool {
	for _, fieldRow := range fields {
		if !fieldRow[index].selected {
			return false
		}
	}

	return true
}

func checkHorizontal(fields [][]Field, index int) bool {
	for _, field := range fields[index] {
		if !field.selected {
			return false
		}
	}

	return true
}

type Field struct {
	number   int
	selected bool
}

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

func inputFromString(inputStr string) (Bingo, error) {
	inputRows := aocconv.StrToStrSlice(inputStr)

	inputs, err := aocconv.StrToIntSlice(inputRows[0], aocconv.WithDelimeter(","))
	if err != nil {
		return Bingo{}, err
	}

	var (
		boards        []Board
		currentFields [][]Field
		y             int
	)

	for _, boardRow := range inputRows[2:] {
		if boardRow == "" && len(currentFields) > 0 {
			boards = append(boards, Board{fields: currentFields})
			currentFields = [][]Field{}
			y = 0
			continue
		}

		currentFields = append(currentFields, []Field{})
		ints, err := aocconv.StrToIntSlice(boardRow, aocconv.WithWhitespace())
		if err != nil {
			return Bingo{}, err
		}

		for _, field := range ints {
			currentFields[y] = append(currentFields[y], Field{
				number: field,
			})
		}

		y++
	}
	// append last one
	if len(currentFields) > 0 {
		boards = append(boards, Board{fields: currentFields})
	}

	return Bingo{
		inputs: inputs,
		boards: boards,
	}, nil
}

func SolvePart1(input Bingo) int {
	winner, finalNumber := findWinner(input)

	if winner.fields == nil {
		panic("couldn't find winner")
	}

	return winner.SumUnmarked() * finalNumber
}

// returns winner with final number
func findWinner(bingo Bingo) (Board, int) {
	for _, number := range bingo.inputs {
		for _, board := range bingo.boards {
			board.SelectNumber(number)
			if board.CheckBingo() {
				return board, number
			}
		}
	}

	return Board{}, 0
}

func SolvePart2(input Bingo) int {
	var lastToWin Board
	var lastNumber int
	for _, number := range input.inputs {
		checkedAtLeastOne := false
		for i := range input.boards {
			if input.boards[i].done {
				continue
			}
			checkedAtLeastOne = true

			input.boards[i].SelectNumber(number)
			if input.boards[i].CheckBingo() {
				input.boards[i].done = true
				lastToWin = input.boards[i]
			}
		}
		if !checkedAtLeastOne {
			break
		}

		lastNumber = number
	}

	return lastToWin.SumUnmarked() * lastNumber
}
