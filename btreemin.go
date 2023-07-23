package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	currentNode := root
	for currentNode.Left != nil {
		currentNode = currentNode.Left
	}
	return currentNode
}
