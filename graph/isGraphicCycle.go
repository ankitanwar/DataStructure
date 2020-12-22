package graph

import (
	"fmt"
	"strconv"
)

//IsGraphCyclic : To check Whether the given graph is cyclic or not
func IsGraphCyclic(graph map[int][][]int) {
	vertices := make(map[int]bool)
	for key := range graph {
		vertices[key] = false
	}
	var ans bool
	//var path string
	//In case the graph is not connected and we need to check wether the graph is cyclic or not
	for key := range vertices {
		if vertices[key] == false {
			ans = dfsApproach(graph, &vertices, -1, key)
			if ans == true {
				fmt.Println("There is a cycle in the graph")
				break
			}
		}
	}
	if ans != true {
		fmt.Println("There is no cycle in the graph")
	}

}

func helperDetectCycle(graph map[int][][]int, vertices *map[int]bool, source int) (string, bool) {
	q := []bfs{}
	t := bfs{
		value: source,
		path:  strconv.Itoa(source) + "--",
	}
	q = append(q, t)
	for len(q) > 0 {
		removed := q[0]
		q = q[1:]
		if (*vertices)[removed.value] == true {
			return removed.path, true
		}
		(*vertices)[removed.value] = true
		current := graph[removed.value]
		for i := 0; i < len(current); i++ {
			if (*vertices)[current[i][0]] == false {
				t := bfs{
					value: current[i][0],
					path:  removed.path + strconv.Itoa(current[i][0]) + "--",
				}
				q = append(q, t)
			}
		}
	}

	return "", false
}

func dfsApproach(graph map[int][][]int, visited *map[int]bool, parent, child int) bool {
	(*visited)[child] = true
	current := graph[child]
	for i := 0; i < len(current); i++ {
		if (*visited)[current[i][0]] == false {
			dfsApproach(graph, visited, child, current[i][0])
		} else if (*visited)[current[i][0]] == true && current[i][0] != parent {
			return true // Yes there exist an cycle in the graph
		}
	}
	return false

}
