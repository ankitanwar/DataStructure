package binarysearchtree

//BSTfromPreOrder : To construct the bst from pre-order traversal
func BSTfromPreOrder(array []interface{}) *Node {
	var maxx int
	maxx = 999999999
	root := &Node{
		Data: array[0],
	}
	treeBuilderHelper(array, root, 1, 0, maxx)
	return root
}

func treeBuilderHelper(array []interface{}, root *Node, currentIndex, small, large int) *Node {
	if currentIndex == len(array) || array[currentIndex].(int) < small || array[currentIndex].(int) > large {
		println("This condition is true", array[currentIndex].(int))
		return nil
	}
	if root.Data.(int) > array[currentIndex].(int) {
		root.Left = &Node{
			Data: array[currentIndex],
		}
		currentIndex++
		treeBuilderHelper(array, root.Left, currentIndex, small, root.Data.(int))
	}
	if root.Data.(int) < array[currentIndex].(int) {
		root.Right = &Node{
			Data: array[currentIndex],
		}
		currentIndex++
		treeBuilderHelper(array, root.Right, currentIndex, root.Data.(int), large)
	}
	return root
}
