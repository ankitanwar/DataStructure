package main

import (
	"fmt"

	binarysearchtree "github.com/ankitanwar/Golang-DataStructure/binarySearchTree"
)

func main() {
	array := make([]interface{}, 0)
	array = append(array, 1, 2, 3)
	// array = append(array, 1, 2, 4, nil, nil, 5, nil, nil, 3, 6, nil, nil, 7, nil, nil)
	// //array = append(array, 20, 8, 5, nil, nil, 3, 10, nil, nil, 14, nil, nil, 22, nil, 25, nil, nil)
	root := binarysearchtree.ConstructTree(array, 0, len(array)-1)
	ans := binarysearchtree.MaxSumBST(root)
	fmt.Println("The ans is ", ans)
	binarysearchtree.LevelOrder(root)
	fmt.Println(ans)

}
