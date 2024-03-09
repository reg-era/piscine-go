package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	s := os.Args[0]
	str := []rune(s)

	for i := 0; i < len(s); i++ {
		if str[i] == '.' || str[i] == '/' {
		} else {
			z01.PrintRune(str[i])
		}
	}
	z01.PrintRune('\n')
}
