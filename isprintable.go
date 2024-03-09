package piscine

func IsPrintable(s string) bool {
	res := false
	str := []rune(s)

	for i := 0; i < len(str); i++ {
		if str[i] > 31 {
			res = true
		} else {
			res = false
			break
		}
	}
	return res
}
