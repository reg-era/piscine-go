package piscine

import (
	"github.com/01-edu/z01"
)

const maxSolutions = 92

var (
	solutions     [maxSolutions]string
	solutionCount int
)

func EightQueens() {
	table := [8][8]bool{}
	solve(table, 0)
	insertionSort(solutions[:solutionCount])
	printSolution(solutions[:solutionCount])
}

func solve(table [8][8]bool, col int) {
	for i := 0; i < 8; i++ {
		if secure(table, i, col) {
			table[i][col] = true
			solve(table, col+1)
			table[i][col] = false
		}
	}
	if col >= 8 {
		addToSolutions(table)
	}
}

func secure(table [8][8]bool, l int, c int) bool {
	var i, j int
	for i = 0; i < c; i++ {
		if table[l][i] {
			return false
		}
	}
	for j = 0; j < l; j++ {
		if table[j][c] {
			return false
		}
	}
	for i, j = l, c; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if table[i][j] {
			return false
		}
	}
	for i, j = l, c; i < 8 && j >= 0; i, j = i+1, j-1 {
		if table[i][j] {
			return false
		}
	}
	for i, j = l, c; i < 8 && j < 8; i, j = i+1, j+1 {
		if table[i][j] {
			return false
		}
	}
	for i, j = l, c; i >= 0 && j < 8; i, j = i-1, j+1 {
		if table[i][j] {
			return false
		}
	}
	return true
}

func addToSolutions(table [8][8]bool) {
	var sol string
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if table[i][j] {
				sol += string(rune(j+1) + '0')
			}
		}
	}
	solutions[solutionCount] = sol
	solutionCount++
}

func insertionSort(arr []string) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

func printSolution(tab []string) {
	for i := 0; i < len(tab); i++ {
		slc := []rune(tab[i])
		for j := 0; j < len(slc); j++ {
			z01.PrintRune(slc[j])
		}
		z01.PrintRune('\n')
	}
}
