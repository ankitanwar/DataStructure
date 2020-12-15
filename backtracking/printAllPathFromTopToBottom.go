package backtracking

import (
	"fmt"
	"strconv"
)

//TopToBottomPath : To print all the path from top to bottom
func TopToBottomPath(arr [][]int) {
	helperpath(0, 0, arr, "")
}

func helperpath(currentRow, currentCol int, arr [][]int, asf string) {
	if currentRow == len(arr)-1 {
		for i := currentCol; i < len(arr[0]); i++ {
			asf += strconv.Itoa(arr[currentRow][i])
		}
		fmt.Println(asf)
		return
	}
	if currentCol == len(arr[0])-1 {
		for i := currentRow; i < len(arr); i++ {
			asf += strconv.Itoa(arr[i][currentCol])
		}
		fmt.Println(asf)
		return
	}
	asf += strconv.Itoa(arr[currentRow][currentCol])
	helperpath(currentRow+1, currentCol, arr, asf)
	helperpath(currentRow, currentCol+1, arr, asf)
}
