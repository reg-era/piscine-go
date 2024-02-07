package main

import (
	"fmt"
)

func IsLower(s string) string {

	res := []byte(s)

	sol := ""

	for i := 0; i < len(s); i++ {
		if res[i] >= 97 && res[i] <= 122 {
			sol = "true"

		} else {
			sol = "false"
		}
	}

	return sol

}

func main() {
	fmt.Println(IsLower("hello"))
	fmt.Println(IsLower("hello!"))

}
