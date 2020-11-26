package binarysearchtree

import "sort"

var data []interface{}

//BinaryTtoBST : To convert normal binary tree to bst
func BinaryTtoBST(root *Node) *Node {
	//-> Idea is to first traverse the tree and store the value of nodes in the array and then sort the array and build the binary tree out of that
	traversingTheTree(root)
	sort.Slice(data, func(i, j int) bool {
		return data[i].(int) < data[j].(int)
	})
	Newroot := ConstructTree(data, 0, len(data)-1)
	return Newroot
}

func traversingTheTree(root *Node) {
	if root == nil {
		return
	}
	traversingTheTree(root.Left)
	data = append(data, root.Data.(int))
	traversingTheTree(root.Right)
}
