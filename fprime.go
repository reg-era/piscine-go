package piscine

import (
	"strconv"
)

func Fprime(nbr int) string {
	res := ""
	for i := 2; i <= nbr; i++ {
		if IsPrime(i) && nbr%i == 0 {
			res += strconv.Itoa(i) + "*"
			nbr /= i
			i = 2
		}
	}
	return res
}
