package dynamicprogramming

//LongestChain : To find out the maximum chain that can be made from the given matrix
func LongestChain(matrix [][]int) int {
	var ans int

	dp := []int{}

	for i := 0; i < len(matrix); i++ {
		dp = append(dp, 1)
	}

	for i := 1; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			if matrix[j][1] < matrix[i][0] {
				dp[i] = maximum(dp[i], 1+dp[j])
			}
		}
	}

	for i := 0; i < len(dp); i++ {
		if dp[i] > ans {
			ans = dp[i]
		}
	}
	return ans
}
