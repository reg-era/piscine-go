package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[0:]

	for i := len(arg) - 1; i > 0; i-- {
		s := arg[i]
		str := []rune(s)
		for j := 0; j < len(str); j++ {
			z01.PrintRune(str[j])
		}
		z01.PrintRune('\n')
	}
}
