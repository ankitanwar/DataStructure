package binarysearchtree

//Sum : it will return the sum of the binary tree
func Sum(root *Node) int {
	if root == nil {
		return 0
	}
	left := Sum(root.Left)
	right := Sum(root.Right)
	return root.Data.(int) + left + right
}
