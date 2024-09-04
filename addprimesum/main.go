package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	} else {
		nb, err := atoi(arg[0])
		if err {
			z01.PrintRune('0')
			z01.PrintRune('\n')
			return
		}
		res := itoa(addprimesum(nb))
		for _, i := range res {
			z01.PrintRune(i)
		}
		z01.PrintRune('\n')
	}
}

func itoa(nb int) string {
	if nb == 0 {
		return "0"
	}
	res := ""
	for nb > 0 {
		n := nb % 10
		res = string(rune('0'+n)) + res
		nb /= 10
	}
	return res
}

func atoi(nb string) (int, bool) {
	res := 0
	if nb == "" {
		return 0, false
	}
	for i := 0; i < len(nb); i++ {
		if nb[i] >= '0' && nb[i] <= '9' {
			res = (res * 10) + (int(nb[i] - '0'))
		} else {
			return 0, true
		}
	}
	return res, false
}

func addprimesum(nbr int) int {
	prim := []int{}
	for i := 1; i <= nbr; i++ {
		if isprim(i) {
			prim = append(prim, i)
		}
	}
	res := 0
	for i := 0; i < len(prim); i++ {
		res += prim[i]
	}
	return res
}

func isprim(nb int) bool {
	if nb == 0 || nb == 1 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb%2 == 0 {
		return false
	}
	for i := 3; i < nb/2; i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
