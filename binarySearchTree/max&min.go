package binarysearchtree

//Max : To return the max value in a binary seach tree
func Max(root *Node) *Node {
	if root.Right != nil {
		return Max(root.Right)
	}
	return root
}

//Min : to return the min value in a binary search tree
func Min(root *Node) *Node {
	if root.Left != nil {
		return Min(root.Left)
	}
	return root
}
