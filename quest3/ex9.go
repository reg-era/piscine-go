package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(BasicAtoi("12345"))
	fmt.Println(BasicAtoi("0000000012345"))
	fmt.Println(BasicAtoi("000000"))
}

func BasicAtoi(s string) int {

	n, _ := strconv.Atoi(s)

	return n

}
