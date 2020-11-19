package binarytree

import "fmt"

var array []int

//PrintAllPath : To print the sum of all the path
func PrintAllPath(root *Node, targetSum int) {
	if root == nil {
		return
	}
	array = append(array, root.Data.(int))
	PrintAllPath(root.Left, targetSum)
	PrintAllPath(root.Right, targetSum)

	currentSum := 0
	for i := len(array) - 1; i >= 0; i-- {
		currentSum += array[i]
		if currentSum == targetSum {
			for j := 0; j < len(array); j++ {
				fmt.Printf("%v ", array[j])
			}
			println("")

		}

	}
	array = array[:len(array)-1]
}
