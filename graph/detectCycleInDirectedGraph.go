package graph

import "fmt"

//DetectCycle : Given a directed graph we need to detect cycle in it
func DetectCycle(directedGraph map[int][][]int) {
	vertices := make(map[int]bool)
	currentStack := make(map[int]bool)
	for key := range directedGraph {
		vertices[key] = false
		currentStack[key] = false
	}
	var check bool = false
	for key := range vertices {
		if vertices[key] == false {
			check = detectCycleHelper(directedGraph, key, &vertices, &currentStack)
			if check == true {
				fmt.Println("Cycle Detected")
				break
			}
		}
	}
	if check == false {
		fmt.Println("There is no cycle in the given directed graph")
	}

}

func detectCycleHelper(graph map[int][][]int, source int, visited, currentStack *map[int]bool) bool {
	(*visited)[source] = true
	(*currentStack)[source] = true
	current := graph[source]
	for i := 0; i < len(current); i++ {
		if (*currentStack)[current[i][0]] == true {
			return true
		}
		detectCycleHelper(graph, current[i][0], visited, currentStack)
	}
	(*currentStack)[source] = false
	return false
}

func detectCycleBFS(vertex int, graph map[int][][]int) {
	indegree := make(map[int]int)
	for key := range graph {
		_, found := indegree[key]
		if !found {
			indegree[key] = 0
		}
		current := graph[key]
		for i := 0; i < len(current); i++ {
			_, found := indegree[current[i][0]]
			if !found {
				indegree[current[i][0]] = 1
			} else {
				indegree[current[i][0]]++
			}
		}
	}
	q := []int{}

	var count int
	for key := range graph {
		if indegree[key] == 0 {
			q = append(q, key)
		}

	}
	for len(q) > 0 {
		count++
		removed := q[0]
		q = q[1:]
		count++
		current := graph[removed]
		for i := 0; i < len(current); i++ {
			indegree[current[i][0]]--
			if indegree[current[i][0]] == 0 {
				q = append(q, current[i][0])
			}
		}
	}
	if count == vertex {
		fmt.Println("There is no cycle in the graph")
	} else {
		println("There exist a cycle in the grah")
	}
}
