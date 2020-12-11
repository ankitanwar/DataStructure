package dynamicprogramming

type Node struct {
	data  interface{}
	Left  *Node
	Right *Node
}

//LargestIndependentSubset : We have to find out the max of total number of nodes such that they are not directly connected to each other
func LargestIndependentSubset(root *Node) int {
	if root == nil {
		return 0
	}
	return maximum(1+LargestIndependentSubset(root.Left.Left)+LargestIndependentSubset(root.Right.Right), LargestIndependentSubset(root.Right))

}
