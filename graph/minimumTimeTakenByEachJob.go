package graph

import "fmt"

type job struct {
	time  int
	value int
}

//MinimumTimeTakenToCompleteJobs : Given  a cyclic graph and we need to determine the time taken by each job to be completed edge U,V represents the jobs U and V such that job V can only be started after completion of job U
func MinimumTimeTakenToCompleteJobs(graph map[int][][]int) {
	visited := make(map[int]bool)
	for key := range graph {
		visited[key] = false
	}

	//We can make the indegree vertex at the time of creating the directed graph also
	inDegre := make(map[int]int)
	for key := range graph {
		_, found := inDegre[key]
		if !found {
			inDegre[key] = 0
		}
		current := graph[key]
		for i := 0; i < len(current); i++ {
			val := current[i][0]
			_, found := inDegre[val]
			if found {
				inDegre[val]++
			} else {
				inDegre[val] = 1
			}
		}
	}
	q := []job{}
	for key := range inDegre {
		if inDegre[key] == 0 {
			t := job{
				time:  1,
				value: key,
			}
			q = append(q, t)
		}
	}
	for len(q) > 0 {
		removed := q[0]
		q = q[1:]
		fmt.Printf("%v %v", removed.value, removed.time)
		println()
		current := graph[removed.value]
		for i := 0; i < len(current); i++ {
			inDegre[current[i][0]]--
			if inDegre[current[i][0]] == 0 {
				t := job{
					time:  removed.time + 1,
					value: current[i][0],
				}
				q = append(q, t)
			}
		}
	}

}
