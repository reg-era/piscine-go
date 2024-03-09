package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	tnb := []int{}

	if n == 0 {
		z01.PrintRune('0')
	}
	for n > 0 {
		i := n % 10
		tnb = append(tnb, i)
		n = n / 10
	}

	for i := 0; i < len(tnb); i++ {
		for j := i + 1; j < len(tnb); j++ {
			if tnb[i] > tnb[j] {
				tnb[i], tnb[j] = tnb[j], tnb[i]
			}
		}
	}

	for i := 0; i < len(tnb); i++ {
		z01.PrintRune(rune(tnb[i] + '0'))
	}
}
