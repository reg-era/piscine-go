package main

import "fmt"

func FirstRune(s string, toFind string) int {

    m := []rune(s)
    f := []rune(toFind)
    
    for i := 0 ; i < len(m) ; i++ {
        
        for j := 0 ; j < len(f) ; j++ {
            
            if m[i] == f[j] {
                return i
            }
    } 

}

func main() {

	fmt.Printf(string(FirstRune("Hello!", "l")))
	fmt.Printf(string(FirstRune("Salut!","alu")))
	fmt.Printf(string(FirstRune("Ola!", "hOl")))

}
