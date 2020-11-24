package binarysearchtree

import "fmt"

//PrintRange : To Print all the values of nodes where data lies between the start and end
func PrintRange(root *Node, start, end int) {
	if root == nil {
		return
	}
	if root.Data.(int) > end {
		PrintRange(root.Left, start, end)
	} else if root.Data.(int) < start {
		PrintRange(root.Right, start, end)
	} else {
		PrintRange(root.Left, start, end)
		fmt.Println(root.Data)
		PrintRange(root.Right, start, end)
	}
}
