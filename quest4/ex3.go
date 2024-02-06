package main

import (
	"fmt"
)

//64

func IterativePower(nb int, power int) int {

	res := 1

	for i := 1; i <= power; i++ {
		res = res * nb
	}

	return res
}

func main() {
	fmt.Println(IterativePower(4, 3))
}
