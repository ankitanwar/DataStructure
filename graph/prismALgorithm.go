package graph

import (
	"container/heap"
	"fmt"
	"strconv"
)

type prism struct {
	weight int
	source string
	value  int
}

type prismHeap []*prism

func (h prismHeap) Len() int           { return len(h) }
func (h prismHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h prismHeap) Less(i, j int) bool { return h[i].weight < h[j].weight }
func (h *prismHeap) Push(value interface{}) {
	x := value.(*prism)
	*h = append(*h, x)
}
func (h *prismHeap) Pop() interface{} {
	x := *h
	val := (x)[len(x)-1]
	*h = x[:len(x)-1]
	return val

}

//PrismAlgorithm : Given a graph we need to tell the minimum weight to connect all the graph
func PrismAlgorithm(graph map[int][][]int, source int) {
	h := &prismHeap{}
	heap.Init(h)
	t := &prism{
		source: strconv.Itoa(source),
	}
	heap.Push(h, t)
	visited := make(map[int]bool)
	for key := range graph {
		visited[key] = false
	}
	var cost int
	for h.Len() > 0 {
		println("The value of heap is ", h.Len())
		removed := heap.Pop(h).(prism)
		if visited[removed.value] == true {
			continue
		} else {
			fmt.Println(removed.source)
			visited[removed.value] = true
			cost += removed.weight
			current := graph[removed.value]
			for i := 0; i < len(current); i++ {
				if visited[current[i][0]] == false {
					t := prism{
						weight: current[i][1],
						source: strconv.Itoa(removed.value),
						value:  current[i][0],
					}
					heap.Push(h, t)
				}
			}
		}
	}
	fmt.Println(cost)

}
