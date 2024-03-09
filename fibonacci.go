package piscine

func Fibonacci(index int) int {
	// n := 0

	if index == 1 {
		return 1
	}
	if index > 0 {
		return (Fibonacci(index-1) + Fibonacci(index-2))
	}
	if index < 0 {
		return -1
	}

	return 0
}
