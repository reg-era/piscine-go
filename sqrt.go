package piscine

func Sqrt(nb int) int {
	res := 0

	for i := 0; i <= nb; i++ {
		rac := i * i

		if rac == nb {
			res = i
			break
		}
	}

	return res
}
