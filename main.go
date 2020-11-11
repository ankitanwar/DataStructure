package main

import (
	"babbar/binarytree"
)

func main() {
	array := make([]interface{}, 0)
	array = append(array, 50, 25, 12, nil, nil, 37, nil, nil, 75, 62, nil, nil, 87, nil, nil)
	root := binarytree.ConstructBinaryTree(array)
	binarytree.LevelOrder(root)
	binarytree.ZigZagTraversal(root)

}
