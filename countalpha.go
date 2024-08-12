package piscine

func CountAlpha(s string) int {
	count := 0
	for _, c := range s {
		if (c >= 65 && c <= 90) || (c >= 97 && c <= 122) {
			count++
		}
	}
	return count
}
