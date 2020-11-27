package binarysearchtree

import (
	"fmt"
)

type maxx struct {
	isBST  bool
	max    int
	min    int
	root   *Node
	height int
}

//MaxSumBST : to find out the max sum of the bst
func MaxSumBST(root *Node) int {
	res := helper(root)
	fmt.Printf("%v ", res)
	return res.height
}

func helper(root *Node) *maxx {
	if root == nil {
		temp := &maxx{
			isBST: true,
			max:   -9999,
			min:   99999,
		}
		return temp
	}
	left := helper(root.Left)
	right := helper(root.Right)
	ans := &maxx{
		isBST: left.isBST && right.isBST && (root.Data.(int) > left.max && root.Data.(int) < right.min),
		max:   max(root.Data.(int), max(left.max, right.max)),
		min:   max(root.Data.(int), max(left.min, right.min)),
	}
	if ans.isBST == true {
		ans.root = root
		ans.height = left.height + right.height + 1
	} else if ans.isBST == false && left.height > right.height {
		ans.root = left.root
		ans.height = left.height
	} else {
		ans.root = right.root
		ans.height = right.height
	}
	return ans
}
