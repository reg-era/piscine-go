package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func QuadD(x, y int) {
	if x > 0 && y > 0 {
		for i := 0; i < y; i++ {
			if i == 0 {
				z01.PrintRune('A')
			}
			if i == y-1 && y-1 != 0 {
				z01.PrintRune('A')
			}
			if i == 0 {
				for j := 0; j < x-1; j++ {
					if j == x-2 {
						z01.PrintRune('C')
					} else {
						z01.PrintRune('B')
					}
				}
				z01.PrintRune('\n')
			}
			if i == y-1 && y-1 != 0 {
				for j := 0; j < x-1; j++ {
					if j == x-2 {
						z01.PrintRune('C')
					} else {
						z01.PrintRune('B')
					}
				}
				z01.PrintRune('\n')
			}
			if i < y-1 && i > 0 {
				for j := 0; j < x; j++ {
					if j == x-1 || j == 0 {
						z01.PrintRune('B')
					} else {
						z01.PrintRune(' ')
					}
				}
				z01.PrintRune('\n')
			}
		}
	}
}

func main() {
	arg := os.Args[1:]

	if len(arg) > 2 {
		fmt.Println("Not a quad function")
		return
	}

	x, err := strconv.Atoi(arg[0])

	if err != nil {
		fmt.Println("Error reading input from stdin:", err)
		os.Exit(1)
	}

	y, err := strconv.Atoi(arg[1])

	if err != nil {
		fmt.Println("Error reading input from stdin:", err)
		os.Exit(1)
	}

	QuadD(x, y)
}
