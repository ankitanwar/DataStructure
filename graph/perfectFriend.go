package graph

import "fmt"

/*
1. You are given a number n (representing the number of students). Each student will have an id
    from 0 to n - 1.
2. You are given a number k (representing the number of clubs)
3. In the next k lines, two numbers are given separated by a space. The numbers are ids of
    students belonging to same club.
4. You have to find in how many ways can we select a pair of students such that both students are
    from different clubs.


*/

//PerfectFriend : To calc the number of ways in which we can pair the friends from different clubs such that no two friends belong to the same club
func PerfectFriend(n int, pairs [][]int) {
	var connections [][]int
	var ans int
	graph := make(map[int][]int)
	for i := 0; i < n; i++ {
		graph[i] = []int{}
	}
	vistied := make(map[int]bool)
	for i := 0; i < len(pairs); i++ {
		graph[pairs[i][0]] = append(graph[pairs[i][0]], pairs[i][1])
		graph[pairs[i][1]] = append(graph[pairs[i][1]], pairs[i][0])
		vistied[pairs[i][0]] = false
		vistied[pairs[i][1]] = false
	}
	for key := range vistied {
		if vistied[key] == false {
			conn := []int{}
			perfectFriendHelper(graph, key, &vistied, &conn)
			connections = append(connections, conn)
		}
	}
	for i := 0; i < len(connections)-1; i++ {
		for j := i + 1; j < len(connections); j++ {
			ans += len(connections[i]) * len(connections[j])
		}
	}
	fmt.Println("The number of ways in which we can pair the friends are", ans)
}

func perfectFriendHelper(graph map[int][]int, source int, visited *map[int]bool, connected *[]int) {
	if (*visited)[source] == true {
		return
	}
	(*visited)[source] = true
	(*connected) = append((*connected), source)
	neighbours := graph[source]
	for i := 0; i < len(neighbours); i++ {
		if (*visited)[neighbours[i]] == false {
			perfectFriendHelper(graph, neighbours[i], visited, connected)
		}
	}
}
