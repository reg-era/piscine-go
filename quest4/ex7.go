package main

import (
	"fmt"
)

//true
//false

func IsPrime(nb int) bool {

	var s bool = true

	for i := 2; i <= nb/2; i++ {

		div := nb % i

		if div == 0 {

			s = false
		}
	}

	return s
}

func main() {
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(4))
}
