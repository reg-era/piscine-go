package piscine

func HashCode(str string) string {
	res := ""
	for i := 0; i < len(str); i++ {
		if str[i] < ' ' {
			res += string(byte((int(byte(str[i])) + 33 + len(str)) % 127))
			continue
		}
		res += string(byte((int(byte(str[i])) + len(str)) % 127))
	}
	return res
}
