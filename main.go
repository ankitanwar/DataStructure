package main

import (
	"babbar/binarytree"
)

func main() {
	array := make([]interface{}, 0)
	// array = append(array, 8, 3, 1, nil, nil, nil, 10, 6, 4, nil, nil, 7, nil, nil, 14, 13, nil, nil, nil)
	array = append(array, 1, 2, 4, nil, nil, 5, nil, nil, 3, 6, nil, nil, 7, nil, nil)
	//array = append(array, 20, 8, 5, nil, nil, 3, 10, nil, nil, 14, nil, nil, 22, nil, 25, nil, nil)
	root := binarytree.ConstructBinaryTree(array)
	binarytree.LevelOrder(root)
	binarytree.ToDoublyLinkedlist(root)

}
