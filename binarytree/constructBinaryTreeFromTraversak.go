package binarytree

//FromInOrderAndPreorder : to construct the tree from inorder and preoder traversal
func FromInOrderAndPreorder(inorder, preoder *[]interface{}) *Node {
	root := constructTreeHelper(inorder, preoder, 0, len(*preoder)-1)
	return root
}

var index int = 0

func constructTreeHelper(inorder, preorder *[]interface{}, start, end int) *Node {
	if start > end {
		return nil
	}
	node := &Node{
		Data: (*preorder)[index],
	}
	index++
	if start == end {
		return node
	}
	posi := findPositionInOrder(inorder, node.Data, start)
	node.Left = constructTreeHelper(inorder, preorder, start, posi-1)
	node.Right = constructTreeHelper(inorder, preorder, posi+1, end)
	return node

}

func findPositionInOrder(inorder *[]interface{}, value interface{}, start int) int {
	for i := start; i < len(*inorder); i++ {
		if (*inorder)[i] == value {
			return i
		}
	}
	return -99999
}
