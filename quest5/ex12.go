package main

import (
	"fmt"
)

func IsUpper(s string) string {

	res := []byte(s)

	sol := ""

	for i := 0; i < len(s); i++ {
		if res[i] >= 65 && res[i] <= 90 {
			sol = "true"

		} else {
			sol = "false"
		}
	}

	return sol

}

func main() {
	fmt.Println(IsUpper("HELLO"))
	fmt.Println(IsUpper("HELLO!"))

}
