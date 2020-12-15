package backtracking

import (
	"fmt"
)

//KnightTour : To help the knigh to tour all the chessboard without visiting the same spot twice
func KnightTour(n int) {
	sol := [][]int{}
	for i := 0; i < n; i++ {
		t := []int{}
		for j := 0; j < n; j++ {
			t = append(t, -1)
		}
		sol = append(sol, t)
	}
	helperKnight(n, 1, 0, 0, &sol)
}

func helperKnight(n, count, row, col int, solution *[][]int) {
	if count == (n * n) {
		fmt.Println(solution)
		return
	}
	if row < 0 || row >= len(*solution) || col < 0 || col >= len((*solution)[0]) || (*solution)[row][col] != -1 {
		return
	}
	(*solution)[row][col] = count
	helperKnight(n, count+1, row-2, col+1, solution)
	helperKnight(n, count+1, row-1, col+2, solution)
	helperKnight(n, count+1, row+1, col+2, solution)
	helperKnight(n, count+1, row+2, col+1, solution)
	helperKnight(n, count+1, row+2, col-1, solution)
	helperKnight(n, count+1, row+1, col-2, solution)
	helperKnight(n, count+1, row-1, col-2, solution)
	helperKnight(n, count+1, row-2, col-1, solution)
	(*solution)[row][col] = -1

}
