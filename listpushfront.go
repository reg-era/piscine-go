package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListPushFront(l *List, data interface{}) {
	wnode := new(NodeL)
	wnode.Data = data
	if l.Head == nil {
		l.Head = wnode
	} else {
		wnode.Next = l.Head
		l.Head = wnode
	}
}
