package main

import (
	"fmt"
)

func main() {
	s := []int{5, 3, 4, 2, 1, 0}
	SortIntegerTable(s)
	fmt.Println(s)
}

func SortIntegerTable(table []int) {

	for i := 0; i < len(table); i++ {
		for j := i + 1; j < len(table); j++ {
			if table[i] > table[j] {
				tmp := table[i]
				table[i] = table[j]
				table[j] = tmp
			}
		}
	}

}
