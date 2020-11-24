package binarysearchtree

//Size : It will return the  size of the binary tree
func Size(root *Node) int {
	if root == nil {
		return 0
	}
	left := Size(root.Left)
	right := Size(root.Right)
	return left + right + 1

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
