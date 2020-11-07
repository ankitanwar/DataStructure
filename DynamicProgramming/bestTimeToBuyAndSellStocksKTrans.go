package dynamicprogramming

import "math"

//KTransaction : It will return the maximum profit that can be made if K transaction is allowed
func KTransaction(prices []int, k int) int {
	dp := [10][10]int{}
	for i := 0; i <= k; i++ {
		for j := 0; j < len(prices); j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else {
				for k := 0; k < j; k++ {
					dp[i][j] = max(dp[i][j-1], dp[i-1][k]+prices[j]-prices[k], dp[i][j])
				}
			}
		}
	}
	return dp[k][len(prices)-1]
}

//KTransaction2 : Optimizing the dp approach instead of n^3 using n^2
func KTransaction2(prices []int, k int) int {
	dp := [10][10]int{}
	for i := 0; i <= k; i++ {
		currentMax := int(math.Inf(-1))
		for j := 0; j < len(prices); j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else {
				currentMax = maximum(currentMax, dp[i-1][j-1]-prices[j-1])
				if prices[j]+currentMax > dp[i][j-1] {
					dp[i][j] = prices[j] + currentMax
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}

	}
	return dp[k][len(prices)-1]
}
