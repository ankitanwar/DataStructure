package binarytree

import (
	"fmt"
	"strconv"
)

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
			lvl += strconv.Itoa(current.Value.Data.(int)) + "->"
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
