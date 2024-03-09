package piscine

func RecursiveFactorial(nb int) int {
	if nb > 0 && nb < 22 {
		return nb * RecursiveFactorial(nb-1)
	} else if nb == 0 {
		return 1
	}
	return 0
}
