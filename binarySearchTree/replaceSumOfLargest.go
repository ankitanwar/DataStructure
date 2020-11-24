package binarysearchtree

var sum int = 0

// ReplaceSumWithLargest : It will replace the data of every node with the sum of largest node
func ReplaceSumWithLargest(root *Node) {
	if root == nil {
		return
	}
	ReplaceSumWithLargest(root.Right)
	oldData := root.Data.(int) //IMP otherwise dusra sum daal jaaeyga leetcode jaise
	root.Data = sum
	sum += oldData

	ReplaceSumWithLargest(root.Left)

}
