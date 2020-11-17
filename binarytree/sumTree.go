package binarytree

//SumTree : To convert the tree where each node contains the sum of the left and right sub trees of the original tree. The values of leaf nodes are changed to 0.
func SumTree(root *Node) int {
	if root == nil {
		return 0
	}
	oldValue := root.Data.(int)
	left := SumTree(root.Left)
	Right := SumTree(root.Right)
	root.Data = left + Right
	return oldValue + root.Data.(int)
}
