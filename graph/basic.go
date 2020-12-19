package graph

import "fmt"

//Edge : To define the edge of the graph
type Edge struct {
	Source      int
	Destination int
	Weight      int
}

//ConstructEdges : To construct the edges from the given array
func ConstructEdges(source, destination, weight []int) []Edge {
	ans := []Edge{}
	for i := 0; i < len(destination); i++ {
		t := &Edge{
			Source:      source[i],
			Destination: destination[i],
			Weight:      weight[i],
		}
		ans = append(ans, *t)
	}
	return ans
}

//ConstructGraph : To construct the grpah from the given edge array
func ConstructGraph(array []Edge) map[int][][]int {
	graph := make(map[int][][]int)
	for i := 0; i < len(array); i++ {
		current := array[i]
		d1 := []int{current.Destination, current.Weight}
		d2 := []int{current.Source, current.Weight}
		graph[current.Source] = append(graph[current.Source], d1)
		graph[current.Destination] = append(graph[current.Destination], d2)
	}
	return graph
}

//ConstructDirectedGraph : Function to construct a directed graph
func ConstructDirectedGraph(array []Edge) map[int][][]int {
	graph := make(map[int][][]int)
	for i := 0; i < len(array); i++ {
		d1 := []int{array[i].Destination, array[i].Weight}
		graph[array[i].Source] = append(graph[array[i].Source], d1)
	}
	fmt.Println(graph)
	return graph
}
