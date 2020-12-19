package graph

import "fmt"

type person struct {
	time  int
	value int
}

//SpreadInfection : given a time and a person(vertex of a graph) we need to tell how many people will get the infection in given time if in one unit to time all the neighbour person will get the infection
func SpreadInfection(graph map[int][][]int, source, time int) {
	var ans int
	vertices := make(map[int]bool)
	for key := range graph {
		vertices[key] = false
	}
	q := []person{}
	t := person{
		time:  0,
		value: source,
	}
	q = append(q, t)
	for len(q) > 0 {
		removed := q[0]
		q = q[1:]
		if removed.time == time {
			break
		}
		if vertices[removed.value] == true {
			continue
		} else {
			vertices[removed.value] = true
			ans++
			current := graph[removed.value]
			for i := 0; i < len(current); i++ {
				if vertices[current[i][0]] == false {
					t := person{
						time:  removed.time + 1,
						value: current[i][0],
					}
					q = append(q, t)
				}
			}

		}
	}
	fmt.Println(ans)
}
