package piscine

import (
	"fmt"
	"strconv"
	"strings"
)

func FindPairs(arg1, arg2 string) string {
	num, _ := strconv.Atoi(arg2)
	if arg1[0] != '[' || arg1[len(arg1)-1] != ']' {
		return "Invalid input."
	} else {
		arg1 = arg1[1:]
		arg1 = arg1[:len(arg1)-1]
	}
	slicestr := strings.Split(arg1, ", ")
	sls := []int{}
	for i := 0; i < len(slicestr); i++ {
		n, err := strconv.Atoi(slicestr[i])
		if err != nil {
			return "Invalid number: " + slicestr[i]
		} else {
			sls = append(sls, n)
		}
	}
	last := [][]int{}
	for i := 0; i < len(sls); i++ {
		for j := i + 1; j < len(sls); j++ {
			if sls[i]+sls[j] == num {
				last = append(last, []int{i, j})
			}
		}
	}
	if len(last) == 0 {
		return "No pairs found."
	}
	res := fmt.Sprintf("Pairs with sum  %v: %v", num, last)
	return res
}
