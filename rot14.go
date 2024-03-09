package piscine

func Rot14(s string) string {
	str := []rune(s)
	res := ""
	for i := 0; i < len(str); i++ {
		if (str[i] >= 'a' && str[i] < 'm') || (str[i] >= 'A' && str[i] < 'M') {
			res += string(str[i] + 14)
		} else if (str[i] >= 'm' && str[i] <= 'z') || (str[i] >= 'M' && str[i] <= 'Z') {
			res += string(str[i] - 12)
		} else {
			res += string(str[i])
		}
	}
	return res
}
