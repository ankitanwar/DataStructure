package binarytree

//PrintKLevelDown : It will print k level down
func PrintKLevelDown(root *Node, k int) {
	if root == nil || k == 0 {
		return
	}
	if k == 0 {
		print(root.Data)
	}
	PrintKLevelDown(root.Left, k-1)
	PrintKLevelDown(root.Right, k-1)
}

//PrintNodesKLevelFar : it will print all the nodes all the nodes k level down and up both
func PrintNodesKLevelFar(node *Node, k int) {

}
