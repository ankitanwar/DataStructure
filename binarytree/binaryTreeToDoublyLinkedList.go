package binarytree

import "fmt"

var prevNode *Node
var head *Node

//ToDoublyLinkedlist : To convert the given binary tree into the doubly linked list
func ToDoublyLinkedlist(root *Node) *Node {
	convertHelper(root)
	print(head)
	return head

}

func convertHelper(root *Node) {
	if root == nil {
		return
	}
	convertHelper(root.Left)
	if prevNode == nil {
		head = root
		prevNode = head
	} else {
		prevNode.Right = root
		root.Left = prevNode
		prevNode = root
	}
	convertHelper(root.Right)

}

func print(root *Node) {
	if root == nil {
		return
	}
	for {
		if root == nil {
			break
		}
		fmt.Println(root.Data)
		root = root.Right
	}
}
