package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	if len(arg) == 0 {
		fmt.Println("File name missing")
		return
	} else if len(arg) > 1 {
		fmt.Println("Too many arguments")
		return
	}
	cont, _ := os.ReadFile(arg[0])

	fmt.Print(string(cont))
}
