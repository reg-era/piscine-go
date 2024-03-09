package piscine

func AlphaCount(s string) int {
	str := []rune(s)
	res := 0
	for i := 0; i < len(str); i++ {
		if (str[i] >= 'a' && str[i] <= 'z') || (str[i] >= 'A' && str[i] <= 'Z') {
			res += 1
		}
	}
	return res
}
