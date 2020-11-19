package binarytree

//DistanceBetween2Node : To find out the distance between the 2 given node
func DistanceBetween2Node(root *Node, node1, node2 int) int {
	//The idea is to find the lowest common ancestor and then calulate the distance between lowest common ancestor to both the node and add the disatnce between them
	lca := LowestCommonAncestor(root, node1, node2)
	first := calulateDistance(lca, node1, 0)
	second := calulateDistance(lca, node2, 0)
	return first + second

}

func calulateDistance(root *Node, node, distance int) int {
	if root == nil {
		return -1
	}
	if root.Data.(int) == node {
		return distance
	}
	d := calulateDistance(root.Left, node, distance+1)
	if d != -1 {
		return d
	}
	d = calulateDistance(root.Right, node, distance+1)
	return d
}
