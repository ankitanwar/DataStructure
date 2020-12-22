package graph

import "fmt"

//TopologicalSort : we have to print all the vertex such that for every vertex u , v if v is dependent on u then u should come first. This algorithm is helpful when we have to print the order of something or one problem is dependent upon other
func TopologicalSort(graph map[int][][]int) {
	//graph is a directed graph

	visited := make(map[int]bool)
	for key := range graph {
		visited[key] = false
	}
	s := []int{}
	for key := range visited {
		if visited[key] == false {
			topologicalSortHelper(graph, key, &visited, &s)
		}
	}
	fmt.Println(s)

}

func topologicalSortHelper(graph map[int][][]int, source int, visited *map[int]bool, stack *[]int) {
	(*visited)[source] = true
	current := graph[source]
	for i := 0; i < len(current); i++ {
		if (*visited)[current[i][0]] == false {
			topologicalSortHelper(graph, current[i][0], visited, stack)
		}
	}
	(*stack) = append(*stack, source)
}
