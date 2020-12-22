package graph

import (
	"fmt"
)

//WiredConnections : given computers some of them are connected to each other and some of them are not and we have to connect them together
func WiredConnections(n int, connections [][]int) int {
	source := []int{}
	destination := []int{}
	for i := 0; i < len(connections); i++ {
		source = append(source, connections[i][0])
		destination = append(destination, connections[i][1])
	}
	g := makeGraph(source, destination)
	fmt.Println("the value of graph is ", g)
	var c int
	visited := make(map[int]bool)
	for i := 0; i < n; i++ {
		visited[i] = false
	}
	for key := range visited {
		if visited[key] == false {
			c++
			getComponents(g, key, &visited)
		}
	}
	if c <= 1 {
		return 0
	}
	r := len(connections) - ((n - 1) - (c - 1))
	if r >= c-1 {
		return c - 1
	}
	return -1

}

func makeGraph(source, destionation []int) map[int][][]int {
	graph := make(map[int][][]int)
	for i := 0; i < len(source); i++ {
		d1 := []int{source[i]}
		d2 := []int{destionation[i]}
		graph[source[i]] = append(graph[source[i]], d2)
		graph[destionation[i]] = append(graph[destionation[i]], d1)
	}
	return graph
}

func getComponents(graph map[int][][]int, source int, visited *map[int]bool) {
	(*visited)[source] = true
	current := graph[source]
	for i := 0; i < len(current); i++ {
		if (*visited)[current[i][0]] == false {
			getComponents(graph, current[i][0], visited)
		}
	}
}
