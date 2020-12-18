package main

import "github.com/ankitanwar/Golang-DataStructure/graph"

func main() {
	p1 := []int{0, 1}
	p2 := []int{2, 3}
	p3 := []int{4, 5}
	p4 := []int{5, 6}
	p5 := []int{4, 6}
	pairs := [][]int{p1, p2, p3, p4, p5}

	graph.PerfectFriend(7, pairs)
}
