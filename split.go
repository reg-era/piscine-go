package piscine

func Split(s, sep string) []string {
	var res []string
	str := make([]string, len(sep)-1)
	for i := 0; i < len(sep); i++ {
		str = append(str, string(sep[i]))
	}

	ind := -1

	for i := 0; i < len(s); i++ {
		if string(s[i]) == str[i] {
			if ind != -1 {
				res = append(res, s[ind:i])
				ind = -1
			}
		} else if ind == -1 {
			ind = i
		}
	}
	if ind != -1 {
		res = append(res, s[ind:])
	}

	return res
}
