package main

import (
	"fmt"
)

func main() {
	arg := 4
	fmt.Println(RecursiveFactorial(arg))
}

func RecursiveFactorial(nb int) int {

	if nb <= 1 {
		return 1
	}

	return nb * RecursiveFactorial(nb-1)

}
