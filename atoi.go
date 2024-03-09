package piscine

func Atoi(s string) int {
	str := []rune(s)
	res := 0
	sign := 1

	for i := 0; i < len(str); i++ {

		if str[i] == ' ' {
			return 0
		}
		if str[i] >= '0' && str[i] <= '9' {
			res = int(str[i]-'0') + res*10
		}

		if str[i] == '+' || str[i] == '-' {

			for j := 0; j < i; j++ {
				if str[j] != 0 {
					return 0
				}
			}
			if str[i] == '-' {
				sign = -1
			}
		}

	}

	return res * sign
}
