package main

import (
	"fmt"
)

func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}

func IterativeFactorial(nb int) int {
    
    res := 1
    
    for i := 1 ; i < nb+1 ; i++ {
        res = i*res
    }
    
    return res
}