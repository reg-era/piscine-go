package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if isEven(len(arg)) {
		printStr("I have an odd number of arguments")
	} else {
		printStr("I have an even number of arguments")
	}
}

func isEven(nbr int) bool {
	if nbr%2 == 1 {
		return true
	} else {
		return false
	}
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
