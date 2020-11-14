package binarytree

import "fmt"

type view struct {
	node  *Node
	level int
}

//LeftViewOfTree : It will print the left view of the tree
func LeftViewOfTree(root *Node) {
	currentLevel := 0
	q := []view{}
	temp := view{
		node:  root,
		level: 1,
	}
	q = append(q, temp)

	for {
		if len(q) == 0 {
			break
		}
		count := len(q)
		for i := 0; i < count; i++ {
			current := q[0]
			q = q[1:]
			if current.level > currentLevel {
				fmt.Println(current.node.Data)
				currentLevel = current.level
			}
			if current.node.Left != nil {
				temp := view{
					node:  current.node.Left,
					level: current.level + 1,
				}
				q = append(q, temp)
			}
			if current.node.Right != nil {
				temp := view{
					node:  current.node.Right,
					level: current.level + 1,
				}
				q = append(q, temp)
			}
		}
	}
}

//RightViewOfTree : It will print the right view of the tree
func RightViewOfTree(root *Node) {
	currentLevel := 0
	q := []view{}
	temp := view{
		node:  root,
		level: 1,
	}
	q = append(q, temp)

	for {
		if len(q) == 0 {
			break
		}
		count := len(q)
		for i := 0; i < count; i++ {
			current := q[0]
			q = q[1:]
			if current.level > currentLevel {
				fmt.Println(current.node.Data)
				currentLevel = current.level
			}
			if current.node.Right != nil {
				temp := view{
					node:  current.node.Right,
					level: current.level + 1,
				}
				q = append(q, temp)
			}
			if current.node.Left != nil {
				temp := view{
					node:  current.node.Left,
					level: current.level + 1,
				}
				q = append(q, temp)
			}
		}
	}
}
