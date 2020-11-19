package main

import (
	"babbar/binarytree"
)

func main() {
	array := make([]interface{}, 0)
	// array = append(array, 8, 3, 1, nil, nil, nil, 10, 6, 4, nil, nil, 7, nil, nil, 14, 13, nil, nil, nil)
	//array = append(array, 1, 2, 4, nil, nil, 5, nil, nil, 3, 6, nil, 8, nil, nil, 7, nil, nil)
	array = append(array, 1, 3, 2, nil, nil, 1, 1, nil, nil, nil, -1, 4, 1, nil, nil, 2, nil, nil, 5, nil, 6, nil, nil)
	root := binarytree.ConstructBinaryTree(array)
	binarytree.LevelOrder(root)
	binarytree.PrintAllPath(root, 5)

}
