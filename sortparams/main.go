package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]

	for i := 0; i < len(arg); i++ {
		for j := i + 1; j < len(arg); j++ {
			if arg[i] > arg[j] {
				arg[i], arg[j] = arg[j], arg[i]
			}
		}
	}
	for i := 0; i < len(arg); i++ {
		s := arg[i]
		str := []rune(s)

		for j := 0; j < len(str); j++ {
			z01.PrintRune(str[j])
		}
		z01.PrintRune('\n')
	}
}
