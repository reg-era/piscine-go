package piscine

func StringToIntSlice(str string) []int {
	var res []int
	s := []rune(str)
	for i := 0; i < len(s); i++ {
		res = append(res, int(rune(s[i])))
	}
	return res
}
