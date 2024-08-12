package piscine

func Append(min, max int) []int {
	if max <= min {
		return []int{}
	}
	res := make([]int, max-min)
	val := min
	for i := 0; i < max-min; i++ {
		res[i] = val
		val++
	}
	return res
}
