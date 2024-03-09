package piscine

func StrRev(s string) string {
	res := ""

	str := []rune(s)

	for i := len(str) - 1; i >= 0; i-- {
		res += string(str[i])
	}
	return res
}
