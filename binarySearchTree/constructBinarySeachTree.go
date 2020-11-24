package binarysearchtree

//Node : Node of a tree
type Node struct {
	Data  interface{}
	Right *Node
	Left  *Node
}

//ConstructTree : To construct the BST from the given array
func ConstructTree(arr []interface{}, start, end int) *Node {
	if start > end {
		return nil
	}
	var mid int
	mid = (start + end) / 2
	root := &Node{
		Data: arr[mid],
	}
	if start == end {
		return root
	}
	root.Left = ConstructTree(arr, start, mid-1)
	root.Right = ConstructTree(arr, mid+1, end)
	return root

}
