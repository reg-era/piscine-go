package piscine

func DescendAppendRange(max, min int) []int {
	res := []int{}
	if max < min {
		return res
	}
	for i := max; i > min; i-- {
		res = append(res, i)
	}
	return res
}
