package piscine

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}

	if IsPrime(nb) {
		return nb
	} else {
		return FindNextPrime(nb + 1)
	}
}
