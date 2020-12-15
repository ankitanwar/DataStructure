package backtracking

import "fmt"

//SudokuSolver : To solve the given sudoku
func SudokuSolver(board [][]int) {
	sudokuHelper(&board, 0, 0)
}

func sudokuHelper(board *[][]int, row, col int) {
	if row == len(*board) {
		fmt.Println(board)
		return
	}
	nextRow := 0
	nextCol := 0
	if col+1 == len((*board)[0]) {
		nextRow = row + 1
		nextCol = 0
	} else {
		nextRow = row
		nextCol = col + 1
	}
	if (*board)[row][col] == 0 {
		for i := 1; i <= 9; i++ {
			(*board)[row][col] = i
			sudokuHelper(board, nextRow, nextCol)
			(*board)[row][col] = 0
		}
	} else {
		sudokuHelper(board, nextRow, nextCol)
	}

}
