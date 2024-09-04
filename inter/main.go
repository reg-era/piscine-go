package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]

	if len(arg) == 2 {
		res := inter(arg[0], arg[1])
		for _, c := range res {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}

func inter(str1, str2 string) string {
	m2 := make(map[rune]int)

	for _, c := range str2 {
		m2[c] = 21
	}

	res := ""
	for i, c := range str1 {
		if _, ok := m2[c]; ok {
			if !contain(c, str1[:i]) {
				res += string(c)
			}
		}
	}
	return res
}

func contain(s rune, ss string) bool {
	for _, c := range ss {
		if c == s {
			return true
		}
	}
	return false
}
