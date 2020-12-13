package dynamicprogramming

import "fmt"

// CountDistinctSubsequences : To count the number of distince subsequneces in a string
func CountDistinctSubsequences(s string) {
	lastOccurance := make(map[string]int)
	dp := make([]int, len(s)+1)
	dp[0] = 1 //empty string ka subsequence 1 hota hai
	for i := 1; i < len(dp); i++ {
		t := 2 * dp[i-1]
		char := string(s[i-1])
		_, found := lastOccurance[char]
		if found {
			t -= dp[lastOccurance[char]-1]
		} else {
			lastOccurance[char] = i
		}
		dp[i] = t
	}
	fmt.Println(dp)
	fmt.Println(dp[len(dp)-1])
}
