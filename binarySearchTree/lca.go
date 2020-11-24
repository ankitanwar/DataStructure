//lowest Common Ancestor

package binarysearchtree

var array []*Node

//LCA : To find the loswest common ancestor
func LCA(root, first, second *Node) int {
	if root.Data.(int) > first.Data.(int) && root.Data.(int) > second.Data.(int) {
		return LCA(root.Left, first, second)
	} else if root.Data.(int) < first.Data.(int) && root.Data.(int) < second.Data.(int) {
		return LCA(root.Right, first, second)
	}
	return root.Data.(int)
}
