package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	for i := 0; i < len(arg); i++ {
		res := reversestrcap(arg[i])
		for _, c := range res {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}

func reversestrcap(str string) string {
	res := ""
	check := false
	for i := len(str) - 1; i >= 0; i-- {
		if check {
			check = false
			continue
		}
		if i == len(str)-1 {
			res = up(str[i]) + res
			continue
		}
		if str[i] == ' ' && i > 0 {
			res = string(str[i]) + res
			res = up(str[i-1]) + res
			check = true
			continue
		}
		res = low(str[i]) + res

	}
	return res
}

func low(c byte) string {
	if c >= 'A' && c <= 'Z' {
		c = 'a' + (c - 'A')
	}
	return string(c)
}

func up(c byte) string {
	if c >= 'a' && c <= 'z' {
		c = 'A' + (c - 'a')
	}
	return string(c)
}
