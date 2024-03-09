package piscine

func AtoiBase(s string, base string) int {

	if len(base) < 2 {
		return 0
	}

	for i := 0; i < len(base); i++ {
		if base[i] == '-' || base[i] == '+' {
			return 0
		}
	}

	for i := 0; i < len(base)-1; i++ {
		for j := i + 1; j < len(base); j++ {
			if base[j] == base[i] {
				return 0
			}
		}
	}

	str := []rune(s)
	bs := []rune(base)
	bsnbr := 1
	res := 0

	for i := 1; i < len(str); i++ {
		bsnbr = bsnbr * len(bs)
	}

	for i := 0; i < len(str); i++ {
		res += bsnbr * srxind(str[i], bs)
		bsnbr /= len(bs)
	}
	return res
}

func srxind(ch rune, base []rune) int {
	ind := 0
	for i := 0; i < len(base); i++ {
		if base[i] == ch {
			ind = i
		}
	}
	return ind
}
