package quest3

import "fmt"

func main() {

	n := 0
	PointOne(&n)
	fmt.Println(n)

}

func PointOne(n *int) {

	a := 1
	*n = a

}
