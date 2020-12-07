package backtracking

import "fmt"

//QueenPermutation : To find out the number of ways in which we can place n Different queens(q1,q2,q3..) in a board of size n
func QueenPermutation(n int) {
	matrix := [][]int{}
	for i := 0; i < n; i++ {
		temp := []int{}
		for j := 0; j < n; j++ {
			temp = append(temp, 0)
		}
		matrix = append(matrix, temp)
	}
	queenPermuhelper(&matrix, n, 0)
}

func queenPermuhelper(board *[][]int, n, currentSelected int) {
	if currentSelected == n {
		for i := 0; i < len(*board); i++ {
			for j := 0; j < len((*board)[0]); j++ {
				if (*board)[i][j] == 0 {
					fmt.Printf("-")
				} else {
					fmt.Printf("%v ", (*board)[i][j])
				}
			}
			fmt.Println()
		}
		fmt.Println()
		fmt.Println()
		return
	}
	for i := 0; i < len(*board); i++ {
		for j := 0; j < len((*board)[0]); j++ {
			if (*board)[i][j] == 0 {
				(*board)[i][j] = currentSelected + 1
				queenPermuhelper(board, n, currentSelected+1)
				(*board)[i][j] = 0
			}
		}
	}
}
