package binarysearchtree

//Baiscally there are two approcaches availale to this solution first is traverse the bst while checking its bst or not and when min and max becomes equal it means its an dead end

// second is we traverse the tree and at the leaf node we check whether the root.Data-1 &&  root.Data+1 is present in the tree or not

var leafNode = make(map[int]int)
var nodes = make(map[int]int)

//ConatainsDeadEnd : To check whether the given tree contains dead end or not
func ConatainsDeadEnd(root *Node) bool {
	traverseTheTree(root)
	for key := range leafNode {
		if _, found := nodes[leafNode[key]-1]; found {
			if _, find := nodes[leafNode[key]+1]; find {
				return true
			}
		}
	}
	return false

}

//traverseTheTree : to traverse the tree
func traverseTheTree(root *Node) {
	if root == nil {
		return
	}
	nodes[root.Data.(int)] = 1
	if root.Left == nil && root.Right == nil {
		leafNode[root.Data.(int)] = 1
	}
	traverseTheTree(root.Left)
	traverseTheTree(root.Right)
}
