package binarytree

//TransformToLeftCloned : It will clone the left root of the tree
func TransformToLeftCloned(root *Node) *Node {
	if root == nil {
		return nil
	}
	left := TransformToLeftCloned(root.Left)
	right := TransformToLeftCloned(root.Right)
	newNode := &Node{
		Data:  root.Data,
		Left:  left,
		Right: nil,
	}
	root.Left = newNode
	root.Right = right
	return root

}

//TransfromLeftClonedToNormal : It will tranform back the left cloned tree to the normal tree
func TransfromLeftClonedToNormal(root *Node) *Node {
	if root == nil {
		return nil
	}
	left := TransformToLeftCloned(root.Left.Left)
	TransformToLeftCloned(root.Right)
	root.Left = left
	return root
}
