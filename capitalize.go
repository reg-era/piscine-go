package piscine

func Capitalize(s string) string {
	res := ""
	if IsLower(string(s[0])) {
		res += ToUpper(string(s[0]))
	} else {
		res += string(s[0])
	}
	for i := 1; i < len(s); i++ {
		if s[i-1] >= 32 && s[i-1] <= 47 || s[i-1] >= 58 && s[i-1] <= 64 || s[i-1] >= 91 && s[i-1] <= 96 || s[i-1] >= 123 && s[i-1] <= 127 {
			// res += string('+')
			res += ToUpper(string(s[i]))
		} else {
			res += ToLower(string(s[i]))
		}
	}
	return res
}
