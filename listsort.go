package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func merge(l1 *NodeI, l2 *NodeI) *NodeI {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var result *NodeI
	if l1.Data < l2.Data {
		result = l1
		result.Next = merge(l1.Next, l2)
	} else {
		result = l2
		result.Next = merge(l1, l2.Next)
	}

	return result
}

func split(l *NodeI) (*NodeI, *NodeI) {
	if l == nil {
		return nil, nil
	}

	slow := l
	fast := l.Next
	for fast != nil {
		fast = fast.Next
		if fast != nil {
			slow = slow.Next
			fast = fast.Next
		}
	}

	l1 := l
	l2 := slow.Next
	slow.Next = nil

	return l1, l2
}

func ListSort(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}

	l1, l2 := split(l)
	l1 = ListSort(l1)
	l2 = ListSort(l2)

	return merge(l1, l2)
}
