package backtracking

import (
	"fmt"
	"strconv"
)

//QueenPermutation2 : Given nxn board and n distinct queen we want to output in how many ways we can place these queens in our boxes
func QueenPermutation2(n int) {
	used := make(map[int]bool)
	queenPermuhelper2(0, 0, 0, n, "", &used)
}

func queenPermuhelper2(row, col, selected, n int, asf string, used *map[int]bool) {
	if row == n {
		if selected == n {
			fmt.Println(asf)
			println()
			println()
		}
		return
	}
	currentRow := 0
	currentCol := 0
	printer := ""
	if col+1 == n {
		currentRow = row + 1
		currentCol = 0
		printer = "\n"
	} else {
		currentRow = row
		currentCol = col + 1
		printer = " "
	}
	for i := 0; i < n; i++ {
		_, found := (*used)[i+1]
		if !found {
			(*used)[i+1] = true
			queenPermuhelper2(currentRow, currentCol, selected+1, n, asf+strconv.Itoa(i+1)+printer, used)
			delete(*used, i+1)
		}
	}
	queenPermuhelper2(currentRow, currentCol, selected, n, asf+"-"+printer, used)

}
