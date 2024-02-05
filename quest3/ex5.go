package main

import "fmt"

func main() {
	PrintStr("Hello World!")
}

func PrintStr(s string) {

	cast := []byte(s)

	for i := 0; i < len(cast); i++ {
		fmt.Printf("%c", cast[i])
	}
}
