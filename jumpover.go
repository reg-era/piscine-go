package piscine

func JumpOver(str string) string {
	res := ""
	s := []rune(str)

	if len(str) < 3 {
		return "\n"
	}
	for i := 2; i < len(str); i += 3 {
		res += string(s[i])
		if i >= len(str)-3 {
			res += "\n"
		}
	}

	return res
}
