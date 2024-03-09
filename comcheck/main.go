package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	match := []string{"01", "galaxy", "galaxy 01"}
	for _, i := range args {
		for _, j := range match {
			if i == j {
				fmt.Println("Alert!!!")
				return
			}
		}
	}
}
