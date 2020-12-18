package graph

import "fmt"

//GetConnectedComponents : Given a graph we need to tell the how many components are connected togther
func GetConnectedComponents(graph map[int][][]int) {
	ans := [][]int{}
	visited := make(map[int]bool)
	for key := range graph {
		visited[key] = false
	}
	for key := range visited {
		if visited[key] == false {
			path := []int{}
			helperGetConnected(graph, key, &visited, &path)
			ans = append(ans, path)
		}
	}
	fmt.Println(ans)
}

func helperGetConnected(graph map[int][][]int, source int, visited *map[int]bool, path *[]int) {
	(*path) = append((*path), source)
	(*visited)[source] = true
	current := graph[source]
	for i := 0; i < len(current); i++ {
		if (*visited)[current[i][0]] == false {
			helperGetConnected(graph, current[i][0], visited, path)
		}
	}
}
