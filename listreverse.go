package piscine

func ListReverse(l *List) {
	if l == nil || l.Head == nil {
		return
	}

	var prev *NodeL
	curr := l.Head
	l.Tail, l.Head = l.Head, l.Tail

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	l.Head = prev
}
