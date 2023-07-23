package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		nodeCount := len(queue)
		for i := 0; i < nodeCount; i++ {
			currentNode := queue[0]
			queue = queue[1:]
			_, err := f(currentNode.Data)
			if err != nil {
				panic(err)
			}
			if currentNode.Left != nil {
				queue = append(queue, currentNode.Left)
			}
			if currentNode.Right != nil {
				queue = append(queue, currentNode.Right)
			}
		}
	}
}
