package piscine

func RecursivePower(nb int, power int) int {
	res := 1

	if power < 0 {
		return 0
	}

	if power == 0 {
		return 1
	}

	if power != 0 {
		res = nb
	}

	return RecursivePower(nb, power-1) * res
}
