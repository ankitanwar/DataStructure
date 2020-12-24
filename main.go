package main

import "github.com/ankitanwar/Golang-DataStructure/graph"

func main() {
	start := []int{1, 1, 3, 2, 2}
	end := []int{3, 4, 4, 4, 3}
	weight := []int{10, 10, 10, 10, 10}
	e := graph.ConstructEdges(start, end, weight)
	g := graph.ConstructGraph(e)
	graph.TotalNumberOfSpanningTree(g, 4)
}
