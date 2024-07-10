package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	tab := make([][]int, 9)

	if len(args) != 9 {
		fmt.Println("Error")
		return
	}

	for i := 0; i < 9; i++ {
		lign := args[i]

		if len(lign) != 9 {
			fmt.Println("Error")
			return
		}

		for j := 0; j < len(lign)-1; j++ {
			for k := 1 + j; k < len(lign); k++ {
				if lign[j] == lign[k] && lign[j] != '.' {

					fmt.Println("Error")
					return
				}
			}
		}

		tab[i] = make([]int, 9)

		for j := 0; j < 9; j++ {
			char := lign[j]

			if char != '.' {
				if char >= '1' && char <= '9' {
					tab[i][j] = int(char - '0')
				} else {
					fmt.Println("Error")
					return
				}
			}
		}
	}

	solve(tab)
}

func solve(tab [][]int) bool {

	lign, col := srchnext(tab)

	if lign == -99 && col == -99 {

		printres(tab)
		return true
	}

	for num := 1; num <= 9; num++ {

		if secure(tab, lign, col, num) {

			tab[lign][col] = num

			if solve(tab) {
				return true
			}
			tab[lign][col] = 0
		}
	}

	return false
}

func srchnext(tab [][]int) (int, int) {

	for i := 0; i < len(tab); i++ {
		for j := 0; j < len(tab[i]); j++ {
			if tab[i][j] == 0 {
				return i, j
			}
		}
	}
	return -99, -99
}

func secure(tab [][]int, lign, col, num int) bool {

	for i := 0; i < 9; i++ {
		if tab[lign][i] == num || tab[i][col] == num {
			return false
		}
	}

	newlign, newcol := 3*(lign/3), 3*(col/3)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if tab[newlign+i][newcol+j] == num {
				return false
			}
		}
	}

	return true
}

func printres(tab [][]int) {
	for i := 0; i < len(tab); i++ {
		for j := 0; j < len(tab[i]); j++ {
			fmt.Print(tab[i][j])

			if j != len(tab[i])-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}