package piscine

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	Nbrs []int
	Res  int
)

func Calcul(op string) error {
	if len(Nbrs) != 2 {
		return fmt.Errorf("err")
	} else {
		if op == "+" {
			Res = Nbrs[0] + Nbrs[1]
		} else if op == "-" {
			Res = Nbrs[0] - Nbrs[1]
		} else if op == "*" {
			Res = Nbrs[0] * Nbrs[1]
		} else if op == "/" {
			Res = Nbrs[0] / Nbrs[1]
		} else if op == "%" {
			Res = Nbrs[0] % Nbrs[1]
		} else {
			return fmt.Errorf("err")
		}
	}
	return nil
}

func RpnCalcProcess(code string) (int, error) {
	tab := strings.Fields(code)
	operation := ""
	for i := 0; i < len(tab); i++ {
		cell, err := strconv.Atoi(tab[i])
		if err != nil && tab[i] != "+" && tab[i] != "-" && tab[i] != "*" && tab[i] != "/" {
			return 0, fmt.Errorf("err")
		} else if tab[i] == "+" || tab[i] == "-" || tab[i] == "*" || tab[i] == "/" {
			operation = tab[i]
		} else {
			Nbrs = append(Nbrs, cell)
		}

		if len(Nbrs) == 2 && operation != "" {
			err := Calcul(operation)
			if err != nil {
				return 0, fmt.Errorf("err")
			}
			Nbrs = []int{}
			operation = ""
			Nbrs = append(Nbrs, Res)
		}
	}

	return Res, nil
}
