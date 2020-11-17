package main

import (
	"babbar/binarytree"
)

func main() {
	array := make([]interface{}, 0)
	// array = append(array, 8, 3, 1, nil, nil, nil, 10, 6, 4, nil, nil, 7, nil, nil, 14, 13, nil, nil, nil)
	array = append(array, 2, 1, 11, nil, nil, nil, 11, 1, nil, nil, nil)
	root := binarytree.ConstructBinaryTree(array)
	binarytree.LevelOrder(root)
	count := binarytree.FindDuplicateSubtrees(root)
	println(count)

}
