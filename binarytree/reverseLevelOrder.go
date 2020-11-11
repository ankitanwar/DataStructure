package binarytree

import "fmt"

//ReverseLevelOrder : To print the level order of the traversal of the tree in reverse order
type queue struct {
	value *Node
}

//ReverseLevelOrder : To reverse the level order transversal of the tree
func ReverseLevelOrder(root *Node) {
	stack := []queue{}
	q := []queue{}

	temp := queue{
		value: root,
	}
	q = append(q, temp)

	for {
		if len(q) == 0 {
			break
		}
		count := len(q)
		for i := 0; i < count; i++ {
			current := q[0]
			stack = append(stack, current)
			q = q[1:]
			if current.value.Left != nil {
				temp := queue{
					value: current.value.Left,
				}
				q = append(q, temp)
			}
			if current.value.Right != nil {
				temp := queue{
					value: current.value.Right,
				}
				q = append(q, temp)
			}

		}

	}

	for {
		if len(stack) == 0 {
			break
		}
		current := stack[len(stack)-1]
		fmt.Printf("%v ", current.value.Data)
		stack = stack[:len(stack)-1]
	}

}
