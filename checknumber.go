package piscine

func CheckNumber(arg string) bool {
	for _, c := range arg {
		if c >= 48 && c <= 57 {
			return true
		}
	}
	return false
}
