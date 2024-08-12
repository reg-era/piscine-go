package piscine

import "github.com/01-edu/z01"

func Revwstr(str string) {
}

// func split(str, sb string) []string {
// 	for i := 0; i < len(str); i++ {
// 		if str[i] == sb[0] && match(str[i:], sb) {
// 		}
// 	}
// 	return
// }

func Match(str1, str2 string) (bool, int) {
	tmp := 1
	for i := 0; i < len(str1); i++ {
		if str1[i] == str2[i] {
			tmp++
			continue
		} else if len(str2) == tmp {
			return true, i
		}
	}
	return false, 0
}

func Pr(str string) {
	for i := 0; i < len(str); i++ {
		z01.PrintRune(rune(str[i]))
	}
}
