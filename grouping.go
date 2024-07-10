package piscine

import "strings"

func FindChar(word string) []string {
	res := []string{}
	if len(word) >= 3 && word[0] == '(' && word[len(word)-1] == ')' {
		part := ""
		for i := 1; i < len(word); i++ {
			if word[i] == '|' || word[i] == ')' {
				res = append(res, part)
				part = ""
			}
			if word[i] == '|' {
				continue
			}
			part += string(word[i])
		}
	}
	return res
}

func Clean(tab []string) []string {
	res := []string{}
	for _, cs := range tab {
		for j := len(cs) - 1; j >= 0; j-- {
			if cs[j] == '.' || cs[j] == ',' {
				cs = cs[:len(cs)-1]
			} else {
				continue
			}
		}
		res = append(res, cs)
	}
	return res
}

func Grouping(part, word string) []string {
	res := []string{}
	slc := Clean(strings.Split(word, " "))
	chars := FindChar(part)
	for i := 0; i < len(slc); i++ {
		for j := 0; j < len(chars); j++ {
			if strings.Contains(slc[i], chars[j]) {
				res = append(res, slc[i])
			}
		}
	}
	return res
}
