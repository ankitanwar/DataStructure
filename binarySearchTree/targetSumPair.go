package binarysearchtree

import "fmt"

//TargetSumPair : To print all the pair of nodes whose sum is equal to the target
func TargetSumPair(root, curretNode *Node, targetSum int) {
	if root == nil {
		return
	}
	toFind := targetSum - curretNode.Data.(int)
	found := Find(root, toFind)
	if found != nil {
		fmt.Println(toFind, found)
	}
	TargetSumPair(root, curretNode.Left, targetSum)
	TargetSumPair(root, curretNode.Right, targetSum)

}
