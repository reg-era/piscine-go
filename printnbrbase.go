package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	bs := []rune(base)

	if len(bs) < 3 && (base != "01" && base != "10") {
		printstr("NV")
		return
	}
	for i := 0; i < len(bs); i++ {
		for j := i + 1; j < len(bs); j++ {
			if bs[i] > bs[j] {
				bs[i], bs[j] = bs[j], bs[i]
			}
		}
	}
	for i := 0; i < len(bs); i++ {
		for j := i + 1; j < len(bs); j++ {
			if bs[i] == bs[j] {
				printstr("NV")
				return
			}
		}
	}

	bs = []rune(base)

	if nbr < 0 {
		printstr("-")
		nbr *= -1
	}
	printstr(bsconv(nbr, bs))
}

func bsconv(nb int, bs []rune) string {
	res := ""
	for nb != 0 {
		div := nb % len(bs)
		res += string(bs[div])
		nb /= len(bs)
	}
	bs = []rune(res)
	res = ""
	for i := 0; i < len(bs)/2; i++ {
		bs[i], bs[len(bs)-1-i] = bs[len(bs)-1-i], bs[i]
	}
	for i := 0; i < len(bs); i++ {
		res += string(bs[i])
	}
	return res
}

func printstr(str string) {
	tab := []rune(str)
	for i := 0; i < len(tab); i++ {
		z01.PrintRune(tab[i])
	}
}
