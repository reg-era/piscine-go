package piscine

func Compact(ptr *[]string) int {
	res := 0
	tab := *ptr
	sol := []string{}
	for i := 0; i < len(*ptr); i++ {
		if tab[i] != "" {
			res++
			sol = append(sol, tab[i])
		}
	}
	*ptr = sol
	return res
}
