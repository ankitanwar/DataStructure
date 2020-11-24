package binarysearchtree

//Find : To find the particaulr value in a binary tree
func Find(root *Node, data interface{}) *Node {
	if root.Data == data {
		return root
	}
	if root.Data.(int) < data.(int) {
		return Find(root.Right, data)
	}
	return Find(root.Left, data)
}
