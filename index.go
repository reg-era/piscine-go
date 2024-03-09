package piscine

func Index(s string, toFind string) int {
	str := []rune(s)
	finder := []rune(toFind)

	if s == "" || toFind == "" {
		return 0
	}

	for i := 0; i < len(str)-len(toFind); i++ {

		match := false

		for j := 0; j < len(finder); j++ {
			if str[i+j] == finder[j] {
				match = true
			} else {
				match = false
			}
		}
		if match {
			return i
		}

	}
	return -1

	// str := []rune(s)
	// res := -1

	// for i := 0; i < len(str); i++ {
	// 	if str[i] == rune(toFind[0]) {
	// 		res = i - 1
	// 	}
	// }
	// return res
}
