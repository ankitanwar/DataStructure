package graph

import "fmt"

//PossibleToFinishAllTask : Given a list of task we need to determine whether there is possible to finsih all task or not task[i][0] can be started after task[i][1]
func PossibleToFinishAllTask(task [][]int) {
	graph := make(map[int][]int)
	for i := 0; i < len(task); i++ {
		graph[task[i][0]] = append(graph[task[i][0]], task[i][1])
	}
	fmt.Println(graph)

	//Now we have to detect cycle in a graph
}
