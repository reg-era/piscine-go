package main

import (
	"fmt"
)

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}

func StrRev(s string) string {

	cast := []byte(s)
	revers := make([]byte, len(cast))

	for i := len(cast) - 1; i >= 0; i-- {

		revers[len(cast)-1-i] = cast[i]
	}

	return string(revers)
}
