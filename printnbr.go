package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	x := 1

	if n < 0 {
		x = -1
		z01.PrintRune('-')
	}

	if n != 0 {

		y := (n / 10) * x
		if y != 0 {
			PrintNbr(y)
		}
		z := (n % 10 * x) + '0'
		z01.PrintRune(rune(z))

	} else {
		z01.PrintRune('0')
	}

}
