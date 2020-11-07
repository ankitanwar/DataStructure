package backtracking

import "fmt"

//To determine whether the word can be placed horizontally or not
func canPlaceHorizonatlly(board [][]string, word string, row, column int) bool {
	if column-1 >= 0 && board[row][column-1] != "+" {
		return false
	}
	if column+len(word) < len(board[0]) && board[row][column+len(word)] != "+" {
		return false
	}
	for i := 0; i < len(word); i++ {
		if column+i >= len(board[0]) {
			return false
		}
		if board[row][column+i] == "-" || board[row][column+i] == string(word[i]) {
			continue
		} else {
			return false
		}
	}
	return true
}

//To determine whether the word can be placed vertically or not
func canPlaceVertically(board *[][]string, word string, row, column int) bool {
	if row-1 >= 0 && (*board)[row-1][column] != "+" {
		return false
	}
	if row+len(word) < len((*board)) && (*board)[row+len(word)][column] != "+" {
		return false
	}
	for i := 0; i < len(word); i++ {
		if row+i >= len((*board)) {
			return false
		}
		if (*board)[row+i][column] == "-" || (*board)[row+i][i] == string(word[i]) {
			continue
		} else {
			return false
		}
	}
	return true
}

func placeHorizontal(board *[][]string, word string, row, column int) *[]bool {
	wePlaced := make([]bool, len(word))
	for i := 0; i < len(word); i++ {
		if (*board)[row][column+i] == "-" {
			wePlaced[i] = true
		}
		(*board)[row][column+i] = string(word[i])
	}

	return &wePlaced
}

func placeVertical(board *[][]string, word string, row, column int) *[]bool {
	wePlaced := make([]bool, len(word))
	for i := 0; i < len(word); i++ {
		if (*board)[row+i][column] == "-" {
			wePlaced[i] = true
		}
		(*board)[row+i][column] = string(word[i])
	}

	return &wePlaced
}

func unPlaceVertical(board *[][]string, wePlace *[]bool, word string, row, column int) {
	for i := 0; i < len(*wePlace); i++ {
		if (*wePlace)[i] == true {
			(*board)[row+i][column] = "-"
		}
	}
}

func unPlaceHorizontal(board *[][]string, wePlace *[]bool, word string, row, column int) {
	for i := 0; i < len(*wePlace); i++ {
		if (*wePlace)[i] == true {
			(*board)[row][column+i] = "-"
		}
	}
}

//CrossWordPuzzle : To solve the cross word puzzle
func CrossWordPuzzle(board [][]string, words []string, currentIndex, boardlength int) {
	if currentIndex == len(words) {
		fmt.Printf("%v", board)
		println("x-x-x-x-x-x-x-x-x--x-x-x-x-x-x--x-x-x-x-x-x--x-x-x-x")
		return
	}
	current := words[currentIndex]
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == "-" || board[i][j] == string(current[0]) {
				if canPlaceHorizonatlly(board, current, i, j) == true {
					placed := placeHorizontal(&board, current, i, j)
					CrossWordPuzzle(board, words, currentIndex+1, boardlength)
					unPlaceHorizontal(&board, placed, current, i, j)
				}
				if canPlaceVertically(&board, current, i, j) == true {
					placed := placeVertical(&board, current, i, j)
					CrossWordPuzzle(board, words, currentIndex+1, boardlength)
					unPlaceVertical(&board, placed, current, i, j)
				}
			}
		}
	}
}
