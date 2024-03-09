package piscine

func TrimAtoi(s string) int {
	str := []rune(s)

	res := 0
	sign := false

	for i := 0; i < len(str); i++ {

		if str[i] == '-' {
			for j := 0; j <= i; j++ {
				if str[j] >= '0' && str[j] <= '9' {
					sign = false
					break
				} else {
					sign = true
				}
			}
		}

		if str[i] >= '0' && str[i] <= '9' {
			if str[i] == ' ' {
				res = 0 - '0'
			}

			res = (int(str[i] - '0')) + res*10
		}
	}

	if sign {
		return res * -1
	} else {
		return res
	}
}
