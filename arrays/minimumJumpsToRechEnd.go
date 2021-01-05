package array

import "fmt"

//MinimumJumpsToReachEnd : Given an array we need to minimze the number of jumps to reach the end
func MinimumJumpsToReachEnd(arr []int) {
	dp := make([]int, len(arr))
	dp[len(arr)-1] = 0
	for i := len(arr) - 2; i >= 0; i-- {
		if i+arr[i] >= len(arr)-1 {
			dp[i] = 1
		} else {
			var ans int = 999999
			for j := 1; j <= arr[i]; j++ {
				ans = min(ans, 1+dp[i+j])
			}
			dp[i] = ans
		}
	}
	fmt.Println(dp[0])
}
