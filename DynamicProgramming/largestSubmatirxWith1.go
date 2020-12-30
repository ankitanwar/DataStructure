package dynamicprogramming

import "fmt"

//LargestSquareMatrix : Given a matrix we need to find out the largest sqaure matrix with all ones
func LargestSquareMatrix(matrix [][]byte) {
	ans := [][]int{}
	for i := 0; i < len(matrix); i++ {
		t := make([]int, len(matrix[0]))
		ans = append(ans, t)
	}
	var res int = 0
	for i := len(matrix) - 1; i >= 0; i-- {
		for j := len(matrix[0]) - 1; j >= 0; j-- {
			if matrix[i][j] == 0 {
				ans[i][j] = 0
			} else if i == len(matrix)-1 {
				ans[i][j] = int(matrix[i][j])
				if ans[i][j] > res {
					res = ans[i][j]
				}
			} else if j == len(matrix[0])-1 {
				ans[i][j] = int(matrix[i][j])
				if ans[i][j] > res {
					res = ans[i][j]
				}
			} else {
				m := min(ans[i][j+1], ans[i+1][j], ans[i+1][j+1])
				ans[i][j] = 1 + m
				if ans[i][j] > res {
					res = ans[i][j]
				}
			}
		}
	}
	fmt.Println(ans)
	fmt.Println(res)
}
