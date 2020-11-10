package binarytree

//FindElement : It will return whether the element is present in the tree or not and path for it
func FindElement(root *Node, i interface{}, ans *[]Node) bool {
	if root == nil {
		return false
	}
	if root.Data == i {
		(*ans) = append((*ans), *root)
		return true
	}
	findInLeft := FindElement(root.Left, i, ans)
	if findInLeft {
		(*ans) = append((*ans), *root)
		return true
	}
	findInRight := FindElement(root.Right, i, ans)
	if findInRight {
		(*ans) = append((*ans), *root)
		return true
	}
	return false
}
