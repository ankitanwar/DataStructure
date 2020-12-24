package graph

import "fmt"

//TotalNumberOfSpanningTree : Given a tree we need to determine the
func TotalNumberOfSpanningTree(graph map[int][][]int, vertices int) {
	matrix := [][]int{}
	for i := 0; i < vertices; i++ {
		t := make([]int, vertices)
		matrix = append(matrix, t)
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if i == j {
				current := graph[i+1] // its only for undirected graph for directed graph calc the indegree and outdegree and add them to calc  the total degree
				matrix[i][j] = len(current)
			} else {
				current := graph[i+1]
				for k := 0; k < len(current); k++ {
					matrix[i][current[k][0]-1] = -1
					matrix[current[k][0]-1][i] = -1
				}
			}
		}
	}
	fmt.Println(matrix)
	//ToDo : Calc the determinant of a matrix

}
