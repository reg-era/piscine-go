package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func QuadC(x, y int) {
	if x > 0 && y > 0 {
		if x == 1 && y == 1 {
			z01.PrintRune('A')
			z01.PrintRune('\n')
			return
		} else if x == 1 {
			z01.PrintRune('A')
			for s := 1; s < y-1; s++ {
				if s <= 0 {
					break
				} else {
					z01.PrintRune('\n')
					z01.PrintRune('B')
				}
				if y == 1 {
					z01.PrintRune('\n')
					z01.PrintRune('A')
					return
				}
			}
			z01.PrintRune('\n')
			z01.PrintRune('C')
			z01.PrintRune('\n')
		} else if y == 1 {
			z01.PrintRune('A')
			for i := 1; i < x-1; i++ {
				z01.PrintRune('B')
			}
			z01.PrintRune('A')
			z01.PrintRune('\n')
			return
		} else {
			z01.PrintRune('A')
			for i := 1; i < x-1; i++ {
				z01.PrintRune('B')
			}
			z01.PrintRune('A')
			z01.PrintRune('\n')
			for s := 1; s < y-1; s++ {
				z01.PrintRune('B')
				for j := 1; j < x-1; j++ {
					z01.PrintRune(' ')
				}
				z01.PrintRune('B')
				z01.PrintRune('\n')
			}
			z01.PrintRune('C')
			for i := 1; i < x-1; i++ {
				z01.PrintRune('B')
			}
			z01.PrintRune('C')
			z01.PrintRune('\n')
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

	QuadC(x, y)
}
