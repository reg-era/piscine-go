package piscine

import "fmt"

func Options(arg []string) {
	alp := make(map[byte]bool)
	x := 26
	ind := byte('z')
	for x > 0 {
		alp[ind] = false
		ind--
		x--
	}

	for i := 0; i < len(arg); i++ {
		if arg[i][0] != '-' {
			fmt.Println("Invalid Option")
			return
		}
		str := arg[i][1:]
		for j := 0; j < len(str); j++ {
			if str[j] < 'a' && str[j] > 'z' {
				fmt.Println("Invalid Option")
				return
			}
			if _, ok := alp[str[j]]; ok {
				alp[str[j]] = true
			}
		}
	}
	res := "000000"
	ind = byte('z')
	for x < 26 {
		if x == 2 || x == 10 || x == 18 {
			res += " "
		}
		if alp[ind] {
			res += "1"
		} else {
			res += "0"
		}
		ind--
		x++
	}
	fmt.Println(res)
}
