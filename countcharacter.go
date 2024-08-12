package piscine

func CountChar(str string, c rune) int {
	count := 0
	for _, ch := range str {
		if ch == c {
			count++
		}
	}
	return count
}
