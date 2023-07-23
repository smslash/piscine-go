package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref}

	if l == nil || data_ref < l.Data {
		newNode.Next = l
		return newNode
	}

	currentNode := l
	for currentNode.Next != nil && data_ref > currentNode.Next.Data {
		currentNode = currentNode.Next
	}

	newNode.Next = currentNode.Next
	currentNode.Next = newNode

	return l
}
