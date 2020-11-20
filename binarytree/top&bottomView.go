package binarytree

import "fmt"

//TopView : To print the top view of the binary tree
func TopView(root *Node) {
	//Its a combination of both level order traversal and the vertical order traversal
	dict := make(map[int][]Node)
	q := []*verticalOrder{}
	first := &verticalOrder{
		horizonatalDistance: 0,
		node:                root,
	}
	q = append(q, first)
	for {
		if len(q) == 0 {
			break
		}
		count := len(q)
		for i := 0; i < count; i++ {
			current := q[0]
			q = q[1:]
			_, found := dict[current.horizonatalDistance]
			if !found {
				fmt.Println(current.node.Data)
				dict[current.horizonatalDistance] = append(dict[current.horizonatalDistance], *current.node)
			}
			if current.node.Left != nil {
				temp := &verticalOrder{
					horizonatalDistance: current.horizonatalDistance - 1,
					node:                current.node.Left,
				}
				q = append(q, temp)
			}
			if current.node.Right != nil {
				temp := &verticalOrder{
					horizonatalDistance: current.horizonatalDistance + 1,
					node:                current.node.Right,
				}
				q = append(q, temp)
			}
		}

	}
}

//BottomView : To print the top view of the binary tree
func BottomView(root *Node) {
	//Its a combination of both level order traversal and the vertical order traversal and the node which have come last in the array is the bottom view
	dict := make(map[int][]Node)
	q := []*verticalOrder{}
	first := &verticalOrder{
		horizonatalDistance: 0,
		node:                root,
	}
	q = append(q, first)
	for {
		if len(q) == 0 {
			break
		}
		count := len(q)
		for i := 0; i < count; i++ {
			current := q[0]
			q = q[1:]
			dict[current.horizonatalDistance] = append(dict[current.horizonatalDistance], *current.node)
			if current.node.Left != nil {
				temp := &verticalOrder{
					horizonatalDistance: current.horizonatalDistance - 1,
					node:                current.node.Left,
				}
				q = append(q, temp)
			}
			if current.node.Right != nil {
				temp := &verticalOrder{
					horizonatalDistance: current.horizonatalDistance + 1,
					node:                current.node.Right,
				}
				q = append(q, temp)
			}
		}

	}
	for key := range dict {
		fmt.Printf("%v ", dict[key][len(dict[key])-1].Data.(int))
	}
}
