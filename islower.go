package piscine

func IsLower(s string) bool {
	res := false
	str := []rune(s)

	for i := 0; i < len(str); i++ {
		if str[i] >= 'a' && str[i] <= 'z' {
			res = true
		} else {
			res = false
			break
		}
	}
	return res
}
