package dynamicprogramming

//LongestCommonSubsequence : To find the common characters between tw0 given strings
func LongestCommonSubsequence(str1, str2 string, m, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	if str1[m-1] == str2[n-1] {
		return 1 + LongestCommonSubsequence(str1, str2, m-1, n-1)
	}
	return maximum(LongestCommonSubsequence(str1, str2, m-1, n), LongestCommonSubsequence(str1, str2, m, n-1))

}

//LongestCommonSubsequence2 : Dp way of doing the lcs
func LongestCommonSubsequence2(str1, str2 string) int {
	dp := [][]int{}
	for i := 0; i <= len(str1); i++ {
		temp := []int{}
		for j := 0; j <= len(str2); j++ {
			temp = append(temp, -1)
		}
		dp = append(dp, temp)
	}

	for i := 0; i <= len(str1); i++ {
		for j := 0; j <= len(str2); j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else if str1[i-1] == str2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = maximum(dp[i-1][j], dp[i][j-1])
			}

		}
	}

	return dp[len(str1)][len(str2)]
}
