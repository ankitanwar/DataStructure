package dynamicprogramming

import "fmt"

//MaximumPathSum : Given a matrix we need to find out the maximum path sum to reach the end row starting for the first row from any column
func MaximumPathSum(matrix [][]int) {
	/*
		Matrix [i+1] [j]
		Matrix [i+1] [j-1]
		Matrix [i+1] [j+1]

		These are the only movies allowed
	*/

	ans := [][]int{}
	for i := 0; i < len(matrix); i++ {
		t := make([]int, len(matrix[0]))
		ans = append(ans, t)
	}
	for row := len(matrix) - 1; row >= 0; row-- {
		for col := len(matrix[0]) - 1; col >= 0; col-- {
			if row == len(matrix)-1 {
				ans[row][col] = matrix[row][col]
			} else if col == len(ans[0])-1 {
				ans[row][col] = matrix[row][col] + maximum(ans[row+1][col], ans[row+1][col-1])
			} else if col == 0 {
				ans[row][col] = matrix[row][col] + maximum(ans[row+1][col], ans[row+1][col+1])
			} else {
				ans[row][col] = matrix[row][col] + max(ans[row+1][col], ans[row+1][col+1], ans[row+1][col-1])
			}
		}
	}
	var res int
	for i := 0; i < len(ans[0]); i++ {
		if ans[0][i] > res {
			res = ans[0][i]
		}
	}
	fmt.Println(res)
}
