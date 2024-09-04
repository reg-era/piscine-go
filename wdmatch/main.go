package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]

	if len(arg) == 2 {
		if res, ok := wdmatch(arg[0], arg[1]); ok {
			for _, c := range res {
				z01.PrintRune(c)
			}
			z01.PrintRune('\n')
		}
	}
}

func wdmatch(str1, str2 string) (string, bool) {
	res := ""

	tab := []rune(str2)
	for _, c := range str1 {
		if nb, ok := match(c, tab); ok {
			res += string(c)
			tab = tab[nb:]
		}
	}

	if res == str1 {
		return str1, true
	} else {
		return "", false
	}
}

func match(r rune, tab []rune) (int, bool) {
	for i, c := range tab {
		if c == r {
			return i, true
		}
	}
	return 0, false
}
