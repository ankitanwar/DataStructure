package binarytree

import "fmt"

//ZigZagTraversal : To print the zig zag traversal of the tree
func ZigZagTraversal(root *Node) {
	currentLevel := []queue{}
	nextLevel := []queue{}
	flag := true // if flag is true then we will append left first and then right
	//else we will append right first and then left

	temp := queue{
		value: root,
	}
	currentLevel = append(currentLevel, temp)
	for {
		if len(currentLevel) == 0 {
			break
		}
		current := currentLevel[len(currentLevel)-1]
		currentLevel = currentLevel[:len(currentLevel)-1]
		fmt.Println(current.value.Data)
		if flag == true {
			if current.value.Left != nil {
				temp := queue{
					value: current.value.Left,
				}
				nextLevel = append(nextLevel, temp)
			}
			if current.value.Right != nil {
				temp := queue{
					value: current.value.Right,
				}
				nextLevel = append(nextLevel, temp)
			}
		} else {
			if current.value.Right != nil {
				temp := queue{
					value: current.value.Right,
				}
				nextLevel = append(nextLevel, temp)
			}
			if current.value.Left != nil {
				temp := queue{
					value: current.value.Left,
				}
				nextLevel = append(nextLevel, temp)
			}

		}
		if len(currentLevel) == 0 {
			currentLevel, nextLevel = nextLevel, currentLevel
			flag = !flag
		}

	}
}
