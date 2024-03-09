package main

import (
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	fnb := os.Args[1]
	lnb := os.Args[3]
	opp := os.Args[2]
	if opp != "+" && opp != "-" && opp != "*" && opp != "/" && opp != "%" {
		return
	}
	if convitoa(lnb) > 2147483647 || convitoa(lnb) < (-2147483647) || convitoa(fnb) > 2147483647 || convitoa(fnb) < (-2147483647) {
		return
	}
	if opp == "/" && lnb == "0" {
		os.Stdout.WriteString("No division by 0\n")
		return
	}
	if opp == "%" && lnb == "0" {
		os.Stdout.WriteString("No modulo by 0\n")
		return
	}

	for _, ch := range fnb {
		if ch > '9' && ch > '0' {
			return
		}
	}
	for _, ch := range lnb {
		if ch > '9' && ch > '0' {
			return
		}
	}
	res := 0
	switch opp {
	case "+":
		res = convitoa(fnb) + convitoa(lnb)
	case "-":
		res = convitoa(fnb) - convitoa(lnb)
	case "*":
		res = convitoa(fnb) * convitoa(lnb)
	case "/":
		res = convitoa(fnb) / convitoa(lnb)
	case "%":
		res = convitoa(fnb) % convitoa(lnb)
	}
	if res >= 2147483647 || res <= (-2147483647) {
		return
	}
	os.Stdout.WriteString(convatoi(res) + "\n")
}

func convitoa(s string) int {
	res := 0
	vd := 1
	nb := []rune(s)
	if s[0] == '-' {
		vd = -1
		for i := 1; i < len(nb); i++ {
			res = int((nb[i] - '0')) + (res * 10)
		}
	} else {
		for i := 0; i < len(nb); i++ {
			res = int((nb[i] - '0')) + (res * 10)
		}
	}
	return res * vd
}

func convatoi(nb int) string {
	res := ""
	vd := false
	if nb < 0 {
		nb = nb * (-1)
		vd = true
	}
	for nb != 0 {
		prn := nb % 10
		res = string(prn+'0') + res
		nb /= 10
	}
	if vd {
		res = "-" + res
	}
	return res
}
