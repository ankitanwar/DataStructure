package binarytree

import "math"

var currentLevel int = int(math.Inf(0))

//AllLeafAtSameLevel : To check whether all the leaf are at same levl or not
func AllLeafAtSameLevel(root *Node, level int) bool {
	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		if currentLevel == int(math.Inf(0)) {
			println("This is working")
			currentLevel = level
			return true
		}
		return level == currentLevel

	}
	return AllLeafAtSameLevel(root.Left, level+1) && AllLeafAtSameLevel(root.Right, level+1)
}
