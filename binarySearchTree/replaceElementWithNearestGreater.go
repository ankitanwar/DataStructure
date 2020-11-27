package binarysearchtree

// we have to replace every element with its nearest greater element

var result *Node

//NearestGreater : To print the nearest greater element
func NearestGreater(array []int) {
	ans := make([]int, len(array))
	root := &Node{
		Data: array[len(array)-1],
	}
	for i := len(array) - 2; i >= 0; i-- {
		result = nil
		insert(root, array[i])
		if result != nil {
			ans = append(ans, result.Data.(int))
		} else {
			ans = append(ans, -1)
		}
	}
}

func insert(root *Node, data int) *Node {
	if root == nil {
		newNode := &Node{
			Data: data,
		}
		return newNode
	}
	if root.Data.(int) > data {
		result = root
		root.Left = insert(root.Left, data)
	} else {
		root.Right = insert(root.Right, data)
	}
	return root
}
