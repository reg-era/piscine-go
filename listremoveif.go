package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	var prev *NodeL

	for newin := l.Head; newin != nil; newin = newin.Next {
		if newin.Data == data_ref {
			if prev == nil {
				l.Head = newin.Next
			} else {
				prev.Next = newin.Next
			}
		} else {
			prev = newin
		}
	}
}
