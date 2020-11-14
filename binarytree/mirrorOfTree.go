package binarytree

//MirrorOfTree : It will create the mirror of the tree left child is gonna be the right child and right child is left
func MirrorOfTree(root *Node) *Node {
	if root == nil {
		return nil
	}
	left := MirrorOfTree(root.Left)
	right := MirrorOfTree(root.Right)
	root.Right = left
	root.Left = right
	return root
}
