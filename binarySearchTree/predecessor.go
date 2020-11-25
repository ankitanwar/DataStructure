package binarysearchtree

var lastRight *Node

//Predecessor : To find out the predecesoor of the binary tree
func Predecessor(root, find *Node) *Node {
	if root == nil {
		return nil
	}
	if root == find {
		if root.Left == nil {
			return lastRight
		}
		current := root.Left
		for {
			if current.Right == nil {
				return current
			}
			current = current.Right
		}

	} else if root.Data.(int) < find.Data.(int) {
		lastRight = root
		return Predecessor(root.Right, find)
	}
	return Predecessor(root.Left, find)
}

var lastLeft *Node

//Successor : To find out the Successor of the binary tree
func Successor(root, find *Node) *Node {
	if root == nil {
		return nil
	}
	if root == find {
		if root.Right == nil {
			return lastLeft
		}
		current := root.Right
		for {
			if current.Left == nil {
				return current
			}
			current = current.Left
		}
	} else if root.Data.(int) < find.Data.(int) {
		lastLeft = root
		return Successor(root.Left, find)
	}
	return Successor(root.Right, find)

}
