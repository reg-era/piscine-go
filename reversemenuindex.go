package piscine

func ReverseMenuIndex(menu []string) []string {
	res := []string{}
	for i := len(menu) - 1; i >= 0; i-- {
		res = append(res, menu[i])
	}
	return res
}
