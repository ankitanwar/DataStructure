package binarytree

import (
	"fmt"
	"strconv"
)

//FindDuplicateSubtrees : To count the number of duplicate sub tree
func FindDuplicateSubtrees(root *Node) []*Node {
	dict := make(map[string]int)
	result := make([]*Node, 0)
	duplicateHelper(root, &dict, &result)
	fmt.Println(len(result))
	return result

}

func duplicateHelper(root *Node, dict *map[string]int, result *[]*Node) string {
	if root == nil {
		return "x"
	}
	left := duplicateHelper(root.Left, dict, result)
	right := duplicateHelper(root.Right, dict, result)
	ans := strconv.Itoa(root.Data.(int)) + "," + left + right
	fmt.Println("The value of ans is ", ans)
	_, found := (*dict)[ans]
	if found {
		fmt.Printf("The Found value is %v ", (*dict)[ans])
		(*dict)[ans]++
		if (*dict)[ans] == 2 {
			(*result) = append((*result), root)
		}
	} else {
		(*dict)[ans] = 1
	}
	return ans
}
