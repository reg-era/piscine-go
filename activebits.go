package piscine

func ActiveBits(n int) int {
	res := 0
	for n > 0 {
		if n%2 != 0 {
			res++
		}
		n = n / 2
	}
	return res
}
