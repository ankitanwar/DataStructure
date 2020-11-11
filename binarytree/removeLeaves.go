package binarytree

//RemoveLeaves : it will remove the leaves s
func RemoveLeaves(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Right == nil && root.Left == nil {
		return nil
	}
	root.Left = RemoveLeaves(root.Left)
	root.Right = RemoveLeaves(root.Right)
	return root
}
