package piscine

func Cleanstr(str string) string {
	res := ""
	block := false
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			block = true
		}
		if block {
			res += string(str[i])
		}
		if str[i] == ' ' {
			block = false
		}
	}
	return res
}
