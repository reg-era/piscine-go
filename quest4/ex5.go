package main

import (
	"fmt"
)

//3

func Fibonacci(index int) int {

	res := 0

	if index == 1 {
		return 1
	}

	if index == 0 {
		return 0
	}

	res = Fibonacci(index-1) + Fibonacci(index-2)
	return res

}

func main() {
	arg1 := 15
	fmt.Println(Fibonacci(arg1))
}
