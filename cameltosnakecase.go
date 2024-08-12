package piscine

func CamelToSnakeCase(str string) string {
	for i := 0; i < len(str); i++ {
		if str[i] < 'a' && str[i] > 'z' && str[i] < 'A' && str[i] > 'Z' && str[i] < '0' && str[i] > '9' {
			return str
		}
		if str[i] >= 'A' && str[i] <= 'Z' && i != len(str)-1 && str[i+1] >= 'A' && str[i+1] <= 'Z' {
			return str
		}
	}
	res := ""
	for i := 0; i < len(str); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' && i != 0 {
			res += "_" + string(str[i])
		} else {
			res += string(str[i])
		}
	}
	return res
}
