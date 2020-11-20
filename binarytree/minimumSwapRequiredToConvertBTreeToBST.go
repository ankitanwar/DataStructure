package binarytree

import (
	"fmt"
	"sort"
)

//BTreeToBST : It will convert the binary tree to bst tree
func BTreeToBST(root *Node) int {
	//Idea is do the inorder traversal and store it into the array and calc the number of swaps required to sort the bst

	arr := []int{}

	var NoOfSwap int

	queue := []*stack{}
	temp := &stack{
		Value:     root,
		Operation: 1,
	}
	queue = append(queue, temp)

	for {
		if len(queue) == 0 {
			break
		}
		current := queue[len(queue)-1]
		if current.Operation == 1 {
			current.Operation++
			if current.Value.Left != nil {
				temp := &stack{
					Value:     current.Value.Left,
					Operation: 1,
				}
				queue = append(queue, temp)
			}
		} else {
			fmt.Println(current.Value.Data.(int))
			arr = append(arr, current.Value.Data.(int))
			queue = queue[:len(queue)-1]
			if current.Value.Right != nil {
				temp := &stack{
					Value:     current.Value.Right,
					Operation: 1,
				}
				queue = append(queue, temp)
			}
		}

	}

	myDict := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		myDict[arr[i]] = i
	}

	copyOfarray := []int{}
	for i := 0; i < len(arr); i++ {
		copyOfarray = append(copyOfarray, arr[i])
	}
	sort.Ints(copyOfarray)
	fmt.Printf("%v ", arr)
	println("")
	fmt.Printf("%v ", copyOfarray)

	for i := 0; i < len(arr); i++ {
		if arr[i] != copyOfarray[i] {
			NoOfSwap++
			init := arr[i]
			arr[i], myDict[copyOfarray[i]] = myDict[copyOfarray[i]], arr[i]
			myDict[init] = myDict[copyOfarray[i]]
			myDict[copyOfarray[i]] = i
		}
	}

	return NoOfSwap

}
