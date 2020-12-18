package graph

import (
	"fmt"
	"strconv"
)

//PrintAllPath : Given a path we have to print all the paths that lead to the given destion
func PrintAllPath(graph map[int][][]int, source, destination int) {
	visited := make(map[int]bool)
	for key := range graph {
		visited[key] = false
	}
	helperPrintAllPath(graph, &visited, 0, 6, strconv.Itoa(source)+"--")
}

func helperPrintAllPath(graph map[int][][]int, visited *map[int]bool, source, destination int, asf string) {
	if source == destination {
		fmt.Println(asf)
		return
	}
	(*visited)[source] = true
	current := graph[source]
	for i := 0; i < len(current); i++ {
		if (*visited)[current[i][0]] != true {
			helperPrintAllPath(graph, visited, current[i][0], destination, asf+strconv.Itoa(current[i][0])+"--")
		}
	}
	(*visited)[source] = false

}
