package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListLast(l *List) interface{} {
	if l.Tail == nil {
		return nil
	}
	return l.Tail.Data
}
