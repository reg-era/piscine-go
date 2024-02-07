package main

import "fmt"

func Index(s string, toFind string) int {
	phrase := []rune(s)
	fin := []rune(toFind)

	for i := 0; i <= len(phrase)-len(fin); i++ {
		match := true
		for j := 0; j < len(fin); j++ {
			if phrase[i+j] != fin[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}
