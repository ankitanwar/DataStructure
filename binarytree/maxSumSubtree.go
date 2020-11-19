package binarytree

//MaxSumSubTree : To find the max sum of all the sub trees
func MaxSumSubTree(root *Node, sum int) int {
	if root == nil {
		return 0
	}
	currentSum := root.Data.(int) + MaxSumSubTree(root.Left, sum) + MaxSumSubTree(root.Right, sum)
	sum = max(sum, currentSum)
	return currentSum
}
