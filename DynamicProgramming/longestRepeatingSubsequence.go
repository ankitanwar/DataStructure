package dynamicprogramming

//LongestRepeatingSubsequence : To print all those character in the string who has occured more than one times
func LongestRepeatingSubsequence(str1, str2 string, m, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if str1[m-1] == str2[n-1] && m != n {
		return 1 + LongestRepeatingSubsequence(str1, str2, m-1, n-1)
	}
	return maximum(LongestRepeatingSubsequence(str1, str2, m-1, n), LongestRepeatingSubsequence(str1, str2, m, n-1))
}

//LongestRepeatingSubsequence2 : To find the out the longest repeating subsequence
func LongestRepeatingSubsequence2(str string) int {
	dp := [][]int{}

	for i := 0; i <= len(str); i++ {
		temp := []int{}
		for j := 0; j <= len(str); j++ {
			temp = append(temp, -1)
		}
		dp = append(dp, temp)
	}

	for i := 0; i <= len(str); i++ {
		for j := 0; j <= len(str); j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else if str[i-1] == str[j-1] && i != j {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = maximum(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(str)][len(str)]
}
