package main

import (
	"fmt"
	"math"
)

//2
//0

func Sqrt(nb int) int {

	var rac float64

	rac = math.Sqrt(float64(nb))

	if rac != float64(int(rac)) {
		return 0
	}

	if rac == 1 {
		return 1
	}

	s := int(rac)
	return s

}

func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(3))
}
