package graph

import (
	"fmt"
	"strconv"
)

//HamiltonianPath : To count all the hamiltonianPath and edges
func HamiltonianPath(graph map[int][][]int, source int) {
	visited := make(map[int]bool)
	for key := range graph {
		visited[key] = false
	}
	hamiltonianPathHelper(graph, &visited, 0, source, source, strconv.Itoa(source)+"--")
}

func hamiltonianPathHelper(graph map[int][][]int, visited *map[int]bool, count, source, IntialSource int, asf string) {
	if count == len(graph)-1 {
		isCycle := false
		current := graph[source]
		for i := 0; i < len(current); i++ {
			if current[i][0] == IntialSource {
				isCycle = true
				break
			}
		}
		if isCycle == true {
			fmt.Println(asf + "*")
		} else {
			fmt.Println(asf + ".")
		}
		return
	}
	(*visited)[source] = true
	current := graph[source]
	for i := 0; i < len(current); i++ {
		if (*visited)[current[i][0]] == false {
			hamiltonianPathHelper(graph, visited, count+1, current[i][0], IntialSource, asf+strconv.Itoa(current[i][0])+"--")
		}
	}
	(*visited)[source] = false

}
