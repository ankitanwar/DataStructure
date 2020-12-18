package graph

import "fmt"

//IsGraphConnected : To verify Whether the given graph is connected or not
func IsGraphConnected(graph map[int][][]int) {
	//The ideas is get to get the length of all connected component and if the length of all connected compontents is 1 then whole grpah is connected otherwise or not
	var connected [][]int
	vertices := make(map[int]bool)
	for key := range graph {
		vertices[key] = false
	}
	for key := range vertices {
		if vertices[key] == false {
			t := []int{}
			helperGetConnected(graph, key, &vertices, &t)
			connected = append(connected, t)
		}
	}
	if len(connected) > 1 {
		fmt.Println("Graph is not connected")
	} else {
		println("Graph is connected")
	}
}
