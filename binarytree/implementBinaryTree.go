package binarytree

import (
	"fmt"
	"strconv"
)

//These can have atmost 2 chil
//Dosent necessary follows binary search tree property

//Node : Node of a binary Tree
type Node struct {
	Data  interface{}
	Left  *Node
	Right *Node
	Lis   int
}

type stack struct {
	Value     *Node
	Operation int // 1 -> push to the left
	// 2 -> push to the right
	//3 - > pop from the stack

}

//BinaryTree : it will implement the binary tree
type BinaryTree struct {
	root *Node
}

//ConstructBinaryTree : To construct the binary tree from the given array
func ConstructBinaryTree(array []interface{}) *Node {
	newStack := []stack{}
	root := &Node{
		Data: array[0],
	}
	temp := &stack{
		Value:     root,
		Operation: 1,
	}
	tree := BinaryTree{
		root: root,
	}
	newStack = append(newStack, *temp)
	i := 0
	for {
		if len(newStack) == 0 {
			break
		}

		current := &newStack[len(newStack)-1]

		if current.Operation == 3 {
			newStack = newStack[:len(newStack)-1]

		} else if current.Operation == 1 {
			current.Operation++
			i++
			if array[i] != nil {

				current.Value.Left = &Node{
					Data: array[i],
				}
				temp := &stack{
					Value:     current.Value.Left,
					Operation: 1,
				}
				newStack = append(newStack, *temp)
			} else {
				current.Value.Left = nil
			}
		} else {
			i++
			if array[i] != nil {

				current.Value.Right = &Node{
					Data: array[i],
				}
				temp := &stack{
					Value:     current.Value.Right,
					Operation: 1,
				}
				newStack = append(newStack, *temp)
			} else {
				current.Value.Right = nil
			}
			current.Operation++
		}
	}
	return tree.root
}

//Display : It will display the binary tree
func Display(n *Node) {
	str := ""
	if n == nil {
		return
	}
	if n.Left != nil {
		str += strconv.Itoa(n.Left.Data.(int))
	} else {
		str += "."
	}
	str += "<--" + strconv.Itoa(n.Data.(int)) + "-->"

	if n.Right != nil {
		str += strconv.Itoa(n.Right.Data.(int))
	} else {
		str += "."
	}
	fmt.Println(str)
	Display(n.Left)
	Display(n.Right)
}
