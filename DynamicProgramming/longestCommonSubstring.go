package dynamicprogramming

//LongestSubstring : To count the length of the max longest substring possible
func LongestSubstring(str1, str2 string, m, n, count int) int {
	if m == 0 || n == 0 {
		return count
	}
	if str1[m-1] == str2[n-1] {
		return LongestSubstring(str1, str2, m-1, n-1, count+1)
	}
	return maximum(count, maximum(LongestSubstring(str1, str2, m-1, n, 0), LongestSubstring(str1, str2, m, n-1, 0)))

}

//LongestSubstring2 : Dp approach for finding the longest substring
func LongestSubstring2(str1, str2 string) int {
	matrix := [100][100]int{}
	var ans int

	for i := 0; i <= len(str1); i++ {
		for j := 0; j <= len(str2); j++ {
			if i == 0 || j == 0 {
				matrix[i][j] = 0
			} else if str1[i-1] == str2[j-1] {
				matrix[i][j] = 1 + matrix[i-1][j-1]
				ans = maximum(ans, matrix[i][j])
			} else {
				matrix[i][j] = 0
			}
		}
	}

	return ans
}
