package graph

import "fmt"

//FloydWarshall : Given a graph we need to find out the minimum distance from evey node to all the other nodes in the graph
func FloydWarshall(edges []Edge, size int) {
	matrix := [][]int{}
	for i := 0; i < size; i++ {
		temp := []int{}
		for j := 0; j < size; j++ {
			if i == j {
				temp = append(temp, 0)
			} else {
				temp = append(temp, 1e7)
			}
		}
		matrix = append(matrix, temp)
	}
	for i := 0; i < len(edges); i++ {
		current := edges[i]
		matrix[current.Source-1][current.Destination-1] = current.Weight
	}
	newMatrix := [][]int{}
	for i := 0; i < size; i++ {
		temp := make([]int, size)
		newMatrix = append(newMatrix, temp)
	}
	for i := 0; i < len(matrix); i++ {
		for row := 0; row < len(newMatrix); row++ {
			for col := 0; col < len(newMatrix[0]); col++ {
				if row == col {
					newMatrix[row][col] = 0
				} else if row == i {
					newMatrix[row][col] = matrix[row][col]
				} else if col == i {
					newMatrix[row][col] = matrix[row][col]
				} else {
					newMatrix[row][col] = min(matrix[row][col], matrix[row][i]+matrix[i][col])
				}
			}
		}
		matrix = newMatrix
	}
	fmt.Println(matrix)

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
