package piscine

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if node == nil {
		return root
	}

	if node.Left == nil {
		root = transplant(root, node, node.Right)
	} else if node.Right == nil {
		root = transplant(root, node, node.Left)
	} else {
		minNode := BTreeMin(node.Right)
		if minNode.Parent != node {
			root = transplant(root, minNode, minNode.Right)
			minNode.Right = node.Right
			minNode.Right.Parent = minNode
		}
		root = transplant(root, node, minNode)
		minNode.Left = node.Left
		minNode.Left.Parent = minNode
	}
	return root
}

func transplant(root, node, rplc *TreeNode) *TreeNode {
	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		node.Parent.Left = rplc
	} else {
		node.Parent.Right = rplc
	}
	if rplc != nil {
		rplc.Parent = node.Parent
	}
	return root
}
