package binarytree

import "fmt"

//PrintKLevelDown : It will print k level down
func PrintKLevelDown(root, blocker *Node, k int) {
	if root == nil || k < 0 || root == blocker {
		return
	}
	if k == 0 {
		fmt.Printf("%v ", root.Data)
	}
	PrintKLevelDown(root.Left, blocker, k-1)
	PrintKLevelDown(root.Right, blocker, k-1)
}

//PrintNodesKLevelFar : it will print all the nodes all the nodes k level down and up both
func PrintNodesKLevelFar(root *Node, position interface{}, k int) {
	Path := []Node{}
	FindElement(root, position, &Path)
	for i := 0; i < len(Path); i++ {
		if i == 0 { //if i==0 then blocker is null we want all child
			PrintKLevelDown(&Path[i], nil, k-i)
		} else {
			PrintKLevelDown(&Path[i], &Path[i-1], k-i)
		}
	}
}
