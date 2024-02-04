package main

import "fmt"

func main6() {

	PrintComb2()

}

func PrintComb2() {

	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {

			for k := 0; k <= 9; k++ {
				for l := 0; l <= 9; l++ {

					fmt.Printf("%d%d %d%d,", i, j, k, l)

				}

			}
		}

	}

}
