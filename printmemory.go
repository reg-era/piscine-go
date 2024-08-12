package piscine

import "fmt"

func PrintMemory(code [10]byte) {
	res := ""

	for i := 0; i < len(code); i++ {
		if code[i] == 0 {
			res += "00"
		}
		res += ToHex(rune(code[i])) + " "
		if i == 3 || i == 7 || i == 9 {
			res += "\n"
		}
	}
	for i := 0; i < len(code); i++ {
		if code[i] < ' ' {
			res += "."
		} else {
			res += string(code[i])
		}
	}
	fmt.Println(res)
}

func ToHex(nbr int32) string {
	bs := "0123456789abcdef"
	res := ""
	for nbr > 0 {
		tmp := nbr % 16
		res = string(bs[tmp]) + res
		nbr /= 16
	}
	return res
}
