package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	in := true
	out := true
	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 {
			in = false
		}
		if f(a[i], a[i+1]) < 0 {
			out = false
		}
	}
	return in || out
}
