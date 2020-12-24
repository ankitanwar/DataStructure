package graph

import (
	"fmt"
)

//BellmanFordAlgorithm : Given a weighted directed graph which may contain some negative edge we need to calc the  shortest path in term of weight from source to the destination
func BellmanFordAlgorithm(graph map[int][][]int, source, vertices int) {
	// we need to relax every edge v-1 times
	distance := make(map[int]int)
	for i := 0; i < vertices; i++ {
		if i == source {
			distance[i] = 0
		} else {
			distance[i] = 999999
		}
	}
	for i := 0; i < vertices-1; i++ {
		for key := range graph {
			s := key
			current := graph[s]
			for j := 0; j < len(current); j++ {
				if distance[s]+current[j][1] < distance[current[j][0]] {
					distance[current[j][0]] = distance[s] + current[j][1]
				}
			}
		}
	}
	//Checking for neagative cycle
	var changed bool = false
	for key := range graph {
		s := key
		current := graph[s]
		for j := 0; j < len(current); j++ {
			if distance[s]+current[j][1] < distance[current[j][0]] {
				fmt.Println("The changed is ", distance[current[j][0]], distance[s], current[j][0])
				changed = true
				distance[current[j][0]] = distance[s] + current[j][1]
			}
		}
	}
	if changed == true {
		fmt.Println("Graph contain negative edge cycle")
	} else {
		for key := range distance {
			fmt.Println(source, "--> ", key, "-Cost- ", distance[key])
		}
	}

}
