package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]

	for i := 0; i < len(arg); i++ {
		s := atoi(arg[i])

		if arg[0] == "--upper" {
			if s >= 1 && s <= 26 {
				z01.PrintRune(rune(s) + 64)
			}
		} else if s >= 1 && s <= 26 {
			z01.PrintRune(rune(s) + 96)
		} else {
			z01.PrintRune(' ')
		}
		if i == len(arg)-1 {
			z01.PrintRune('\n')
		}
	}
}

func atoi(s string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0
		}
		result = result*10 + int(c-'0')
	}
	return result
}
