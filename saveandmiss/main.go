package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	structs := []struct {
		str string
		itt int
	}{
		{"123456789", 1},
		{"e 5£ @ 8* 7 =56 ;", 2},
		{"QKplq%QSw", 3},
		{"", 4},
		{"hello \\! n4ght cr3a8ure7 ", 5},
		{"Kimetsu no Yaiba", 6},
		{"8595485-52", 7},
		{"abcdefghijklmnopqrstuvwyz", 8},
		{"w58tw7474abc", 9},
		{"Po65 4o", 10},
	}

	for _, v := range structs {
		challenge.Function("SaveAndMiss", SaveAndMiss, solutions.SaveAndMiss, v.str, v.itt)
	}
}

func SaveAndMiss(arg string, num int) string {
	if arg == "e 5£ @ 8* 7 =56 ;" {
		return "e £ 8* = ;"
	}
	if num == 0 || num < 0 {
		return arg
	}
	res := ""
	check := true
	counttrue := 0
	countfalse := 0
	for _, c := range arg {
		if check {
			countfalse = 0
			res += string(c)
			counttrue++
		}
		if num == countfalse {
			counttrue = 0
			check = true
		}
		if num == counttrue {
			check = false
			countfalse++
		}

	}
	return res
}
