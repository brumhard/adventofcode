package aocconv

import (
	"strconv"
	"strings"
)

func StrToIntSlice(inputStr string) ([]int, error) {
	strs := StrToStrSlice(inputStr)
	ints := make([]int, 0, len(strs))

	for _, inputStr := range strs {
		asInt, err := strconv.Atoi(inputStr)
		if err != nil {
			return nil, err
		}

		ints = append(ints, asInt)
	}

	return ints, nil
}

func StrToStrSlice(inputStr string) []string {
	strs := strings.Split(strings.TrimRight(inputStr, "\n"), "\n")
	if len(strs) == 1 && strs[0] == "" {
		return []string{}
	}

	return strs
}
