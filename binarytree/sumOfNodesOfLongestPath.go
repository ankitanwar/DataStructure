package binarytree

import "math"

var maxpath int = int(math.Inf(-1))
var maxsum int = int(math.Inf(-1))

//SumOfLongestPathNode : It will return the sum of Longest path from root to leaf if tree have 2 common longest path then we will take the path with largest sum
func SumOfLongestPathNode(root *Node) int {
	sumHelper(root, 0, 0)
	return maxsum
}

func sumHelper(root *Node, path, sum int) {
	if root == nil {
		if path > maxpath {
			maxpath = path
			maxsum = sum
			return
		} else if maxsum < sum {
			maxsum = sum
			return
		}
		return
	}
	sumHelper(root.Left, path+1, sum+root.Data.(int))
	sumHelper(root.Right, path+1, sum+root.Data.(int))
}
