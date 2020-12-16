package backtracking

import (
	"fmt"
	"os"
)

//ShortestSafestPath : To print the safest and shortest path from col[0] to last col
func ShortestSafestPath(board [][]int) {
	var ans int = 999999
	final := markUnsafe(board)
	sol := [][]bool{}
	for i := 0; i < len(board); i++ {
		t := make([]bool, len(board[0]))
		sol = append(sol, t)
	}
	for i := 0; i < len(board); i++ {
		var temp int = 0
		found := calcDistance(final, &sol, 0, i, &temp)
		if found && temp < ans {
			ans = temp
		}
	}
	fmt.Println(ans)
}

func markUnsafe(board [][]int) [][]int {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 0 {
				if i-1 >= 0 && j >= 0 {
					board[i][j] = 0
				}
				if i >= 0 && j-1 >= 0 {
					board[i][j] = 0
				}
				if i >= 0 && j+1 < len(board[0]) {
					board[i][j] = 0
				}
				if i+1 < len(board) && j >= 0 {
					board[i][j] = 0
				}
			}
		}
	}
	return board
}

func calcDistance(board [][]int, sol *[][]bool, col, row int, distance *int) bool {
	if col == len(board[0])-1 {
		fmt.Println(sol)
		fmt.Print(ans)
		os.Exit(0)
		return true
	}
	if col < 0 || col >= len(board[0]) || row < 0 || row >= len(board) || board[row][col] == 0 || (*sol)[row][col] == true {
		return false
	}
	(*sol)[row][col] = true
	(*distance) = *distance + 1
	left := calcDistance(board, sol, col-1, row, distance)
	right := calcDistance(board, sol, col+1, row, distance)
	up := calcDistance(board, sol, col, row-1, distance)
	down := calcDistance(board, sol, col, row+1, distance)
	(*sol)[row][col] = false
	return left || right || up || down
}
