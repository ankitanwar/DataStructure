package backtracking

import "fmt"

// QueenCombinationAs1D : To place n similar queens in nxn board
func QueenCombinationAs1D(n int) {
	board := [][]bool{}
	for i := 0; i < n; i++ {
		t := make([]bool, n)
		board = append(board, t)
	}
	helperQueenCombination1D(&board, 0, n, -1)
}
func helperQueenCombination1D(board *[][]bool, selectedQueen, n, lastPlaced int) {
	if n == selectedQueen {
		for i := 0; i < len(*board); i++ {
			for j := 0; j < len((*board)[0]); j++ {
				if (*board)[i][j] == false {
					fmt.Print("q ")
				} else {
					fmt.Print("- ")
				}
			}
		}
		fmt.Println()
		return
	}
	for i := lastPlaced + 1; i < len(*board)*len(*board); i++ {
		row := i / n
		col := i % n
		(*board)[row][col] = true
		helperQueenCombination1D(board, selectedQueen+1, n, i)
		(*board)[row][col] = false
	}
}
