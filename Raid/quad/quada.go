package piscine

import "fmt"

func QuadA(x, y int) {
	if x >= 0 && y >= 0 {
		for i := 1; i <= y; i++ {
			for j := 1; j <= x; j++ {
				if x == y {
					fmt.Print("o")
					break
				}
				if i == 1 || i == y {
					if i == 1 && j == 1 {
						fmt.Print("o")
					} else if i == 1 && j == x {
						fmt.Print("o")
					} else if i == y && j == 1 {
						fmt.Print("o")
					} else if i == y && j == x {
						fmt.Print("o")
					} else {
						fmt.Print("-")
					}
				} else if j == 1 || j == x {
					fmt.Print("|")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
	} else {
		fmt.Println("Please choose another number :)")
		fmt.Println("REMINDER! : Make sure the number is positive ;)")
	}
}
