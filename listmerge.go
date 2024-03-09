package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1 == nil || l1.Head == nil {
		l1.Head = l2.Head
		l1.Tail = l2.Tail
		return
	}
	if l2 == nil || l2.Head == nil {
		return
	}
	l1.Tail.Next = l2.Head
	l1.Tail = l2.Tail
}
