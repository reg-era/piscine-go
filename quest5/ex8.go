package main

import (
	"fmt"
)

func ToLower(s string) string {

	res := []byte(s)

	for i := 0; i < len(s); i++ {
		if res[i] >= 32 && res[i] <= 47 && res[i+1] >= 97 && res[i+1] <= 122 {
			res[i+1] -= 32

		}
	}

	return string(res)
}

func main() {
	fmt.Println(ToLower("Hello! How are you? How+are+things+4you?"))
}
