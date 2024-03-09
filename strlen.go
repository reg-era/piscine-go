package piscine

func StrLen(s string) int {
	res := 0

	tabstr := []rune(s)

	res = len(tabstr)

	// for i := 0; i < len(s); i++ {
	// if (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= 'a' && s[i] <= 'z') {

	//esp := 0
	//if s[i] == ' ' {
	//	esp += 1
	//}
	//res = i + esp + 1
	//}
	//}

	return res
}
