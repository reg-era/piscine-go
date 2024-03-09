package piscine

func IterativeFactorial(nb int) int {
	res := nb
	if nb > 0 && nb < 22 {
		for i := nb - 1; i >= 1; i-- {
			res *= i
		}
		return res
	} else if nb == 0 {
		return 1
	}
	return 0
}
