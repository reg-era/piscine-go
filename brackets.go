package piscine

func Same(a, b byte) bool {
	if (a == '[' && b == ']') || (a == '(' && b == ')') || (a == '{' && b == '}') {
		return true
	} else {
		return false
	}
}

func BracketsProcess(code string) bool {
	str := ""
	for i := 0; i < len(code); i++ {
		if code[i] == '(' || code[i] == '[' || code[i] == '{' || code[i] == ')' || code[i] == ']' || code[i] == '}' {
			str += string(code[i])
		}
		if len(str) >= 2 && Same(str[len(str)-2], str[len(str)-1]) {
			str = str[:len(str)-2]
		}
	}
	if len(str) == 0 {
		return true
	} else {
		return false
	}
}
