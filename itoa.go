package piscine

func Itoa(num int) string {
	res := ""
	sign := false

	if num == 0 {
		return "0"
	}
	if num < 0 {
		sign = true
		num = -num
	}
	for num > 0 {
		digit := num % 10
		res = string(digit+'0') + res
		num /= 10
	}
	if sign {
		res = "-" + res
	}

	return res
}
