package piscine

func CountIf(f func(string) bool, tab []string) int {
	res := 0
	for _, str := range tab {
		if f(str) {
			res++
		}
	}
	return res
}
