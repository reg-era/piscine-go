package main

import "fmt"

func printCombinations(n int) {
	if n <= 0 || n > 10 {
		fmt.Println("Invalid value of n. Please provide a value such that 0 < n <= 10.")
		return
	}

	digits := make([]int, n)
	used := make([]bool, 10) // To keep track of used digits

	generateCombinations(digits, used, 0)
}

func generateCombinations(digits []int, used []bool, index int) {
	if index == len(digits) {
		printDigits(digits)
		return
	}

	for i := 0; i < 10; i++ {
		if !used[i] {
			digits[index] = i
			used[i] = true
			generateCombinations(digits, used, index+1)
			used[i] = false // Backtrack
		}
	}
}

func printDigits(digits []int) {
	for i, digit := range digits {
		if i > 0 {
			fmt.Print("")
		}
		fmt.Print(digit)
	}
	fmt.Printf(",")
}

func main() {

	printCombinations(9)
}
