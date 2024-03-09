package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	var res *NodeL
	n := l
	count := 0

	for n != nil && count < pos {
		count++
		n = n.Next
	}

	if count == pos && n != nil {
		res = n
	}

	return res
}
