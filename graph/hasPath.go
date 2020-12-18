package graph

import "fmt"

//HasPath : We have to determine whether there is path from source to the destination
func HasPath(graph map[int][][]int, source, target int) {
	nodes := make(map[int]bool)
	for key := range graph {
		nodes[key] = false
	}
	ans := helperHasPath(graph, &nodes, 6, 0)
	fmt.Println(ans)
}

func helperHasPath(graph map[int][][]int, visited *map[int]bool, target, source int) bool {
	if source == target {
		return true
	}
	(*visited)[source] = true
	current := graph[source]
	for i := 0; i < len(current); i++ {
		if (*visited)[current[i][0]] != true {
			found := helperHasPath(graph, visited, target, current[i][0])
			if found {
				return true
			}
		}
	}
	return false

}
