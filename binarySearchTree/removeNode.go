package binarysearchtree

//if it has no child directly remove the node
//if it has one child attach that child to parent of that node
//it it has two child remove the calc the max value from "left" and then remove the max value from the tree and replace the value of removed node with maxleft value

//RemoveNode : To remove the node from the BST
func RemoveNode(root, remove *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Data.(int) < remove.Data.(int) {
		root.Right = RemoveNode(root.Right, remove)
	} else if root.Data.(int) > remove.Data.(int) {
		root.Left = RemoveNode(root.Left, remove)
	} else {
		if root.Left != nil && root.Right != nil {
			max := Max(root.Left)
			root.Data = max.Data
			root.Left = RemoveNode(root.Left, max)
			return root
		} else if root.Left != nil {
			return root.Left
		} else if root.Right != nil {
			return root.Right
		} else {
			return nil
		}
	}

	return root
}
