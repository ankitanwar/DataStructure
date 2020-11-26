package binarysearchtree

var res interface{}
var count int

//KthLargest : To find the kth largest element in the tree
func KthLargest(root *Node, K int) interface{} {
	if root == nil {
		return nil
	}
	KthLargest(root.Right, K)
	count++
	if K == count {
		res = root.Data
	}

	KthLargest(root.Left, K)
	return res
}
