package main

import (
	"fmt"
)

//64

func RecursivePower(nb int, power int) int {

	res := 0
	if power == 1 {
		return nb
	} else if power == 0 {
		return 1
	}

	res = nb * RecursivePower(nb, power-1)
	return res
}

func main() {
	fmt.Println(RecursivePower(4, 3))
}
