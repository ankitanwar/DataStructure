package main

import "github.com/ankitanwar/Golang-DataStructure/graph"

func main() {
	arr := [][]int{{1, 0, 0, 0},
		{1, 1, 0, 1},
		{1, 1, 0, 0},
		{0, 1, 1, 1}}
	graph.RatInAMaze(arr)
}
