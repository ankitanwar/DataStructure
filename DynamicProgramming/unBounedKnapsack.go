package dynamicprogramming

//UnbounedKnapsack : To maximize the profit such that repetation of item is allowed
func UnbounedKnapsack(wt, val []int, capacity, n int) int {
	if n == 0 {
		return 0
	}
	if wt[n-1] <= capacity {
		return maximum(val[n-1]+UnbounedKnapsack(wt, val, capacity-wt[n-1], n), UnbounedKnapsack(wt, val, capacity, n-1))
	}
	return UnbounedKnapsack(wt, val, capacity, n-1)
}

//UnbounedKnapsack2 : dp way of unbouned knapsack
func UnbounedKnapsack2(wt, val []int, capacity int) int {
	matrix := [][]int{}

	for i := 0; i <= len(wt); i++ {
		temp := []int{}
		for j := 0; j <= capacity; j++ {
			temp = append(temp, -1)
		}
		matrix = append(matrix, temp)
	}

	for i := 0; i <= len(wt); i++ {
		for j := 0; j <= capacity; j++ {
			if i == 0 || j == 0 {
				matrix[i][j] = 0
			} else if wt[i-1] <= j {
				matrix[i][j] = maximum(val[i-1]+matrix[i][j-wt[i-1]], matrix[i-1][j])
			} else {
				matrix[i][j] = matrix[i-1][j]
			}
		}
	}

	return matrix[len(wt)][capacity]
}
