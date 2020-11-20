package main

import (
	"babbar/binarytree"
	"fmt"
)

func main() {
	array := make([]interface{}, 0)
	array = append(array, 5, 6, 8, nil, nil, 9, nil, nil, 7, 10, nil, nil, 11, nil, nil)
	// array = append(array, 1, 2, 4, nil, nil, 5, nil, nil, 3, 6, nil, nil, 7, nil, nil)
	// //array = append(array, 20, 8, 5, nil, nil, 3, 10, nil, nil, 14, nil, nil, 22, nil, 25, nil, nil)
	root := binarytree.ConstructBinaryTree(array)
	binarytree.LevelOrder(root)
	ans := binarytree.BTreeToBST(root)
	fmt.Println("The number of swaps required is ", ans)

}
