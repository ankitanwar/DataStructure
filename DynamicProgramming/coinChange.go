package dynamicprogramming

import "fmt"

//CoinChange : To count the numbers of way in which we can form the following sum
func CoinChange(array []int, total, length int) int {
	// Given a value N, find the number of ways to make change for N cents, if we have infinite supply of each of S = { S1, S2, .. , SM } valued coins

	if length < 0 && total != 0 {
		return 0
	}
	if total == 0 && length >= 0 {
		fmt.Println("the value of length is ", length)
		return 1
	}
	if array[length] > total {
		return CoinChange(array, total, length-1)
	}
	return CoinChange(array, total-array[length], length) + CoinChange(array, total, length-1)

}

//CoinChange2 : CoinChange using matrix
func CoinChange2(array []int, total int) int {

	//Preparing the matrix
	matrix := [][]int{}
	for i := 0; i <= len(array); i++ {
		temp := []int{}
		for j := 0; j <= total; j++ {
			temp = append(temp, -1)
		}
		matrix = append(matrix, temp)
	}

	for i := 0; i <= len(array); i++ {
		matrix[i][0] = 1
	}
	for i := 1; i < len(matrix[0]); i++ {
		matrix[0][i] = 0
	}

	for i := 1; i <= len(array); i++ {
		current := array[i-1]
		for j := 1; j < len(matrix[0]); j++ {
			if j >= current {
				matrix[i][j] = matrix[i][j-current] + matrix[i-1][j]
			} else {
				matrix[i][j] = matrix[i-1][j]
			}
		}
	}
	return matrix[len(array)][total]

}
