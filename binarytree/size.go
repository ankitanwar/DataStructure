package binarytree

//Size : it will return the size of the binary tree
func Size(root *Node) int {
	if root == nil {
		return 0
	}
	left := Size(root.Left)
	right := Size(root.Right)
	size := left + right + 1
	return size
}

//Sum : It will return the sum of the tree
func Sum(root *Node) int {
	if root == nil {
		return 0
	}
	left := Sum(root.Left)
	right := Sum(root.Right)
	ans := left + right + root.Data.(int)
	return ans
}

//Height : It will return the height of the tree  if edges -> return -1  if node -> return 0
func Height(root *Node) int {
	if root == nil {
		return 0
	}
	down := max(Height(root.Left), Height(root.Right))
	return down + 1

}

//max : it will return the maximum of 2 number
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
