//Same As BinarySearch Tree

package binarysearchtree

import (
	"fmt"
	"strconv"
)

//DiagonalOrder : struct for diagonal order traversal
type DiagonalOrder struct {
	root  *Node
	order int
}
type stack struct {
	Value     *Node
	Operation int // 1 -> push to the left
	// 2 -> push to the right
	//3 - > pop from the stack

}

//Preorder : It will give the pre oder traversal of the tree
func Preorder(roots *Node) {
	//->Node ->Left ->Right
	if roots == nil {
		return
	}
	fmt.Println(roots.Data)
	Preorder(roots.Left)
	Preorder(roots.Right)
}

//Inorder : It will give the inorder traversal of the tree
func Inorder(root *Node) {
	if root == nil {
		return
	}
	Preorder(root.Left)
	fmt.Println(root.Data)
	Preorder(root.Right)
}

//Postorder : It will return the post oder traversal of the tree
func Postorder(root *Node) {
	if root == nil {
		return
	}
	Preorder(root.Left)
	Preorder(root.Right)
	fmt.Println(root.Data)
}

//IterativeTraversal : It will give you the iterative version of the traversal
func IterativeTraversal(root *Node) {
	p := []stack{}
	temp := &stack{
		Value:     root,
		Operation: 1,
	}
	p = append(p, *temp)

	pre := ""
	in := ""
	post := ""

	for {
		if len(p) == 0 {
			break
		}
		current := &p[len(p)-1]
		if current.Operation == 1 {
			pre += strconv.Itoa(current.Value.Data.(int)) + "-->"
			current.Operation++
			if current.Value.Left != nil {
				temp := &stack{
					Value:     current.Value.Left,
					Operation: 1,
				}
				p = append(p, *temp)
			}
		} else if current.Operation == 2 {
			in += strconv.Itoa(current.Value.Data.(int)) + "-->"
			current.Operation++
			if current.Value.Right != nil {
				temp := &stack{
					Value:     current.Value.Right,
					Operation: 1,
				}
				p = append(p, *temp)
			}
		} else {
			post += strconv.Itoa(current.Value.Data.(int)) + "-->"
			p = p[:len(p)-1]
		}
	}
	fmt.Println(pre)
	fmt.Println(in)
	fmt.Println(post)

}

//LevelOrder : To give the value of all the traversal Preoder Inorder Postoder
func LevelOrder(root *Node) {
	println("Level order invoked")
	q := []stack{}

	lvl := ""

	temp := &stack{
		Value: root,
	}
	q = append(q, *temp)

	for {
		if len(q) == 0 {
			break
		}

		count := len(q)
		for i := 0; i < count; i++ {
			current := q[0]
			lvl += strconv.Itoa(current.Value.Data.(int)) + "-> "
			q = q[1:]
			if current.Value.Left != nil {
				temp := &stack{
					Value: current.Value.Left,
				}
				q = append(q, *temp)
			}
			if current.Value.Right != nil {
				temp := &stack{
					Value: current.Value.Right,
				}
				q = append(q, *temp)
			}

		}
	}

	fmt.Println(lvl)
}

//DiagonalTraversal : It will do the diagonal order traversal of the tree
func DiagonalTraversal(root *Node) {
	if root == nil {
		return
	}

	node := &DiagonalOrder{
		root:  root,
		order: 0,
	}
	q := []DiagonalOrder{}
	myDict := make(map[int][]DiagonalOrder)
	q = append(q, *node)
	for {
		if len(q) == 0 {
			break
		}
		count := len(q)
		for i := 0; i < count; i++ {
			current := q[0]
			q = q[1:]
			myDict[current.order] = append(myDict[current.order], current)
			if current.root.Left != nil {
				temp := &DiagonalOrder{
					root:  current.root.Left,
					order: current.order + 1,
				}
				q = append(q, *temp)
			}
			if current.root.Right != nil {
				temp := &DiagonalOrder{
					root:  current.root.Right,
					order: current.order,
				}
				q = append(q, *temp)
			}
		}
	}
	for key, value := range myDict {
		for i := 0; i < len(value); i++ {
			current := myDict[key][i]
			fmt.Printf("%v ", current.root.Data)
		}
		println("")
	}

}

//BoundaryTraversal : IT will give the boundary traversal of the tree
func BoundaryTraversal(root *Node) {
	leftBoundary(root)
	PrintLeaves(root)
	rightBoundary(root.Right)
}

func leftBoundary(root *Node) {
	if root != nil {
		if root.Left != nil {
			fmt.Println(root.Data)
			leftBoundary(root.Left)
		} else if root.Right != nil {
			fmt.Println(root.Data)
			leftBoundary(root.Left)
		}
	}
}

func rightBoundary(root *Node) {
	if root != nil {
		if root.Right != nil {
			rightBoundary(root.Right)
			fmt.Println(root.Data)
		} else if root.Left != nil {
			rightBoundary(root.Left)
			fmt.Println(root.Data)
		}
	}
}

//PrintLeaves : To print the leave
func PrintLeaves(root *Node) {
	if root != nil {
		PrintLeaves(root.Left)
		if root.Left == nil && root.Right == nil {
			fmt.Println(root.Data)
		}
		PrintLeaves(root.Right)
	}
}

type verticalOrder struct {
	horizonatalDistance int
	node                *Node
}

//VerticalOrderTraversal : to traverse the tree in the vertical order
func VerticalOrderTraversal(root *Node) {
	//we will assign the horizontal distance to each node
	//-> for root 0
	//-> for left its horizontal distance of root -1
	//-> for right its horizontal distance of root +1

	//Do the level order traversal and assign the values of the horizonatal distance and add the values with same horizonatal distance in the dict

	q := []*verticalOrder{}
	first := &verticalOrder{
		horizonatalDistance: 0,
		node:                root,
	}
	dict := make(map[int][]*Node)
	q = append(q, first)
	for {
		if len(q) == 0 {
			break
		}
		count := len(q)
		for i := 0; i < count; i++ {
			current := q[0]
			q = q[1:]
			dict[current.horizonatalDistance] = append(dict[current.horizonatalDistance], current.node)
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
		for i := 0; i < len(dict[key]); i++ {
			fmt.Printf("%v ", dict[key][i].Data.(int))
		}
		println("")

	}

}
