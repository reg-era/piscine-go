package piscine

func Itoa(nbr int) string {
	res := ""
	sol := ""
	tmp := false
	if nbr < 0 {
		nbr = nbr * -1
		tmp = true
	}
	for nbr > 0 {
		res += string(byte(nbr%10) + '0')
		nbr /= 10
	}
	for i := len(res) - 1; i >= 0; i-- {
		sol += string(res[i])
	}
	if tmp {
		sol = "-" + sol
	}
	if sol == "" {
		return "0"
	}
	return sol
}
