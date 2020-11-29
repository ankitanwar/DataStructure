package main

import (
	binarysearchtree "github.com/ankitanwar/Golang-DataStructure/binarySearchTree"
)

func main() {
	array := make([]interface{}, 0)
	array = append(array, 8, 5, 1, 7, 10, 12)
	// array = append(array, 1, 2, 4, nil, nil, 5, nil, nil, 3, 6, nil, nil, 7, nil, nil)
	// //array = append(array, 20, 8, 5, nil, nil, 3, 10, nil, nil, 14, nil, nil, 22, nil, 25, nil, nil)
	root := binarysearchtree.BSTfromPreOrder(array)
	binarysearchtree.LevelOrder(root)

}
