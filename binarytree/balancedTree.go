package binarytree

import "math"

//Isbalanced : To check whether the given tree is balanced or not
var Isbalanced bool = true

//BalanceTree : A tree is a balanced tree if difference between heights of left and right subtrees is not more than one for all nodes of tree.
func BalanceTree(root *Node) int {
	println("The value of is balanced is ", Isbalanced)
	if root == nil {
		return 0
	}
	leftHeight := BalanceTree(root.Left)
	rightHeight := BalanceTree(root.Right)

	gap := int(math.Abs(float64(leftHeight - rightHeight)))
	println("The value of gap is ", gap)
	if gap > 1 {
		Isbalanced = false
	}

	return max(leftHeight, rightHeight) + 1
}
