package dynamicprogramming

import "fmt"

//CountPalindromeSubstring : To count the number of palindrome substring in a given string
func CountPalindromeSubstring(s string) {
	var count int
	dp := [][]bool{}
	for i := 0; i < len(s); i++ {
		t := make([]bool, len(s))
		dp = append(dp, t)
	}
	for g := 0; g < len(s); g++ {
		col := g
		row := 0
		for {
			if row == len(s) || col == len(s) {
				break
			}
			if g == 0 {
				dp[row][col] = true
			} else if g == 1 {
				if string(s[row]) == string(s[col]) {
					dp[row][col] = true
					count++

				} else {
					dp[row][col] = false
				}
			} else {
				if string(s[row]) == string(s[col]) {
					if dp[row+1][col-1] == false {
						dp[row][col] = false
					} else {
						dp[row][col] = true
						count++
					}
				} else {
					dp[row][col] = false
				}
			}

			row++
			col++

		}
	}
	fmt.Println("The count is ", count)
}
