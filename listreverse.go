package piscine

func ListReverse(l *List) {
	var Prev *NodeL
	var Next *NodeL
	newin := l.Head
	for newin != nil {
		Next = newin.Next
		newin.Next = Prev
		Prev = newin
		newin = Next
	}
	l.Head, l.Tail = l.Tail, l.Head
}
