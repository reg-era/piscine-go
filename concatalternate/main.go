package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	args := [][][]int{
		{{1, 2, 3}, {4, 5, 6}},
		{{1, 2, 3}, {4, 5}},
		{{}, {4, 5, 6}},
		{{1, 2, 3}, {}},
		{{}, {}},
		{{1, 2, 4}, {10, 20, 30, 40, 50}},
	}
	for _, arg := range args {
		challenge.Function("ConcatAlternate", ConcatAlternate, solutions.ConcatAlternate, arg[0], arg[1])
	}
}

func ConcatAlternate(slice1, slice2 []int) []int {
	var res []int
	len1, len2 := len(slice1), len(slice2)

	var long, short []int
	if len1 >= len2 {
		long = slice1
		short = slice2
	} else {
		long = slice2
		short = slice1
	}

	minLen := len(short)
	for i := 0; i < minLen; i++ {
		res = append(res, long[i])
		res = append(res, short[i])
	}

	res = append(res, long[minLen:]...)

	return res
}
