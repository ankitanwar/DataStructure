package dynamicprogramming

import (
	"fmt"

	"github.com/ankitanwar/Golang-DataStructure/binarytree"
)

//LargestIndependentSubset : We have to find out the max of total number of nodes such that they are not directly connected to each other
func LargestIndependentSubset(root *binarytree.Node) {
	r := lisDP(root)
	fmt.Println(r)
}

func lisRecurrsive(root *binarytree.Node) int {
	if root == nil {
		return 0
	}
	excluded := lisRecurrsive(root.Right) + lisRecurrsive(root.Left)

	included := 1
	if root.Left != nil {
		included += lisRecurrsive(root.Left.Left) + lisRecurrsive(root.Left.Right)
	}
	if root.Right != nil {
		included += lisRecurrsive(root.Right.Left) + lisRecurrsive(root.Right.Right)
	}
	return maximum(excluded, included)
}

func lisDP(root *binarytree.Node) int {
	if root == nil {
		return 0
	}
	if root.Lis != 0 {
		return root.Lis
	}
	excluded := lisRecurrsive(root.Right) + lisRecurrsive(root.Left)

	included := 1
	if root.Left != nil {
		included += lisRecurrsive(root.Left.Left) + lisRecurrsive(root.Left.Right)
	}
	if root.Right != nil {
		included += lisRecurrsive(root.Right.Left) + lisRecurrsive(root.Right.Right)
	}
	root.Lis = maximum(excluded, included)
	return root.Lis

}
