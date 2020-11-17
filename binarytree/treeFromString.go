package binarytree

import (
	"strconv"
	"unicode"
)

//ConstructBinaryTreeFromString : To construct the binary tree from the string
func ConstructBinaryTreeFromString(s string) *Node {
	root := helper(s, 0)
	return root
}

func helper(s string, index int) *Node {
	if index >= len(s) {
		return nil
	}
	negative := false
	if string(s[index]) == "-" {
		negative = true
		index++
	}
	num := 0
	for {
		if index >= len(s) || unicode.IsDigit(rune(s[index])) == false {
			break
		}
		temp, _ := strconv.Atoi(string(s[index]))
		println(temp)
		num = num*10 + temp
		index++
	}
	if negative == true {
		num = -num
	}
	root := &Node{
		Data: num,
	}
	if index >= len(s) {
		return root
	}
	if index < len(s) && string(s[index]) == "(" {
		index++
		root.Left = helper(s, index)
	}
	if index < len(s) && string(s[index]) == ")" {
		index++
		return root
	}
	if index < len(s) && string(s[index]) == "(" {
		index++
		root.Right = helper(s, index)
	}
	if index < len(s) && string(s[index]) == ")" {
		index++
		return root
	}
	println("The value at current index is ", string(s[index]), index)
	return root

}
