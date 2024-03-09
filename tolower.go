package piscine

func ToLower(s string) string {
	res := ""
	str := []rune(s)

	for i := 0; i < len(str); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			res += string(str[i] + 32)
		} else {
			res += string(str[i])
		}
	}
	return res
}
