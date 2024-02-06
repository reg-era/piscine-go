package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(BasicAtoi2("12345"))
	fmt.Println(BasicAtoi2("0000000012345"))
	fmt.Println(BasicAtoi2("012 345"))
	fmt.Println(BasicAtoi2("Hello World!"))
}

func BasicAtoi2(s string) int {

	n, _ := strconv.Atoi(s)

	return n

}
