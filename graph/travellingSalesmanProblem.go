package graph

import "fmt"

var totalcost int = 1e7

//TravellingSalesmanProblem : Given a graph we need to select the shortest route(in terms of weight) possible such that evey city is exactly once
func TravellingSalesmanProblem(graph map[int][][]int, vertices int) {
	visited := make(map[int]bool)
	for key := range graph {
		visited[key] = false
	}
	for key := range visited {
		if visited[key] == false {
			backTracking(graph, &visited, 0, key, key, 0, vertices)
		}
	}
	fmt.Println(totalcost)
}

func backTracking(graph map[int][][]int, visited *map[int]bool, cost, source, current, count, n int) {
	if (*visited)[current] == true {
		return
	}
	if count == n {
		if totalcost > cost {
			cost = totalcost
		}
		return
	}

	(*visited)[current] = true
	curr := graph[current]
	for i := 0; i < len(curr); i++ {
		if (*visited)[curr[i][0]] == false {
			backTracking(graph, visited, cost+curr[i][1], source, curr[i][0], count+1, n)
		}
	}
	(*visited)[current] = false
}
