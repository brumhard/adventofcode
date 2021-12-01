package aocconv

import (
	"strconv"
	"strings"
)

func StrToIntSlice(inputBytes string) ([]int, error) {
	inputs := strings.Split(string(inputBytes), "\n")
	inputsInt := make([]int, 0, len(inputs))

	for _, inputStr := range inputs {
		if inputStr == "" {
			continue
		}

		asInt, err := strconv.Atoi(inputStr)
		if err != nil {
			return nil, err
		}

		inputsInt = append(inputsInt, asInt)
	}

	return inputsInt, nil
}
