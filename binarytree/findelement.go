package binarytree

import "fmt"

//FindElement : It will return whether the element is present in the tree or not and path for it
func FindElement(root *Node, i interface{}, ans *[]int) bool {
	if root == nil {
		return false
	}
	fmt.Printf("The value is %v %T", root.Data, root.Data)
	println("")
	if root.Data == i {
		println("Workd for true")
		(*ans) = append((*ans), root.Data.(int))
		return true
	}
	findInLeft := FindElement(root.Left, i, ans)
	if findInLeft {
		(*ans) = append((*ans), root.Data.(int))
		return true
	}
	findInRight := FindElement(root.Right, i, ans)
	if findInRight {
		(*ans) = append((*ans), root.Data.(int))
		return true
	}
	return false
}
