package main

import (
	"fmt"
	"os"
)

func main() {
	result := ""
	if len(os.Args) == 3 {
		first := os.Args[1]
		for _, v := range first {
			found := false
			for _, r := range result {
				if v == r {
					found = true
					break
				}
			}
			if !found {
				result += string(v)
			}
		}

		second := os.Args[2]
		for _, v := range second {
			found := false
			for _, r := range result {
				if v == r {
					found = true
					break
				}
			}
			if !found {
				result += string(v)
			}
		}
		fmt.Printf("%s\n", result)
	}
}
