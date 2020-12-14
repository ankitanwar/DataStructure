package backtracking

import "fmt"

//RatInAMaze : To calc the number of ways in which a rat can reach the end of the board
func RatInAMaze(board [][]int) {
	sol := [][]int{}
	for i := 0; i < len(board); i++ {
		t := make([]int, len(board[0]))
		sol = append(sol, t)
	}
	ratMazeHelper(board, 0, 0, &sol)
}

func ratMazeHelper(board [][]int, currentRow, currentCol int, solution *[][]int) {
	if currentRow == len(board)-1 && currentCol == len(board[0])-1 {
		(*solution)[currentRow][currentCol] = 1
		printSolution(*solution)
		return
	}
	if currentRow < 0 || currentRow >= len(board) || currentCol < 0 || currentCol >= len(board[0]) || (*solution)[currentRow][currentCol] == 1 || board[currentRow][currentCol] == 0 {
		return
	}
	(*solution)[currentRow][currentCol] = 1
	ratMazeHelper(board, currentRow+1, currentCol, solution)
	ratMazeHelper(board, currentRow, currentCol+1, solution)
	ratMazeHelper(board, currentRow-1, currentCol, solution)
	ratMazeHelper(board, currentRow, currentCol-1, solution)
	(*solution)[currentRow][currentCol] = 0

}

func printSolution(sol [][]int) {
	fmt.Println(sol)
}
