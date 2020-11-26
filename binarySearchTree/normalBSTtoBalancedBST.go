package binarysearchtree

var traversal []interface{}

//BalanceBST : To balance the binary seach tree
func BalanceBST(root *Node) *Node {
	// -> Idea is to do the inorder traversal and store it in array and then create the balanced BST out of it
	Newroot := ConstructTree(traversal, 0, len(traversal)-1)
	return Newroot
}

func doTheInorder(root *Node) {
	if root == nil {
		return
	}
	doTheInorder(root.Left)
	traversal = append(traversal, root.Data)
	doTheInorder(root.Right)
}
