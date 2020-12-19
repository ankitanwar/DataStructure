package graph

import "fmt"

//RatInAMaze : Given a matix and we need to reach the end of the matix and have to print the possible paths to do do
func RatInAMaze(matrix [][]int) {
	visited := [][]bool{}
	for i := 0; i < len(matrix); i++ {
		t := make([]bool, len(matrix[0]))
		visited = append(visited, t)
	}
	ratInMazeHelper(matrix, 0, 0, &visited, "")
}

func ratInMazeHelper(matrix [][]int, row, col int, visited *[][]bool, asf string) {
	if row == len(matrix)-1 && col == len(matrix[0])-1 {
		fmt.Println(asf)
		return
	}
	if row < 0 || row >= len(matrix) || col < 0 || col >= len(matrix[0]) || matrix[row][col] == 0 || (*visited)[row][col] == true {
		return
	}
	(*visited)[row][col] = true
	ratInMazeHelper(matrix, row, col-1, visited, asf+"L")
	ratInMazeHelper(matrix, row, col+1, visited, asf+"R")
	ratInMazeHelper(matrix, row-1, col, visited, asf+"U")
	ratInMazeHelper(matrix, row+1, col, visited, asf+"D")
	(*visited)[row][col] = false
}
