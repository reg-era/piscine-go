package piscine

func SplitWhiteSpaces(s string) []string {
	var res []string
	ind := -1

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' || s[i] == '\n' {
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
