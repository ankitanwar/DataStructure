package backtracking

import "fmt"

//NQUeenPermutation : To place n different queens in a board of nxn such that no queen can kill each other
func NQUeenPermutation(n int) {
	board := [][]bool{}
	for i := 0; i < n; i++ {
		t := make([]bool, n)
		board = append(board, t)
	}
	nQueenPermutationHelper(&board, 0, n)
}

func nQueenPermutationHelper(board *[][]bool, selected int, n int) {
	if selected == n {
		for i := 0; i < len(*board); i++ {
			for j := 0; j < len(*board); j++ {
				if (*board)[i][j] == false {
					fmt.Print("- ")
				} else {
					fmt.Print("q ")
				}
			}
			fmt.Println()
		}
		fmt.Println()
		fmt.Println()
		return
	}
	for i := 0; i < len(*board)*len(*board); i++ {
		row := i / 4
		col := i % 4
		if (*board)[row][col] == false {
			safe := isSafeForQueen(*board, row, col)
			if safe {
				(*board)[row][col] = true
				nQueenPermutationHelper(board, selected+1, n)
				(*board)[row][col] = false
			}

		}
	}
}
