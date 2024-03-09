package piscine

func isnegative(a []string, nbr int) int {
	if nbr < 0 {
		nbr = len(a) + nbr
	}
	if nbr < 0 {
		nbr = 0
	} else if nbr > len(a) {
		nbr = len(a)
	}
	return nbr
}

func Slice(a []string, nbrs ...int) []string {
	start := 0
	end := 0

	if len(nbrs) == 1 {
		if nbrs[0] < 0 {
			start = len(a) + nbrs[0]
		} else {
			start = nbrs[0]
		}
		return a[start:]
	}

	start = isnegative(a, nbrs[0])
	end = isnegative(a, nbrs[1])
	if start > end {
		return nil
	}
	return a[start:end]
}
