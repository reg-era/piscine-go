package main

import "fmt"

func main4() {
	IsNegative(1)
	IsNegative(0)
	IsNegative(-1)
}

func IsNegative(nbr int) {

	if nbr >= 0 {
		fmt.Printf("F\n")
	}

	if nbr < 0 {
		fmt.Printf("T\n")
	}

}
