package graph

import (
	"container/heap"
	"fmt"
	"math"
	"strconv"

	"github.com/ankitanwar/Golang-DataStructure/heaps"
)

/*

1. You are given a graph, a src vertex and a destination vertex.
2. You are give a number named "weight" and a number "kth largest".
3. You are required to find and print the values of
3.1 Smallest path and it's weight separated by an "@"
3.2 Largest path and it's weight separated by an "@"
3.3 Just Larger path (in terms of given weight) and it's weight separated by an "@"
3.4 Just smaller path (in terms of given weight weight) and it's weight separated by an "@"
3.5 Kth largest path and it's weight separated by an "@"

*/

var criteria int

var smallestPathWeight int = 9999999999
var smallestPath string

var longestPathWeight int = int(math.Inf(-1))
var longestPath string

var justSmallerWeight int = int(math.Inf(-1))
var justSmallerPath string

var justGreaterWeight int = 99999999
var justGreaterPath string

//MultiQuery : To quey the multiple things
func MultiQuery(graph map[int][][]int, source, destination, givenWeight, k int) {
	criteria = givenWeight
	h := &heaps.MinHeap{}
	heap.Init(h)
	vistied := make(map[int]bool)
	for key := range graph {
		vistied[key] = false
	}
	var current int
	multiQueryHelper(graph, source, destination, k, 0, &vistied, h, strconv.Itoa(source)+"--")
	current = heap.Pop(h).(int)
	fmt.Println("The smallest Path by weight is ", smallestPathWeight, smallestPath)
	fmt.Println("The largest path by weight is ", longestPathWeight, longestPath)
	fmt.Println("The just larger path than the given weight is ", justGreaterWeight, justGreaterPath)
	fmt.Println("The just smaller path then the given weight is ", justSmallerWeight, justSmallerPath)
	fmt.Println("The Kth largest weight to reach destionation is ", current)
}

func multiQueryHelper(graph map[int][][]int, source, destinantion, k, currentWeight int, visited *map[int]bool, h *heaps.MinHeap, asf string) {
	if source == destinantion {
		if h.Len() == k {
			current := heap.Pop(h).(int)
			if current < currentWeight {
				heap.Push(h, currentWeight)
			} else {
				heap.Push(h, current)
			}
		} else {
			heap.Push(h, currentWeight)
		}
		if smallestPathWeight > currentWeight {
			smallestPathWeight = currentWeight
			smallestPath = asf
		}
		if longestPathWeight < currentWeight {
			longestPathWeight = currentWeight
			longestPath = asf
		}
		if currentWeight > criteria && currentWeight < justGreaterWeight {
			justGreaterWeight = currentWeight
			justGreaterPath = asf
		}
		if currentWeight < criteria && currentWeight > justSmallerWeight {
			justSmallerWeight = currentWeight
			justSmallerPath = asf
		}

		return
	}
	(*visited)[source] = true
	current := graph[source]
	for i := 0; i < len(current); i++ {
		if (*visited)[current[i][0]] == false {
			multiQueryHelper(graph, current[i][0], destinantion, k, currentWeight+current[i][1], visited, h, asf+strconv.Itoa(current[i][0])+"--")
		}
	}
	(*visited)[source] = false

}
