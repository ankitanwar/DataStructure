package graph

import (
	"fmt"
	"strconv"
)

//DepthFirstSearch : To impelement the depth first search of the given graph
func DepthFirstSearch(graph map[int][][]int) {
	visited := make(map[int]bool)
	for key := range graph {
		visited[key] = false
	}
	for key := range visited {
		if visited[key] == false {
			depthFirstSearchHelper(graph, &visited, key, strconv.Itoa(key)+"--")
		}
	}

}

func depthFirstSearchHelper(graph map[int][][]int, visited *map[int]bool, source int, asf string) {
	fmt.Println(asf)
	(*visited)[source] = true
	current := graph[source]
	for i := 0; i < len(current); i++ {
		if (*visited)[current[i][0]] == false {
			depthFirstSearchHelper(graph, visited, current[i][0], asf+strconv.Itoa(current[i][0])+"--")
		}
	}
}
