package main

import (
	"github.com/01-edu/go-tests/lib/challenge"
	"github.com/01-edu/go-tests/lib/chars"
	"github.com/01-edu/go-tests/lib/random"
	"github.com/01-edu/go-tests/solutions"
)

func main() {
	elems := [][]interface{}{
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			1,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			-3,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			2, 4,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			-2, -1,
		},
		{
			[]string{"coding", "algorithm", "ascii", "package", "golang"},
			2, 0,
		},
	}

	s := random.StrSlice(chars.Words)

	elems = append(elems, []interface{}{s, -len(s) - 10, -len(s) - 5})

	for i := 0; i < 3; i++ {
		s = random.StrSlice(chars.Words)
		elems = append(elems, []interface{}{s, random.IntBetween(-len(s)-10, len(s)+10), random.IntBetween(-len(s)-8, len(s)+10)})
	}

	for _, a := range elems {
		challenge.Function("Slice", Slice, solutions.Slice, a...)
	}
}

func isnegative(a []string, nbr int) int {
	if nbr < 0 {
		nbr = len(a) + nbr
	}
	if nbr < 0 {
		nbr = 0
	} else if nbr > len(a) {
		nbr = len(a)
	}
	return nbr
}

func Slice(a []string, nbrs ...int) []string {
	start := 0
	end := 0

	if len(nbrs) == 1 {
		if nbrs[0] < 0 {
			start = len(a) + nbrs[0]
		} else {
			start = nbrs[0]
		}
		return a[start:]
	}

	start = isnegative(a, nbrs[0])
	end = isnegative(a, nbrs[1])
	if start > end {
		return nil
	}
	return a[start:end]
}
