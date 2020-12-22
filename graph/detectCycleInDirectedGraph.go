package graph

import "fmt"

//DetectCycle : Given a directed graph we need to detect cycle in it
func DetectCycle(directedGraph map[int][][]int) {
	vertices := make(map[int]bool)
	for key := range directedGraph {
		vertices[key] = false
	}
	var check bool = false
	for key := range vertices {
		if vertices[key] == false {
			check = detectCycleHelper(directedGraph, key, &vertices)
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

func detectCycleHelper(graph map[int][][]int, source int, visited *map[int]bool) bool {
	(*visited)[source] = true
	current := graph[source]
	for i := 0; i < len(current); i++ {
		if (*visited)[current[i][0]] == true {
			return true
		}
		detectCycleHelper(graph, current[i][0], visited)
	}
	(*visited)[source] = false
	return false
}
