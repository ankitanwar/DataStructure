package binarytree

//BTreeToBST : It will convert the binary tree to bst tree
func BTreeToBST(root *Node) int {
	//Idea is do the inorder traversal and store it into the array and calc the number of swaps required to sort the bst

	arr := []int{}

	var NoOfSwap int

	queue := []Node{}

	temp := &Node{
		Data: root,
	}
	queue = append(queue, *temp)

	for {
		if len(queue) == 0 {
			break
		}
		count := len(queue)
		for i := 0; i < count; i++ {
			current := queue[0]
			queue = queue[1:]
			arr = append(arr, current.Data.(int))
			if current.Left != nil {
				temp := &Node{
					Data: current.Left,
				}
				queue = append(queue, *temp)
			}
			if current.Right != nil {
				temp := &Node{
					Data: current.Right,
				}
				queue = append(queue, *temp)
			}
		}
	}

	myDict := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		myDict[i] = arr[i]
	}

	return NoOfSwap

}
