package binarytree

//PrintChildWithOneNode : it will print all the child who has single node
func PrintChildWithOneNode(child, parent *Node) {
	if child == nil {
		return
	}
	if parent != nil && parent.Left == child && parent.Right == nil {
		println(child.Data.(int))
	} else if parent != nil && parent.Left == child && parent.Right == nil {
		println(child.Data.(int))
	}

	PrintChildWithOneNode(child.Left, child) // child will become parent now nd we will call for left sub tree and then right sub tree
	PrintChildWithOneNode(child.Right, child)
}
