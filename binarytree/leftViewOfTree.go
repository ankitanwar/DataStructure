package binarytree

import "fmt"

//LeftViewOfTree : It will give you the left view of the tree --> left nodes only
func LeftViewOfTree(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.Data)
	LeftViewOfTree(root.Left)
}
