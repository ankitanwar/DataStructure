package dynamicprogramming

//NumberOfWaysToReachScore : It will return the maximum number of ways possible to reach the given socre
func NumberOfWaysToReachScore(arr []int, score, n int) int {
	if n == 0 {
		return 0
	}
	if score == 0 && n != 0 {
		return 1
	}
	if score >= arr[n-1] {
		return NumberOfWaysToReachScore(arr, score-arr[n-1], n) + NumberOfWaysToReachScore(arr, score, n-1)
	}
	return NumberOfWaysToReachScore(arr, score, n-1)
}

//NumberOfWaysToReachScore2 : It will return the maximum number of ways possible to reach the given socre
func NumberOfWaysToReachScore2(arr []int, score int) int {
	matrix := [100][100]int{}

	for i := 0; i <= len(arr); i++ {
		for j := 0; j <= score; j++ {
			if i == 0 {
				matrix[i][j] = 0
			} else if j == 0 {
				matrix[i][j] = 1
			} else if j >= arr[i-1] {
				matrix[i][j] = matrix[i][j-arr[i-1]] + matrix[i-1][j]
			} else {
				matrix[i][j] = matrix[i-1][j]
			}
		}
	}

	return matrix[len(arr)][score]

}
