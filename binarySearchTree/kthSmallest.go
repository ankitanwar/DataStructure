package binarysearchtree

var counter int
var ans int

//KthSmallest : to find the kth smallest element in the bst
func KthSmallest(root *Node, K int) int {
	if root == nil {
		return -1
	}
	KthSmallest(root.Left, K)
	counter++
	if counter == K {
		ans = root.Data.(int)
	}
	KthSmallest(root.Right, K)
	return ans
}
