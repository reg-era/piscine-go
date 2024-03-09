package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for i := 0; i < len(a); i++ {

		b := []rune(a[i])

		for j := 0; j < len(a[i]); j++ {
			z01.PrintRune(b[j])
		}
		z01.PrintRune('\n')
	}
}
