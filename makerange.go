package piscine

func MakeRange(min, max int) []int {
	len := max - min
	if max <= min {
		return nil
	}
	res := make([]int, max-min)

	inc := min
	for i := 0; i < len; i++ {
		// if min != max {
		res[i] = inc
		inc++
		//}
	}

	return res
}
