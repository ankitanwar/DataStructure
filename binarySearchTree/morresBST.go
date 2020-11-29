package binarysearchtree

import "fmt"

//MorresTraversal : to traversal the in inorder manner and
func MorresTraversal(root *Node) {
	if root == nil {
		return

	}
	for {
		if root == nil {
			break
		}
		if root.Left == nil {
			fmt.Println(root.Data)
			root = root.Right
		} else if root.Left != nil {
			temp := root.Left
			for {
				if temp.Right == nil || temp.Right == root {
					break

				}
				temp = temp.Right
			}
			if temp.Right == nil {
				temp.Right = root
				root = root.Left
			} else {
				temp.Right = nil
				fmt.Println(root.Data)
				root = root.Right
			}
		}
	}
}
