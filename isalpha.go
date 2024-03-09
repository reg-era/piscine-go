package piscine

func IsAlpha(s string) bool {
	res := false
	str := []rune(s)

	for i := 0; i < len(str); i++ {
		if (str[i] >= 'a' && str[i] <= 'z') || (str[i] >= 'A' && str[i] <= 'Z') || (str[i] >= '0' && str[i] <= '9') {
			res = true
		} else {
			res = false
			break
		}
	}
	return res
}
