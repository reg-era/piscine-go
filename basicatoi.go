package piscine

/*func BasicAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	return BasicAtoirec(s, 0)
}

func BasicAtoirec(s string, res int) int {
	if len(s) == 0 {
		return res
	}
	if s[0] >= '0' && s[0] <= '9' {
		toint := int(s[0] - '0')
		return BasicAtoirec(s[1:], res*10+toint)
	}
	return res
}*/

func BasicAtoi(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			res = int(s[i]-'0') + res*10
		}
	}
	return res
}
