package dynamicprogramming

import (
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

	for i := 0; i <= len(weight); i++ {
		for j := 0; j <= capacity; j++ {
			if i == 0 || j == 0 {
				matrix[i][j] = 0
			} else if weight[i-1] <= j {
				matrix[i][j] = maximum(value[i-1]+matrix[i-1][j-weight[i-1]], matrix[i-1][j])
			} else {
				matrix[i][j] = matrix[i-1][j]
			}
		}
	}

	return matrix[len(weight)][capacity]
}
