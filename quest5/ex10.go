package main

import (
	"fmt"
)

// and && or ||

func IsNumeric(s string) string {

	res := []byte(s)
	for i := 0; i < len(s); i++ {
		if res[i] < 48 || res[i] > 57 {
			return "false"
		}
	}
	return "true"

}

func main() {
	fmt.Println(IsNumeric("010203"))
	fmt.Println(IsNumeric("01a,02,03"))

}
