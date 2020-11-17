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

//CheckMirror : To check whether the given two trees are mirror of each other or not
func CheckMirror(root1, root2 *Node) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}
	return root1.Data == root2.Data && CheckMirror(root1.Left, root2.Right) && CheckMirror(root1.Right, root2.Left)
}
