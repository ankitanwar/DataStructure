package main

import "github.com/ankitanwar/Golang-DataStructure/graph"

func main() {
	start := []int{1, 1, 1, 2, 2, 2, 3, 4, 4, 5, 6, 7, 8}
	end := []int{3, 4, 5, 3, 8, 9, 6, 6, 8, 8, 7, 8, 10}
	weight := []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	e := graph.ConstructEdges(start, end, weight)
	g := graph.ConstructDirectedGraph(e)
	graph.MinimumTimeTakenToCompleteJobs(g)

}
