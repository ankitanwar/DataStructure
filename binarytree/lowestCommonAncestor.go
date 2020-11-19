package binarytree

//LowestCommonAncestor : To find the lost common parent between node1 and node2
func LowestCommonAncestor(root *Node, node1, node2 int) *Node {
	path1 := []Node{}
	first := FindPath(root, node1, &path1)
	path2 := []Node{}
	second := FindPath(root, node2, &path2)
	if first == false || second == false {
		return nil
	}
	i := 0
	for {
		if i >= len(path1) || i >= len(path2) || path1[i] != path2[i] {
			break
		}
		i++
	}
	return &path1[i-1]
}

//FindPath : to find the path from root to the given node
func FindPath(root *Node, toFind int, path *[]Node) bool {
	if root == nil {
		return false
	}
	(*path) = append((*path), *root)
	if root.Data.(int) == toFind {

		return true
	}

	if (root.Left != nil && FindPath(root.Left, toFind, path)) || (root.Right != nil && FindPath(root.Right, toFind, path)) {
		return true
	}
	(*path) = (*path)[:len(*path)-1]
	return false

}
