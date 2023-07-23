package piscine

func BTreeMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	currentNode := root
	for currentNode.Right != nil {
		currentNode = currentNode.Right
	}
	return currentNode
}
