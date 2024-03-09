package piscine

func IterativePower(nb int, power int) int {
	res := 1

	if power < 0 {
		return 0
	}

	for power != 0 {
		res = res * nb
		power = power - 1
	}

	return res
}
