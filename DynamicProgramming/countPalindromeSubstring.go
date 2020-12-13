package dynamicprogramming

import "fmt"

//CountPalindromeSubstring : To count the number of palindrome substring in a given string
func CountPalindromeSubstring(s string) {
	dp := [][]bool{}
	for i := 0; i < len(s); i++ {
		t := make([]bool, len(s))
		dp = append(dp, t)
	}
	row := 0
	for i := 0; i < len(s); i++ {
		newRow := row
		newCol := i
		for newCol < len(s) {
			if i == 0 {
				dp[newRow][newCol] = true
			} else if i == 1 {
				if s[newRow] == s[newCol] {
					dp[newRow][newCol] = true
				} else {
					dp[newRow][newCol] = false
				}
			} else {
				if s[newRow] == s[newCol] && dp[newRow-1][newCol-1] == true {
					dp[newRow][newCol] = true
				} else {
					dp[newRow][newCol] = false
				}
			}

			newRow++
			newCol++
		}

	}
	fmt.Println(dp)
}
