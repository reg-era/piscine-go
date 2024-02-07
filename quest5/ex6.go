package main

import (
	"fmt"
)

func ToUpper(s string) string {

	res := []byte(s)

	for i := 0; i < len(s); i++ {
		if res[i] >= 97 && res[i] <= 122 {
			res[i] -= 32
		}
	}

	return string(res)
}

func main() {
	fmt.Println(ToUpper("Hello! How are you?"))
}
