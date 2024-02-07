package main

import (
	"fmt"
)

func IsAlpha(s string) string {

	res := []byte(s)

	for i := 0; i < len(s); i++ {
		if res[i] >= 65 && res[i] <= 90 && res[i] >= 97 && res[i] <= 122 {
			return "true"

		}
	}

	return "false"
}

func main() {
	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("HelloHowareyou"))
	fmt.Println(IsAlpha("What's this 4?"))
	fmt.Println(IsAlpha("Whatsthis4"))
}
