package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	newin := l.Head
	var res *interface{}
	for newin != nil {
		if comp(newin.Data, ref) {
			res = &newin.Data
		}
		newin = newin.Next
	}
	return res
}
