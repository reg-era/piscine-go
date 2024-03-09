package piscine

func NRune(s string, n int) rune {
	res := []rune(s)

	if n <= 0 || n > len(res) {
		return 0
	}
	return res[n-1]
}
