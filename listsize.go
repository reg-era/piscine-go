package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListSize(l *List) int {
	res := 0
	n := l.Head
	for n != nil {
		res++
		n = n.Next
	}
	return res
}
