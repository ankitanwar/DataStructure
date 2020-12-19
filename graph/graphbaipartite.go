package graph

import (
	"fmt"
	"strconv"
)

type bipartite struct {
	level int
	path  string
	value int
}

//IsBipartite : a bipartite graph is a graph whose vertices can be divided into two disjoint and independent sets and such that every edge connects a vertex in to one in.
func IsBipartite(graph map[int][][]int) {
	vertices := make(map[int]bool)
	visitedLevel := make(map[int]int)
	for key := range graph {
		vertices[key] = false
		visitedLevel[key] = -1
	}
	var check bool = true
	for key := range vertices {
		if vertices[key] == false {
			check = bipartiteHelper(graph, key, 1, &visitedLevel, &vertices)
			if check == false {
				fmt.Println("Graph is not bipartite")
				break
			}
		}
	}
	if check == true {
		fmt.Println("Yes Graph is a bipartite")
	}
}

func bipartiteHelper(graph map[int][][]int, source, level int, visitedLevel *map[int]int, vertices *map[int]bool) bool {
	q := []bipartite{}
	t := bipartite{
		value: source,
		path:  strconv.Itoa(source) + "--",
		level: level,
	}
	q = append(q, t)
	for len(q) > 0 {
		removed := q[0]
		q = q[1:]
		if (*vertices)[removed.value] == true {
			if (*visitedLevel)[removed.value] != removed.level {
				return false
			}
		} else {
			(*vertices)[removed.value] = true
			current := graph[removed.value]
			for i := 0; i < len(current); i++ {
				if (*vertices)[current[i][0]] == false {
					t := bipartite{
						value: current[i][0],
						path:  strconv.Itoa(current[i][0]) + "--",
						level: removed.level + 1,
					}
					q = append(q, t)
				}
			}
		}
	}

	return true
}
