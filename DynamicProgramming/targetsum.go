package dynamicprogramming

import "fmt"

//TargetSum : To find out all the subsequence whose sum is equal to the target
func TargetSum(arr []int, target int) {
	dp := [][]bool{}

	for i := 0; i <= len(arr); i++ {
		temp := make([]bool, target+1)
		dp = append(dp, temp)
	}

	for i := 1; i < len(dp[0]); i++ {
		dp[0][i] = false
	}
	for i := 0; i < len(dp); i++ {
		dp[i][0] = true
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if j < arr[i-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-arr[i-1]] //bachi hui team and bache hue run
			}
		}
	}
	fmt.Println(dp)
}
