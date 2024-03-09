package piscine

func Any(f func(string) bool, a []string) bool {
	res := false
	for _, str := range a {
		if f(string(str)) {
			res = true
			break
		}
	}
	return res
}
