package main

import (
	"fmt"
)

//5
//5

func FindNextPrime(nb int) int {

	pr := nb
	var s bool = true

	for i := 2; i <= pr/2; i++ {
		div := pr % i
		if div == 0 {
			s = false
		}
	}

	if s == false {
		return FindNextPrime(pr + 1)
	}

	return pr

}

func main() {
	fmt.Println(FindNextPrime(5))
	fmt.Println(FindNextPrime(4))
}
