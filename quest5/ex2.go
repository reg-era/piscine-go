package main

import "fmt"

func FirstRune(s string, n int) rune {

	w := []rune(s)

	if n <= 0 || n > len(w) {

		return 0
	}

	return w[n-1]

}

func main() {

	fmt.Printf(string(FirstRune("Hello!", 3)))
	fmt.Printf(string(FirstRune("Salut!", 2)))
	fmt.Printf(string(FirstRune("Bye!", -1)))
	fmt.Printf(string(FirstRune("Bye!", 5)))
	fmt.Printf(string(FirstRune("Ola!", 4)))
	fmt.Printf("\n")

}
