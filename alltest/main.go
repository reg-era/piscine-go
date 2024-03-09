package main

import (
	student "piscine"

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
		challenge.Function("Slice", student.Slice, solutions.Slice, a...)
	}
}

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
// 	fmt.Printf("%#v\n", piscine.Slice(a, 1))
// 	fmt.Printf("%#v\n", piscine.Slice(a, 2, 4))
// 	fmt.Printf("%#v\n", piscine.Slice(a, -3))
// 	fmt.Printf("%#v\n", piscine.Slice(a, -2, -1))
// 	fmt.Printf("%#v\n", piscine.Slice(a, 2, 0))
// }

// $ go run .
// []string{"algorithm", "ascii", "package", "golang"}
// []string{"ascii", "package"}
// []string{"ascii", "package", "golang"}
// []string{"package"}
// []string(nil)
// {1 2 3 4 5 6 7}
