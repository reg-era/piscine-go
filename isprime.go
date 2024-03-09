package piscine

func IsPrime(nb int) bool {
	res := true

	if nb < -6000 {
		return false
	}
	if nb == 0 || nb == 1 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb%2 == 0 {
		return false
	}

	for i := 3; i*i <= nb; i = i + 2 {
		prm := nb % i
		if prm == 0 {
			res = false
		}
	}

	return res
}
