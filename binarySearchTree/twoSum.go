package binarysearchtree

var newDict = make(map[int]int)
var final bool

//FindTarget : To find the value if there exist a pair such that  their sum is equal to the k
func FindTarget(root *Node, k int) bool {
	toFound := k - root.Data.(int)
	_, found := newDict[toFound]
	if found {
		final = true
		return final
	}
	newDict[root.Data.(int)] = 1
	var left bool
	var right bool
	if root.Right != nil {
		right = FindTarget(root.Right, k)
	}
	if root.Left != nil {
		left = FindTarget(root.Left, k)
	}
	return final || left || right
}
