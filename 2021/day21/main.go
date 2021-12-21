package main

import (
	_ "embed"
	"fmt"
	"os"
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
	p1 int
	p2 int
}

func inputFromString(_ string) (input, error) {
	return input{
		p1: 7,
		p2: 1,
	}, nil

}

func SolvePart1(input input) int {
	var score1, score2 int
	die := 0
	round := 0
	for score1 < 1000 && score2 < 1000 {
		round++

		roll := 0
		for i := 0; i < 3; i++ {
			die = ((die) % 100) + 1
			roll += die
		}

		if round%2 == 1 {
			input.p1 = ((input.p1 + roll - 1) % 10) + 1
			score1 += input.p1
		} else {
			input.p2 = ((input.p2 + roll - 1) % 10) + 1
			score2 += input.p2
		}
	}

	loser := score1
	if score2 < score1 {
		loser = score2
	}

	return loser * (round * 3)
}

var possibleRolls = calcPossibleRolls()

func calcPossibleRolls() []int {
	possibleDieValues := []int{1, 2, 3}

	var possibleRolls []int
	for i := range possibleDieValues {
		for y := range possibleDieValues {
			for x := range possibleDieValues {
				possibleRolls = append(possibleRolls, possibleDieValues[i]+possibleDieValues[y]+possibleDieValues[x])
			}
		}
	}

	return possibleRolls
}

type roundDef struct {
	input
	score1, score2 int
	roll           int
	round          int
}

var resMap = map[roundDef]map[int]int{}

func checkWin(in input, score1, score2, round, roll int) (res map[int]int) {
	roundDef := roundDef{in, score1, score2, roll, round % 2}
	if res, ok := resMap[roundDef]; ok {
		return res
	}

	defer func() {
		resMap[roundDef] = res
	}()

	if round%2 == 1 {
		in.p1 = ((in.p1 + roll - 1) % 10) + 1
		score1 += in.p1
		if score1 >= 21 {
			return map[int]int{1: 1}
		}
	} else {
		in.p2 = ((in.p2 + roll - 1) % 10) + 1
		score2 += in.p2
		if score2 >= 21 {
			return map[int]int{2: 1}
		}
	}

	// map from roll to resultMap
	roundCache := map[int]map[int]int{}
	resultMap := map[int]int{}
	for _, r := range possibleRolls {
		rollResult, ok := roundCache[r]
		if !ok {
			rollResult = checkWin(in, score1, score2, round+1, r)
			roundCache[r] = rollResult
		}
		for k, v := range rollResult {
			current := resultMap[k]
			resultMap[k] = current + v
		}
	}

	return resultMap
}

func SolvePart2(input input) int {
	// map from roll to resultMap
	roundCache := map[int]map[int]int{}
	resultMap := map[int]int{}
	for _, roll := range possibleRolls {
		res, ok := roundCache[roll]
		if !ok {
			res = checkWin(input, 0, 0, 1, roll)
			roundCache[roll] = res
		}
		for k, v := range res {
			current := resultMap[k]
			resultMap[k] = current + v
		}
	}

	maxWins := resultMap[1]
	if resultMap[2] > resultMap[1] {
		maxWins = resultMap[2]
	}

	return maxWins
}
