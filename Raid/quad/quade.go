package piscine

import "fmt"

func QuadE(x, y int) {
	if x >= 0 && y >= 0 {
		for i := 1; i <= y; i++ {
			for j := 1; j <= x; j++ {
				if x == y {
					fmt.Print("A")
					break
				}
				if i == 1 || i == y {
					if i == 1 && j == 1 {
						fmt.Print("A")
					} else if i == 1 && j == x {
						fmt.Print("C")
					} else if i == y && j == 1 {
						fmt.Print("C")
					} else if i == y && j == x {
						fmt.Print("A")
					} else {
						fmt.Print("B")
					}
				} else if j == 1 || j == x {
					fmt.Print("B")
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
