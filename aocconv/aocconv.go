package aocconv

import (
	"strconv"
	"strings"
)

type ParseOpts struct {
	delimeter  *string
	whitespace bool
}

type Option func(*ParseOpts)

func WithDelimeter(delimeter string) Option {
	return func(po *ParseOpts) {
		po.delimeter = &delimeter
	}
}

func WithWhitespace() Option {
	return func(po *ParseOpts) {
		po.whitespace = true
	}
}

func StrToIntSlice(inputStr string, opts ...Option) ([]int, error) {
	strs := StrToStrSlice(inputStr, opts...)
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

func StrToStrSlice(inputStr string, opts ...Option) []string {
	options := ParseOpts{}

	for _, opt := range opts {
		opt(&options)
	}

	delimeter := "\n"
	if options.delimeter != nil {
		delimeter = *options.delimeter
	}

	splitterFn := func(s string) []string {
		return strings.Split(s, delimeter)
	}
	if options.whitespace {
		splitterFn = strings.Fields
	}

	strs := splitterFn(strings.TrimRight(inputStr, "\n"))
	if len(strs) == 1 && strs[0] == "" {
		return []string{}
	}

	return strs
}
