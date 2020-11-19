package binarytree

import "fmt"

//GivenSumFromRootToLeaf : To calculate the path from root to node where sum of all the nodes from root to the node is equal to the sum
func GivenSumFromRootToLeaf(root *Node, targetSum int) {
	stack := []int{}
	rootSumHepler(root, targetSum, 0, &stack)

}

func rootSumHepler(root *Node, targetSum, currentSum int, stack *[]int) {
	if root == nil {
		return
	}
	currentSum += root.Data.(int)
	(*stack) = append((*stack), root.Data.(int))
	if currentSum == targetSum {
		for i := 0; i < len(*stack); i++ {
			fmt.Printf("%v ", (*stack)[i])
		}
		println("")
	}
	rootSumHepler(root.Left, targetSum, currentSum, stack)
	rootSumHepler(root.Right, targetSum, currentSum, stack)
	currentSum -= root.Data.(int)
	*stack = (*stack)[:len(*stack)-1]

}
