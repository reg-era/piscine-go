package main

import "fmt"

func FirstRune(s string) rune {

	w := []rune(s)

	return w[len(w)-1]

}

func main() {

	fmt.Printf(string(FirstRune("Hello!")))
	fmt.Printf(string(FirstRune("Salut!")))
	fmt.Printf(string(FirstRune("Ola!")))
	fmt.Printf("\n")

}
