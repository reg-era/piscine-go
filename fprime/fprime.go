package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	nbr, err := strconv.Atoi(os.Args[1])
	res := ""

	if err != nil || nbr == 0 || nbr == 1 || len(os.Args) != 2 {
		return
	}

	for i := 2; i <= nbr; i++ {
		for nbr%i == 0 {
			if i != nbr {
				res += strconv.Itoa(i) + "*"
			} else {
				res += strconv.Itoa(i)
			}
			nbr = nbr / i
		}
	}

	fmt.Println(res)
}
