package binarysearchtree

// we have given a binary search tree we need to calulate the maximum bst possible

type maxBST struct {
	isBST bool
	max   int
	min   int
	root  *Node
	Size  int
}

//MAXBSTPossible : it will return the max binary search tree possible if given aa binary tree
func MAXBSTPossible(root *Node) *maxBST {
	if root == nil {
		temp := &maxBST{
			isBST: true,
			max:   -9999999,
			min:   99999999,
			root:  nil,
			Size:  0,
		}
		return temp
	}
	left := MAXBSTPossible(root.Left)
	right := MAXBSTPossible(root.Right)
	newmaxBST := &maxBST{
		isBST: (left.isBST) && (right.isBST) && (root.Data.(int) > left.max && root.Data.(int) < right.min),
		max:   max(root.Data.(int), max(left.max, right.max)),
		min:   min(root.Data.(int), min(left.min, right.min)),
	}
	if newmaxBST.isBST == true {
		newmaxBST.root = root
		newmaxBST.Size = right.Size + left.Size + 1
	} else if newmaxBST.isBST == false && left.Size > right.Size {
		newmaxBST.root = left.root
		newmaxBST.Size = left.Size
	} else {
		newmaxBST.root = right.root
		newmaxBST.Size = right.Size
	}
	return newmaxBST
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
