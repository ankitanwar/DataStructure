package graph

import (
	"fmt"
	"strconv"
)

type bfs struct {
	value int
	path  string
}

//BreadthFirstSearch : BreadthFirst search of a graph in breadth first search we visit all the neightbour first before moving to their neightbours
func BreadthFirstSearch(graph map[int][][]int, source int) {
	queue := []bfs{}
	t := &bfs{
		value: source,
		path:  strconv.Itoa(source) + "--",
	}
	queue = append(queue, *t)
	visited := make(map[int]bool)
	for key := range graph {
		visited[key] = false
	}

	for len(queue) >= 1 {
		removed := queue[0]
		queue = queue[1:]
		if visited[removed.value] == true {
			continue
		}
		visited[removed.value] = true
		fmt.Println(removed.path)
		current := graph[removed.value]
		for i := 0; i < len(current); i++ {
			if visited[current[i][0]] == false {
				t := bfs{
					value: current[i][0],
					path:  removed.path + strconv.Itoa(current[i][0]) + "--",
				}
				queue = append(queue, t)
			}
		}

	}
}
