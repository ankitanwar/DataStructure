package dynamicprogramming

import (
	"fmt"
	"math"
)

//KnapSack : 0-1 KnapSack
func KnapSack(value, weight []int, capacity, length int) int {
	if capacity == 0 {
		return 0
	}
	if length < 0 {
		return 0
	}
	if weight[length] <= capacity {
		return int(math.Max(float64(value[length]+KnapSack(value, weight, capacity-weight[length], length-1)), float64(KnapSack(value, weight, capacity, length-1))))
	}
	return KnapSack(value, weight, capacity, length-1)

}

//KnapSack2 : Matrix Solution
func KnapSack2(value, weight []int, capacity int) int {
	matrix := [][]int{}

	for i := 0; i <= len(value); i++ {
		temp := []int{}
		for j := 0; j <= capacity; j++ {
			temp = append(temp, -1)
		}
		matrix = append(matrix, temp)
	}

	for i := 0; i <= len(value); i++ {
		matrix[i][0] = 0
	}
	for i := 1; i <= capacity; i++ {
		matrix[0][i] = 0
	}

	for i := 1; i <= len(weight); i++ {
		current := value[i-1]
		for j := 1; j <= capacity; j++ {
			if j <= capacity {
				matrix[i][j] = int(math.Max(float64(current+matrix[i-1][capacity-j]), float64(matrix[i-1][j])))
			} else {
				matrix[i][j] = matrix[i-1][j]
			}
		}
	}

	println("The value of matrix is ")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Printf("%v ", matrix[i][j])
		}
		println("")
	}

	return matrix[len(weight)][capacity]
}
