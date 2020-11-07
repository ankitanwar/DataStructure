package backtracking

import (
	"fmt"
	"strconv"
)

//Nqueen : To place queeen in the NxN chess board chess that no queen can kill each other
func Nqueen(chess [][]int, ans string, row int) {
	if row == len(chess) {
		fmt.Println(ans)
		return
	}
	for col := 0; col < len(chess[0]); col++ {
		if isSafeToPlace(chess, row, col) == true {
			chess[row][col] = 1
			Nqueen(chess, ans+strconv.Itoa(row)+"-"+strconv.Itoa(col)+",", row+1)
			chess[row][col] = 0
		}
	}
}

func isSafeToPlace(chess [][]int, row, col int) bool {
	i := row
	j := col

	//vertical check
	for {
		if i < 0 {
			break
		}
		if chess[i][j] == 1 {
			return false
		}
		i--
	}

	//left diagonal check
	i = row
	j = col
	for {
		if i < 0 || j < 0 {
			break
		}
		if chess[i][j] == 1 {
			return false
		}
		i--
		j--
	}

	//Right diagonal check
	i = row
	j = col
	for {
		if i < 0 || j >= len(chess[0]) {
			break
		}
		if chess[i][j] == 1 {
			return false
		}
		i--
		j++
	}

	return true

}
