package graph

import (
	"container/heap"
	"fmt"
	"strconv"
)

type dijkstra struct {
	value int
	cost  int
	path  string
}

type array []dijkstra

func (a array) Len() int           { return len(a) }
func (a array) Less(i, j int) bool { return a[i].cost < a[j].cost }
func (a array) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (a *array) Push(value interface{}) {
	*a = append(*a, value.(dijkstra))
}
func (a *array) Pop() interface{} {
	t := *a
	x := t[len(t)-1]
	*a = t[:len(t)-1]
	return x

}

//Dijkstra : Given a weighted graph we need to find out the shortest path between source to destination in terms of weight
func Dijkstra(graph map[int][][]int, source int) {
	q := &array{}
	t := dijkstra{
		value: source,
		cost:  0,
		path:  strconv.Itoa(source) + "--",
	}
	visited := make(map[int]bool)
	for key := range graph {
		visited[key] = false
	}
	heap.Init(q)
	heap.Push(q, t)
	for q.Len() > 0 {
		removed := heap.Pop(q).(dijkstra)
		if visited[removed.value] == true {
			continue
		} else {
			visited[removed.value] = true
			fmt.Println(removed.path, removed.cost)
			current := graph[removed.value]
			for i := 0; i < len(current); i++ {
				if visited[current[i][0]] == false {
					t := dijkstra{
						value: current[i][0],
						cost:  removed.cost + current[i][1],
						path:  removed.path + strconv.Itoa(current[i][0]) + "--",
					}
					heap.Push(q, t)
				}
			}

		}

	}
}
