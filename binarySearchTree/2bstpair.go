package binarysearchtree

//TwoBSTPair : Given two bst we have to find the pair of such that node1 + node 2 == target --> basicallly we follow 2 pointer approach
func TwoBSTPair(root1, root2 *Node, target int) int {
	if root1 == nil || root2 == nil {
		return 0

	}
	var ans int
	stack1 := []*Node{}
	stack2 := []*Node{}
	for {
		for {
			if root1 == nil {
				break
			}
			stack1 = append(stack1, root1)
			root1 = root1.Left
		}
		for {
			if root2 == nil {
				break
			}
			stack1 = append(stack2, root2)
			root1 = root2.Right
		}
		if len(stack1) == 0 || len(stack2) == 0 {
			break
		}
		topStack1 := stack1[len(stack1)-1]
		topStack2 := stack2[len(stack2)-1]
		if topStack1.Data.(int)+topStack2.Data.(int) == target {
			stack1 = stack1[:len(stack1)-1]
			stack2 = stack2[:len(stack2)-1]
			ans++
			root1 = topStack1.Right
			root2 = topStack2.Left
		} else if topStack1.Data.(int)+topStack2.Data.(int) > target {
			stack2 = stack2[:len(stack2)-1]
			root2 = topStack2.Left
		} else {
			stack1 = stack1[:len(stack1)-1]
			root1 = topStack1.Right
		}

	}
	return count
}
