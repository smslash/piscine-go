package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	node := l.Head
	var prev *NodeL

	for node != nil {
		if node.Data == data_ref {
			if prev == nil {
				l.Head = node.Next
			} else {
				prev.Next = node.Next
			}
			if node.Next == nil {
				l.Tail = prev
			}
		} else {
			prev = node
		}
		node = node.Next
	}
}
