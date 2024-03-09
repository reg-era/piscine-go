package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	var res, next, prevres, current *NodeI
	if l == nil || l.Next == nil {
		return l
	}
	for l != nil {
		current = l
		l = l.Next
		if res == nil || current.Data <= res.Data {
			current.Next = res
			res = current
		} else {
			prevres = res
			next = res.Next
			for next != nil && current.Data > next.Data {
				prevres = next
				next = next.Next
			}
			prevres.Next = current
			current.Next = next
		}
	}
	return res
}
