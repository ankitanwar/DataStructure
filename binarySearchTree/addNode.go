package binarysearchtree

//AddNode : To add a node in a BST
func AddNode(root, add *Node) *Node {
	if root == nil {
		return add
	}
	if add.Data.(int) > root.Data.(int) {
		root.Right = AddNode(root.Right, add)
	} else if root.Data.(int) < add.Data.(int) {
		root.Right = AddNode(root.Left, add)
	} else {
		return nil
	}
	return root

}
