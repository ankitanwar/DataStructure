package backtracking

import "fmt"

//NKnightCombination : To place n Knight in nxn board such that no knight can kill each other
func NKnightCombination(n int) {
	board := [][]bool{}
	for i := 0; i < n; i++ {
		t := make([]bool, n)
		board = append(board, t)
	}
	helperPlaceKnight(&board, 0, n, -1)
}

func helperPlaceKnight(board *[][]bool, selected, n, lastPlaced int) {
	if selected == n {
		for i := 0; i < len(*board); i++ {
			for j := 0; j < len(*board); j++ {
				if (*board)[i][j] == false {
					fmt.Print("- ")
				} else {
					fmt.Print("k ")
				}
			}
			fmt.Println()
		}
		fmt.Println()
		fmt.Println()
		return
	}
	for i := lastPlaced + 1; i < len(*board)*len(*board); i++ {
		row := i / n
		col := i % n
		safe := isKnightSafe(*board, row, col)
		if safe {
			(*board)[row][col] = true
			helperPlaceKnight(board, selected+1, n, i)
			(*board)[row][col] = false
		}
	}
}

func isKnightSafe(board [][]bool, row, col int) bool {
	if row-2 >= 0 && col+1 < len(board[0]) {
		if board[row-2][col+1] == true {
			return false
		}
	}
	if row-1 >= 0 && col+2 < len(board[0]) {
		if board[row-1][col+2] == true {
			return false
		}
	}
	if row-2 >= 0 && col-1 >= 0 {
		if board[row-2][col-1] == true {
			return false
		}
	}
	if row-2 >= 0 && col-2 >= 0 {
		if board[row-2][col-2] == true {
			return false
		}
	}
	return true
}
