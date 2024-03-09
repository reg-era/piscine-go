package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printint(num int) {
	a := '0'
	b := '0'

	if num < 0 {
		z01.PrintRune('-')
	}
	for i := 0; i < num/10; i++ {
		a++
	}
	z01.PrintRune(a)
	for i := 0; i < num%10; i++ {
		b++
	}
	z01.PrintRune(b)
}

func printstr(s string) {
	for _, str := range s {
		z01.PrintRune(str)
	}
}

func main() {
	points := &point{}

	setPoint(points)
	printstr("x = ")
	printint(points.x)
	printstr(", y = ")
	printint(points.y)
	printstr("\n")
}
