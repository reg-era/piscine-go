package piscine

func Gcd(in1, in2 int) int {
	finish := 0
	if in1 > in2 {
		finish = in1
	} else {
		finish = in2
	}
	for finish > 0 {
		tmp1 := float64(in1) / float64(finish)
		tmp2 := float64(in2) / float64(finish)

		if tmp1 == float64(in1/finish) && tmp2 == float64(in2/finish) {
			return finish
		}
		finish--
	}
	return 0
}
