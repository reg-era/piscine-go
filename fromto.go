package piscine

import (
	"strconv"
)

func FromTo(start, end int) string {
	if (start > 10 || start < 0) || (end > 10 || end < 0) {
		return "Invalid"
	}
	res := ""
	if start <= end {
		for i := start; i <= end; i++ {
			if i > start {
				res += ", "
			}
			if i == 10 {
				res += strconv.Itoa(i)
			} else {
				res += "0" + strconv.Itoa(i)
			}
		}
	} else {
		for i := start; i >= end; i-- {
			if i < start {
				res += ", "
			}
			if i == 10 {
				res += strconv.Itoa(i)
			} else {
				res += "0" + strconv.Itoa(i)
			}
		}
	}
	return res + "\n"
}
