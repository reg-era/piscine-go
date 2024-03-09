package piscine

func IsUpper(s string) bool {
	res := false
	str := []rune(s)

	for i := 0; i < len(str); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			res = true
		} else {
			res = false
			break
		}
	}

	return res
}
