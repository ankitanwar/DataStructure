package binarytree

import "fmt"

var k int

//PrintKthAncestor : To print the KthAncestor of the tree
func PrintKthAncestor(root *Node, target, number int) bool {
	k = number
	ans := ancestorHelper(root, target)
	return ans

}

func ancestorHelper(root *Node, target int) bool {
	if root == nil {
		return false
	}
	if root.Data.(int) == target {
		k--
		return true
	}
	ans := ancestorHelper(root.Left, target) || ancestorHelper(root.Right, target)
	if ans == true {
		k = k - 1
		if k == 0 {
			fmt.Println(root.Data)
			return false
		}
		return true

	}
	return false
}
