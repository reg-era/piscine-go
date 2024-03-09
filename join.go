package piscine

func Join(strs []string, sep string) string {
	res := ""

	for i := 0; i < len(strs); i++ {
		if i == len(strs)-1 {
			res += strs[i]
			break
		}
		res += strs[i] + sep
	}

	return res
}
