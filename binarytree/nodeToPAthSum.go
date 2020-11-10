package binarytree

import "strconv"

//PathToNodeSum : It will return the all the path possible from node till end of the tree where sum is equal less than or greater than or equal to the given data
func PathToNodeSum(root *Node, path string, low, high, currentSum int) {
	if root == nil {
		return
	}
	if root.Right == nil && root.Left == nil {
		currentSum += root.Data.(int)
		if currentSum >= low && currentSum <= high {
			println(path + strconv.Itoa(root.Data.(int)) + " ")
		}
	}
	PathToNodeSum(root.Left, path+strconv.Itoa(root.Data.(int))+" ", low, high, currentSum+root.Data.(int))
	PathToNodeSum(root.Right, path+strconv.Itoa(root.Data.(int))+" ", low, high, currentSum+root.Data.(int))

}
