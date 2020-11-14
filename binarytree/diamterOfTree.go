package binarytree

//Diameter : it will return the diameter of the tree
func Diameter(root *Node) int {
	if root == nil {
		return 0
	}
	left := Diameter(root.Left)
	right := Diameter(root.Right)
	factor := Height(root.Left) + Height(root.Right) + 2
	ans := maximum(left, right, factor)
	return ans

}

func maximum(a, b, c int) int {
	if a > b && a > c {
		return a
	} else if b > a && b > c {
		return b
	}
	return c

}

//Efficient Solution to find the diameter

//DiaHeight : It contains pair of diameter and height both
type DiaHeight struct {
	diameter int
	height   int
}

func diameterOfBinaryTree(root *Node) int {
	ans := solve(root)
	return ans.diameter
}

func solve(root *Node) *DiaHeight {
	if root == nil {
		temp := DiaHeight{
			diameter: 0,
			height:   -1,
		}
		return &temp
	}
	left := solve(root.Left)
	right := solve(root.Right)
	temp := &DiaHeight{
		diameter: maximum(left.diameter, right.diameter, left.height+right.height+2),
		height:   max(left.height, right.height) + 1,
	}
	return temp
}
