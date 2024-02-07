package main

import (
	"fmt"
)

func BasicJoin(elems []string) string {

	somme := ""
	for i := 0; i < len(elems); i++ {
		somme += elems[i]
	}

	return somme
}

func main() {
	elems := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(BasicJoin(elems))
}
