package piscine

func Map(f func(int) bool, a []int) []bool {
	sol := make([]bool, len(a))
	for i, res := range a {
		sol[i] = f(res)
	}
	return sol
}
