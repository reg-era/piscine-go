package main

import "fmt"

func main() {

	a := 13
	b := 2

	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}

func UltimateDivMod(a *int, b *int) {
	var x, y int

	x = *a / *b
	y = *a % *b

	*a = x
	*b = y
}
