package dynamicprogramming

import "fmt"

//MinimumJumpsToReachEnd : it will return the minimum numbers of steps required to reach the end
func MinimumJumpsToReachEnd(array []int) int {
	dp := make([]int, len(array))
	path := make([]int, len(array))
	for i := 0; i < len(array); i++ {
		dp[i] = 9999
	}
	path[0] = 0

	dp[0] = 0
	for i := 1; i < len(array); i++ {
		for j := 0; j < i; j++ {
			if j+array[j] >= i {
				if dp[i] > 1+dp[j] {
					dp[i] = 1 + dp[j]
					path[i] = array[j]
				}
			}
		}
	}

	fmt.Printf("%v ", path)

	return dp[len(array)-1]
}
