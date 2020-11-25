package binarysearchtree

import (
	"fmt"
)

//ValidateBST : To validate whether the given tree is bst or not
func ValidateBST(root *Node) bool {
	maxx := 99999999999
	minn := -999999999
	return inorderTraversal(root, maxx, minn)
}

func inorderTraversal(root *Node, maxx, minn int) bool {
	if root == nil {
		return true
	}
	if root.Data.(int) < minn || root.Data.(int) > maxx {
		fmt.Println("This condition is true", root.Data.(int) > maxx, root.Data.(int), maxx)
		return false
	}
	return inorderTraversal(root.Left, root.Data.(int), minn) && inorderTraversal(root.Right, maxx, root.Data.(int))
}
