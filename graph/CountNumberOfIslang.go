package graph

import "fmt"

//NumberOfIsland : We have given a matrix of 0 and 1 and we need to count the number of island 0 is land and 1 is water
func NumberOfIsland(matrix [][]int) {
	var count int
	visited := [][]bool{}
	for i := 0; i < len(matrix); i++ {
		t := make([]bool, len(matrix[0]))
		visited = append(visited, t)
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 1 && visited[i][j] == false {
				helperCountIsland(matrix, i, j, &visited)
				count++
			}
		}
	}
	fmt.Println("The number of islands are ", count)
}

func helperCountIsland(matrix [][]int, row, col int, visited *[][]bool) {
	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) || (*visited)[row][col] == true || matrix[row][col] == 0 {
		return
	}
	(*visited)[row][col] = true
	helperCountIsland(matrix, row, col-1, visited)
	helperCountIsland(matrix, row, col+1, visited)
	helperCountIsland(matrix, row-1, col, visited)
	helperCountIsland(matrix, row+1, col, visited)

}
