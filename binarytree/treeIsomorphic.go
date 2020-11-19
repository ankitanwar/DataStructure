package binarytree

//TreeIsomorphic : To determine whether the Two Trees are isomorphic or not
func TreeIsomorphic(root1, root2 *Node) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}

	//There are two possible cases at the node
	//1 -> They have been flipped
	//2 -> They have not been flipped
	return root1.Data == root2.Data && (TreeIsomorphic(root1.Left, root2.Left) && TreeIsomorphic(root1.Right, root2.Right)) && (TreeIsomorphic(root1.Left, root2.Right) && TreeIsomorphic(root1.Right, root2.Left))
}
