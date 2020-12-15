package backtracking

import "fmt"

//PlaceQueenAs1D : We have to place n similar queens in a board of nxn such that they cannot kill each other
func PlaceQueenAs1D(n int) {
	board := [][]bool{}
	for i := 0; i < n; i++ {
		t := make([]bool, n)
		board = append(board, t)
	}
	placeQueenHelper(&board, 0, n, -1)
}

func placeQueenHelper(board *[][]bool, selected, n, lastPlaced int) {
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
	for i := lastPlaced + 1; i < len(*board)*len(*board); i++ {
		row := i / n
		col := i % n
		isSafe := isSafeForQueen(*board, row, col)
		if isSafe {
			(*board)[row][col] = true
			placeQueenHelper(board, selected+1, n, i)
			(*board)[row][col] = false
		}
	}
}

func isSafeForQueen(board [][]bool, row, col int) bool {
	//check horizontal(in same row)
	for i := 0; i < len(board[0]); i++ {
		if board[row][i] == true {
			return false
		}
	}

	// check vertical(in same column)
	for i := 0; i < len(board); i++ {
		if board[i][col] == true {
			return false
		}
	}

	//check up  left
	i := row
	j := col
	for {
		if i < 0 || j < 0 {
			break
		}
		if board[i][j] == true {
			return false
		}
		i--
		j--
	}

	// check up right
	i = row
	j = col
	for {
		if i < 0 || j >= len(board[0]) {
			break
		}
		if board[i][j] == true {
			return false
		}
		i--
		j++
	}

	//check down left
	i = row
	j = col
	for {
		if i >= len(board) || j < 0 {
			break
		}
		if board[i][j] == true {
			return false
		}
		i++
		j--
	}

	//check down right
	i = row
	j = col
	for {
		if i >= len(board) || j >= len(board[0]) {
			break
		}
		if board[i][j] == true {
			return false
		}
		i++
		j++
	}

	return true

}
