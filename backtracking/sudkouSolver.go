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
			place := canPlaceNumber(*board, row, col, i)
			if place {
				(*board)[row][col] = i
				sudokuHelper(board, nextRow, nextCol)
				(*board)[row][col] = 0
			}
		}
	} else {
		sudokuHelper(board, nextRow, nextCol)
	}

}

func canPlaceNumber(board [][]int, row, col, val int) bool {
	for i := 0; i < len(board[0]); i++ {
		if board[row][i] == val {
			return false
		}
	}
	for i := 0; i < len(board); i++ {
		if board[i][col] == val {
			return false
		}
	}
	startRow := 3 * (row / 3)
	starCol := 3 * (col / 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[startRow+i][starCol+j] == val {
				return false
			}
		}
	}

	return true
}
